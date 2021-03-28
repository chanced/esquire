package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Script filters documents based on a provided script. The script query is
// typically used in a filter context.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-query.html
type Script struct {
	Source   string                 `json:"source" bson:"source"`
	Language string                 `json:"lang,omitempty" bson:"lang,omitempty"`
	Params   map[string]interface{} `json:"params,omitempty" bson:"params,omitempty"`
}
type script Script

func (s Script) MarshalJSON() ([]byte, error) {
	if (s.Language == "" || s.Language == "painless") && len(s.Params) == 0 {
		return json.Marshal(s.Source)
	}
	return json.Marshal(script(s))
}
func (s *Script) UnmarshalJSON(data []byte) error {
	s.Language = ""
	s.Params = map[string]interface{}{}
	s.Source = ""
	d := dynamic.JSON(data)
	if d.IsString() {
		var str string
		err := json.Unmarshal(d, &str)
		if err != nil {
			return err
		}
		s.Source = d.UnquotedString()
		return nil
	}
	m := map[string]dynamic.JSON{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	if src, ok := m["source"]; ok {
		var str string
		err := json.Unmarshal(src, &str)
		if err != nil {
			return err
		}
		s.Source = str
	}
	if lang, ok := m["lang"]; ok {
		var str string
		err := json.Unmarshal(lang, &str)
		if err != nil {
			return err
		}
		s.Language = str
	}
	if params, ok := m["params"]; ok {
		var p map[string]interface{}
		err := json.Unmarshal(params, &p)
		if err != nil {
			return err
		}
		s.Params = p
	}

	return nil
}

// ScriptQuery filters documents based on a provided script. The script query is
// typically used in a filter context.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-query.html
type ScriptQuery struct {
	ScriptValue *Script `json:"script,omitempty" bson:"script,omitempty"`
}

func (s ScriptQuery) ScriptLanguage() string {
	if s.ScriptValue == nil {
		return ""
	}
	return s.ScriptValue.Language
}

func (s ScriptQuery) ScriptParams() map[string]interface{} {
	if s.ScriptValue == nil {
		return nil
	}
	return s.ScriptValue.Params
}
func (s ScriptQuery) ScriptSource() string {
	if s.ScriptValue == nil {
		return ""
	}
	return s.ScriptValue.Source
}

func (s *ScriptQuery) SetScriptLanguage(v string) {
	s.ScriptValue.Language = v
}
func (s *ScriptQuery) SetScriptSource(v string) {
	s.ScriptValue.Source = v
}
func (s *ScriptQuery) SetScriptParams(v map[string]interface{}) {
	s.ScriptValue.Params = v
}
func (s *ScriptQuery) SetScriptParam(key string, value interface{}) {
	if s.ScriptValue.Params == nil {
		s.ScriptValue.Params = map[string]interface{}{}
	}
	s.ScriptValue.Params[key] = value
}
func (s *ScriptQuery) RemoveScriptParam(key string) {
	delete(s.ScriptValue.Params, key)
}
