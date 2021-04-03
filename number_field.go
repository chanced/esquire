package picker

import "encoding/json"

// TODO: this needs refactoring. I initially had numberFieldParams produce all types but converted it to the more copypasta approach.
// however, having to have a seperate numberField kiiiiinda sucks

type NumberField interface {
	Field
	WithCoerce
	WithIgnoreMalformed
	WithDocValues
	WithNullValue
	WithMeta
	WithIndex
	WithStore
}

type numberField struct {
	Coerce          interface{} `json:"coerce,omitempty"`
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	DocValues       interface{} `json:"doc_values,omitempty"`
	Index           interface{} `json:"index,omitempty"`
	NullValue       interface{} `json:"null_value,omitempty"`
	Store           interface{} `json:"store,omitempty"`
	Meta            Meta        `json:"meta,omitempty"`
	Boost           interface{} `json:"boost,omitempty"`
	Type            FieldType   `json:"type"`
}

type numberFieldParams struct {
	// Coercion attempts to clean up dirty values to fit the data type of a
	// field. (Optional, bool) Defaults to false.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
	Coerce interface{} `json:"coerce,omitempty"`
	// IgnoreMalformed determines if malformed numbers are ignored. If true,
	// malformed numbers are ignored. If false (default), malformed numbers
	// throw an exception and reject the whole document.
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	// (Optional, bool, default: true)
	//
	// Most fields are indexed by default, which makes them searchable. The
	// inverted index allows queries to look up the search term in unique sorted
	// list of terms, and from that immediately have access to the list of
	// documents that contain the term.
	//
	// Sorting, aggregations, and access to field values in scripts requires a
	// different data access pattern. Instead of looking up the term and finding
	// documents, we need to be able to look up the document and find the terms
	// that it has in a field.
	//
	// Doc values are the on-disk data structure, built at document index time,
	// which makes this data access pattern possible. They store the same values
	// as the _source but in a column-oriented fashion that is way more
	// efficient for sorting and aggregations. Doc values are supported on
	// almost all field types, with the notable exception of text and
	// annotated_text fields.
	//
	// All fields which support doc values have them enabled by default. If you
	// are sure that you don’t need to sort or aggregate on a field, or access
	// the field value from a script, you can disable doc values in order to
	// save disk space
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/doc-values.html
	DocValues interface{} `json:"doc_values,omitempty"`
	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool)
	Index interface{} `json:"index,omitempty"`
	// NullValue parameter allows you to replace explicit null values with the
	// specified value so that it can be indexed and searched
	NullValue interface{} `json:"null_value,omitempty"`
	// WithStore is a mapping with a store paramter.
	//
	// By default, field values are indexed to make them searchable, but they are
	// not stored. This means that the field can be queried, but the original field
	// value cannot be retrieved.
	//
	// Usually this doesn’t matter. The field value is already part of the _source
	// field, which is stored by default. If you only want to retrieve the value of
	// a single field or of a few fields, instead of the whole _source, then this
	// can be achieved with source filtering.
	//
	// In certain situations it can make sense to store a field. For instance, if
	// you have a document with a title, a date, and a very large content field, you
	// may want to retrieve just the title and the date without having to extract
	// those fields from a large _source field
	//
	// Stored fields returned as arrays
	//
	// For consistency, stored fields are always returned as an array because there
	// is no way of knowing if the original field value was a single value, multiple
	// values, or an empty array.
	//
	// The original value can be retrieved from the _source field instead.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-store.html
	Store interface{} `json:"store,omitempty"`
	// Metadata attached to the field. This metadata is opaque to Elasticsearch, it
	// is only useful for multiple applications that work on the same indices to
	// share meta information about fields such as units
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
	Meta Meta `json:"meta,omitempty"`
	// Deprecated
	Boost interface{} `json:"boost,omitempty"`
}

func (p numberFieldParams) numberField(f NumberField) error {
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		return err
	}
	err = f.SetDocValues(p.DocValues)
	if err != nil {
		return err
	}
	err = f.SetIgnoreMalformed(p.IgnoreMalformed)
	if err != nil {
		return err
	}

	err = f.SetIndex(p.Index)
	if err != nil {
		return err
	}

	err = f.SetMeta(p.Meta)
	if err != nil {
		return err
	}
	f.SetNullValue(p.NullValue)
	err = f.SetStore(p.Store)
	if err != nil {
		return err
	}
	return nil
}

// LongFieldParams - field for a signed 64-bit integer with a minimum value of
// -263 and a maximum value of 263-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type LongFieldParams numberFieldParams

func (LongFieldParams) Type() FieldType {
	return FieldTypeLong
}

func (l LongFieldParams) Field() (Field, error) {
	return l.Long()
}

func (l LongFieldParams) Long() (*LongField, error) {
	f := &LongField{}
	err := numberFieldParams(l).numberField(f)
	return f, err
}

func NewLongField(params LongFieldParams) (*LongField, error) {
	return params.Long()
}

// A LongField is a signed 64-bit integer with a minimum value of -263
// and a maximum value of 263-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type LongField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (LongField) Type() FieldType {
	return FieldTypeLong
}
func (l *LongField) Field() (Field, error) {
	return l, nil
}
func (l *LongField) UnmarshalJSON(data []byte) error {

	var params LongFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Long()
	if err != nil {
		return err
	}
	*l = *v
	return nil
}

func (l LongField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          l.coerce.Value(),
		IgnoreMalformed: l.ignoreMalformed.Value(),
		DocValues:       l.docValues.Value(),
		Index:           l.index.Value(),
		NullValue:       l.nullValue,
		Store:           l.store.Value(),
		Meta:            l.meta,
		Boost:           l.boost.Value(),
		Type:            l.Type(),
	})
}

type IntegerFieldParams numberFieldParams

func (IntegerFieldParams) Type() FieldType {
	return FieldTypeInteger
}

func (i IntegerFieldParams) Field() (Field, error) {
	return i.Integer()
}
func (p IntegerFieldParams) Integer() (*IntegerField, error) {
	f := &IntegerField{}
	err := numberFieldParams(p).numberField(f)
	return f, err
}

func NewIntegerField(params IntegerFieldParams) (*IntegerField, error) {
	return params.Integer()
}

// An IntegerField is a signed 64-bit integer with a minimum value of -263
// and a maximum value of 263-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type IntegerField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (l *IntegerField) Field() (Field, error) {
	return l, nil
}
func (IntegerField) Type() FieldType {
	return FieldTypeInteger
}

func (i *IntegerField) UnmarshalJSON(data []byte) error {

	var params IntegerFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Integer()
	if err != nil {
		return err
	}
	*i = *v
	return nil
}

func (i IntegerField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          i.coerce.Value(),
		IgnoreMalformed: i.ignoreMalformed.Value(),
		DocValues:       i.docValues.Value(),
		Index:           i.index.Value(),
		NullValue:       i.nullValue,
		Store:           i.store.Value(),
		Meta:            i.meta,
		Boost:           i.boost.Value(),
		Type:            i.Type(),
	})
}

type ShortFieldParams numberFieldParams

func (ShortFieldParams) Type() FieldType {
	return FieldTypeShort
}

func (s ShortFieldParams) Field() (Field, error) {
	return s.Short()
}

func (p ShortFieldParams) Short() (*ShortField, error) {
	f := &ShortField{}
	err := numberFieldParams(p).numberField(f)
	return f, err
}
func NewShortField(params ShortFieldParams) (*ShortField, error) {
	return params.Short()
}

// A ShortField is signed 16-bit integer with a minimum value of -32,768
// and a maximum value of 32,767.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type ShortField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (ShortField) Type() FieldType {
	return FieldTypeShort
}

func (s *ShortField) Field() (Field, error) {
	return s, nil
}
func (s *ShortField) UnmarshalJSON(data []byte) error {

	var params ShortFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Short()
	if err != nil {
		return err
	}
	*s = *v
	return nil
}

func (s ShortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          s.coerce.Value(),
		IgnoreMalformed: s.ignoreMalformed.Value(),
		DocValues:       s.docValues.Value(),
		Index:           s.index.Value(),
		NullValue:       s.nullValue,
		Store:           s.store.Value(),
		Meta:            s.meta,
		Boost:           s.boost.Value(),
		Type:            s.Type(),
	})
}

type DoubleFieldParams numberFieldParams

func (DoubleFieldParams) Type() FieldType {
	return FieldTypeDouble
}

func (l DoubleFieldParams) Field() (Field, error) {
	return l.Double()
}
func (p DoubleFieldParams) Double() (*DoubleField, error) {
	f := &DoubleField{}
	err := numberFieldParams(p).numberField(f)
	return f, err
}

func NewDoubleField(params DoubleFieldParams) (*DoubleField, error) {
	return params.Double()
}

// A DoubleField is a double-precision 64-bit IEEE 754 floating point number, restricted to finite values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type DoubleField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (d *DoubleField) Field() (Field, error) {
	return d, nil
}
func (DoubleField) Type() FieldType {
	return FieldTypeDouble
}

func (d *DoubleField) UnmarshalJSON(data []byte) error {

	var params DoubleFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Double()
	if err != nil {
		return err
	}
	*d = *v
	return nil
}

func (d DoubleField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          d.coerce.Value(),
		IgnoreMalformed: d.ignoreMalformed.Value(),
		DocValues:       d.docValues.Value(),
		Index:           d.index.Value(),
		NullValue:       d.nullValue,
		Store:           d.store.Value(),
		Meta:            d.meta,
		Boost:           d.boost.Value(),
		Type:            d.Type(),
	})
}

type ByteFieldParams numberFieldParams

func (ByteFieldParams) Type() FieldType {
	return FieldTypeByte
}

func (b ByteFieldParams) Field() (Field, error) {
	return b.Byte()
}
func (p ByteFieldParams) Byte() (*ByteField, error) {
	f := &ByteField{}
	err := numberFieldParams(p).numberField(f)
	return f, err
}

func NewByteField(params ByteFieldParams) (*ByteField, error) {
	return params.Byte()
}

// A ByteField is a signed 8-bit integer with a minimum value of -128 and a maximum value of 127.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type ByteField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (b *ByteField) Field() (Field, error) {
	return b, nil
}

func (ByteField) Type() FieldType {
	return FieldTypeByte
}

func (b *ByteField) UnmarshalJSON(data []byte) error {

	var params ByteFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Byte()
	if err != nil {
		return err
	}
	*b = *v
	return nil
}

func (b ByteField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          b.coerce.Value(),
		IgnoreMalformed: b.ignoreMalformed.Value(),
		DocValues:       b.docValues.Value(),
		Index:           b.index.Value(),
		NullValue:       b.nullValue,
		Store:           b.store.Value(),
		Meta:            b.meta,
		Boost:           b.boost.Value(),
		Type:            b.Type(),
	})
}

type FloatFieldParams numberFieldParams

func (FloatFieldParams) Type() FieldType {
	return FieldTypeFloat
}

func (p FloatFieldParams) Field() (Field, error) {
	return p.Float()
}

func (p FloatFieldParams) Float() (*FloatField, error) {
	f := &FloatField{}
	err := numberFieldParams(p).numberField(f)
	return f, err
}

func NewFloatField(params FloatFieldParams) (*FloatField, error) {
	return params.Float()
}

// A FloatField is a single-precision 32-bit IEEE 754 floating point
// number, restricted to finite values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type FloatField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (FloatField) Type() FieldType {
	return FieldTypeFloat
}
func (f *FloatField) Field() (Field, error) {
	return f, nil
}
func (f *FloatField) UnmarshalJSON(data []byte) error {

	var params FloatFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Float()
	if err != nil {
		return err
	}
	*f = *v
	return nil
}

func (f FloatField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          f.coerce.Value(),
		IgnoreMalformed: f.ignoreMalformed.Value(),
		DocValues:       f.docValues.Value(),
		Index:           f.index.Value(),
		NullValue:       f.nullValue,
		Store:           f.store.Value(),
		Meta:            f.meta,
		Boost:           f.boost.Value(),
		Type:            f.Type(),
	})
}

type HalfFloatFieldParams numberFieldParams

func (HalfFloatFieldParams) Type() FieldType {
	return FieldTypeHalfFloat
}

func (l HalfFloatFieldParams) Field() (Field, error) {
	return l.HalfFloat()
}

func (l HalfFloatFieldParams) HalfFloat() (*HalfFloatField, error) {
	f := &HalfFloatField{}
	err := numberFieldParams(l).numberField(f)
	return f, err
}

func NewHalfFloatField(params HalfFloatFieldParams) (*HalfFloatField, error) {
	return params.HalfFloat()
}

// A HalfFloatField is a half-precision 16-bit IEEE 754 floating point
// number, restricted to finite values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type HalfFloatField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (f *HalfFloatField) Field() (Field, error) {
	return f, nil
}
func (HalfFloatField) Type() FieldType {
	return FieldTypeHalfFloat
}

func (hf *HalfFloatField) UnmarshalJSON(data []byte) error {

	var params HalfFloatFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.HalfFloat()
	if err != nil {
		return err
	}
	*hf = *v
	return nil
}

func (hf HalfFloatField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          hf.coerce.Value(),
		IgnoreMalformed: hf.ignoreMalformed.Value(),
		DocValues:       hf.docValues.Value(),
		Index:           hf.index.Value(),
		NullValue:       hf.nullValue,
		Store:           hf.store.Value(),
		Meta:            hf.meta,
		Boost:           hf.boost.Value(),
		Type:            hf.Type(),
	})
}

/// UnsignedLongFieldParams are params for an UnsignedLongField which is an unsigned 64-bit integer with a minimum value
// of 0 and a maximum value of 264-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type UnsignedLongFieldParams numberFieldParams

func (UnsignedLongFieldParams) Type() FieldType {
	return FieldTypeUnsignedLong
}

func (l UnsignedLongFieldParams) Field() (Field, error) {
	return l.UnsignedLong()
}

func (l UnsignedLongFieldParams) UnsignedLong() (*UnsignedLongField, error) {
	f := &UnsignedLongField{}
	err := numberFieldParams(l).numberField(f)
	// TODO: wrap error in a FieldError
	return f, err
}

func NewUnsignedLongField(params UnsignedLongFieldParams) (*UnsignedLongField, error) {
	return params.UnsignedLong()
}

// An UnsignedLongField is an unsigned 64-bit integer with a minimum value
// of 0 and a maximum value of 264-1.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
type UnsignedLongField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (ul *UnsignedLongField) Field() (Field, error) {
	return ul, nil
}
func (UnsignedLongField) Type() FieldType {
	return FieldTypeUnsignedLong
}

func (ul *UnsignedLongField) UnmarshalJSON(data []byte) error {

	var params UnsignedLongFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.UnsignedLong()
	if err != nil {
		return err
	}
	*ul = *v
	return nil
}

func (ul UnsignedLongField) MarshalJSON() ([]byte, error) {
	return json.Marshal(numberField{
		Coerce:          ul.coerce.Value(),
		IgnoreMalformed: ul.ignoreMalformed.Value(),
		DocValues:       ul.docValues.Value(),
		Index:           ul.index.Value(),
		NullValue:       ul.nullValue,
		Store:           ul.store.Value(),
		Meta:            ul.meta,
		Boost:           ul.boost.Value(),
		Type:            ul.Type(),
	})
}

type scaledFloatField struct {
	Coerce          interface{} `json:"coerce,omitempty"`
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	DocValues       interface{} `json:"doc_values,omitempty"`
	Index           interface{} `json:"index,omitempty"`
	NullValue       interface{} `json:"null_value,omitempty"`
	Store           interface{} `json:"store,omitempty"`
	Meta            Meta        `json:"meta,omitempty"`
	Boost           interface{} `json:"boost,omitempty"`
	ScalingFactor   interface{} `scaling_factor`
	Type            FieldType   `json:"type"`
}

type ScaledFloatFieldParams struct {

	// The scaling factor to use when encoding values. Values will be multiplied by
	// this factor at index time and rounded to the closest long value. For
	// instance, a scaled_float with a scaling_factor of 10 would internally store
	// 2.34 as 23 and all search-time operations (queries, aggregations, sorting)
	// will behave as if the document had a value of 2.3. High values of
	// scaling_factor improve accuracy but also increase space requirements. (Required)
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html#scaled-float-params

	// A ScaledFloatField is a floating point number that is backed by a long,
	// scaled by a fixed double scaling factor.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	ScalingFactor interface{} `scaling_factor`
	// Coercion attempts to clean up dirty values to fit the data type of a
	// field. (Optional, bool) Defaults to false.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
	Coerce interface{} `json:"coerce,omitempty"`
	// IgnoreMalformed determines if malformed numbers are ignored. If true,
	// malformed numbers are ignored. If false (default), malformed numbers
	// throw an exception and reject the whole document.
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	// (Optional, bool, default: true)
	//
	// Most fields are indexed by default, which makes them searchable. The
	// inverted index allows queries to look up the search term in unique sorted
	// list of terms, and from that immediately have access to the list of
	// documents that contain the term.
	//
	// Sorting, aggregations, and access to field values in scripts requires a
	// different data access pattern. Instead of looking up the term and finding
	// documents, we need to be able to look up the document and find the terms
	// that it has in a field.
	//
	// Doc values are the on-disk data structure, built at document index time,
	// which makes this data access pattern possible. They store the same values
	// as the _source but in a column-oriented fashion that is way more
	// efficient for sorting and aggregations. Doc values are supported on
	// almost all field types, with the notable exception of text and
	// annotated_text fields.
	//
	// All fields which support doc values have them enabled by default. If you
	// are sure that you don’t need to sort or aggregate on a field, or access
	// the field value from a script, you can disable doc values in order to
	// save disk space
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/doc-values.html
	DocValues interface{} `json:"doc_values,omitempty"`
	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool)
	Index interface{} `json:"index,omitempty"`
	// NullValue parameter allows you to replace explicit null values with the
	// specified value so that it can be indexed and searched
	NullValue interface{} `json:"null_value,omitempty"`
	// WithStore is a mapping with a store paramter.
	//
	// By default, field values are indexed to make them searchable, but they are
	// not stored. This means that the field can be queried, but the original field
	// value cannot be retrieved.
	//
	// Usually this doesn’t matter. The field value is already part of the _source
	// field, which is stored by default. If you only want to retrieve the value of
	// a single field or of a few fields, instead of the whole _source, then this
	// can be achieved with source filtering.
	//
	// In certain situations it can make sense to store a field. For instance, if
	// you have a document with a title, a date, and a very large content field, you
	// may want to retrieve just the title and the date without having to extract
	// those fields from a large _source field
	//
	// Stored fields returned as arrays
	//
	// For consistency, stored fields are always returned as an array because there
	// is no way of knowing if the original field value was a single value, multiple
	// values, or an empty array.
	//
	// The original value can be retrieved from the _source field instead.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-store.html
	Store interface{} `json:"store,omitempty"`
	// Metadata attached to the field. This metadata is opaque to Elasticsearch, it
	// is only useful for multiple applications that work on the same indices to
	// share meta information about fields such as units
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
	Meta Meta `json:"meta,omitempty"`

	// Deprecated
	Boost interface{} `json:"boost,omitempty"`
	// Ignore this field
	FieldType FieldType `json:"type"`
}

func (ScaledFloatFieldParams) Type() FieldType {
	return FieldTypeScaledFloat
}

func (p ScaledFloatFieldParams) Field() (Field, error) {
	return p.ScaledFloat()
}

func (p ScaledFloatFieldParams) ScaledFloat() (*ScaledFloatField, error) {
	f := &ScaledFloatField{}
	err := numberFieldParams{
		Coerce:          p.Coerce,
		IgnoreMalformed: p.IgnoreMalformed,
		DocValues:       p.DocValues,
		Index:           p.Index,
		NullValue:       p.NullValue,
		Store:           p.Store,
		Meta:            p.Meta,
		Boost:           p.Boost,
	}.numberField(f)
	if err != nil {
		return f, err
	}
	err = f.SetScalingFactor(p.ScalingFactor)
	if err != nil {
		return f, err
	}
	return f, nil
}

func NewScaledFloatField(params ScaledFloatFieldParams) (*ScaledFloatField, error) {
	return params.ScaledFloat()
}

type ScaledFloatField struct {
	coerceParam
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	scalingFactorParam
	boostParam
}

func (sf *ScaledFloatField) Field() (Field, error) {
	return sf, nil
}
func (ScaledFloatField) Type() FieldType {
	return FieldTypeScaledFloat
}

func (sf *ScaledFloatField) UnmarshalJSON(data []byte) error {

	var params ScaledFloatFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.ScaledFloat()
	if err != nil {
		return err
	}
	*sf = *v
	return nil
}

func (sf ScaledFloatField) MarshalJSON() ([]byte, error) {
	return json.Marshal(scaledFloatField{
		ScalingFactor:   sf.scalingFactor,
		Coerce:          sf.coerce.Value(),
		IgnoreMalformed: sf.ignoreMalformed.Value(),
		DocValues:       sf.docValues.Value(),
		Index:           sf.index.Value(),
		NullValue:       sf.nullValue,
		Store:           sf.store.Value(),
		Meta:            sf.meta,
		Boost:           sf.boost.Value(),
		Type:            sf.Type(),
	})
}
