package picker

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultFuzzyPrefixLength = int(0)

type WithFuzzyPrefixLength interface {
	FuzzyPrefixLength() int
	SetFuzzyPrefixLength(v interface{}) error
}
type fuzzyPrefixLengthParam struct {
	fuzzyPrefixLength dynamic.Number
}

func (me fuzzyPrefixLengthParam) FuzzyPrefixLength() int {
	if me.fuzzyPrefixLength.HasValue() {
		if i, ok := me.fuzzyPrefixLength.Int(); ok {
			return i
		}
		if f, ok := me.fuzzyPrefixLength.Float64(); ok {
			return int(f)
		}
	}
	return DefaultFuzzyPrefixLength
}
func (me *fuzzyPrefixLengthParam) SetFuzzyPrefixLength(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if n.IsNil() {
		_ = me.fuzzyPrefixLength.Set(nil)
		return nil
	}
	if i, ok := n.Int(); ok {
		_ = me.fuzzyPrefixLength.Set(i)
		return nil
	}
	return fmt.Errorf("%w <%s>", errors.New("picker: invalid fuzzy prefix length"), v)
}
func unmarshalFuzzyPrefixLengthParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithFuzzyPrefixLength); ok {
		n, err := dynamic.NewNumber(data.UnquotedString())
		if err != nil {
			return err
		}
		if v, ok := n.Int(); ok {
			return a.SetFuzzyPrefixLength(v)
		}
		return nil
	}
	return nil
}
func marshalFuzzyPrefixLengthParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithFuzzyPrefixLength); ok {
		if b.FuzzyPrefixLength() != DefaultFuzzyPrefixLength {
			return json.Marshal(b.FuzzyPrefixLength())
		}
	}
	return nil, nil
}
