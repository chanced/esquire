package mapping

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
//  	- Strings will be coerced to numbers.
//
//  	- Floating points will be truncated for integer values.
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
	SetCoerce(v bool)
}

// CoerceParam is a mixin for Field mapping types with coerce
//
// Data is not always clean. Depending on how it is produced a number might be
// rendered in the JSON body as a true JSON number, e.g. 5, but it might also be
// rendered as a string, e.g. "5". Alternatively, a number that should be an
// integer might instead be rendered as a floating point, e.g. 5.0, or even
// "5.0".
//
// Coercion attempts to clean up dirty values to fit the data type of a field.
// For instance:
//
//  	- Strings will be coerced to numbers.
//
//  	- Floating points will be truncated for integer values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
type CoerceParam struct {
	CoerceValue *bool `bson:"coerce,omitempty" json:"coerce,omitempty"`
}

// Coerce attempts to clean up dirty values to fit the data type of a field.
func (cp CoerceParam) Coerce() bool {
	if cp.CoerceValue != nil {
		return *cp.CoerceValue
	}
	return false
}

// SetCoerce sets CoerceParam to v
func (cp *CoerceParam) SetCoerce(v bool) {
	if cp.Coerce() != v {
		cp.CoerceValue = &v
	}
}
