package search

import (
	"math"

	"github.com/tidwall/gjson"
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
func unmarshalMaxBoostParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithMaxBoost); ok {
		a.SetMaxBoost(value.Float())
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
