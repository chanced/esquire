package mapping

// As far as integer types (byte, short, integer and long) are concerned, you
// should pick the smallest type which is enough for your use-case. This will
// help indexing and searching be more efficient. Note however that storage is
// optimized based on the actual values that are stored, so picking one type
// over another one will have no impact on storage requirements.
//
// For floating-point types, it is often more efficient to store floating-point
// data into an integer using a scaling factor, which is what the scaled_float
// type does under the hood. For instance, a price field could be stored in a
// scaled_float with a scaling_factor of 100. All APIs would work as if the
// field was stored as a double, but under the hood Elasticsearch would be
// working with the number of cents, price*100, which is an integer. This is
// mostly helpful to save disk space since integers are way easier to compress
// than floating points. scaled_float is also fine to use in order to trade
// accuracy for disk space. For instance imagine that you are tracking cpu
// utilization as a number between 0 and 1. It usually does not matter much
// whether cpu utilization is 12.7% or 13%, so you could use a scaled_float with
// a scaling_factor of 100 in order to round cpu utilization to the closest
// percent in order to save space.
//
// If scaled_float is not a good fit, then you should pick the smallest type
// that is enough for the use-case among the floating-point types: double, float
// and half_float. Here is a table that compares these types in order to help
// make a decision.
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html

// A LongField is a signed 64-bit integer with a minimum value of -263
// and a maximum value of 263-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type LongField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f LongField) Clone() Field {
	n := NewLongField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewLongField() *LongField {
	return &LongField{BaseField: BaseField{MappingType: FieldTypeLong}}
}

// An IntegerField is a signed 64-bit integer with a minimum value of -263
// and a maximum value of 263-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type IntegerField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f IntegerField) Clone() Field {
	n := NewIntegerField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewIntegerField() *IntegerField {
	return &IntegerField{BaseField: BaseField{MappingType: FieldTypeInteger}}
}

// A ShortField is signed 16-bit integer with a minimum value of -32,768
// and a maximum value of 32,767.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type ShortField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f ShortField) Clone() Field {
	n := NewShortField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewShortField() *ShortField {
	return &ShortField{BaseField: BaseField{MappingType: FieldTypeShort}}
}

// A ByteField is a signed 8-bit integer with a minimum value of -128 and a maximum value of 127.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type ByteField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f ByteField) Clone() Field {
	n := NewByteField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewByteField() *ByteField {
	return &ByteField{BaseField: BaseField{MappingType: FieldTypeByte}}
}

// A DoubleField is a double-precision 64-bit IEEE 754 floating point number, restricted to finite values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type DoubleField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f DoubleField) Clone() Field {
	n := NewDoubleField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewDoubleField() *DoubleField {
	return &DoubleField{BaseField: BaseField{MappingType: FieldTypeDouble}}
}

// A FloatField is a single-precision 32-bit IEEE 754 floating point
// number, restricted to finite values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type FloatField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f FloatField) Clone() Field {
	n := NewFloatField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewFloatField() *FloatField {
	return &FloatField{
		BaseField: BaseField{
			MappingType: FieldTypeFloat,
		},
	}
}

// A HalfFloatField is a half-precision 16-bit IEEE 754 floating point
// number, restricted to finite values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type HalfFloatField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f HalfFloatField) Clone() Field {
	n := NewHalfFloatField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}

func NewHalfFloatField() *HalfFloatField {
	return &HalfFloatField{
		BaseField: BaseField{MappingType: FieldTypeHalfFloat},
	}
}

// A ScaledFloatField is a floating point number that is backed by a long,
// scaled by a fixed double scaling factor.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type ScaledFloatField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
	ScalingFactorParam   `bson:",inline" json:",inline"`
}

func (f ScaledFloatField) Clone() Field {
	n := NewScaledFloatField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	n.SetScalingFactor(f.ScalingFactor())
	return n
}

func NewScaledFloatField() *ScaledFloatField {
	return &ScaledFloatField{BaseField: BaseField{MappingType: FieldTypeScaledFloat}}
}

// An UnsignedLongField is an unsigned 64-bit integer with a minimum value
// of 0 and a maximum value of 264-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type UnsignedLongField struct {
	BaseField            `bson:",inline" json:",inline"`
	coerceParam          `bson:",inline" json:",inline"`
	docValuesParam       `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	indexParam           `bson:",inline" json:",inline"`
	nullValueParam       `bson:",inline" json:",inline"`
	storeParam           `bson:",inline" json:",inline"`
	MetaParam            `bson:",inline" json:",inline"`
}

func (f UnsignedLongField) Clone() Field {
	n := NewUnsignedLongField()
	n.SetCoerce(f.Coerce())
	n.SetDocValues(f.DocValues())
	n.SetIgnoreMalformed(f.IgnoreMalformed())
	n.SetIndex(f.Index())
	n.SetMeta(f.Meta().Clone())
	n.SetStore(f.Store())
	return n
}
func NewUnsignedLongField() *UnsignedLongField {
	return &UnsignedLongField{BaseField: BaseField{MappingType: FieldTypeUnsignedLong}}
}
