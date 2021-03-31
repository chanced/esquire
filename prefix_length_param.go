package picker

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

const DefaultPrefixLength = 0

// WithPrefixLength is an interface for a type with the PrefixLength and
// SetPrefixLength methods
//
// PrefixLength is the umber of beginning characters left unchanged when fuzzy
// matching. Defaults to 0.
type WithPrefixLength interface {
	PrefixLength() int
	SetPrefixLength(v interface{}) error
}

// prefixLengthParam is a mixin that adds the prefix_length param
//
// PrefixLength is the number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
type prefixLengthParam struct {
	prefixLength *int
}

func (pl prefixLengthParam) PrefixLength() int {
	if pl.prefixLength == nil {
		return DefaultPrefixLength
	}
	return *pl.prefixLength
}

func (pl *prefixLengthParam) SetPrefixLength(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if n.IsNil() {
		pl.prefixLength = nil
		return nil
	}
	if i, ok := n.Int(); ok {
		iv := int(i)
		pl.prefixLength = &iv
		return nil
	}
	return fmt.Errorf("%w <%s>", ErrInvalidMaxExpansions, v)
}
func unmarshalPrefixLengthParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithPrefixLength); ok {
		n, err := dynamic.NewNumber(data.UnquotedString())
		if err != nil {
			return &json.UnmarshalTypeError{Value: string(data), Type: typeFloat64}
		}
		if v, ok := n.Int(); ok {
			a.SetPrefixLength(v)
		}
	}
	return nil
}

func marshalPrefixLengthParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithPrefixLength); ok {
		if b.PrefixLength() != DefaultPrefixLength {
			data["prefix_length"] = b.PrefixLength()
		}
	}
	return data, nil
}
