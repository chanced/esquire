package picker

import "encoding/json"

type BooleanFieldParams struct {

	// DocValues sets doc_values (Optional, bool or string that can be parsed as a bool)
	//
	// Most fields are indexed by default, which makes them searchable. The inverted
	// index allows queries to look up the search term in unique sorted list of
	// terms, and from that immediately have access to the list of documents that
	// contain the term.
	//
	// Sorting, aggregations, and access to field values in scripts requires a
	// different data access pattern. Instead of looking up the term and finding
	// documents, we need to be able to look up the document and find the terms that
	// it has in a field.
	//
	// Doc values are the on-disk data structure, built at document index time,
	// which makes this data access pattern possible. They store the same values as
	// the _source but in a column-oriented fashion that is way more efficient for
	// sorting and aggregations. Doc values are supported on almost all field types,
	// with the notable exception of text and annotated_text fields.
	//
	// All fields which support doc values have them enabled by default. If you are
	// sure that you don’t need to sort or aggregate on a field, or access the field
	// value from a script, you can disable doc values in order to save disk space
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/doc-values.html
	DocValues interface{} `json:"doc_values,omitempty" bson:"doc_values,omitempty"`

	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool or string that can be parsed as a bool)
	Index interface{} `bson:"index,omitempty" json:"index,omitempty"`
	// A null value cannot be indexed or searched. When a field is set to null, (or
	// an empty array or an array of null values) it is treated as though that field
	// has no values.
	//
	// The null_value parameter allows you to replace explicit null values with the
	// specified value so that it can be indexed and searched
	//
	// The null_value needs to be the same data type as the field. For instance, a
	// long field cannot have a string null_value.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/null-value.html
	NullValue interface{} `json:"null_value,omitempty"`

	// By default, field values are indexed to make them searchable, but they
	// are not stored. This means that the field can be queried, but the
	// original field value cannot be retrieved.
	//
	// (Optional, bool or string that can be parsed as a bool)
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

	// Metadata attached to the field. This metadata is opaque to Elasticsearch, it
	// is only useful for multiple applications that work on the same indices to
	// share meta information about fields such as units
	//
	// map[string]string or picker.Meta
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
	Meta map[string]string `json:"meta,omitempty"`
}

func (BooleanFieldParams) Type() FieldType {
	return FieldTypeBoolean
}
func (b BooleanFieldParams) Field() (Field, error) {
	return b.Boolean()
}

func (b BooleanFieldParams) Boolean() (*BooleanField, error) {
	f := &BooleanField{}
	err := f.SetDocValues(b.DocValues)
	if err != nil {
		return f, err
	}
	err = f.SetIndex(b.Index)
	if err != nil {
		return f, err
	}
	err = f.SetMeta(b.Meta)
	if err != nil {
		return f, err
	}
	err = f.SetStore(b.Store)
	if err != nil {
		return f, err
	}
	f.SetNullValue(b.NullValue)
	return f, nil
}

func NewBooleanField() *BooleanField {

	return &BooleanField{}
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
	docValuesParam `bson:",inline" json:",inline"`
	indexParam     `bson:",inline" json:",inline"`
	nullValueParam `bson:",inline" json:",inline"`
	storeParam     `bson:",inline" json:",inline"`
	metaParam      `bson:",inline" json:",inline"`
}

func (BooleanField) Type() FieldType {
	return FieldTypeBoolean
}

func (b *BooleanField) UnmarshalJSON(data []byte) error {

	var params BooleanFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Boolean()
	if err != nil {
		return err
	}
	*b = *v
	return nil
}

func (b BooleanField) MarshalJSON() ([]byte, error) {
	return json.Marshal(BooleanFieldParams{
		DocValues: b.docValues.Value(),
		Index:     b.index.Value(),
		NullValue: b.nullValue,
		Store:     b.store.Value(),
		Meta:      b.meta,
	})
}
