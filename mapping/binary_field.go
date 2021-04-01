package mapping

type BinaryFieldParams struct {
}

func (BinaryField) Type() FieldType {
	return FieldTypeBinary
}
func NewBinaryField(params BinaryFieldParams) *BinaryField {

}

// Whether the field value should be stored and retrievable separately from the _source field. Accepts true or false (default).
// Binary

// BinaryField is a value encoded as a Base64 string.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/binary.html
type BinaryField struct {
	docValuesParam `json:",inline" bson:",inline"`
	storeParam     `json:",inline" bson:",inline"`
}

func (BinaryField) Type() FieldType {
	return FieldTypeBinary
}
