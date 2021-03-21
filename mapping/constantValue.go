package mapping

// WithConstantValue is a mapping with the value parameter
//
// The value to associate with all documents in the index. If this parameter is
// not provided, it is set based on the first document that gets indexed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-params
type WithConstantValue interface {
	// ConstantValue is the constant value used in a ConstantKeywordField
	// ConstantValue is the value to associate with all documents in the index.
	// If this parameter is not provided, it is set based on the first document
	// that gets indexed.
	ConstantValue() interface{}
	// SetConstantValue sets the Constant Value to v
	SetConstantValue(v interface{})
}

// FieldWithConstantValue is a Field mapping with the value parameter
type FieldWithConstantValue interface {
	Field
	WithConstantValue
}

// ConstantValueParam is a mapping with the value parameter
//
// The value to associate with all documents in the index. If this parameter is
// not provided, it is set based on the first document that gets indexed.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-params
type ConstantValueParam struct {
	ConstantValueValue interface{} `bson:"value,omitempty" json:"value,omitempty"`
}

// ConstantValue is the constant value used in a ConstantKeywordField
// ConstantValue is the value to associate with all documents in the index.
// If this parameter is not provided, it is set based on the first document
// that gets indexed.
func (cv ConstantValueParam) ConstantValue() interface{} {
	return cv.ConstantValue
}

// SetConstantValue sets the ConstantValue to v
func (cv *ConstantValueParam) SetConstantValue(v interface{}) {
	cv.ConstantValueValue = v
}
