package search

import (
	"encoding/json"
	"math"

	"github.com/chanced/dynamic"
)

const DefaultMaxBoost = math.MaxFloat32

type WithMaxBoost interface {
	MaxBoost() float64
	SetMaxBoost(v float64)
}

type maxBoostParam struct {
	maxBoostValue *float64
}

func (mb maxBoostParam) MaxBoost() float64 {
	if mb.maxBoostValue == nil {
		return DefaultMaxBoost
	}
	return *mb.maxBoostValue
}

func (mb maxBoostParam) SetMaxBoost(v float64) {
	if mb.MaxBoost() != v {
		mb.maxBoostValue = &v
	}
}
func unmarshalMaxBoostParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithMaxBoost); ok {
		n := dynamic.NewNumber()
		if v, ok := n.Float(); ok {
			a.SetMaxBoost(v)
			return nil
		}
		return &json.UnmarshalTypeError{Value: data.String()}
	}
	return nil
}
func marshalMaxBoostParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithMaxBoost); ok {
		if b.MaxBoost() != DefaultMaxBoost {
			data[paramMaxBoost] = b.MaxBoost()
		}
	}
	return data, nil
}
