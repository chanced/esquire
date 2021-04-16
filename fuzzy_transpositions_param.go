package picker

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
	SetFuzzyTranspositions(v interface{}) error
}

type fuzzyTranspositionsParam struct {
	fuzzyTranspositions dynamic.Bool
}

// FuzzyTranspositions edits for fuzzy matching include transpositions of
// two adjacent characters (ab → ba). Defaults to true
func (ft fuzzyTranspositionsParam) FuzzyTranspositions() bool {
	if b, ok := ft.fuzzyTranspositions.Bool(); ok {
		return b
	}
	return DefaultFuzzyTranspositions
}

// SetFuzzyTranspositions sets FuzzyTranspositions to v
func (ft *fuzzyTranspositionsParam) SetFuzzyTranspositions(fuzzyTranspositions interface{}) error {
	return ft.fuzzyTranspositions.Set(fuzzyTranspositions)
}

func unmarshalFuzzyTranspositionsParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithFuzzyTranspositions); ok {
		if data.IsNull() {
			return nil
		}
		var err error
		var b dynamic.Bool
		if data.IsBool() {
			b, err = dynamic.NewBool(string(data))
			if err != nil {
				return err
			}
			bv, _ := b.Bool()
			return a.SetFuzzyTranspositions(bv)
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
			return a.SetFuzzyTranspositions(v)
		}

		return &json.UnmarshalTypeError{Value: string(data), Type: typeString}

	}
	return nil
}
func marshalFuzzyTranspositionsParam(source interface{}) (dynamic.JSON, error) {
	if a, ok := source.(WithFuzzyTranspositions); ok {
		if !a.FuzzyTranspositions() {
			return json.Marshal(a.FuzzyTranspositions())
		}
	}
	return nil, nil
}
