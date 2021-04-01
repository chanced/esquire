package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type WeightFunc struct {
	// Weight
	Weight float64
	Filter CompleteClauser
}

func (WeightFunc) FuncKind() FuncKind {
	return FuncKindWeight
}
func (w WeightFunc) Function() (Function, error) {
	f := &WeightFunction{}
	if w.Weight == 0 {
		return f, ErrWeightRequired
	}
	err := f.SetWeight(w.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetFilter(w.Filter)
	return f, err
}

type WeightFunction struct {
	weightParam
	filter QueryClause
}

func (WeightFunction) marshalParams(dynamic.JSONObject) error {
	return nil
}
func (WeightFunction) unmarshalParams(data dynamic.JSON) error {
	return nil
}

func (WeightFunction) FuncKind() FuncKind {
	return FuncKindWeight
}
func (w WeightFunction) Filter() QueryClause {
	return w.filter
}
func (w *WeightFunction) SetFilter(c CompleteClauser) error {
	if c == nil {
		w.filter = nil
		return nil
	}
	qc, err := c.Clause()
	if err != nil {
		return err
	}
	w.filter = qc
	return nil
}

func (w WeightFunction) MarshalJSON() ([]byte, error) {
	data := dynamic.JSONObject{}
	if w.filter != nil {
		filter, err := w.filter.MarshalJSON()
		if err != nil {
			return nil, err
		}
		data["filter"] = filter
	}
	if w.weight != nil {
		weight, err := json.Marshal(*w.weight)
		if err != nil {
			return nil, err
		}
		data["weight"] = weight
	}
	return json.Marshal(data)
}
func (w *WeightFunction) UnmarshalJSON(data []byte) error {
	*w = WeightFunction{}
	var obj dynamic.JSONObject
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	filter, err := unmarshalQueryClause(obj["filter"])
	if err != nil {
		return err
	}
	w.filter = filter
	var weight *float64
	err = json.Unmarshal(obj["weight"], &weight)
	if err != nil {
		return err
	}
	w.weight = weight
	return nil
}
