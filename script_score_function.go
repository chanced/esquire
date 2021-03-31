package picker

import (
	"github.com/chanced/dynamic"
)

type ScriptScoreFunc struct {
	Filter CompleteClauser
	// float
	Weight interface{}
	Script *Script
}

func (ss *ScriptScoreFunc) Function() (Function, error) {
	if ss == nil {
		return nil, nil
	}
	return ss.ScriptScore()
}
func (ss *ScriptScoreFunc) ScriptScore() (*ScriptScoreFunction, error) {
	if ss == nil {
		return nil, nil
	}
	f := &ScriptScoreFunction{}
	err := f.SetWeight(ss.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetScript(ss.Script)
	if err != nil {
		return f, err
	}
	err = f.SetFilter(ss.Filter)
	if err != nil {
		return f, err
	}
	return f, nil
}

type ScriptScoreFunction struct {
	filter QueryClause
	weightParam
	scriptParams
}

func (ss *ScriptScoreFunction) SetFilter(filter CompleteClauser) error {
	if filter == nil {
		ss.filter = nil
		return nil
	}
	c, err := filter.Clause()
	if err != nil {
		return err
	}
	ss.filter = c
	return nil
}

func (ss *ScriptScoreFunction) Filter() QueryClause {
	return ss.filter
}

func (ScriptScoreFunction) FuncKind() FuncKind {
	return FuncKindScriptScore
}

func (ss ScriptScoreFunction) MarshalJSON() ([]byte, error) {
	return marshalFunction(&ss)

}
func (ss *ScriptScoreFunction) unmarshalParams(data []byte) error {
	return ss.scriptParams.unmarshalScriptParams(data)
}

func (ss *ScriptScoreFunction) marshalParams(data dynamic.JSONObject) error {
	script, err := ss.scriptParams.marshalScriptParams()
	if err != nil {
		return err
	}
	data["script"] = script
	return nil
}
