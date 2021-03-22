package mapping

// A RangeField type represent a continuous range of values between an upper and
// lower bound. For example, a range can represent any date in October or any
// integer from 0 to 9. They are defined using the operators gt or gte for the
// lower bound, and lt or lte for the upper bound. They can be used for
// querying, and have limited support for aggregations. The only supported
// aggregations are histogram, cardinality.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html

// IntegerRangeField is a range  of signed 32-bit integers with a minimum value
// of -231 and maximum of 231-1.
type IntegerRangeField struct {
	BaseField   `bson:",inline" json:",inline"`
	CoerceParam `bson:",inline" json:",inline"`
	IndexParam  `bson:",inline" json:",inline"`
	StoreParam  `bson:",inline" json:",inline"`
}

func (f IntegerRangeField) Clone() Field {
	n := NewIntegerRangeField()
	n.SetCoerce(f.Coerce())
	n.SetIndex(f.Index())
	n.SetStore(f.Store())
	return n
}

func NewIntegerRangeField() *IntegerRangeField {
	return &IntegerRangeField{BaseField: BaseField{MappingType: TypeIntegerRange}}
}

// FloatRangeField is a range of single-precision 32-bit IEEE 754 floating point
// values.
type FloatRangeField struct {
	BaseField   `bson:",inline" json:",inline"`
	CoerceParam `bson:",inline" json:",inline"`
	IndexParam  `bson:",inline" json:",inline"`
	StoreParam  `bson:",inline" json:",inline"`
}

func (f FloatRangeField) Clone() Field {
	n := NewFloatRangeField()
	n.SetCoerce(f.Coerce())
	n.SetIndex(f.Index())
	n.SetStore(f.Store())
	return n
}

func NewFloatRangeField() *FloatRangeField {
	return &FloatRangeField{BaseField: BaseField{MappingType: TypeFloatRange}}
}

// LongRangeField is a range of signed 64-bit integers with a minimum value of
// -263 and maximum of 263-1.
type LongRangeField struct {
	BaseField   `bson:",inline" json:",inline"`
	CoerceParam `bson:",inline" json:",inline"`
	IndexParam  `bson:",inline" json:",inline"`
	StoreParam  `bson:",inline" json:",inline"`
}

func (f LongRangeField) Clone() Field {
	n := NewLongRangeField()
	n.SetCoerce(f.Coerce())
	n.SetIndex(f.Index())
	n.SetStore(f.Store())
	return n
}
func NewLongRangeField() *LongRangeField {
	return &LongRangeField{BaseField: BaseField{MappingType: TypeLongRange}}
}

// DoubleRangeField is a range of double-precision 64-bit IEEE 754 floating
// point values.
type DoubleRangeField struct {
	BaseField   `bson:",inline" json:",inline"`
	CoerceParam `bson:",inline" json:",inline"`
	IndexParam  `bson:",inline" json:",inline"`
	StoreParam  `bson:",inline" json:",inline"`
}

func (f DoubleRangeField) Clone() Field {
	n := NewDoubleRangeField()
	n.SetCoerce(f.Coerce())
	n.SetIndex(f.Index())
	n.SetStore(f.Store())
	return n
}
func NewDoubleRangeField() *DoubleRangeField {
	return &DoubleRangeField{BaseField: BaseField{MappingType: TypeDoubleRange}}
}

// DateRangeField is a range of date values. Date ranges support various date
// formats through the format mapping parameter. Regardless of the format used,
// date values are parsed into an unsigned 64-bit integer representing
// milliseconds since the Unix epoch in UTC. Values containing the now date math
// expression are not supported.
type DateRangeField struct {
	BaseField   `bson:",inline" json:",inline"`
	CoerceParam `bson:",inline" json:",inline"`
	IndexParam  `bson:",inline" json:",inline"`
	StoreParam  `bson:",inline" json:",inline"`
}

func (f DateRangeField) Clone() Field {
	n := NewDateRangeField()
	n.SetCoerce(f.Coerce())
	n.SetIndex(f.Index())
	n.SetStore(f.Store())
	return n
}
func NewDateRangeField() *DateRangeField {
	return &DateRangeField{BaseField: BaseField{MappingType: TypeDateRange}}
}

// IPRangeField is a range of ip values supporting either IPv4 or IPv6 (or
// mixed) addresses.
type IPRangeField struct {
	BaseField   `bson:",inline" json:",inline"`
	CoerceParam `bson:",inline" json:",inline"`
	IndexParam  `bson:",inline" json:",inline"`
	StoreParam  `bson:",inline" json:",inline"`
}

func (f IPRangeField) Clone() Field {
	n := NewIPRangeField()
	n.SetCoerce(f.Coerce())
	n.SetIndex(f.Index())
	n.SetStore(f.Store())
	return n
}
func NewIPRangeField() *IPRangeField {
	return &IPRangeField{BaseField: BaseField{MappingType: TypeIPRange}}
}
