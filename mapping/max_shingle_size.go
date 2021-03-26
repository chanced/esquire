package mapping

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
	MaxShingleSize() int64
	// SetMaxShingleSize sets the MaxShingleSize to v
	//
	// Valid values are 2 (inclusive) to 4 (inclusive). Defaults to 3.
	SetMaxShingleSize(v int64)
}

// FieldWithMaxShingleSize is a Field mapping with the max_shingle_size parameter
type FieldWithMaxShingleSize interface {
	Field
	WithMaxShingleSize
}

// MaxShingleSizeParam is a mixin that adds the max_shingle_size parameter
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
type MaxShingleSizeParam struct {
	MaxShingleSizeValue *int64 `bson:"max_shingle_size,omitempty" json:"max_shingle_size,omitempty"`
}

// MaxShingleSize is the largest shingle size to create. Valid values are 2
// (inclusive) to 4 (inclusive). Defaults to 3.
func (mss MaxShingleSizeParam) MaxShingleSize() int64 {
	if mss.MaxShingleSizeValue == nil {
		return 3
	}
	return *mss.MaxShingleSizeValue
}

// SetMaxShingleSize sets the MaxShingleSize to v
//
// Valid values are 2 (inclusive) to 4 (inclusive). Defaults to 3.
func (mss *MaxShingleSizeParam) SetMaxShingleSize(v int64) {
	if mss.MaxShingleSize() != v {
		mss.MaxShingleSizeValue = &v
	}
}
