package mapping

// WithScalingFactor is a mapping with the scaling_factor param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html#scaled-float-params
type WithScalingFactor interface {
	ScalingFactor() float32
}

// ScalingFactorParam is a mapping with the scaling_factor param
//
// The scaling factor to use when encoding values. Values will be multiplied by
// this factor at index time and rounded to the closest long value. For
// instance, a scaled_float with a scaling_factor of 10 would internally store
// 2.34 as 23 and all search-time operations (queries, aggregations, sorting)
// will behave as if the document had a value of 2.3. High values of
// scaling_factor improve accuracy but also increase space requirements. This
// parameter is required.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html#scaled-float-params
type ScalingFactorParam struct {
	ScalingFactorValue float32 `json:"scaling_factor" bson:"scaling_factor"`
}

// ScalingFactor to use when encoding values. Values will be multiplied by this
// factor at index time and rounded to the closest long value. For instance, a
// scaled_float with a scaling_factor of 10 would internally store 2.34 as 23
// and all search-time operations (queries, aggregations, sorting) will behave
// as if the document had a value of 2.3. High values of scaling_factor improve
// accuracy but also increase space requirements. This parameter is required.
func (sf ScalingFactorParam) ScalingFactor() float32 {
	return sf.ScalingFactorValue
}

// SetScalingFactor sets the ScalingFactorValue to v
func (sf *ScalingFactorParam) SetScalingFactor(v float32) {
	if sf.ScalingFactor() != v {
		sf.ScalingFactorValue = v
	}
}
