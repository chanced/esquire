package picker

import "github.com/chanced/dynamic"

const DefaultEagerGlobalOrdinals = false

type WithEagerGlobalOrdinals interface {
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
	EagerGlobalOrdinals() bool
	// SetEagerGlobalOrdinals sets the EagerGlobalOrdinals value to v
	SetEagerGlobalOrdinals(v interface{}) error
}

type eagerGlobalOrdinalsParam struct {
	eagerGlobalOrdinals dynamic.Bool
}

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
func (ego eagerGlobalOrdinalsParam) EagerGlobalOrdinals() bool {
	if b, ok := ego.eagerGlobalOrdinals.Bool(); ok {
		return b
	}
	return DefaultEagerGlobalOrdinals
}

// SetEagerGlobalOrdinals sets the eager_global_ordinals value to v
func (ego *eagerGlobalOrdinalsParam) SetEagerGlobalOrdinals(v interface{}) error {
	return ego.eagerGlobalOrdinals.Set(v)
}
