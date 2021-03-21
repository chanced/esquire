package mapping

// WithDocValues is a mapping with a DocValues parameter
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
type WithDocValues interface {
	// SetDocValues sets Value to v
	SetDocValues(v bool)
	// DocValues returns DocValues, defaulted to true
	DocValues() bool
}

// FieldWithDocValues is a Field with a DocValues parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/doc-values.html
type FieldWithDocValues interface {
	Field
}

// DocValuesParam is a mixin for Mappings that have doc_values
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
type DocValuesParam struct {
	DocValuesValue *bool `json:"doc_values,omitempty" bson:"doc_values,omitempty"`
}

// SetDocValues sets Value to v
func (dv *DocValuesParam) SetDocValues(v bool) {
	if dv.DocValues() != v {
		dv.DocValuesValue = &v
	}
}

// DocValues returns DocValues, defaulted to true
func (dv DocValuesParam) DocValues() bool {
	if dv.DocValuesValue != nil {
		return *dv.DocValuesValue
	}
	return true
}
