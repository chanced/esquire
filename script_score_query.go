package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type ScriptScorer interface {
	ScriptScore() (*ScriptScoreQuery, error)
}

// ScriptScoreQueryParams uses a script to provide a custom score for returned documents.
//
// The script_score query is useful if, for example, a scoring function is
// expensive and you only need to calculate the score of a filtered set of
// documents.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html
type ScriptScoreQueryParams struct {
	// Query used to return documents. (Required)
	Query *QueryParams
	// Documents with a score lower than this floating point number are excluded
	// from the search results. (Optional)
	MinScore float64
	// Documents scores produced by script are multiplied by boost to produce
	// final documents' scores. Defaults to 1.0. (Optional)
	Boost  interface{}
	Name   string
	Script *Script
}

func (s ScriptScoreQueryParams) Clause() (QueryClause, error) {
	return s.ScriptScore()
}
func (s ScriptScoreQueryParams) ScriptScore() (*ScriptScoreQuery, error) {
	q := &ScriptScoreQuery{}

	err := q.setQuery(s.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindScriptScore)
	}

	if err != nil {
		return q, newQueryError(err, QueryKindScriptScore)
	}

	err = q.SetBoost(s.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindScriptScore)
	}
	q.SetMinScore(s.MinScore)

	err = q.SetScript(s.Script)
	if err != nil {
		return q, newQueryError(err, QueryKindScriptScore)
	}
	q.SetName(s.Name)

	return q, nil
}

func (ScriptScoreQueryParams) Kind() QueryKind {
	return QueryKindScriptScore
}

type ScriptScoreQuery struct {
	query *Query
	scriptParams
	boostParam
	minScoreParam
	nameParam
	completeClause
}

func (ScriptScoreQuery) Kind() QueryKind {
	return QueryKindScriptScore
}

// Set sets the ScriptScoreQuery
// Options include:
//   - picker.ScriptScoreQuery
func (s *ScriptScoreQuery) Set(scriptScore *ScriptScoreQueryParams) error {
	if scriptScore == nil {
		*s = ScriptScoreQuery{}
		return nil
	}
	scr, err := scriptScore.ScriptScore()
	if err != nil {
		return newQueryError(err, QueryKindScriptScore)
	}
	*s = *scr
	return nil
}
func (s *ScriptScoreQuery) Clause() (QueryClause, error) {
	return s, nil
}
func (s ScriptScoreQuery) MarshalBSON() ([]byte, error) {
	return s.MarshalJSON()
}

func (s ScriptScoreQuery) MarshalJSON() ([]byte, error) {
	if s.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalClauseParams(&s)
	if err != nil {
		return nil, err
	}
	script, err := s.scriptParams.marshalScriptParams()
	if err != nil {
		return nil, err
	}
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

func (s *ScriptScoreQuery) UnmarshalBSON(data []byte) error {
	return s.UnmarshalJSON(data)
}

func (s *ScriptScoreQuery) UnmarshalJSON(data []byte) error {
	s = &ScriptScoreQuery{}
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

func (s *ScriptScoreQuery) IsEmpty() bool {
	return s == nil || s.scriptParams.IsEmpty()
}

func (s *ScriptScoreQuery) setQuery(query *QueryParams) error {
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

func (s *ScriptScoreQuery) Clear() {
	*s = ScriptScoreQuery{}
}
