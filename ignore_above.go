package picker

import "github.com/chanced/dynamic"

var DefaultIgnoreAbove = float64(2147483647)

// WithIgnoreAbove is a mapping the ignore_above parameter
//
// Strings longer than the ignore_above setting will not be indexed or stored.
// For arrays of strings, ignore_above will be applied for each array element
// separately and string elements longer than ignore_above will not be indexed or
// stored.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-above.html
type WithIgnoreAbove interface {
	// IgnoreAbove signiall to not index any string longer than this value.
	// Defaults to 2147483647 so that all values would be accepted. Please
	// however note that default dynamic mapping rules create a sub keyword
	// field that overrides this default by setting ignore_above: 256.
	IgnoreAbove() float64
	// SetIgnoreAbove sets the IgnoreAbove value to v
	SetIgnoreAbove(v interface{}) error
}

// FieldWithIgnoreAbove is a Field with the IgnoreAbove param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-above.html
type FieldWithIgnoreAbove interface {
	Field
	WithIgnoreAbove
}

// ignoreAboveParam is a mixin for mappings that adds the IgnoreAbove

type ignoreAboveParam struct {
	ignoreAbove dynamic.Number
}

// IgnoreAbove signals to not index any string longer than this value.
// Defaults to 2147483647 so that all values would be accepted. Please
// however note that default dynamic mapping rules create a sub keyword
// field that overrides this default by setting ignore_above: 256.
func (ia ignoreAboveParam) IgnoreAbove() float64 {
	if f, ok := ia.ignoreAbove.Float64(); ok {
		return f
	}
	return DefaultIgnoreAbove
}

// SetIgnoreAbove sets the IgnoreAbove value to v
func (ia ignoreAboveParam) SetIgnoreAbove(v interface{}) error {
	return ia.ignoreAbove.Set(v)
}
