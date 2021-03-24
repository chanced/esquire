package search

import "github.com/tidwall/gjson"

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

type FuzzyTranspositionsParam struct {
	FuzzyTranspositionsValue *bool `json:"fuzzy_transpositions,omitempty" bson:"fuzzy_transpositions,omitempty"`
}

// FuzzyTranspositions edits for fuzzy matching include transpositions of
// two adjacent characters (ab → ba). Defaults to true
func (ft FuzzyTranspositionsParam) FuzzyTranspositions() bool {
	if ft.FuzzyTranspositionsValue == nil {
		return true
	}
	return *ft.FuzzyTranspositionsValue
}

// SetFuzzyTranspositions sets FuzzyTranspositions to v
func (ft *FuzzyTranspositionsParam) SetFuzzyTranspositions(v bool) {
	if ft.FuzzyTranspositions() != v {
		ft.FuzzyTranspositionsValue = &v
	}
}

func unmarshalFuzzyTranspositionsParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithFuzzyTranspositions); ok {
		a.SetFuzzyTranspositions(value.Bool())
	}
	return nil
}
