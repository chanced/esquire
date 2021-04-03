package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Scripter interface {
	Script() (*ScriptQuery, error)
}

// ScriptQueryParams filters documents based on a provided script. The script query is typically used in a filter context.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-query.html
type ScriptQueryParams struct {
	// Name of the query (Optional)
	Name string
	// Script to run as a query. This script must return a boolean value, true or false. (Required)
	Lang   string
	Source string
	Params interface{}
}

func (s ScriptQueryParams) Clause() (QueryClause, error) {
	return s.Script()
}
func (s ScriptQueryParams) Script() (*ScriptQuery, error) {
	q := &ScriptQuery{}
	q.SetName(s.Name)
	err := q.SetScript(&Script{
		Lang:   s.Lang,
		Source: s.Source,
		Params: s.Params,
	})
	if err != nil {
		return q, err
	}

	return q, nil
}

func (ScriptQueryParams) Kind() QueryKind {
	return KindScript
}

type ScriptQuery struct {
	scriptParams
	nameParam
	completeClause
}

func (ScriptQuery) Kind() QueryKind {
	return KindScript
}

// Set sets the ScriptQuery
// Options include:
//   - picker.ScriptQuery
func (s *ScriptQuery) Set(script Scripter) error {

	if script == nil {
		*s = ScriptQuery{}
		return nil
	}
	scr, err := script.Script()
	if err != nil {
		return NewQueryError(err, KindScript)
	}
	*s = *scr
	return nil
}
func (s *ScriptQuery) Clause() (QueryClause, error) {
	return s, nil
}
func (s ScriptQuery) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(data)
}

func (s *ScriptQuery) UnmarshalJSON(data []byte) error {
	s = &ScriptQuery{}
	params, err := unmarshalClauseParams(data, s)
	if err != nil {
		return err
	}
	err = s.scriptParams.unmarshalScriptParams(params["script"])
	if err != nil {
		return err
	}
	return nil
}

func (s *ScriptQuery) IsEmpty() bool {
	return s == nil || s.scriptParams.IsEmpty()
}

func (s *ScriptQuery) Clear() {
	*s = ScriptQuery{}
}
