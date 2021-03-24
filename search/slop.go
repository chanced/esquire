package search

import "github.com/tidwall/gjson"

const DefaultSlop = 0

type WithSlop interface {
	Slop() int
	SetSlop(v int)
}

type SlopParam struct {
	SlopValue *int `json:"slop,omitempty" bson:"slop,omitempty"`
}

func (s SlopParam) Slop() int {
	if s.SlopValue == nil {
		return DefaultSlop
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
func marshalSlopParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithSlop); ok {
		if b.Slop() != DefaultSlop {
			data[paramSlop] = b.Slop()
		}
	}
	return data, nil
}
