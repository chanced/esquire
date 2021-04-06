package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultSlop = 0

type WithSlop interface {
	Slop() int
	SetSlop(v interface{}) error
}

type slopParam struct {
	slop dynamic.Number
}

func (s slopParam) Slop() int {
	if i, ok := s.slop.Int(); ok {
		return i
	}
	if f, ok := s.slop.Float64(); ok {
		return int(f)
	}

	return DefaultSlop
}

func (s *slopParam) SetSlop(v interface{}) error {
	return s.slop.Set(v)
}

func unmarshalSlopParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithSlop); ok {
		n, err := dynamic.NewNumber(data.UnquotedString())
		if err != nil {
			return err
		}
		if i, ok := n.Int(); ok {
			a.SetSlop(int(i))
		}
	}
	return nil
}
func marshalSlopParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithSlop); ok {
		if b.Slop() != DefaultSlop {
			return json.Marshal(b.Slop())
		}
	}
	return nil, nil
}
