package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Scripter interface {
	Script() (*ScriptClause, error)
}

// ScriptQuery filters documents based on a provided script. The script query is typically used in a filter context.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-query.html
type ScriptQuery struct {
	// Name of the query (Optional)
	Name string
	// Script to run as a query. This script must return a boolean value, true or false. (Required)
	Lang   string
	Source string
	Params interface{}
}

func (s ScriptQuery) Clause() (QueryClause, error) {
	return s.Script()
}
func (s ScriptQuery) Script() (*ScriptClause, error) {
	q := &ScriptClause{}
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

func (ScriptQuery) Kind() QueryKind {
	return KindScript
}

type ScriptClause struct {
	scriptParams
	nameParam
	completeClause
}

func (ScriptClause) Kind() QueryKind {
	return KindScript
}

// Set sets the ScriptQuery
// Options include:
//   - picker.ScriptQuery
func (s *ScriptClause) Set(script Scripter) error {

	if script == nil {
		*s = ScriptClause{}
		return nil
	}
	scr, err := script.Script()
	if err != nil {
		return NewQueryError(err, KindScript)
	}
	*s = *scr
	return nil
}
func (s *ScriptClause) Clause() (QueryClause, error) {
	return s, nil
}
func (s ScriptClause) MarshalJSON() ([]byte, error) {
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

func (s *ScriptClause) UnmarshalJSON(data []byte) error {
	s = &ScriptClause{}
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

func (s *ScriptClause) IsEmpty() bool {
	return s == nil || s.scriptParams.IsEmpty()
}

func (s *ScriptClause) Clear() {
	*s = ScriptClause{}
}
