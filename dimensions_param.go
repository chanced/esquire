package picker

import "github.com/chanced/dynamic"

// WithDimensions is a mapping with the dims parameter
//
// dims is the number of dimensions in the vector, required parameter for fields
// that have it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
type WithDimensions interface {
	// Dimensions is the number of dimensions in the vector, required parameter.
	Dimensions() int
	// SetDimensions sets the dimensions to v
	SetDimensions(v interface{}) error
}

// dimensionsParam is a mapping with the dims parameter
//
// dims is the number of dimensions in the vector, required parameter for fields
// that have it.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
type dimensionsParam struct {
	dimensions dynamic.Number
}

func (d dimensionsParam) Dimensions() int {
	if i, ok := d.dimensions.Int(); ok {
		return i
	}
	return -1
}

func (d *dimensionsParam) SetDimensions(v interface{}) error {
	return d.dimensions.Set(v)

}
