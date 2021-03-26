package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultLenient = false

// WithLenient is a query with the lenient param
//
// Lenient determines whether format-based errors, such as providing a text
// query value for a numeric field, are ignored. Defaults to false.
type WithLenient interface {
	// Lenient determines whether format-based errors, such as providing a text
	// query value for a numeric field, are ignored. Defaults to false.
	Lenient() bool
	// SetLenient sets Lenient to v
	SetLenient(v bool)
}

// LenientParam is a query mixin that adds the lenient para
//
// if true, format-based errors, such as providing a text query value for a
// numeric field, are ignored. Defaults to false.
type LenientParam struct {
	// Lenient determines whether format-based errors, such as providing a text
	// query value for a numeric field, are ignored. Defaults to false.
	LenientValue *bool `json:"lenient,omitempty" bson:"lenient,omitempty"`
}

// Lenient determines whether format-based errors, such as providing a text
// query value for a numeric field, are ignored. Defaults to false.
func (l LenientParam) Lenient() bool {
	if l.LenientValue != nil {
		return *l.LenientValue
	}
	return DefaultLenient
}

// SetLenient sets Lenient to v
func (l *LenientParam) SetLenient(v bool) {
	l.LenientValue = &v
}
func unmarshalLenientParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithLenient); ok {
		b := dynamic.NewBool(data.UnquotedString())
		if v, ok := b.Bool(); ok {
			a.SetLenient(v)
			return nil
		}
		if !ok {
			return &json.UnmarshalTypeError{Value: data.String()}
		}
	}
	return nil
}
func marshalLenientParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithLenient); ok {
		if b.Lenient() != DefaultLenient {
			data[paramLenient] = b.Lenient()
		}
	}
	return data, nil
}
