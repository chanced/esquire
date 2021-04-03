package picker

// FieldDataFrequencyFilter is utilized with mappings
type FieldDataFrequencyFilter struct {
	Min            float64 `bson:"min" json:"min"`
	Max            float64 `bson:"max" json:"max"`
	MinSegmentSize int     `bson:"min_segment_size" json:"min_segment_size"`
}

// WithFieldDataFrequencyFilter is a mapping with the
// field_data_frequency_filter
//
// FieldDataFrequencyFilter is an expert settings which allow to decide which
// values to load in memory when fielddata is enabled. By default all values are
// loaded.
//
// FieldData filtering can be used to reduce the number of terms loaded into
// memory, and thus reduce memory usage. Terms can be filtered by frequency:
//
// The frequency filter allows you to only load terms whose document frequency
// falls between a min and max value, which can be expressed an absolute number
// (when the number is bigger than 1.0) or as a percentage (eg 0.01 is 1% and
// 1.0 is 100%). Frequency is calculated per segment. Percentages are based on
// the number of docs which have a value for the field, as opposed to all docs
// in the segment.
//
// Small segments can be excluded completely by specifying the minimum number of
// docs that the segment should contain with min_segment_size
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html#field-data-filtering
type WithFieldDataFrequencyFilter interface {
	// FieldDataFrequencyFilter can be used to reduce the number of terms loaded
	// into memory, and thus reduce memory usage when using FieldData filtering.
	FieldDataFrequencyFilter() *FieldDataFrequencyFilter
	// SetFieldDataFrequencyFilter sets the FieldDataFrequencyFilter value to v
	SetFieldDataFrequencyFilter(v *FieldDataFrequencyFilter) error
}

// FieldWithFieldDataFrequencyFilter is a Field with a FieldDataFrequencyFilter
// param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html#field-data-filtering
type FieldWithFieldDataFrequencyFilter interface {
	Field
	WithFieldDataFrequencyFilter
}

// fieldDataFrequencyFilterParam is a mixin that adds the
// field_data_frequency_filter param
//
// FieldDataFrequencyFilter is an expert settings which allow to decide which
// values to load in memory when fielddata is enabled. By default all values are
// loaded.
//
// FieldData filtering can be used to reduce the number of terms loaded into
// memory, and thus reduce memory usage. Terms can be filtered by frequency:
//
// The frequency filter allows you to only load terms whose document frequency
// falls between a min and max value, which can be expressed an absolute number
// (when the number is bigger than 1.0) or as a percentage (eg 0.01 is 1% and
// 1.0 is 100%). Frequency is calculated per segment. Percentages are based on
// the number of docs which have a value for the field, as opposed to all docs
// in the segment.
//
// Small segments can be excluded completely by specifying the minimum number of
// docs that the segment should contain with min_segment_size
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html#field-data-filtering
type fieldDataFrequencyFilterParam struct {
	fieldDataFrequencyFilter *FieldDataFrequencyFilter `bson:"fielddata_frequency_filter,omitempty" json:"fielddata_frequency_filter,omitempty"`
}

// FieldDataFrequencyFilter can be used to reduce the number of terms loaded into memory, and
// thus reduce memory usage when using FieldData filtering.
func (fd fieldDataFrequencyFilterParam) FieldDataFrequencyFilter() *FieldDataFrequencyFilter {
	return fd.fieldDataFrequencyFilter
}

// SetFieldDataFrequencyFilter sets the FieldDataFrequencyFilter value to v
func (fd *fieldDataFrequencyFilterParam) SetFieldDataFrequencyFilter(v *FieldDataFrequencyFilter) error {
	fd.fieldDataFrequencyFilter = v
	// TODO: Validation
	return nil
}
