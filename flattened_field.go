package picker

// FlattenedField maps an entire object as a single field.
//
// By default, each subfield in an object is mapped and indexed separately. If
// the names or types of the subfields are not known in advance, then they are
// mapped dynamically.
//
// The flattened type provides an alternative approach, where the entire object
// is mapped as a single field. Given an object, the flattened mapping will
// parse out its leaf values and index them into one field as keywords. The
// object’s contents can then be searched through simple queries and
// aggregations.
//
// This data type can be useful for indexing objects with a large or unknown
// number of unique keys. Only one field mapping is created for the whole JSON
// object, which can help prevent a mappings explosion from having too many
// distinct field mappings.
//
// On the other hand, flattened object fields present a trade-off in terms of search functionality. Only basic queries are allowed, with no support for numeric range queries or highlighting. Further information on the limitations can be found in the Supported operations section.
//
// ! X-Pack
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/flattened.html
type FlattenedField struct {
	BaseField `json:",inline" bson:",inline"`
	// DepthLimitParam - i didnt write this out yet
	docValuesParam                `json:",inline" bson:",inline"`
	EagerGlobalOrdinalsParam      `json:",inline" bson:",inline"`
	IgnoreAboveParam              `json:",inline" bson:",inline"`
	indexParam                    `json:",inline" bson:",inline"`
	IndexOptionsParam             `json:",inline" bson:",inline"`
	nullValueParam                `json:",inline" bson:",inline"`
	SimilarityParam               `json:",inline" bson:",inline"`
	SplitQueriesOnWhitespaceParam `json:",inline" bson:",inline"`
}

func (f FlattenedField) Clone() Field {
	n := NewFlattenedField()
	n.SetDocValues(f.DocValues())
	n.SetEagerGlobalOrdinals(f.EagerGlobalOrdinals())
	n.SetIgnoreAbove(f.IgnoreAbove())
	n.SetIndexOptions(f.IndexOptions())
	n.SetIndex(f.Index())
	n.SetSimilarity(f.Similarity())
	n.SetNullValue(f.NullValue())
	n.SetSplitQueriesOnWhitespace(f.SplitQueriesOnWhitespace())
	return n
}
func NewFlattenedField() *FlattenedField {
	return &FlattenedField{BaseField: BaseField{MappingType: FieldTypeFlattened}}
}
