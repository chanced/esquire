package search

import "github.com/tidwall/gjson"

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

type TranspositionsParam struct {
	TranspositionsValue *bool `json:"transpositions,omitempty" bson:"transpositions,omitempty"`
}

// Transpositions indicates whether edits include transpositions of two
// adjacent characters (ab → ba). Defaults to true.
func (t TranspositionsParam) Transpositions() bool {
	if t.TranspositionsValue == nil {
		return DefaultTranspositions
	}
	return *t.TranspositionsValue
}

// SetTranspositions sets the value of Transpositions to v
func (t *TranspositionsParam) SetTranspositions(v bool) {
	t.TranspositionsValue = &v
}
func unmarshalTranspositionsParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithTranspositions); ok {
		a.SetTranspositions(value.Bool())
	}
	return nil
}
func marshalTranspositionsParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithTranspositions); ok {
		if b.Transpositions() != DefaultTranspositions {
			data[paramTranspositions] = b.Transpositions()
		}
	}
	return data, nil
}
