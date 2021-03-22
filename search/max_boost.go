package search

import "math"

type WithMaxBoost interface {
	MaxBoost() float32
	SetMaxBoost(v float32)
}

type MaxBoostParam struct {
	MaxBoostValue *float32 `json:"max_boost,omitempty" bson:"max_boost,omitempty"`
}

func (mb MaxBoostParam) MaxBoost() float32 {
	if mb.MaxBoostValue == nil {
		return math.MaxFloat32
	}
	return *mb.MaxBoostValue
}

func (mb MaxBoostParam) SetMaxBoost(v float32) {
	if mb.MaxBoost() != v {
		mb.MaxBoostValue = &v
	}
}
