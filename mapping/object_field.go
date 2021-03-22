package mapping

// An ObjectField is a field type that contains other documents
//
// JSON documents are hierarchical in nature: the document may contain inner
// objects which, in turn, may contain inner objects themselve
type ObjectField struct {
	BaseField       `json:",inline" bson:",inline"`
	PropertiesParam `json:",inline" bson:",inline"`
	EnabledParam    `json:",inline" bson:",inline"`
	DynamicParam    `json:",inline" bson:",inline"`
}

func (f ObjectField) Clone() Field {
	n := NewObjectField()
	n.SetEnabled(f.Enabled())
	n.SetDynamic(f.Dynamic())
	n.SetProperties(f.Properties().Clone())
	return n
}

func NewObjectField() *ObjectField {
	return &ObjectField{BaseField: BaseField{MappingType: TypeObject}}
}
