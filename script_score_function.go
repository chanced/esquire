package picker

import (
	"encoding/json"

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
	data := ss.scriptParams.marshalScriptParams()
	if ss.weight != nil {
		data["weight"] = ss.Weight()
	}
	if ss.filter != nil {
		filter, err := ss.filter.MarshalJSON()
		if err != nil {
			return nil, err
		}
		data["filter"] = dynamic.JSON(filter)
	}
	return json.Marshal(data)
}

func (ss *ScriptScoreFunction) UnmarshalJSON(data []byte) error {
	*ss = ScriptScoreFunction{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	err = ss.scriptParams.unmarshalScriptParams(obj["script"])
	if err != nil {
		return err
	}
	if wd, ok := obj["weight"]; ok {
		err = unmarshalWeightParam(wd, ss)
		if err != nil {
			return err
		}
	}
	if fd, ok := obj["filter"]; ok {
		filter, err := unmarshalQueryClause(fd)
		if err != nil {
			return err
		}
		ss.filter = filter
	}
	return nil
}
