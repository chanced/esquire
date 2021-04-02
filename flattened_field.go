package picker

type FlattenedFieldParams struct {
	DepthLimit interface{} `json:"depth_limit,omitempty"`
	DocValues  interface{} `json:"doc_values,omitempty"`
	// To support aggregations and other operations that require looking up field
	// values on a per-document basis, Elasticsearch uses a data structure called
	// doc values. Term-based field types such as keyword store their doc values
	// using an ordinal mapping for a more compact representation. This mapping
	// works by assigning each term an incremental integer or ordinal based on its
	// lexicographic order. The field’s doc values store only the ordinals for each
	// document instead of the original terms, with a separate lookup structure to
	// convert between ordinals and terms.
	//
	// When used during aggregations, ordinals can greatly improve performance. As
	// an example, the terms aggregation relies only on ordinals to collect
	// documents into buckets at the shard-level, then converts the ordinals back to
	// their original term values when combining results across shards.
	//
	// Each index segment defines its own ordinal mapping, but aggregations collect
	// data across an entire shard. So to be able to use ordinals for shard-level
	// operations like aggregations, Elasticsearch creates a unified mapping called
	// global ordinals. The global ordinal mapping is built on top of segment
	// ordinals, and works by maintaining a map from global ordinal to the local
	// ordinal for each segment.
	//
	// Global ordinals are used if a search contains any of the following
	// components:
	//
	// - Certain bucket aggregations on keyword, ip, and flattened fields. This
	// includes terms aggregations as mentioned above, as well as composite,
	// diversified_sampler, and significant_terms.
	//
	// - Bucket aggregations on text fields that require fielddata to be enabled.
	//
	// - Operations on parent and child documents from a join field, including
	// has_child queries and parent aggregations.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eager-global-ordinals.html
	EagerGlobalOrdinals interface{} `json:"eager_global_ordinals,omitempty"`
	ignoreAboveParam
	indexParam
	IndexOptionsParam
	nullValueParam
	SimilarityParam
	SplitQueriesOnWhitespaceParam
}

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
	depthLimitParam
	docValuesParam
	eagerGlobalOrdinalsParam
	ignoreAboveParam
	indexParam
	IndexOptionsParam
	nullValueParam
	SimilarityParam
	SplitQueriesOnWhitespaceParam
}

func NewFlattenedField() *FlattenedField {
	return &FlattenedField{BaseField: BaseField{MappingType: FieldTypeFlattened}}
}
