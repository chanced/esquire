package picker

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/chanced/dynamic"
)

const DefaultBoost = float64(1)

// WithBoost is an interface for queries with the boost parameter
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
	SetBoost(v interface{}) error
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
func (b *boostParam) SetBoost(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if f, ok := n.Float(); ok {
		b.boostValue = &f
	} else if n.IsNil() {
		b.boostValue = nil
	} else {
		return fmt.Errorf("%w <%s>", ErrInvalidBoost, v)
	}
	return nil
}

func marshalBoostParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithBoost); ok {
		if b.Boost() != DefaultBoost {
			return json.Marshal(b.Boost())
		}
	}
	return []byte{}, nil
}
func unmarshalBoostParam(data dynamic.JSON, target interface{}) error {
	if r, ok := target.(WithBoost); ok {
		if data.IsNumber() {
			n, err := dynamic.NewNumber(string(data))
			if err != nil {
				return err
			}
			f, ok := n.Float()
			if !ok {
				return &json.UnmarshalTypeError{
					Value: string(data),
					Type:  reflect.TypeOf(float64(0)),
				}
			}
			r.SetBoost(f)
			return nil
		}
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			if len(data.UnquotedString()) == 0 {
				return nil
			}
			var str string
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}
			n, err := dynamic.NewNumber(str)
			if err != nil {
				return err
			}
			f, ok := n.Float()
			if ok {
				r.SetBoost(f)
			}
			return nil
		}
	}
	return nil
}
