package mapping

// WithIgnoreAbove is a mapping the ignore_above parameter
//
//Strings longer than the ignore_above setting will not be indexed or stored.
//For arrays of strings, ignore_above will be applied for each array element
//separately and string elements longer than ignore_above will not be indexed or
//stored.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-above.html
type WithIgnoreAbove interface {
	// IgnoreAbove signiall to not index any string longer than this value.
	// Defaults to 2147483647 so that all values would be accepted. Please
	// however note that default dynamic mapping rules create a sub keyword
	// field that overrides this default by setting ignore_above: 256.
	IgnoreAbove() uint
	// SetIgnoreAbove sets the IgnoreAbove value to v
	SetIgnoreAbove(v uint)
}

// FieldWithIgnoreAbove is a Field with the IgnoreAbove param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-above.html
type FieldWithIgnoreAbove interface {
	Field
	WithIgnoreAbove
}

// IgnoreAboveParam is a mixin for mappings that adds the IgnoreAbove
// (ignore_above) parameter
//
// Strings longer than the ignore_above setting will not be indexed or stored.
// For arrays of strings, ignore_above will be applied for each array element
// separately and string elements longer than ignore_above will not be indexed
// or stored.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-above.html
type IgnoreAboveParam struct {
	IgnoreAboveValue *uint `bson:"ignore_above,omitempty" json:"ignore_above,omitempty"`
}

// IgnoreAbove signiall to not index any string longer than this value.
// Defaults to 2147483647 so that all values would be accepted. Please
// however note that default dynamic mapping rules create a sub keyword
// field that overrides this default by setting ignore_above: 256.
func (ia IgnoreAboveParam) IgnoreAbove() uint {
	if ia.IgnoreAboveValue == nil {
		return 0
	}
	return *ia.IgnoreAboveValue
}

// SetIgnoreAbove sets the IgnoreAbove value to v
func (ia IgnoreAboveParam) SetIgnoreAbove(v uint) {
	ia.IgnoreAboveValue = &v
}
