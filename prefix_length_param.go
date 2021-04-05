package picker

import (
	"fmt"

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
	PrefixLength() int
	SetPrefixLength(v interface{}) error
}

// prefixLengthParam is a mixin that adds the prefix_length param
//
// PrefixLength is the number of beginning characters left unchanged for fuzzy matching. Defaults to 0.
type prefixLengthParam struct {
	prefixLength dynamic.Number
}

func (pl prefixLengthParam) PrefixLength() int {
	if i, ok := pl.prefixLength.Int(); ok {
		return i
	}
	if f, ok := pl.prefixLength.Float64(); ok {
		return int(f)
	}
	return DefaultPrefixLength
}

func (pl *prefixLengthParam) SetPrefixLength(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if n.IsNil() {
		pl.prefixLength.Set(nil)
		return nil
	}
	if i, ok := n.Int(); ok {
		pl.prefixLength.Set(i)
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
			return a.SetPrefixLength(v)
		}
	}
	return nil
}

func marshalPrefixLengthParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithPrefixLength); ok {
		if b.PrefixLength() != DefaultPrefixLength {
			return json.Marshal(b.PrefixLength())
		}
	}
	return nil, nil
}
