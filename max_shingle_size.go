package picker

import "fmt"

const DefaultMaxShingleSize = 3

// WithMaxShingleSize is a mapping with the max_shingle_size parameter
//
// (Optional, integer) Largest shingle size to create. Valid values are 2
// (inclusive) to 4 (inclusive). Defaults to 3.
//
// A subfield is created for each integer between 2 and this value. For example,
// a value of 3 creates two subfields: my_field._2gram and my_field._3gram
//
// More subfields enables more specific queries but increases index size.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html#specific-params
type WithMaxShingleSize interface {
	// MaxShingleSize is the largest shingle size to create. Valid values are 2
	// (inclusive) to 4 (inclusive). Defaults to 3.
	MaxShingleSize() int
	// SetMaxShingleSize sets the MaxShingleSize to v
	//
	// Valid values are 2 (inclusive) to 4 (inclusive). Defaults to 3.
	SetMaxShingleSize(v int) error
}

// FieldWithMaxShingleSize is a Field mapping with the max_shingle_size parameter
type FieldWithMaxShingleSize interface {
	Field
	WithMaxShingleSize
}

// maxShingleSizeParam is a mixin that adds the max_shingle_size parameter
//
// (Optional, integer) Largest shingle size to create. Valid values are 2
// (inclusive) to 4 (inclusive). Defaults to 3.
//
// A subfield is created for each integer between 2 and this value. For example,
// a value of 3 creates two subfields: my_field._2gram and my_field._3gram
//
// More subfields enables more specific queries but increases index size.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html#specific-params
type maxShingleSizeParam struct {
	maxShingleSize int
}

// MaxShingleSize is the largest shingle size to create. Valid values are 2
// (inclusive) to 4 (inclusive). Defaults to 3.
func (mss maxShingleSizeParam) MaxShingleSize() int {
	if mss.maxShingleSize == 0 {
		return DefaultMaxShingleSize
	}
	return mss.maxShingleSize
}

// SetMaxShingleSize sets the MaxShingleSize to v
//
// Valid values are 2 (inclusive) to 4 (inclusive). Defaults to 3.
func (mss *maxShingleSizeParam) SetMaxShingleSize(v int) error {
	if v == 0 {
		mss.maxShingleSize = 0
		return nil
	}
	if v == 1 || v > 4 {
		return fmt.Errorf("%w; received %d", ErrInvalidMaxShingleSize, v)
	}
	mss.maxShingleSize = v
	return nil
}
