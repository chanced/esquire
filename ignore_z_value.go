package picker

// WithIgnoreZValue is a mapping with the ignore_z_value parameter
//
// If IgnoreZValue is true (default) three dimension points will be accepted
// (stored in source) but only latitude and longitude values will be indexed;
// the third dimension is ignored. If false, geo-points containing any more than
// latitude and longitude (two dimensions) values throw an exception and reject
// the whole document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html#geo-point-params
type WithIgnoreZValue interface {
	// IgnoreZValue determines whether three dimension points will be indexed.
	//
	// If true (default) three dimension points will be accepted (stored in
	// source) but only latitude and longitude values will be indexed; the third
	// dimension is ignored. If false, geo-points containing any more than
	// latitude and longitude (two dimensions) values throw an exception and
	// reject the whole document.
	IgnoreZValue() bool
	// SetIgnoreZValue sets the IgnoreZValue Value to v
	SetIgnoreZValue(v bool)
}

// FieldWithIgnoreZValue is a Field mapping with the ignore_z_value parameter
//
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html#geo-point-params
type FieldWithIgnoreZValue interface {
	Field
	WithIgnoreZValue
}

// IgnoreZValueParam is a mixin that adds the ignore_z_value parameter
//
// If IgnoreZValue is true (default) three dimension points will be accepted
// (stored in source) but only latitude and longitude values will be indexed;
// the third dimension is ignored. If false, geo-points containing any more than
// latitude and longitude (two dimensions) values throw an exception and reject
// the whole document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html#geo-point-params
type IgnoreZValueParam struct {
	IgnoreZValueValue *bool `bson:"ignore_z_value,omitempty" json:"ignore_z_value,omitempty"`
}

// IgnoreZValue determines whether three dimension points will be indexed.
//
// If true (default) three dimension points will be accepted (stored in
// source) but only latitude and longitude values will be indexed; the third
// dimension is ignored. If false, geo-points containing any more than
// latitude and longitude (two dimensions) values throw an exception and
// reject the whole document.
func (izv IgnoreZValueParam) IgnoreZValue() bool {
	if izv.IgnoreZValueValue == nil {
		return true
	}
	return *izv.IgnoreZValueValue
}

// SetIgnoreZValue sets the IgnoreZValue Value to v
func (izv *IgnoreZValueParam) SetIgnoreZValue(v bool) {
	if izv.IgnoreZValue() != v {
		izv.IgnoreZValueValue = &v
	}
}
