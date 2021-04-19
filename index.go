package picker

import (
	"encoding/json"
)

type IndexSettingsParams struct {
	// The number of primary shards that an index should have. Defaults to 1.
	// This setting can only be set at index creation time. It cannot be changed
	// on a closed index.
	NumberOfShards interface{}

	// Number of routing shards used to split an index.
	//
	// This setting’s default value depends on the number of primary shards in
	// the index. The default is designed to allow you to split by factors of 2
	// up to a maximum of 1024 shards.
	//
	// In Elasticsearch 7.0.0 and later versions, this setting affects how
	// documents are distributed across shards. When reindexing an older index
	// with custom routing, you must explicitly set
	// index.number_of_routing_shards to maintain the same document
	// distribution.
	NumberOfRoutingShards interface{}

	// Whether or not shards should be checked for corruption before opening.
	// When corruption is detected, it will prevent the shard from being opened.
	// Accepts:
	//
	//  "false"
	// (default) Don’t check for corruption when opening a shard.
	//  "checksum"
	// Check for physical corruption.
	//  "true"
	// Check for both physical and logical corruption. This is much more
	// expensive in terms of CPU and memory usage.
	ShardCheckOnStartup string

	// The default value compresses stored data with LZ4 compression, but this
	// can be set to best_compression which uses DEFLATE for a higher
	// compression ratio, at the expense of slower stored fields performance. If
	// you are updating the compression type, the new one will be applied after
	// segments are merged. Segment merging can be forced using force merge.
	Codec string

	// The number of shards a custom routing value can go to. Defaults to 1 and
	// can only be set at index creation time. This value must be less than the
	// index.number_of_shards unless the index.number_of_shards value is also 1.
	// See Routing to an index partition for more details about how this setting
	// is used.
	RoutingPartitionSize interface{}

	// Indicates whether soft deletes are enabled on the index. Soft deletes can
	// only be configured at index creation and only on indices created on or
	// after Elasticsearch 6.5.0. Defaults to true.
	//
	// Deprecated
	SoftDeletesEnabled interface{}

	// The maximum period to retain a shard history retention lease before it is
	// considered expired. Shard history retention leases ensure that soft
	// deletes are retained during merges on the Lucene index. If a soft delete
	// is merged away before it can be replicated to a follower the following
	// process will fail due to incomplete history on the leader. Defaults to
	// 12h.
	SoftDeletesRetentionLeasePeriod string

	// Indicates whether cached filters are pre-loaded for nested queries.
	// Possible values are true (default) and false.
	LoadFixedBitsetFiltersEagerly interface{}

	// Indicates whether the index should be hidden by default. Hidden indices
	// are not returned by default when using a wildcard expression. This
	// behavior is controlled per request through the use of the
	// expand_wildcards parameter. Possible values are true and false (default).
	Hidden interface{}

	// The number of replicas each primary shard has. Defaults to 1.
	NumberOfReplicas interface{}

	// Auto-expand the number of replicas based on the number of data nodes in
	// the cluster. Set to a dash delimited lower and upper bound (e.g. 0-5) or
	// use all for the upper bound (e.g. 0-all). Defaults to false (i.e.
	// disabled). Note that the auto-expanded number of replicas only takes
	// allocation filtering rules into account, but ignores any other allocation
	// rules such as shard allocation awareness and total shards per node, and
	// this can lead to the cluster health becoming YELLOW if the applicable
	// rules prevent all the replicas from being allocated.
	AutoExpandReplicas interface{}

	// How long a shard can not receive a search or get request until it’s
	// considered search idle. (default is 30s)
	SearchIdleAfter string

	// How often to perform a refresh operation, which makes recent changes to
	// the index visible to search. Defaults to 1s. Can be set to -1 to disable
	// refresh. If this setting is not explicitly set, shards that haven’t seen
	// search traffic for at least index.search.idle.after seconds will not
	// receive background refreshes until they receive a search request.
	// Searches that hit an idle shard where a refresh is pending will wait for
	// the next background refresh (within 1s). This behavior aims to
	// automatically optimize bulk indexing in the default case when no searches
	// are performed. In order to opt out of this behavior an explicit value of
	// 1s should set as the refresh interval.
	RefreshInterval string

	// The maximum value of from + size for searches to this index. Defaults to
	// 10000. Search requests take heap memory and time proportional to from +
	// size and this limits that memory.
	MaxResultWindow interface{}

	// The maximum value of from + size for inner hits definition and top hits
	// aggregations to this index. Defaults to 100. Inner hits and top hits
	// aggregation take heap memory and time proportional to from + size and
	// this limits that memory.
	MaxInnerResultWindow interface{}

	// The maximum value of window_size for rescore requests in searches of this
	// index. Defaults to index.max_result_window which defaults to 10000.
	// Search requests take heap memory and time proportional to
	// max(window_size, from + size) and this limits that memory.
	MaxRescoreWindow interface{}

	// The maximum number of docvalue_fields that are allowed in a query.
	// Defaults to 100. Doc-value fields are costly since they might incur a
	// per-field per-document seek.
	MaxDocvalueFieldsSearch interface{}

	// The maximum number of script_fields that are allowed in a query. Defaults
	// to 32.
	MaxScriptFields interface{}

	// The maximum allowed difference between min_gram and max_gram for
	// NGramTokenizer and NGramTokenFilter. Defaults to 1.
	MaxNgramDiff interface{}

	// The maximum allowed difference between max_shingle_size and
	// min_shingle_size for the shingle token filter. Defaults to 3.
	MaxShingleDiff interface{}

	// Maximum number of refresh listeners available on each shard of the index.
	// These listeners are used to implement refresh=wait_for.
	MaxRefreshListeners interface{}

	// The maximum number of tokens that can be produced using _analyze API.
	// Defaults to 10000.
	AnalyzeMaxTokenCount interface{}

	// The maximum number of characters that will be analyzed for a highlight
	// request. This setting is only applicable when highlighting is requested
	// on a text that was indexed without offsets or term vectors. Defaults to
	// 1000000.
	HighlightMaxAnalyzedOffset interface{}

	// The maximum number of terms that can be used in Terms Query. Defaults to
	// 65536.
	MaxTermsCount interface{}

	// The maximum length of regex that can be used in Regexp Query. Defaults to
	// 1000.
	MaxRegexLength interface{}

	// (string or array of strings) Wildcard (*) patterns matching one or more
	// fields. The following query types search these matching fields by
	// default:
	//
	// - More like this
	//
	// - Multi-match
	//
	// - Query string
	//
	// - Simple query string
	//
	// Defaults to *, which matches all fields eligible for term-level queries,
	// excluding metadata fields.
	QueryDefaultField interface{}

	// Controls shard allocation for this index. It can be set to:
	//
	//  "all" // (default) Allows shard allocation for all shards.
	//  "primaries" // Allows shard allocation only for primary shards.
	//  "new_primaries" // Allows shard allocation only for newly-created primary shards.
	//  "none" // No shard allocation is allowed.
	RoutingAllocationEnable string

	// Enables shard rebalancing for this index. It can be set to:
	//
	//  "all" // (default) - Allows shard rebalancing for all shards.
	//  "primaries" // Allows shard rebalancing only for primary shards.
	//  "replicas" // Allows shard rebalancing only for replica shards.
	//  "none" // No shard rebalancing is allowed.
	RoutingRebalanceEnable string

	// The length of time that a deleted document’s version number remains
	// available for further versioned operations. Defaults to 60s.
	GCDeletes interface{}

	// The default ingest node pipeline for this index. Index requests will fail
	// if the default pipeline is set and the pipeline does not exist. The
	// default may be overridden using the pipeline parameter. The special
	// pipeline name _none indicates no ingest pipeline should be run.
	DefaultPipeline string

	// The final ingest node pipeline for this index. Index requests will fail
	// if the final pipeline is set and the pipeline does not exist. The final
	// pipeline always runs after the request pipeline (if specified) and the
	// default pipeline (if it exists). The special pipeline name _none
	// indicates no ingest pipeline will run.
	FinalPipeline string
}

type IndexParams struct {
	Mappings Mappings `json:"mappings"`
	Settings map[string]interface{}
}

func (p IndexParams) Index() (*Index, error) {

	fm, err := p.Mappings.FieldMappings()
	i := &Index{}
	if err != nil {
		return i, err
	}
	i.Mappings = fm
	return i, nil
}

type Index struct {
	Mappings FieldMappings `json:"mappings"`
}

//easyjson:json
type index struct {
	Mappings FieldMappings          `json:"mappings"`
	Settings map[string]interface{} `json:"settings"`
}

func (i Index) MarshalBSON() ([]byte, error) {
	return i.MarshalJSON()
}

func (i Index) MarshalJSON() ([]byte, error) {
	return index{Mappings: i.Mappings}.MarshalJSON()
}

func (i *Index) UnmarshalBSON(data []byte) error {
	return i.UnmarshalJSON(data)
}

func (i *Index) UnmarshalJSON(data []byte) error {
	var idx index
	err := json.Unmarshal(data, &idx)
	if err != nil {
		return err
	}
	i.Mappings = idx.Mappings
	return nil
}

func NewIndex(params IndexParams) (*Index, error) {
	return params.Index()
}
