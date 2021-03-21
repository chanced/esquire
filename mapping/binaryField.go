package mapping

func NewBinaryField() *BinaryField {
	return &BinaryField{
		BaseField: BaseField{
			MappingType: TypeBinary,
		},
	}
}

// Whether the field value should be stored and retrievable separately from the _source field. Accepts true or false (default).
// Binary

// BinaryField is a value encoded as a Base64 string.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/binary.html
type BinaryField struct {
	BaseField      `json:",inline" bson:",inline"`
	DocValuesParam `json:",inline" bson:",inline"`
	StoreParam     `json:",inline" bson:",inline"`
}

func (b BinaryField) Clone() Field {
	n := NewBinaryField()
	n.SetDocValues(b.DocValues())
	return n
}
