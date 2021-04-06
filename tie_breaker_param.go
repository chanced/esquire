package picker

import (
	"fmt"

	"github.com/chanced/dynamic"
)

type WithTieBreaker interface {
	TieBreaker() float64
	SetTieBreaker(v interface{}) error
}

type tieBreakerParam struct {
	tieBreaker dynamic.Number
}

func (tb *tieBreakerParam) SetTieBreaker(tieBreaker interface{}) error {
	if tieBreaker == nil {
		return tb.tieBreaker.Set(nil)
	}
	n, err := dynamic.NewNumber(tieBreaker)
	if err != nil {
		return err
	}
	v, ok := n.Float64()
	if !ok {
		return ErrInvalidTieBreaker
	}
	if v > 1 || v < 0 {
		return fmt.Errorf("%w <%f>; valid values are 0.0 (inclusive) through 1.0 (inclusive)", ErrInvalidTieBreaker, tieBreaker)
	}
	tb.tieBreaker = n
	return nil
}
