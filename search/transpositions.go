package search

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
		return true
	}
	return *t.TranspositionsValue
}

// SetTranspositions sets the value of Transpositions to v
func (t *TranspositionsParam) SetTranspositions(v bool) {
	t.TranspositionsValue = &v
}
