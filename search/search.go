package search

import (
	"time"

	"github.com/chanced/dynamic"
)

func NewSearch(Params) (*Search, error) {
	s := &Search{}
	return s, nil
}

var (
	DefaultExplain = false
	DefaultFrom    = int64(0)
	DefaultSize    = int64(10)
)

// Params are the initial params passed to NewSearch
type Params struct {
	// Array of wildcard (*) patterns. The request returns doc values for field
	// names matching these patterns in the hits.fields property of the response
	// (Optional) .
	//
	// You can specify items in the array as a string or object.
	DocValueFields DocValueFields

	// Array of wildcard (*) patterns. The request returns values for field
	// names matching these patterns in the hits.fields property of the response
	// (Optional).
	//
	// You can specify items in the array as a string or object.
	Fields Fields

	// If true, returns detailed information about score computation as part of
	// a hit. Defaults to false. (Optional)
	Explain bool

	// Starting document offset. Defaults to 0.
	//
	// By default, you cannot page through more than 10,000 hits using the from
	// and size parameters. To page through more hits, use the search_after
	// parameter. (Optional)
	From int64

	// Boosts the _score of documents from specified indices (Optional).
	IndicesBoost IndicesBoost

	// Minimum _score for matching documents. Documents with a lower _score are
	// not included in the search results (Optional).
	MinScore float64

	// Limits the search to a point in time (PIT). If you provide a pit, you
	// cannot specify a <target> in the request path. (Optional)
	PointInTime *PointInTime

	//  Defines a runtime field in the search request that exist only as part of
	//  the query. Fields defined in the search request take precedence over
	//  fields defined with the same name in the index mappings. (Optional)
	RuntimeMappings RuntimeMappings

	// If true, returns sequence number and primary term of the last
	// modification of each hit. See Optimistic concurrency control. (Optional)
	SeqNoPrimaryTerm bool

	// The number of hits to return. Defaults to 10. (Optional)
	//
	// By default, you cannot page through more than 10,000 hits using the from
	// and size parameters. To page through more hits, use the search_after
	// parameter.
	Size int64

	// Indicates which source fields are returned for matching documents. These
	// fields are returned in the hits._source property of the search response.
	// Defaults to true. (Optional)
	Source *Source

	// Stats groups to associate with the search. Each group maintains a
	// statistics aggregation for its associated searches. You can retrieve
	// these stats using the indices stats API (Optional).
	Stats []string

	// The maximum number of documents to collect for each shard, upon reaching
	// which the query execution will terminate early. (Optional)
	//
	// Defaults to 0, which does not terminate query execution early.
	TerminateAfter int64

	// Specifies the period of time to wait for a response. If no response is
	// received before the timeout expires, the request fails and returns an
	// error. Defaults to no timeout. (Optional)
	Timeout time.Duration

	// If true, returns document version as part of a hit. Defaults to false. (Optional)
	Version bool

	// Term returns documents that contain an exact term in a provided field.
	//
	// You can use the term query to find documents based on a precise value such as
	// a price, a product ID, or a username.
	//
	// Avoid using the term query for text fields.
	//
	// By default, Elasticsearch changes the values of text fields as part of
	// analysis. This can make finding exact matches for text field values
	// difficult.
	//
	// To search text field values, use the match query instead.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
	Term *Term

	// Terms returns documents that contain one or more exact terms in a provided
	// field.
	//
	// The terms query is the same as the term query, except you can search for
	// multiple values.
	Terms *Terms

	// Match returns documents that match a provided text, number, date or boolean
	// value. The provided text is analyzed before matching.
	//
	// The match query is the standard query for performing a full-text search,
	// including options for fuzzy matching.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
	Match *Match

	// Boolean is a query that matches documents matching boolean combinations
	// of other queries. The bool query maps to Lucene BooleanQuery. It is built
	// using one or more boolean clauses, each clause with a typed occurrence.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
	Boolean *Boolean
}

type Search struct {
	// Defines the search definition using the Query DSL. (Optional)
	query            Query           // query
	docValueFields   DocValueFields  // docvalue_fields
	fields           Fields          // fields
	explain          bool            // explain
	from             int64           // from
	indicesBoost     IndicesBoost    // indices_boost
	minScore         float64         // min_score
	pointInTime      *PointInTime    // pit
	runtimeMappings  RuntimeMappings // runtime_mappings
	seqNoPrimaryTerm bool            // seq_no_primary_term
	size             dynamic.Number  // size
	source           *Source         // _source
	stats            []string        // stats
	terminateAfter   int64           // terminate_after
	timeout          time.Duration   // timeout
	version          bool            // version
}

func (s *Search) UnmarshalJSON(data []byte) error {
	panic("not impl")
}

func (s Search) MarshalJSON() ([]byte, error) {
	panic("not impl")
}

// DocValueFields is used to return  return doc values for one or more fields in
// the search response.
//
// (Optional, array of strings and objects) Array of wildcard (*) patterns. The
// request returns doc values for field names matching these patterns in the
// hits.fields property of the response.
//
// You can specify items in the array as a string or object.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#docvalue-fields
func (s Search) DocValueFields() DocValueFields {
	if s.docValueFields == nil {
		s.docValueFields = DocValueFields{}
	}
	return s.docValueFields
}

// SetDocValueFields sets DocValueFieldsValue to v
func (s *Search) SetDocValueFields(v DocValueFields) *Search {
	s.docValueFields = v
	return s
}

// Fields allows for retrieving a list of document fields in the search
// response. It consults both the document _source and the index mappings to
// return each value in a standardized way that matches its mapping type. By
// default, date fields are formatted according to the date format parameter in
// their mappings. You can also use the fields parameter to retrieve runtime
// field values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#search-fields-param
func (s *Search) Fields() Fields {
	if s.fields == nil {
		s.fields = Fields{}
	}
	return s.fields

}

// SetFields sets the FieldsValue to v
func (s *Search) SetFields(v Fields) *Search {
	s.fields = v
	return s
}

// Explain indicates whether the search returns detailed information about score
// computation as part of a hit. Defaults to false.
func (s Search) Explain() bool {
	return s.explain
}

// SetExplain sets the ExplainValue to v
func (s *Search) SetExplain(v bool) *Search {
	s.explain = v
	return s
}

// From sets the starting document offset. Defaults to 0.
//
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters. To page through more hits, use the search_after parameter.
func (s Search) From() int64 {
	return s.from
}

// SetFrom sets the FromValue to v
func (s *Search) SetFrom(v int64) *Search {
	s.from = v
	return s
}

// IndicesBoost buusts the _score of documents from specified indices
func (s *Search) IndicesBoost() IndicesBoost {
	if s.indicesBoost == nil {
		s.indicesBoost = IndicesBoost{}
	}
	return s.indicesBoost
}

// SetIndicesBoost sets IndicesBoostValue to v
func (s *Search) SetIndicesBoost(v IndicesBoost) *Search {
	s.indicesBoost = v
	return s
}

// MinScore is the minimum _score for matching documents. Documents with a lower
// _score are not included in the search results.
func (s *Search) MinScore() float64 {
	return s.minScore
}

// SetMinScore sets the MinScoreValue to v
func (s *Search) SetMinScore(v float64) *Search {
	s.minScore = v
	return s
}

// SetPointInTime sets the PointInTimeValue to v
func (s *Search) SetPointInTime(v *PointInTime) *Search {
	s.pointInTime = v
	return s
}

// PointInTime is a lightweight view into the state of the data as it existed
// when initiated
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
func (s Search) PointInTime() *PointInTime {
	return s.pointInTime
}

func (s Search) PointInTimeID() string {
	pit := s.PointInTime()

	if pit == nil {
		return ""
	}
	return pit.ID
}

func (s Search) PITID() string {

	return s.PointInTimeID()
}

func (s Search) PointInTimeKeepAlive() *time.Time {
	pit := s.PointInTime()
	if pit == nil {
		return nil
	}
	return pit.KeepAlive
}
func (s Search) PITKeepAlive() *time.Time {
	return s.PointInTimeKeepAlive()
}

// PIT is an alias for PointInTime
func (s Search) PIT() *PointInTime {
	return s.PointInTime()
}

// SetPIT is an alias for SetPointInTime
func (s *Search) SetPIT(v *PointInTime) *Search {
	return s.SetPointInTime(v)
}

// SetQuery sets QueryValue to v
func (s *Search) SetQuery(v *Query) *Search {
	if v == nil {
		s.query = Query{}
	}
	s.query = *v
	return s
}

// Query defines the search definition using the Query DSL.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
func (s Search) Query() *Query {
	return &s.query
}

func (s Search) RuntimeMappings() RuntimeMappings {
	if s.runtimeMappings == nil {
		s.runtimeMappings = RuntimeMappings{}
	}
	return s.runtimeMappings
}

func (s *Search) SetRuntimeMappings(v RuntimeMappings) *Search {
	s.runtimeMappings = v
	return s
}

// SeqNoPrimaryTerm https://www.elastic.co/guide/en/elasticsearch/reference/current/optimistic-concurrency-control.html
func (s Search) SeqNoPrimaryTerm() bool {
	return s.seqNoPrimaryTerm
}

func (s *Search) SetSeqNoPrimaryTerm(v bool) *Search {
	s.seqNoPrimaryTerm = v
	return s
}

// Size is number of hits to return. Defaults to 10.
//
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters. To page through more hits, use the search_after parameter.
func (s Search) Size() int64 {
	if i, ok := s.size.Int(); ok {
		return i
	}
	return DefaultSize
}

func (s *Search) SetSize(v int64) *Search {
	s.size.Set(v)
	return s
}

// Source indicates which source fields are returned for matching documents.
// These fields are returned in the hits._source property of the search
// response. Defaults to true.
func (s Search) Source() *Source {
	return s.source
}

// SetSource sets the value of Source
//
// The options are:
//  search.Source, *search.Source,
//  string, []string,
//  dynamic.StringOrArrayOfStrings, *dynamic.StringOrArrayOfStrings,
//  search.SourceSpecifications
//  *search.SourceSpecifications
//  bool, *bool
//  nil
// Note, "true" || "false" get parsed as boolean
//
// SetSource panics if v is not one of the types listed above.
//
// You can explicitly set the source, such as:
//  s := NewSearch()
//  src := &Source{}
//  err := src.SetValue(v)
//  _ = err // handle err
//  s.SourceValue = src
func (s *Search) SetSource(v interface{}) *Search {
	switch t := v.(type) {
	case *Source:
		ts := *t
		s.source = &ts
	case Source:
		s.source = &t
	default:
		s.source = &Source{}
		err := s.source.SetValue(v)
		if err != nil {
			panic(err)
		}
	}
	return s
}

// Stats groups to associate with the search. Each group maintains a statistics
// aggregation for its associated searches. You can retrieve these stats using
// the indices stats API (Optional).
func (s Search) Stats() []string {
	return s.stats
}

func (s *Search) SetStats(v []string) *Search {
	s.stats = v
	return s
}

// TerminateAfter is maximum number of documents to collect for each shard, upon
// reaching which the query execution will terminate early.
//
// Defaults to 0, which does not terminate query execution early.
func (s Search) TerminateAfter() int64 {
	return s.terminateAfter
}

func (s *Search) SetTerminateAfter(v int64) *Search {
	s.terminateAfter = v
	return s
}

// Timeout specifies the period of time to wait for a  response. If no response
// is received before the timeout expires, the request fails and returns an
// error. Defaults to no timeout.
func (s Search) Timeout() time.Duration {
	return s.timeout
}

func (s *Search) SetTimeout(v time.Duration) *Search {
	if s.Timeout() != v {
		s.timeout = v
	}
	return s
}

// Version determines whether the document version should be returned as part a
// hit. Default: false
func (s Search) Version() bool {
	return s.version
}

func (s *Search) SetVersion(v bool) *Search {
	s.version = v
	return s
}

func (s *Search) Clone() (*Search, error) {
	n, _ := NewSearch(Params{})
	n.SetDocValueFields(s.DocValueFields().Clone())
	n.SetExplain(s.Explain())
	n.SetFields(s.Fields().Clone())
	n.SetFrom(s.From())
	n.SetIndicesBoost(s.IndicesBoost().Clone())
	n.SetMinScore(s.MinScore())
	n.SetPointInTime(s.PointInTime().Clone())
	n.SetQuery(s.Query().Clone())
	n.SetRuntimeMappings(s.RuntimeMappings().Clone())
	n.SetSeqNoPrimaryTerm(s.SeqNoPrimaryTerm())
	n.SetSize(s.Size())
	n.SetSource(s.Source())
	n.SetStats(s.Stats())
	n.SetTerminateAfter(s.TerminateAfter())
	n.SetTimeout(s.Timeout())
	n.SetVersion(s.Version())
	return n, nil
}
