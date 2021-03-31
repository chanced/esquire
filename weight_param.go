package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

var DefaultWeight = float64(0)

type WithWeight interface {
	Weight() float64
	SetWeight(v interface{}) error
}

type weightParam struct {
	weight *float64
}

func (b *weightParam) Weight() float64 {
	if b == nil {
		return 0
	}
	if b.weight == nil {
		return 0
	}
	return *b.weight
}

// SetWeight sets Weight to v
func (b *weightParam) SetWeight(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if f, ok := n.Float(); ok {
		b.weight = &f
	} else if n.IsNil() {
		b.weight = nil
	}
	return nil
}

func marshalWeightParam(source interface{}) (dynamic.JSON, error) {

	if b, ok := source.(WithWeight); ok {
		if b.Weight() != DefaultWeight {
			return json.Marshal(b.Weight())
		}
	}
	return nil, nil
}
func unmarshalWeightParam(data dynamic.JSON, target interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if r, ok := target.(WithWeight); ok {
		var n dynamic.Number
		err := json.Unmarshal(data, &n)
		if err != nil {
			return err
		}
		r.SetWeight(n.Value())
	}
	return nil
}
