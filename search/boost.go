package search

import (
	"github.com/chanced/dynamic"
)

const DefaultBoost = float64(1)

// WithBoost is an interface with the Boost and SetBoost methods
//
// Boost is a floating point number used to decrease or increase the relevance
// scores of a query. Defaults to 1.0.
//
// You can use the boost parameter to adjust relevance scores for searches
// containing two or more queries.
//
// Boost values are relative to the default value of 1.0. A boost value between
// 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases
// the relevance score.

type WithBoost interface {
	// Boost is a floating point number used to decrease or increase the relevance
	// scores of a query. Defaults to 1.0.
	//
	// You can use the boost parameter to adjust relevance scores for searches
	// containing two or more queries.
	//
	// Boost values are relative to the default value of 1.0. A boost value between
	// 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases
	// the relevance score.
	Boost() float64
	SetBoost(v float64)
}

type boostParam struct {
	// boostValue is a floating point number used to decrease or increase the relevance
	// scores of a query. Defaults to 1.0.
	//
	// You can use the boost parameter to adjust relevance scores for searches
	// containing two or more queries.
	//
	// Boost values are relative to the default value of 1.0. A boost value between
	// 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases
	// the relevance score.
	boostValue *float64
}

// Boost is a floating point number used to decrease or increase the relevance
// scores of a query. Defaults to 1.0.
//
// You can use the boost parameter to adjust relevance scores for searches
// containing two or more queries.
//
// Boost values are relative to the default value of 1.0. A boost value between
// 0 and 1.0 decreases the relevance score. A value greater than 1.0 increases
// the relevance score.
func (b boostParam) Boost() float64 {
	if b.boostValue == nil {
		return DefaultBoost
	}
	return *b.boostValue
}

// SetBoost sets Boost to v
func (b *boostParam) SetBoost(v float64) {
	if b.Boost() != v {
		b.boostValue = &v
	}
}

func marshalBoostParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithBoost); ok {
		if b.Boost() != DefaultBoost {
			data[paramBoost] = b.Boost()
		}
	}
	return data, nil
}
func unmarshalBoostParam(data dynamic.RawJSON, target interface{}) error {
	if r, ok := target.(WithBoost); ok {
		if data.IsNumber() {
			f, _ := dynamic.NewNumber(data.String()).Float()
			r.SetBoost(f)
			return nil
		}
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			if data.UnquotedString() == "" {
				return nil
			}
			str := dynamic.NewString(data.UnquotedString())
			f, err := str.Float64()
			if err != nil {
				return err
			}
			r.SetBoost(f)
			return nil
		}
	}
	return nil
}
