package picker

import "encoding/json"

type RangeField interface {
	Field
	WithCoerce
	WithIndex
	WithStore
}

//easyjson:json
type numericRangeField struct {
	Coerce interface{} `json:"coerce,omitempty"`
	Index  interface{} `json:"index,omitempty"`
	Store  interface{} `json:"store,omitempty"`
	Type   FieldType   `json:"type"`
}

type rangeFieldParams struct {
	// Coercion attempts to clean up dirty values to fit the data type of a
	// field. (Optional, bool) Defaults to false.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
	Coerce interface{} `json:"coerce,omitempty"`
	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool)
	Index interface{} `json:"index,omitempty"`
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
}

// A RangeField type represent a continuous range of values between an upper and
// lower bound. For example, a range can represent any date in October or any
// integer from 0 to 9. They are defined using the operators gt or gte for the
// lower bound, and lt or lte for the upper bound. They can be used for
// querying, and have limited support for aggregations. The only supported
// aggregations are histogram, cardinality.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html

type IntegerRangeFieldParams rangeFieldParams

func (p IntegerRangeFieldParams) Field() (Field, error) {
	return p.IntegerRange()
}
func (IntegerRangeFieldParams) Type() FieldType {
	return FieldTypeIntegerRange
}

func (p IntegerRangeFieldParams) IntegerRange() (*IntegerRangeField, error) {
	f := &IntegerRangeField{}
	e := &MappingError{}
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewIntegerRangeField(params IntegerRangeFieldParams) (*IntegerRangeField, error) {
	return params.IntegerRange()
}

type IntegerRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (r *IntegerRangeField) Field() (Field, error) {
	return r, nil
}
func (IntegerRangeField) Type() FieldType {
	return FieldTypeIntegerRange
}

func (r IntegerRangeField) MarshalJSON() ([]byte, error) {
	return numericRangeField{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
		Type:   r.Type(),
	}.MarshalJSON()
}

func (r *IntegerRangeField) UnmarshalJSON(data []byte) error {
	*r = IntegerRangeField{}
	p := IntegerRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.IntegerRange()
	*r = *v
	return err
}

type FloatRangeFieldParams rangeFieldParams

func (FloatRangeFieldParams) Type() FieldType {
	return FieldTypeFloatRange
}

func (p FloatRangeFieldParams) Field() (Field, error) {
	return p.FloatRange()
}
func (p FloatRangeFieldParams) FloatRange() (*FloatRangeField, error) {
	f := &FloatRangeField{}
	e := &MappingError{}
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	return f, nil
}

func NewFloatRangeField(params FloatRangeFieldParams) (*FloatRangeField, error) {
	return params.FloatRange()
}

// FloatRangeField is a range of single-precision 32-bit IEEE 754 floating point
// values.
type FloatRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (r *FloatRangeField) Field() (Field, error) {
	return r, nil
}
func (FloatRangeField) Type() FieldType {
	return FieldTypeFloatRange
}

func (r FloatRangeField) MarshalJSON() ([]byte, error) {
	return numericRangeField{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
		Type:   r.Type(),
	}.MarshalJSON()
}

func (r *FloatRangeField) UnmarshalJSON(data []byte) error {
	*r = FloatRangeField{}
	p := FloatRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.FloatRange()
	*r = *v
	return err
}

type LongRangeFieldParams rangeFieldParams

func (LongRangeFieldParams) Type() FieldType {
	return FieldTypeLongRange
}

func (p LongRangeFieldParams) Field() (Field, error) {
	return p.LongRange()
}
func (p LongRangeFieldParams) LongRange() (*LongRangeField, error) {
	f := &LongRangeField{}
	e := &MappingError{}
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}

	return f, e.ErrorOrNil()
}

func NewLongRangeField(params LongRangeFieldParams) (*LongRangeField, error) {
	return params.LongRange()
}

// LongRangeField is a range of signed 64-bit integers with a minimum value of
// -263 and maximum of 263-1.
type LongRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (r *LongRangeField) Field() (Field, error) {
	return r, nil
}
func (LongRangeField) Type() FieldType {
	return FieldTypeLongRange
}

func (r LongRangeField) MarshalJSON() ([]byte, error) {
	return numericRangeField{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
		Type:   r.Type(),
	}.MarshalJSON()
}

func (r *LongRangeField) UnmarshalJSON(data []byte) error {
	*r = LongRangeField{}
	p := LongRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.LongRange()
	*r = *v
	return err
}

type DoubleRangeFieldParams rangeFieldParams

func (DoubleRangeFieldParams) Type() FieldType {
	return FieldTypeDoubleRange
}
func (p DoubleRangeFieldParams) Field() (Field, error) {
	return p.DoubleRange()
}
func (p DoubleRangeFieldParams) DoubleRange() (*DoubleRangeField, error) {
	f := &DoubleRangeField{}
	e := &MappingError{}
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewDoubleRangeField(params DoubleRangeFieldParams) (*DoubleRangeField, error) {
	return params.DoubleRange()
}

// DoubleRangeField is a range of double-precision 64-bit IEEE 754 floating
// point values.
type DoubleRangeField struct {
	coerceParam
	indexParam
	storeParam
}

func (r *DoubleRangeField) Field() (Field, error) {
	return r, nil
}

func (DoubleRangeField) Type() FieldType {
	return FieldTypeDoubleRange
}

func (r DoubleRangeField) MarshalJSON() ([]byte, error) {
	return numericRangeField{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
		Type:   r.Type(),
	}.MarshalJSON()
}

func (r *DoubleRangeField) UnmarshalJSON(data []byte) error {
	*r = DoubleRangeField{}
	p := DoubleRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.DoubleRange()
	*r = *v
	return err
}

type DateRangeFieldParams struct {
	// Coercion attempts to clean up dirty values to fit the data type of a
	// field. (Optional, bool) Defaults to false.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/coerce.html
	Coerce interface{} `json:"coerce,omitempty"`
	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool)
	Index interface{} `json:"index,omitempty"`
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
	Store  interface{} `json:"store,omitempty"`
	Format string      `json:"format,omitempty"`
}

//easyjson:json
type dateRangeField struct {
	Coerce interface{} `json:"coerce,omitempty"`
	Index  interface{} `json:"index,omitempty"`
	Store  interface{} `json:"store,omitempty"`
	Format string      `json:"format,omitempty"`
	Type   FieldType   `json:"type"`
}

func (DateRangeFieldParams) Type() FieldType {
	return FieldTypeDateRange
}

func (p DateRangeFieldParams) Field() (Field, error) {
	return p.DateRange()
}
func (p DateRangeFieldParams) DateRange() (*DateRangeField, error) {
	f := &DateRangeField{}
	e := &MappingError{}
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	f.SetFormat(p.Format)
	return f, e.ErrorOrNil()
}

func NewDateRangeField(params DateRangeFieldParams) (*DateRangeField, error) {
	return params.DateRange()
}

// DateRangeField is a range of date values. Date ranges support various date
// formats through the format mapping parameter. Regardless of the format used,
// date values are parsed into an unsigned 64-bit integer representing
// milliseconds since the Unix epoch in UTC. Values containing the now date math
// expression are not supported.
type DateRangeField struct {
	coerceParam
	indexParam
	storeParam
	formatParam
}

func (r *DateRangeField) Field() (Field, error) {
	return r, nil
}
func (DateRangeField) Type() FieldType {
	return FieldTypeDateRange
}

func (r DateRangeField) MarshalJSON() ([]byte, error) {
	return dateRangeField{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
		Format: r.format,
		Type:   r.Type(),
	}.MarshalJSON()
}

func (r *DateRangeField) UnmarshalJSON(data []byte) error {
	*r = DateRangeField{}
	p := DateRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.DateRange()
	*r = *v
	return err

}

type IPRangeFieldParams rangeFieldParams

func (IPRangeFieldParams) Type() FieldType {
	return FieldTypeIPRange
}
func (p IPRangeFieldParams) Field() (Field, error) {
	return p.IPRange()
}
func (p IPRangeFieldParams) IPRange() (*IPRangeField, error) {
	f := &IPRangeField{}
	e := &MappingError{}
	err := f.SetCoerce(p.Coerce)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewIPRangeField(params IPRangeFieldParams) (*IPRangeField, error) {
	return params.IPRange()
}

// IPRangeField is a range of ip values supporting either IPv4 or IPv6 (or
// mixed) addresses.
type IPRangeField struct {
	coerceParam
	indexParam
	storeParam
}

//easyjson:json
type ipRangeField struct {
	Coerce interface{} `json:"coerce,omitempty"`
	Index  interface{} `json:"index,omitempty"`
	Store  interface{} `json:"store,omitempty"`
	Type   FieldType   `json:"type"`
}

func (r *IPRangeField) Field() (Field, error) {
	return r, nil
}
func (IPRangeField) Type() FieldType {
	return FieldTypeIPRange
}

func (r IPRangeField) MarshalJSON() ([]byte, error) {
	return ipRangeField{
		Coerce: r.coerce.Value(),
		Index:  r.index.Value(),
		Store:  r.store.Value(),
		Type:   r.Type(),
	}.MarshalJSON()
}

func (r *IPRangeField) UnmarshalJSON(data []byte) error {
	*r = IPRangeField{}
	p := IPRangeFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	v, err := p.IPRange()
	*r = *v
	return err

}
