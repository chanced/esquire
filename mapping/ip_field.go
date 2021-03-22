package mapping

// An IPField can index/store either IPv4 or IPv6 addresses.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ip.html
type IPField struct {
	BaseField            `json:",inline" bson:",inline"`
	DocValuesParam       `json:",inline" bson:",inline"`
	IgnoreMalformedParam `json:",inline" bson:",inline"`
	IndexParam           `json:",inline" bson:",inline"`
	NullValueParam       `json:",inline" bson:",inline"`
	StoreParam           `json:",inline" bson:",inline"`
}

func (f IPField) Clone() Field {
	n := NewIPField()
	n.SetDocValues(f.DocValues())
	n.SetIndex(f.Index())
	n.SetNullValue(f.NullValue())
	n.SetStore(f.Store())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	return n
}
func NewIPField() *IPField {
	return &IPField{BaseField: BaseField{MappingType: TypeIP}}
}
