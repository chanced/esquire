package picker

import "github.com/chanced/dynamic"

const DefaultIgnoreZ = true

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
	SetIgnoreZValue(v interface{}) error
}

type ignoreZValueParam struct {
	ignoreZ dynamic.Bool
}

// IgnoreZValue determines whether three dimension points will be indexed.
//
// If true (default) three dimension points will be accepted (stored in
// source) but only latitude and longitude values will be indexed; the third
// dimension is ignored. If false, geo-points containing any more than
// latitude and longitude (two dimensions) values throw an exception and
// reject the whole document.
func (z ignoreZValueParam) IgnoreZValue() bool {
	if b, ok := z.ignoreZ.Bool(); ok {
		return b
	}
	return DefaultIgnoreZ
}

// SetIgnoreZValue sets the IgnoreZValue Value to v
func (z *ignoreZValueParam) SetIgnoreZValue(v interface{}) error {
	return z.ignoreZ.Set(v)
}
