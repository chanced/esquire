package search

import "github.com/tidwall/gjson"

const DefaultPrefixLength = 0

// WithPrefixLength is an interface for a type with the PrefixLength and
// SetPrefixLength methods
//
// PrefixLength is the umber of beginning characters left unchanged when fuzzy
// matching. Defaults to 0.
type WithPrefixLength interface {
	PrefixLength() int
	SetPrefixLength(v int)
}

// PrefixLengthParam is a mixin that adds the prefix_length param
//
// PrefixLength is the number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
type PrefixLengthParam struct {
	PrefixLengthValue *int `json:"prefix_length,omitempty" bson:"prefix_length,omitempty"`
}

func (pl PrefixLengthParam) PrefixLength() int {
	if pl.PrefixLengthValue == nil {
		return DefaultPrefixLength
	}
	return *pl.PrefixLengthValue
}

func (pl *PrefixLengthParam) SetPrefixLength(v int) {
	pl.PrefixLengthValue = &v
}
func unmarshalPrefixLengthParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithPrefixLength); ok {
		a.SetPrefixLength(int(value.Int()))
	}
	return nil
}

func marshalPrefixLengthParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithPrefixLength); ok {
		if b.PrefixLength() != DefaultPrefixLength {
			data[paramPrefixLength] = b.PrefixLength()
		}
	}
	return data, nil
}
