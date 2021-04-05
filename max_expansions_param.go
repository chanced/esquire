package picker

import (
	"fmt"

	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultMaxExpansions = int(50)

// WithMaxExpansions is a query with the max_expansions param
//
// Maximum number of variations created. Defaults to 50.
//
// Warning
//
// Avoid using a high value in the max_expansions parameter, especially if the
// prefix_length parameter value is 0. High values in the max_expansions
// parameter can cause poor performance due to the high number of variations
// examined.
type WithMaxExpansions interface {
	// MaxExpansions is the maximum number of variations created. Defaults to 50.
	MaxExpansions() int
	// SetMaxExpansions sets the max_expansions param
	//
	// Maximum number of variations created. Defaults to 50.
	//
	// Warning
	//
	// Avoid using a high value in the max_expansions parameter, especially if the
	// prefix_length parameter value is 0. High values in the max_expansions
	// parameter can cause poor performance due to the high number of variations
	// examined.
	SetMaxExpansions(v interface{}) error
}

// maxExpansionsParam is a mixin that adds the max_expansions param to queries
//
// Maximum number of variations created. Defaults to 50.
type maxExpansionsParam struct {
	maxExpansions dynamic.Number
}

// MaxExpansions is the maximum number of variations created. Defaults to 50.
func (me maxExpansionsParam) MaxExpansions() int {
	if me.maxExpansions.HasValue() {
		if i, ok := me.maxExpansions.Int(); ok {
			return i
		}
		if f, ok := me.maxExpansions.Float64(); ok {
			return int(f)
		}
	}
	return DefaultMaxExpansions
}
func (me *maxExpansionsParam) SetMaxExpansions(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if n.IsNil() {
		_ = me.maxExpansions.Set(nil)
		return nil
	}
	if i, ok := n.Int(); ok {
		_ = me.maxExpansions.Set(i)
		return nil
	}
	return fmt.Errorf("%w <%s>", ErrInvalidMaxExpansions, v)
}
func unmarshalMaxExpansionsParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithMaxExpansions); ok {
		n, err := dynamic.NewNumber(data.UnquotedString())
		if err != nil {
			return err
		}
		if v, ok := n.Int(); ok {
			return a.SetMaxExpansions(v)
		}
		return nil
	}
	return nil
}
func marshalMaxExpansionsParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithMaxExpansions); ok {
		if b.MaxExpansions() != DefaultMaxExpansions {
			return json.Marshal(b.MaxExpansions())
		}
	}
	return nil, nil
}
