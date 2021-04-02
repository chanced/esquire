package picker

import (
	"encoding/json"
	"math"

	"github.com/chanced/dynamic"
)

const DefaultMaxBoost = float64(math.MaxFloat32)

type WithMaxBoost interface {
	MaxBoost() float64
	SetMaxBoost(v float64)
}

type maxBoostParam struct {
	maxBoost *float64
}

func (mb maxBoostParam) MaxBoost() float64 {
	if mb.maxBoost == nil {
		return DefaultMaxBoost
	}
	return *mb.maxBoost
}

func (mb maxBoostParam) SetMaxBoost(v float64) {
	if mb.MaxBoost() != v && v != 0 {
		mb.maxBoost = &v
	}
}
func unmarshalMaxBoostParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithMaxBoost); ok {
		var n dynamic.Number
		err := json.Unmarshal(data, &n)
		if err != nil {
			return err
		}
		if v, ok := n.Float64(); ok {
			a.SetMaxBoost(v)
			return nil
		}
		return nil
	}
	return nil
}
func marshalMaxBoostParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithMaxBoost); ok {
		if b.MaxBoost() != DefaultMaxBoost {
			return json.Marshal(b.MaxBoost())
		}
	}
	return nil, nil
}
