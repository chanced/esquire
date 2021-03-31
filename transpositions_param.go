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
	SetTranspositions(v bool)
}

type transpositionsParam struct {
	transpositions *bool
}

// Transpositions indicates whether edits include transpositions of two
// adjacent characters (ab → ba). Defaults to true.
func (t transpositionsParam) Transpositions() bool {
	if t.transpositions == nil {
		return DefaultTranspositions
	}
	return *t.transpositions
}

// SetTranspositions sets the value of Transpositions to v
func (t *transpositionsParam) SetTranspositions(v bool) {
	t.transpositions = &v
}
func unmarshalTranspositionsParam(value dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithTranspositions); ok {
		b, err := dynamic.NewBool(value)
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			a.SetTranspositions(v)
		}
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
