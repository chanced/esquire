package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type ScriptScorer interface {
	ScriptScore() (*ScriptScoreClause, error)
}

// ScriptScoreQuery uses a script to provide a custom score for returned documents.
//
// The script_score query is useful if, for example, a scoring function is
// expensive and you only need to calculate the score of a filtered set of
// documents.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html
type ScriptScoreQuery struct {
	// Query used to return documents. (Required)
	Query *Query
	// Documents with a score lower than this floating point number are excluded
	// from the search results. (Optional)
	MinScore float64
	// Documents scores produced by script are multiplied by boost to produce
	// final documents' scores. Defaults to 1.0. (Optional)
	Boost  interface{}
	Name   string
	Script *Script
}

func (s ScriptScoreQuery) Clause() (Clause, error) {
	return s.ScriptScore()
}
func (s ScriptScoreQuery) ScriptScore() (*ScriptScoreClause, error) {
	q := &ScriptScoreClause{}

	err := q.setQuery(s.Query)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}

	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}

	err = q.SetBoost(s.Boost)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}
	q.SetMinScore(s.MinScore)

	err = q.SetScript(s.Script)
	if err != nil {
		return q, NewQueryError(err, KindScriptScore)
	}
	q.SetName(s.Name)

	return q, nil
}

func (ScriptScoreQuery) Kind() Kind {
	return KindScriptScore
}

type ScriptScoreClause struct {
	query *QueryValues
	scriptParams
	boostParam
	minScoreParam
	nameParam
}

func (ScriptScoreClause) Kind() Kind {
	return KindScriptScore
}

// Set sets the ScriptScoreQuery
// Options include:
//   - picker.ScriptScoreQuery
func (s *ScriptScoreClause) Set(scriptScore *ScriptScoreQuery) error {
	if scriptScore == nil {
		*s = ScriptScoreClause{}
		return nil
	}
	scr, err := scriptScore.ScriptScore()
	if err != nil {
		return NewQueryError(err, KindScriptScore)
	}
	*s = *scr
	return nil
}

func (s ScriptScoreClause) MarshalJSON() ([]byte, error) {
	if s.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalClauseParams(&s)
	if err != nil {
		return nil, err
	}
	script := s.scriptParams.marshalScriptParams()
	if script != nil {
		data["script"] = script
	}
	var query dynamic.JSON
	query, err = s.query.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if query != nil && !query.IsNull() {
		data["query"] = query
	}
	return json.Marshal(data)
}

func (s *ScriptScoreClause) UnmarshalJSON(data []byte) error {
	s = &ScriptScoreClause{}
	params, err := unmarshalClauseParams(data, s)
	if err != nil {
		return err
	}
	err = s.scriptParams.unmarshalScriptParams(params["script"])
	if err != nil {
		return err
	}
	err = s.query.UnmarshalJSON(params["query"])
	if err != nil {
		return err
	}
	return nil
}

func (s *ScriptScoreClause) IsEmpty() bool {
	return s == nil || s.scriptParams.IsEmpty()
}

func (s *ScriptScoreClause) setQuery(query *Query) error {
	if query == nil {
		return ErrQueryRequired
	}
	qv, err := query.Query()
	if err != nil {
		return err
	}
	if qv.IsEmpty() {
		return ErrQueryRequired
	}
	s.query = qv
	return nil
}

func (s *ScriptScoreClause) Clear() {
	*s = ScriptScoreClause{}
}
