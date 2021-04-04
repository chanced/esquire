package picker

import "fmt"

type WithNegativeBoost interface {
	// Floating point number between 0 and 1.0 used to
	// decrease the relevance scores of documents matching the negative query. (Required, float)
	NegativeBoost() float64
	// Sets negative_boost param
	SetNegativeBoost(negativeBoost interface{}) error
}

type negativeBoostParam struct {
	negativeBoost float64
}

// NegativeBoost is a floating point number between 0 and 1.0 used to decrease
// the relevance scores of documents matching the negative query. (Required)
func (nb negativeBoostParam) NegativeBoost() float64 {
	return nb.negativeBoost
}

// SetNegativeBoost sets negative_boost
func (nb *negativeBoostParam) SetNegativeBoost(negativeBoost float64) error {
	if 0 >= negativeBoost || negativeBoost > 1 {
		return fmt.Errorf("%w; received %f", ErrInvalidNegativeBoost, negativeBoost)
	}
	nb.negativeBoost = negativeBoost
	return nil
}
