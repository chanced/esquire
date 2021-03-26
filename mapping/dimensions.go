package mapping

// WithDimensions is a mapping with the dims parameter
//
// dims is the number of dimensions in the vector, required parameter for fields
// that have it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
type WithDimensions interface {
	// Dims is the number of dimensions in the vector, required parameter.
	Dims() int64
	// SetDims sets the dimensions to v
	SetDims(v int64)
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
	DimensionsValue int64 `bson:"dims,omitempty" json:"dims,omitempty"`
}

func (d DimensionsParam) Dimensions() int64 {
	return d.DimensionsValue
}

// Dims is the number of dimensions in the vector, required parameter.
func (d DimensionsParam) Dims() int64 {
	return d.DimensionsValue
}

// SetDims sets the dimensions to v
func (d *DimensionsParam) SetDims(v int64) {
	d.DimensionsValue = v
}
func (d *DimensionsParam) SetDimensions(v int64) {
	d.DimensionsValue = v
}
