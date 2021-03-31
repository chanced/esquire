package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultSlop = 0

type WithSlop interface {
	Slop() int
	SetSlop(v int)
}

type slopParam struct {
	slop *int
}

func (s slopParam) Slop() int {
	if s.slop == nil {
		return DefaultSlop
	}
	return *s.slop
}

func (s *slopParam) SetSlop(v int) {
	s.slop = &v
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
