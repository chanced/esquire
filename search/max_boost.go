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
	maxBoost *float64
}

func (mb maxBoostParam) MaxBoost() float64 {
	if mb.maxBoost == nil {
		return DefaultMaxBoost
	}
	return *mb.maxBoost
}

func (mb maxBoostParam) SetMaxBoost(v float64) {
	if mb.MaxBoost() != v {
		mb.maxBoost = &v
	}
}
func unmarshalMaxBoostParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithMaxBoost); ok {
		n := dynamic.NewNumber(data.UnquotedString())
		if v, ok := n.Float(); ok {
			a.SetMaxBoost(v)
			return nil
		}
		return &json.UnmarshalTypeError{Value: data.String(), Type: typeFloat64}
	}
	return nil
}
func marshalMaxBoostParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithMaxBoost); ok {
		if b.MaxBoost() != DefaultMaxBoost {
			data[paramMaxBoost] = b.MaxBoost()
		}
	}
	return data, nil
}
