package mapping

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

func NewBinaryField() *BinaryField {
	return &BinaryField{
		BaseField: BaseField{
			MappingType: TypeBinary,
		},
	}
}

func (b *BinaryField) SetDocValues(v bool) *BinaryField {
	b.DocValuesParam.SetDocValues(v)
	return b
}
func (b *BinaryField) SetStore(v bool) *BinaryField {
	b.StoreParam.SetStore(v)
	return b
}
