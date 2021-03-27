package search

import (
	"github.com/chanced/dynamic"
)

const DefaultSlop = 0

type WithSlop interface {
	Slop() int64
	SetSlop(v int64)
}

type slopParam struct {
	slop *int64
}

func (s slopParam) Slop() int64 {
	if s.slop == nil {
		return DefaultSlop
	}
	return *s.slop
}

func (s *slopParam) SetSlop(v int64) {
	s.slop = &v
}

func unmarshalSlopParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithSlop); ok {
		n := dynamic.NewNumber(data.UnquotedString())
		if i, ok := n.Int(); ok {
			a.SetSlop(i)
		}
	}
	return nil
}
func marshalSlopParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithSlop); ok {
		if b.Slop() != DefaultSlop {
			data[paramSlop] = b.Slop()
		}
	}
	return data, nil
}
