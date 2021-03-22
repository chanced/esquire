package mapping

// ConstantField is a specialization of the Keyword field for the case
// that all documentsin the index have the same value.
//
// ! X-Pack
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-field-type
type ConstantField struct {
	BaseField          `bson:",inline" json:",inline"`
	ConstantValueParam `bson:",inline" json:",inline"`
}

func (f ConstantField) Clone() Field {
	n := NewConstantField()
	n.SetConstantValue(f.ConstantValue())
	return n
}
func NewConstantField() *ConstantField {
	return &ConstantField{BaseField: BaseField{MappingType: TypeConstant}}
}
