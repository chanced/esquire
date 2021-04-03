package picker

import "encoding/json"

type dateField struct {
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	DocValues       interface{} `json:"doc_values,omitempty"`
	Index           interface{} `json:"index,omitempty"`
	NullValue       interface{} `json:"null_value,omitempty"`
	Store           interface{} `json:"store,omitempty"`
	Meta            Meta        `json:"meta,omitempty"`
	Format          string      `json:"format,omitempty"`
	Boost           interface{} `json:"boost,omitempty"`
	Type            FieldType   `json:"type"`
}

type DateFieldParams struct {
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
	// By default, field values are indexed to make them searchable, but they
	// are not stored. This means that the field can be queried, but the
	// original field value cannot be retrieved.
	//
	// Usually this doesn’t matter. The field value is already part of the
	// _source field, which is stored by default. If you only want to retrieve
	// the value of a single field or of a few fields, instead of the whole
	// _source, then this can be achieved with source filtering.
	//
	// In certain situations it can make sense to store a field. For instance,
	// if you have a document with a title, a date, and a very large content
	// field, you may want to retrieve just the title and the date without
	// having to extract those fields from a large _source field
	//
	// Stored fields returned as arrays
	//
	// For consistency, stored fields are always returned as an array because
	// there is no way of knowing if the original field value was a single
	// value, multiple values, or an empty array.
	//
	// The original value can be retrieved from the _source field instead.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-store.html
	Store interface{} `json:"store,omitempty"`
	// Metadata attached to the field. This metadata is opaque to Elasticsearch,
	// it is only useful for multiple applications that work on the same indices
	// to share meta information about fields such as units
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
	Meta Meta `json:"meta,omitempty"`
	//Format is the format(d) that the that can be parsed. Defaults to
	//strict_date_optional_time||epoch_millis.
	//
	// Multiple formats can be seperated by ||
	Format string `json:"format,omitempty"`

	// Deprecated
	Boost interface{} `json:"boost,omitempty"`
}

func (DateFieldParams) Type() FieldType {
	return FieldTypeDate
}

func (p DateFieldParams) Field() (Field, error) {
	return p.Date()
}

func (p DateFieldParams) Date() (*DateField, error) {
	f := &DateField{}
	e := &MappingError{}
	err := f.SetDocValues(p.DocValues)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIgnoreMalformed(p.IgnoreMalformed)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetMeta(p.Meta)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	err = f.SetBoost(p.Boost)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewDateField(params DateFieldParams) (*DateField, error) {
	return params.Date()
}

type DateField struct {
	docValuesParam
	formatParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (d *DateField) Field() (Field, error) {
	return d, nil
}
func (DateField) Type() FieldType {
	return FieldTypeDate
}

func (d *DateField) UnmarshalJSON(data []byte) error {

	var params DateFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Date()
	*d = *v
	return err
}

func (d DateField) MarshalJSON() ([]byte, error) {
	return json.Marshal(dateField{
		Format:          d.format,
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

type DateNanoSecFieldParams DateFieldParams

func (p DateNanoSecFieldParams) Type() FieldType {
	return FieldTypeDateNanos
}
func (p DateNanoSecFieldParams) Field() (Field, error) {
	return p.DateNanoSec()
}
func (p DateNanoSecFieldParams) DateNanoSec() (*DateNanoSecField, error) {
	f := &DateNanoSecField{}
	e := &MappingError{}
	err := f.SetDocValues(p.DocValues)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIgnoreMalformed(p.IgnoreMalformed)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetMeta(p.Meta)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewDateNanoSecField(params DateNanoSecFieldParams) (*DateNanoSecField, error) {
	return params.DateNanoSec()
}

// DateNanoSecField is an addition to the DateField data type.
//
// However there is an important distinction between the two. The existing date
// data type stores dates in millisecond resolution. The date_nanos data type
// stores dates in nanosecond resolution, which limits its range of dates from
// roughly 1970 to 2262, as dates are still stored as a long representing
// nanoseconds since the epoch.
//
// Queries on nanoseconds are internally converted to range queries on this long
// representation, and the result of aggregations and stored fields is converted
// back to a string depending on the date format that is associated with the
// field.
//
// Date formats can be customised, but if no format is specified then it uses
// the default:
//
//  "strict_date_optional_time||epoch_millis"
//
// This means that it will accept dates with optional timestamps, which conform
// to the formats supported by strict_date_optional_time including up to nine
// second fractionals or milliseconds-since-the-epoch (thus losing precision on
// the nano second part). Using strict_date_optional_time will format the result
// up to only three second fractionals. To print and parse up to nine digits of
// resolution, use strict_date_optional_time_nanos.
//
// Limitations
//
// Aggregations are still on millisecond resolution, even when using a
// date_nanos field. This limitation also affects transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/date_nanos.html
type DateNanoSecField struct {
	docValuesParam
	formatParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	metaParam
	boostParam
}

func (d DateNanoSecField) Type() FieldType {
	return FieldTypeDateNanos
}
func (d *DateNanoSecField) Field() (Field, error) {
	return d, nil
}
func (d *DateNanoSecField) UnmarshalJSON(data []byte) error {

	var params DateNanoSecFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.DateNanoSec()
	*d = *v
	return err
}

func (d DateNanoSecField) MarshalJSON() ([]byte, error) {
	return json.Marshal(dateField{
		Format:          d.format,
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

var (
	_ WithDocValues       = (*DateField)(nil)
	_ WithFormat          = (*DateField)(nil)
	_ WithIgnoreMalformed = (*DateField)(nil)
	_ WithIndex           = (*DateField)(nil)
	_ WithNullValue       = (*DateField)(nil)
	_ WithStore           = (*DateField)(nil)
	_ WithMeta            = (*DateField)(nil)
)
