package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultFuzzyTranspositions = true

// WithFuzzyTranspositions is a query with the fuzzy_transpositions param
//
// If true, edits for fuzzy matching include transpositions of two adjacent
// characters (ab → ba). Defaults to true
type WithFuzzyTranspositions interface {
	// FuzzyTranspositions edits for fuzzy matching include transpositions of
	// two adjacent characters (ab → ba). Defaults to true
	FuzzyTranspositions() bool
	// SetFuzzyTranspositions sets FuzzyTranspositions to v
	SetFuzzyTranspositions(v bool)
}

type fuzzyTranspositionsParam struct {
	fuzzyTranspositionsValue *bool
}

// FuzzyTranspositions edits for fuzzy matching include transpositions of
// two adjacent characters (ab → ba). Defaults to true
func (ft fuzzyTranspositionsParam) FuzzyTranspositions() bool {
	if ft.fuzzyTranspositionsValue == nil {
		return DefaultFuzzyTranspositions
	}
	return *ft.fuzzyTranspositionsValue
}

// SetFuzzyTranspositions sets FuzzyTranspositions to v
func (ft *fuzzyTranspositionsParam) SetFuzzyTranspositions(v bool) {
	if ft.FuzzyTranspositions() != v {
		ft.fuzzyTranspositionsValue = &v
	}
}

func unmarshalFuzzyTranspositionsParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithFuzzyTranspositions); ok {
		b := dynamic.NewBool(data.UnquotedString())
		if v, ok := b.Bool(); ok {
			a.SetFuzzyTranspositions(v)
			return nil
		}
		if !ok {
			return &json.UnmarshalTypeError{Value: data.String()}
		}
	}
	return nil
}
func marshalFuzzyTranspositionsParam(data M, source interface{}) (M, error) {
	if a, ok := source.(WithFuzzyTranspositions); ok {
		if a.FuzzyTranspositions() != DefaultFuzzyTranspositions {
			data[paramFuzzyTranspositions] = a.FuzzyTranspositions()
		}
	}
	return data, nil
}
