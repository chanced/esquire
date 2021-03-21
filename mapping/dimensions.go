package mapping

// WithDimensions is a mapping with the dims parameter
//
// dims is the number of dimensions in the vector, required parameter for fields
// that have it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
type WithDimensions interface {
	// Dims is the number of dimensions in the vector, required parameter.
	Dims() uint8
	// SetDims sets the dimensions to v
	SetDims(v uint8)
}

// FieldWithDimensions is a Field mapping with the dims parameter
type FieldWithDimensions interface {
	Field
	WithDimensions
}

// DimensionsParam is a mapping with the dims parameter
//
// dims is the number of dimensions in the vector, required parameter for fields
// that have it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
type DimensionsParam struct {
	DimensionsValue uint8 `bson:"dims,omitempty" json:"dims,omitempty"`
}

// Dims is the number of dimensions in the vector, required parameter.
func (d DimensionsParam) Dims() uint8 {
	return d.DimensionsValue
}

// SetDims sets the dimensions to v
func (d *DimensionsParam) SetDims(v uint8) {
	d.DimensionsValue = v
}
