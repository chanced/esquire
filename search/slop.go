package search

import "github.com/tidwall/gjson"

type WithSlop interface {
	Slop() int
	SetSlop(v int)
}

type SlopParam struct {
	SlopValue *int `json:"slop,omitempty" bson:"slop,omitempty"`
}

func (s SlopParam) Slop() int {
	if s.SlopValue == nil {
		return 0
	}
	return *s.SlopValue
}

func (s *SlopParam) SetSlop(v int) {
	s.SlopValue = &v
}

func unmarshalSlopParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithSlop); ok {
		a.SetSlop(int(value.Int()))
	}
	return nil
}
