package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultTranspositions = true

// WithTranspositions is an interface for queries with the transpositions param
//
// (Optional, Boolean) Indicates whether edits include transpositions of two
// adjacent characters (ab → ba). Defaults to true.
type WithTranspositions interface {
	// Transpositions indicates whether edits include transpositions of two
	// adjacent characters (ab → ba). Defaults to true.
	Transpositions() bool
	// SetTranspositions sets the value of Transpositions to v
	SetTranspositions(v interface{}) error
}

type transpositionsParam struct {
	transpositions dynamic.Bool
}

// Transpositions indicates whether edits include transpositions of two
// adjacent characters (ab → ba). Defaults to true.
func (t transpositionsParam) Transpositions() bool {
	if b, ok := t.transpositions.Bool(); ok {
		return b
	}
	return DefaultTranspositions

}

// SetTranspositions sets the value of Transpositions to v
func (t *transpositionsParam) SetTranspositions(v interface{}) error {
	return t.transpositions.Set(v)
}
func unmarshalTranspositionsParam(value dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithTranspositions); ok {
		b, err := dynamic.NewBool(value)
		if err != nil {
			return err
		}
		a.SetTranspositions(b)
	}
	return nil
}
func marshalTranspositionsParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithTranspositions); ok {
		if !b.Transpositions() {
			return json.Marshal(b.Transpositions())
		}
	}
	return nil, nil
}
