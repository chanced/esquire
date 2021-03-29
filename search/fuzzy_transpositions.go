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

func unmarshalFuzzyTranspositionsParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithFuzzyTranspositions); ok {
		if data.IsNull() {
			return nil
		}
		var err error
		var b dynamic.Bool
		if data.IsBool() {
			b, err = dynamic.NewBool(data.String())
			if err != nil {
				return err
			}
			bv, _ := b.Bool()
			a.SetFuzzyTranspositions(bv)
			return nil
		}
		var str string
		err = json.Unmarshal(data, &str)
		if err != nil {
			return err
		}
		b, err = dynamic.NewBool(str)
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			a.SetFuzzyTranspositions(v)
			return nil
		} else {
			return &json.UnmarshalKindError{Value: data.String(), Kind: typeString}
		}

	}
	return nil
}
func marshalFuzzyTranspositionsParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if a, ok := source.(WithFuzzyTranspositions); ok {
		if !a.FuzzyTranspositions() {
			data[paramFuzzyTranspositions] = a.FuzzyTranspositions()
		}
	}
	return data, nil
}
