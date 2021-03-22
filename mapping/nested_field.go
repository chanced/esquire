package mapping

// NestedField is a specialised version of the object data type that allows
// arrays of objects to be indexed in a way that they can be queried
// independently of each other.
//
// When ingesting key-value pairs with a large, arbitrary set of keys, you might
// consider modeling each key-value pair as its own nested document with key and
// value fields. Instead, consider using the flattened data type, which maps an
// entire object as a single field and allows for simple searches over its
// contents. Nested documents and queries are typically expensive, so using the
// flattened data type for this use case is a better option.
//
// Interacting with nested documents
//
// Nested documents can be:
//
// - queried with the nested query.
//
// - analyzed with the nested and reverse_nested aggregations.
//
// - sorted with nested sorting.
//
// - retrieved and highlighted with nested inner hits.
//
// IMPORTANT
//
// Because nested documents are indexed as separate documents, they can only be
// accessed within the scope of the nested query, the nested/reverse_nested
// aggregations, or nested inner hits.
//
// For instance, if a string field within a nested document has index_options
// set to offsets to allow use of the postings during the highlighting, these
// offsets will not be available during the main highlighting phase. Instead,
// highlighting needs to be performed via nested inner hits. The same
// consideration applies when loading fields during a search through
// docvalue_fields or stored_fields.
//
//
// Limits on nested mappings and objects
//
// As described earlier, each nested object is indexed as a separate Lucene
// document. Continuing with the previous example, if we indexed a single
// document containing 100 user objects, then 101 Lucene documents would be
// created: one for the parent document, and one for each nested object. Because
// of the expense associated with nested mappings, Elasticsearch puts settings
// in place to guard against performance problems:
//
// 	index.mapping.nested_fields.limit
// The maximum number of distinct nested mappings in an index. The nested type
// should only be used in special cases, when arrays of objects need to be
// queried independently of each other. To safeguard against poorly designed
// mappings, this setting limits the number of unique nested types per index.
// Default is 50.
//
// 	index.mapping.nested_objects.limit
// The maximum number of nested JSON objects that a single document can contain
// across all nested types. This limit helps to prevent out of memory errors
// when a document contains too many nested objects. Default is 10000.
//
// To illustrate how this setting works, consider adding another nested type
// called comments to the previous example mapping. For each document, the
// combined number of user and comment objects it contains must be below the
// limit.
//
// See Settings to prevent mapping explosion regarding additional settings for
// preventing mappings explosion.
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html
//
type NestedField struct {
	BaseField            `json:",inline" bson:",inline"`
	DynamicParam         `json:",inline" bson:",inline"`
	PropertiesParam      `json:",inline" bson:",inline"`
	IncludeInParentParam `json:",inline" bson:",inline"`
	IncludeInRootParam   `json:",inline" bson:",inline"`
}

func (f NestedField) Clone() Field {
	n := NewNestedField()
	n.SetDynamic(f.Dynamic())
	n.SetProperties(f.Properties().Clone())
	n.SetIncludeInParent(f.IncludeInParent())
	n.SetIncludeInRoot(f.IncludeInRoot())
	return n
}

func NewNestedField() *NestedField {
	return &NestedField{BaseField: BaseField{MappingType: TypeNested}}
}
