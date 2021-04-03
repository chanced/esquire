package picker

// SearchAsYouTypeField is a text-like field that is optimized to provide
// out-of-the-box support for queries that serve an as-you-type completion use
// case. It creates a series of subfields that are analyzed to index terms that
// can be efficiently matched by a query that partially matches the entire
// indexed text value. Both prefix completion (i.e matching terms starting at
// the beginning of the input) and infix completion (i.e. matching terms at any
// position within the input) are supported.
//
// The size of shingles in subfields can be configured with the max_shingle_size
// mapping parameter. The default is 3, and valid values for this parameter are
// integer values 2 - 4 inclusive. Shingle subfields will be created for each
// shingle size from 2 up to and including the max_shingle_size. The
// my_field._index_prefix subfield will always use the analyzer from the shingle
// subfield with the max_shingle_size when constructing its own analyzer.
//
// Increasing the max_shingle_size will improve matches for queries with more
// consecutive terms, at the cost of larger index size. The default
// max_shingle_size should usually be sufficient.
//
// The same input text is indexed into each of these fields automatically, with
// their differing analysis chains, when an indexed document has a value for the
// root field my_field.
//
//
// The most efficient way of querying to serve a search-as-you-type use case is
// usually a multi_match query of type bool_prefix that targets the root
// search_as_you_type field and its shingle subfields. This can match the query
// terms in any order, but will score documents higher if they contain the terms
// in order in a shingle subfield.
//
// To search for documents that strictly match the query terms in order, or to
// search using other properties of phrase queries, use a match_phrase_prefix
// query on the root field. A match_phrase query can also be used if the last
// term should be matched exactly, and not as a prefix. Using phrase queries may
// be less efficient than using the match_bool_prefix query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html
type SearchAsYouTypeField struct {
	MaxShingleSizeParam `bson:",inline" json:",inline"`
	analyzerParam       `bson:",inline" json:",inline"`
	indexParam          `bson:",inline" json:",inline"`
	indexOptionsParam   `bson:",inline" json:",inline"`
	NormsParam          `bson:",inline" json:",inline"`
	storeParam          `bson:",inline" json:",inline"`
	similarityParam     `bson:",inline" json:",inline"`
	TermVectorParam     `bson:",inline" json:",inline"`
}

func NewSearchAsYouTypeField() *SearchAsYouTypeField {
	return &SearchAsYouTypeField{BaseField: BaseField{MappingType: FieldTypeSearchAsYouType}}
}
