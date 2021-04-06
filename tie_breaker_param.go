package picker

import (
	"fmt"
)

type WithTieBreaker interface {
	TieBreaker() float64
	SetTieBreaker(float64) error
}

type tieBreakerParam struct {
	tieBreaker float64
}

func (tb *tieBreakerParam) SetTieBreaker(tieBreaker float64) error {
	if tieBreaker > 1 || tieBreaker < 0 {
		return fmt.Errorf("picker: invalid tie breaker value %f; valid values are 0.0 (inclusive) through 1.0 (inclusive)", tieBreaker)
	}
	tb.tieBreaker = tieBreaker
	return nil
}
