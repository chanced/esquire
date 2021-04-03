package picker

import "github.com/chanced/dynamic"

const DefaultPositionIncrementGap = float64(100)

// WithPositionIncrementGap is a mapping with the position_increment_gap
// parameter
//
// Analyzed text fields take term positions into account, in order to be able to
// support proximity or phrase queries. When indexing text fields with multiple
// values a "fake" gap is added between the values to prevent most phrase
// queries from matching across the values. The size of this gap is configured
// using position_increment_gap and defaults to 100.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/position-increment-gap.html
type WithPositionIncrementGap interface {
	// PositionIncrementGap is the number of fake term position which should be
	// inserted between each element of an array of strings. Defaults to the
	// position_increment_gap configured on the analyzer which defaults to 100.
	// 100 was chosen because it prevents phrase queries with reasonably large
	// slops (less than 100) from matching terms across field values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/position-increment-gap.html
	PositionIncrementGap() int
	// SetPositionIncrementGap sets the PositionIncrementGap Value to v
	SetPositionIncrementGap(v interface{}) error
}

// FieldWithPositionIncrementGap is a Field mapping with the
// position_increment_gap parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/position-increment-gap.html
type FieldWithPositionIncrementGap interface {
	Field
	WithPositionIncrementGap
}

// positionIncrementGapParam is a mixin that adds the position_increment_gap
// parameter
//
// Analyzed text fields take term positions into account, in order to be able to
// support proximity or phrase queries. When indexing text fields with multiple
// values a "fake" gap is added between the values to prevent most phrase
// queries from matching across the values. The size of this gap is configured
// using position_increment_gap and defaults to 100.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/position-increment-gap.html
type positionIncrementGapParam struct {
	positionIncrementGap dynamic.Number `json:"position_increment_gap,omitempty"`
}

// PositionIncrementGap is the number of fake term position which should be
// inserted between each element of an array of strings. Defaults to the
// position_increment_gap configured on the analyzer which defaults to 100.
// 100 was chosen because it prevents phrase queries with reasonably large
// slops (less than 100) from matching terms across field values.
func (pig positionIncrementGapParam) PositionIncrementGap() float64 {
	if f, ok := pig.positionIncrementGap.Float64(); ok {
		return f
	}
	return DefaultPositionIncrementGap
}

// SetPositionIncrementGap sets the PositionIncrementGap Value to v
func (pig *positionIncrementGapParam) SetPositionIncrementGap(v interface{}) error {
	return pig.positionIncrementGap.Set(v)
}
