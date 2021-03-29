package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Scripter interface {
	ScriptQuery() (*ScriptQuery, error)
}

// Script uses a script to provide a custom score for returned documents.
//
// The script_score query is useful if, for example, a scoring function is
// expensive and you only need to calculate the score of a filtered set of
// documents.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html
type Script struct {
	// Query used to return documents. (Required)
	Query Query
	// Script used to compute the score of documents returned by the query.
	// (Required)
	Script string
	// Documents with a score lower than this floating point number are excluded
	// from the search results. (Optional)
	MinScore float64
	// Documents scores produced by script are multiplied by boost to produce
	// final documents' scores. Defaults to 1.0. (Optional)
	Boost  interface{}
	Params interface{}
	Name   string
}

func (s Script) Clause() (Clause, error) {
	return s.ScriptQuery()
}
func (s Script) ScriptQuery() (*ScriptQuery, error) {
	q := &ScriptQuery{}

	err := q.setQuery(s.Query)
	if err != nil {
		return q, NewQueryError(err, KindScript)
	}

	err = q.setScript(s.Script)
	if err != nil {
		return q, NewQueryError(err, KindScript)
	}

	err = q.SetBoost(s.Boost)
	if err != nil {
		return q, NewQueryError(err, KindScript)
	}
	q.SetMinScore(s.MinScore)

	err = q.SetParams(s.Params)
	if err != nil {
		return q, NewQueryError(err, KindScript)
	}

	return q, nil
}

func (Script) Kind() Kind {
	return KindScript
}

type ScriptQuery struct {
	query  QueryValues
	params dynamic.JSON
	script string
	boostParam
	minScoreParam
	nameParam
}

func (ScriptQuery) Kind() Kind {
	return KindScript
}

func (s ScriptQuery) DecodeParams(val interface{}) error {

	return json.Unmarshal(s.params, val)
}

func (s *ScriptQuery) SetParams(params interface{}) error {
	if params == nil {
		s.params = []byte{}
		return nil
	}
	d, err := json.Marshal(params)
	if err != nil {
		return err
	}
	data := dynamic.JSON(d)
	if data.IsNull() {
		s.params = []byte{}
		return nil
	}
	if !data.IsObject() {
		return ErrInvalidParams
	}
	s.params = data
	return nil
}

func (s *ScriptQuery) setScript(script string) error {
	if len(script) == 0 {
		return ErrScriptRequired
	}
	s.script = script
	return nil
}
func (s *ScriptQuery) setQuery(query Query) error {
	qv, err := newQuery(query)
	if err != nil {
		return err
	}
	if qv.IsEmpty() {
		return ErrQueryRequired
	}
	s.query = *qv
	return nil
}
