package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type ScriptScorer interface {
	ScriptScore() (*ScriptScoreQuery, error)
}

// ScriptScore uses a script to provide a custom score for returned documents.
//
// The script_score query is useful if, for example, a scoring function is
// expensive and you only need to calculate the score of a filtered set of
// documents.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html
type ScriptScore struct {
	// Query used to return documents. (Required)
	Query *Query
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

func (s ScriptScore) Clause() (Clause, error) {
	return s.ScriptScore()
}
func (s ScriptScore) ScriptScore() (*ScriptScoreQuery, error) {
	q := &ScriptScoreQuery{}

	err := q.setQuery(s.Query)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}

	err = q.setScript(s.Script)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}

	err = q.SetBoost(s.Boost)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}
	q.SetMinScore(s.MinScore)

	err = q.SetParams(s.Params)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}

	return q, nil
}

func (ScriptScore) Kind() Kind {
	return KindScriptScore
}

type ScriptScoreQuery struct {
	query  *QueryValues
	params dynamic.JSON
	script string
	boostParam
	minScoreParam
	nameParam
}

func (ScriptScoreQuery) Kind() Kind {
	return KindScriptScore
}

func (s ScriptScoreQuery) DecodeParams(val interface{}) error {

	return json.Unmarshal(s.params, val)
}

func (s *ScriptScoreQuery) SetParams(params interface{}) error {
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

func (s *ScriptScoreQuery) setScript(script string) error {
	if len(script) == 0 {
		return ErrScriptRequired
	}
	s.script = script
	return nil
}
func (s *ScriptScoreQuery) setQuery(query *Query) error {
	if query == nil {
		return ErrQueryRequired
	}
	qv, err := newQuery(*query)
	if err != nil {
		return err
	}
	if qv.IsEmpty() {
		return ErrQueryRequired
	}
	s.query = qv
	return nil
}
