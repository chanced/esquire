package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultPrefixLength = 0

// WithPrefixLength is an interface for a type with the PrefixLength and
// SetPrefixLength methods
//
// PrefixLength is the umber of beginning characters left unchanged when fuzzy
// matching. Defaults to 0.
type WithPrefixLength interface {
	PrefixLength() int64
	SetPrefixLength(v int64)
}

// prefixLengthParam is a mixin that adds the prefix_length param
//
// PrefixLength is the number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
type prefixLengthParam struct {
	prefixLengthValue *int64
}

func (pl prefixLengthParam) PrefixLength() int64 {
	if pl.prefixLengthValue == nil {
		return DefaultPrefixLength
	}
	return *pl.prefixLengthValue
}

func (pl *prefixLengthParam) SetPrefixLength(v int64) {
	pl.prefixLengthValue = &v
}
func unmarshalPrefixLengthParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithPrefixLength); ok {
		n := dynamic.NewNumber(data.UnquotedString())
		if v, ok := n.Int(); ok {
			a.SetPrefixLength(v)
			return nil
		}
		return &json.UnmarshalTypeError{Value: data.String()}
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
