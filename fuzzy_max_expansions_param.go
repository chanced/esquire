package picker

import (
	"fmt"

	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultFuzzyMaxExpansions = int(50)

// WithFuzzyMaxExpansions is a query with the max_expansions param
//
// Maximum number of variations created. Defaults to 50.
//
// Warning
//
// Avoid using a high value in the max_expansions parameter, especially if the
// prefix_length parameter value is 0. High values in the max_expansions
// parameter can cause poor performance due to the high number of variations
// examined.
type WithFuzzyMaxExpansions interface {
	// FuzzyMaxExpansions is the maximum number of variations created. Defaults to 50.
	FuzzyMaxExpansions() int
	// SetFuzzyMaxExpansions sets the max_expansions param
	//
	// Maximum number of variations created. Defaults to 50.
	//
	// Warning
	//
	// Avoid using a high value in the max_expansions parameter, especially if the
	// prefix_length parameter value is 0. High values in the max_expansions
	// parameter can cause poor performance due to the high number of variations
	// examined.
	SetFuzzyMaxExpansions(v interface{}) error
}

// fuzzyMaxExpansionsParam is a mixin that adds the max_expansions param to queries
//
// Maximum number of variations created. Defaults to 50.
type fuzzyMaxExpansionsParam struct {
	fuzzyMaxExpansions dynamic.Number
}

// FuzzyMaxExpansions is the maximum number of variations created. Defaults to 50.
func (me fuzzyMaxExpansionsParam) FuzzyMaxExpansions() int {
	if me.fuzzyMaxExpansions.HasValue() {
		if i, ok := me.fuzzyMaxExpansions.Int(); ok {
			return i
		}
		if f, ok := me.fuzzyMaxExpansions.Float64(); ok {
			return int(f)
		}
	}
	return DefaultFuzzyMaxExpansions
}
func (me *fuzzyMaxExpansionsParam) SetFuzzyMaxExpansions(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if n.IsNil() {
		_ = me.fuzzyMaxExpansions.Set(nil)
		return nil
	}
	if i, ok := n.Int(); ok {
		_ = me.fuzzyMaxExpansions.Set(i)
		return nil
	}
	return fmt.Errorf("%w <%s>", ErrInvalidFuzzyMaxExpansions, v)
}
func unmarshalFuzzyMaxExpansionsParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithFuzzyMaxExpansions); ok {
		n, err := dynamic.NewNumber(data.UnquotedString())
		if err != nil {
			return err
		}
		if v, ok := n.Int(); ok {
			return a.SetFuzzyMaxExpansions(v)
		}
		return nil
	}
	return nil
}
func marshalFuzzyMaxExpansionsParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithFuzzyMaxExpansions); ok {
		if b.FuzzyMaxExpansions() != DefaultFuzzyMaxExpansions {
			return json.Marshal(b.FuzzyMaxExpansions())
		}
	}
	return nil, nil
}
