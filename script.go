package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Script struct {
	Lang   string
	Source string
	Params interface{}
}

func (s Script) Script() (*ScriptClause, error) {
	sq := &ScriptQuery{
		Lang:   s.Lang,
		Source: s.Source,
		Params: s.Params,
	}
	return sq.Script()
}

func (s *Script) script() (*scriptParams, error) {
	if s == nil {
		return nil, nil
	}
	sp := &scriptParams{}
	err := sp.setSource(s.Source)
	if err != nil {
		return sp, err
	}
	err = sp.SetParams(s.Params)
	if err != nil {
		return sp, err
	}
	sp.setLang(s.Lang)
	return sp, nil
}

type scriptParams struct {
	params dynamic.JSON
	source string
	lang   string
}

func (s *scriptParams) DecodeParams(val interface{}) error {
	return json.Unmarshal(s.params, val)
}

func (s *scriptParams) marshalScriptParams() (dynamic.JSON, error) {
	if s == nil || s.IsEmpty() {
		return nil, nil
	}
	source, err := json.Marshal(s.source)

	if err != nil {
		return nil, err
	}
	data := dynamic.JSONObject{"source": source}
	if s.params != nil {
		params, err := json.Marshal(s.params)
		if err != nil {
			return nil, err
		}
		data["params"] = params
	}
	if len(s.lang) > 0 {
		lang, err := json.Marshal(s.lang)
		if err != nil {
			return nil, err
		}
		data["lang"] = lang
	}

	return json.Marshal(data)
}

func (s *scriptParams) unmarshalScriptParams(data []byte) error {
	*s = scriptParams{}
	if len(data) == 0 {
		return nil
	}
	var obj dynamic.JSONObject
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	if d, ok := obj["source"]; ok {
		var source string
		err := json.Unmarshal(d, &source)
		if err != nil {
			return err
		}
		s.source = source
	}
	if d, ok := obj["lang"]; ok {
		var lang string
		err := json.Unmarshal(d, &lang)
		if err != nil {
			return err
		}
		s.lang = lang
	}
	if d, ok := obj["params"]; ok {
		s.params = d
	}
	return nil
}

func (s *scriptParams) SetScript(script *Script) error {
	if script == nil {
		return ErrScriptRequired
	}
	scr, err := script.script()
	if err != nil {
		return err
	}
	if scr == nil {
		return ErrScriptRequired
	}
	*s = *scr
	return nil
}

func (s *scriptParams) setLang(lang string) {
	s.lang = lang
}

func (s *scriptParams) SetParams(params interface{}) error {
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

func (s *scriptParams) setSource(script string) error {
	if len(script) == 0 {
		return ErrScriptRequired
	}
	s.source = script
	return nil
}

func (s *scriptParams) IsEmpty() bool {
	return s == nil || len(s.source) == 0
}
