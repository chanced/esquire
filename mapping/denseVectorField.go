package mapping

// DenseVectorField stores dense vectors of float values. The maximum number of
// dimensions that can be in a vector should not exceed 2048. A dense_vector
// field is a single-valued field.
//
//! X-Pack
//
// These vectors can be used for document scoring. For example, a document score
// can represent a distance between a given query vector and the indexed
// document vector.
//
// You index a dense vector as an array of floats.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
type DenseVectorField struct {
	BaseField       `json:",inline" bson:",inline"`
	DimensionsParam `json:",inline" bson:",inline"`
}

func (f DenseVectorField) Clone() Field {
	n := NewDenseVectorField()
	n.SetDims(f.Dims())
	return n
}
func NewDenseVectorField() *DenseVectorField {
	return &DenseVectorField{BaseField: BaseField{MappingType: TypeDenseVector}}
}
