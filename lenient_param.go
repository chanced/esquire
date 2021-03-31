package picker

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

// lenientParam is a query mixin that adds the lenient para
//
// if true, format-based errors, such as providing a text query value for a
// numeric field, are ignored. Defaults to false.
type lenientParam struct {
	// Lenient determines whether format-based errors, such as providing a text
	// query value for a numeric field, are ignored. Defaults to false.
	lenient *bool
}

// Lenient determines whether format-based errors, such as providing a text
// query value for a numeric field, are ignored. Defaults to false.
func (l lenientParam) Lenient() bool {
	if l.lenient != nil {
		return *l.lenient
	}
	return DefaultLenient
}

// SetLenient sets Lenient to v
func (l *lenientParam) SetLenient(v bool) {
	l.lenient = &v
}
func unmarshalLenientParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithLenient); ok {
		var b dynamic.Bool
		err := json.Unmarshal(data, &b)
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			a.SetLenient(v)
			return nil
		}
		if !ok {
			return &json.UnmarshalTypeError{Value: string(data), Type: typeBool}
		}
	}
	return nil
}
func marshalLenientParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithLenient); ok {
		if b.Lenient() {
			return json.Marshal(b.Lenient())
		}
	}
	return nil, nil
}
