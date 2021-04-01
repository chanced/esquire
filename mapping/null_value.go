package mapping

// WithNullValue is a mapping with the null_value parameter
//
// A null value cannot be indexed or searched. When a field is set to null, (or
// an empty array or an array of null values) it is treated as though that field
// has no values.
//
// The null_value parameter allows you to replace explicit null values with the
// specified value so that it can be indexed and searched
//
// The null_value needs to be the same data type as the field. For instance, a
// long field cannot have a string null_value.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/null-value.html
type WithNullValue interface {
	// NullValue parameter allows you to replace explicit null values with the
	// specified value so that it can be indexed and searched
	NullValue() interface{}
	// SetNullValue sets the NullValue value to v
	SetNullValue(v interface{})
}

type nullValueParam struct {
	nullValue interface{}
}

// NullValue parameter allows you to replace explicit null values with the
// specified value so that it can be indexed and searched
func (nv nullValueParam) NullValue() interface{} {
	return nv.nullValue
}

// SetNullValue sets the NullValue value to v
func (nv *nullValueParam) SetNullValue(v interface{}) {
	nv.nullValue = v
}
