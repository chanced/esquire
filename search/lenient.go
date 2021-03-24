package search

import "github.com/tidwall/gjson"

// WithLenient is a query with the lenient param
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
	LenientValue *bool `json:"lenient,omitempty" bson:"lenient,omitempty"`
}

// Lenient determines whether format-based errors, such as providing a text
// query value for a numeric field, are ignored. Defaults to false.
func (l LenientParam) Lenient() bool {
	if l.LenientValue != nil {
		return *l.LenientValue
	}
	return false
}

// SetLenient sets Lenient to v
func (l *LenientParam) SetLenient(v bool) {
	l.LenientValue = &v
}
func unmarshalLenientParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithLenient); ok {
		a.SetLenient(value.Bool())
	}
	return nil
}
