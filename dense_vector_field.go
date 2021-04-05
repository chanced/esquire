package picker

import "encoding/json"

//easyjson:json
type denseVectorField struct {
	Dimensions interface{} `json:"dims,omitempty"`
	Type       FieldType   `json:"type"`
}

type DenseVectorFieldParams struct {
	// Dimensions is the number of dimensions in the vector, required parameter.
	Dimensions interface{} `json:"dims,omitempty"`
}

func (DenseVectorFieldParams) Type() FieldType {
	return FieldTypeDenseVector
}

func (p DenseVectorFieldParams) Field() (Field, error) {
	return p.DenseVector()
}

func (p DenseVectorFieldParams) DenseVector() (*DenseVectorField, error) {
	f := &DenseVectorField{}
	e := &MappingError{}
	err := f.SetDimensions(p.Dimensions)
	e.Append(err)
	return f, e.ErrorOrNil()
}

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
	dimensionsParam
}

func (DenseVectorField) Type() FieldType {
	return FieldTypeDenseVector
}
func (dv *DenseVectorField) UnmarshalJSON(data []byte) error {
	var p DenseVectorFieldParams
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.DenseVector()
	*dv = *n
	return err
}
func (dv DenseVectorField) MarshalJSON() ([]byte, error) {
	return json.Marshal(denseVectorField{
		Dimensions: dv.dimensions.Value(),
		Type:       dv.Type(),
	})
}
func (dv *DenseVectorField) Field() (Field, error) {
	return dv, nil
}

func NewDenseVectorField(params DenseVectorFieldParams) (*DenseVectorField, error) {
	return params.DenseVector()
}
