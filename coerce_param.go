package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

var DefaultCoerce = false

// FieldWithCoerce has a Coerce param
//
// Data is not always clean. Depending on how it is produced a number might be
// rendered in the JSON body as a true JSON number, e.g. 5, but it might also be
// rendered as a string, e.g. "5". Alternatively, a number that should be an
// integer might instead be rendered as a floating point, e.g. 5.0, or even "5.0".
//
// Coercion attempts to clean up dirty values to fit the data type of a field.
// For instance:
//
//      - Strings will be coerced to numbers.
//
//      - Floating points will be truncated for integer values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
type FieldWithCoerce interface {
	Field
	WithCoerce
}

// WithCoerce has a Coerce param
type WithCoerce interface {
	// Coerce attempts to clean up dirty values to fit the data type of a field.
	// For instance:
	//
	// Default: false
	Coerce() bool
	// SetCoerce sets CoerceParam to v
	SetCoerce(v interface{}) error
}

// coerceParam is a mixin for Field mapping types with coerce
//
// Data is not always clean. Depending on how it is produced a number might be
// rendered in the JSON body as a true JSON number, e.g. 5, but it might also be
// rendered as a string, e.g. "5". Alternatively, a number that should be an
// integer might instead be rendered as a floating point, e.g. 5.0, or even
// "5.0".
//
// Coercion attempts to clean up dirty values to fit the data type of a field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
type coerceParam struct {
	coerce dynamic.Bool
}

// Coerce attempts to clean up dirty values to fit the data type of a field.
func (cp coerceParam) Coerce() bool {
	if v, ok := cp.coerce.Bool(); ok {
		return v
	}
	return DefaultCoerce
}

// SetCoerce sets CoerceParam to v
func (cp *coerceParam) SetCoerce(v interface{}) error {
	return cp.coerce.Set(v)
}

func unmarshalCoerceParam(value dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithCoerce); ok {
		b, err := dynamic.NewBool(value)
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			a.SetCoerce(v)
		}
	}
	return nil
}
func marshalCoerceParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithCoerce); ok {
		if !b.Coerce() {
			return json.Marshal(b.Coerce())
		}
	}
	return nil, nil
}
