package picker

type IPFieldParams struct {
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
}

// An IPField can index/store either IPv4 or IPv6 addresses.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ip.html
type IPField struct {
	docValuesParam
	ignoreMalformedParam
	indexParam
	nullValueParam
	storeParam
	boostParam
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
	return &IPField{BaseField: BaseField{MappingType: FieldTypeIP}}
}
