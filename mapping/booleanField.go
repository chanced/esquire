package mapping

// BooleanField is a Field Type
//
// Boolean fields accept JSON true and false values, but can also accept strings which are interpreted as either true or false:
//
// False values
// 	false, "false", "" (empty string)
//
// True values
// 	true, "true"
type BooleanField struct {
	BaseField      `bson:",inline" json:",inline"`
	DocValuesParam `bson:",inline" json:",inline"`
	IndexParam     `bson:",inline" json:",inline"`
	NullValueParam `bson:",inline" json:",inline"`
	StoreParam     `bson:",inline" json:",inline"`
	MetaParam      `bson:",inline" json:",inline"`
}

func NewBooleanField() *BooleanField {
	return &BooleanField{
		BaseField: BaseField{MappingType: TypeBoolean},
	}
}

func (b *BooleanField) SetDocValues(v bool) *BooleanField {
	b.DocValuesParam.SetDocValues(v)
	return b
}

func (b *BooleanField) SetIndex(v bool) *BooleanField {
	b.IndexParam.SetIndex(v)
	return b
}

func (b *BooleanField) SetNullValue(v interface{}) *BooleanField {
	b.NullValueParam.SetNullValue(v)
	return b
}

func (b *BooleanField) SetStore(v bool) *BooleanField {
	b.StoreParam.SetStore(v)
	return b
}

func (b *BooleanField) SetMetaParam(v map[string]string) *BooleanField {
	b.MetaParam.SetMeta(v)
	return b
}
