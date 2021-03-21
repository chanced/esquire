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

func NewConstantField() *ConstantField {
	return &ConstantField{BaseField: BaseField{MappingType: TypeConstant}}
}

// SetConstantValue sets the ConstantValue to v
func (c *ConstantField) SetConstantValue(v interface{}) *ConstantField {
	c.ConstantValueParam.SetConstantValue(v)
	return c
}
