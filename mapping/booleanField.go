package mapping

func NewBooleanField() *BooleanField {
	return &BooleanField{
		BaseField: BaseField{MappingType: TypeBoolean},
	}
}

// BooleanField accepts JSON true and false values, but can also accept strings
// which are interpreted as either true or false:
//
// False values
//  false, "false", "" (empty string)
//
// True values
//  true, "true"
type BooleanField struct {
	BaseField      `bson:",inline" json:",inline"`
	DocValuesParam `bson:",inline" json:",inline"`
	IndexParam     `bson:",inline" json:",inline"`
	NullValueParam `bson:",inline" json:",inline"`
	StoreParam     `bson:",inline" json:",inline"`
	MetaParam      `bson:",inline" json:",inline"`
}

func (b BooleanField) Clone() Field {
	n := NewBooleanField()
	n.SetDocValues(b.DocValues())
	n.SetIndex(b.Index())
	n.SetMeta(b.Meta())
	n.SetMetricType(b.MetricType())
	n.SetNullValue(b.NullValue())
	n.SetStore(b.Store())
	n.SetUnit(b.Unit())
	return n
}
