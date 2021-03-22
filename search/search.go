package search

import "time"

func NewSearch() *Search {
	return &Search{}
}

type Search struct {
	// Array of wildcard (*) patterns. The request returns doc values for field
	// names matching these patterns in the hits.fields property of the response
	// (Optional) .
	//
	// You can specify items in the array as a string or object.
	DocValueFieldsValue DocValueFields `bson:"docvalue_fields,omitempty" json:"docvalue_fields,omitempty"`

	// Array of wildcard (*) patterns. The request returns values for field
	// names matching these patterns in the hits.fields property of the response
	// (Optional).
	//
	// You can specify items in the array as a string or object.
	FieldsValue Fields `bson:"fields,omitempty" json:"fields,omitempty"`

	// (Optional, Boolean) If true, returns detailed information about score
	// computation as part of a hit. Defaults to false.
	ExplainValue *bool `bson:"epxlain,omitempty" json:"explain,omitempty"`

	// Starting document offset. Defaults to 0.
	//
	// By default, you cannot page through more than 10,000 hits using the from
	// and size parameters. To page through more hits, use the search_after
	// parameter.
	FromValue *int `bson:"from,omitempty" json:"from,omitempty"`

	// Boosts the _score of documents from specified indices (Optional).
	IndicesBoostValue IndicesBoost `bson:"indices_boost,omitempty"`

	// Minimum _score for matching documents. Documents with a lower _score are
	// not included in the search results (Optional, float).
	MinScoreValue *float32 `bson:"min_score,omitempty" json:"min_score,omitempty"`

	// Limits the search to a point in time (PIT). If you provide a pit, you
	// cannot specify a <target> in the request path. (Optional)
	PointInTimeValue *PointInTime `bson:"pit,omitempty" json:"pit,omitempty"`

	// Defines the search definition using the Query DSL. (Optional)
	QueryValue *Query `bson:"query,omitempty" json:"query,omitempty"`

	//  Defines a runtime field in the search request that exist only as part of
	//  the query. Fields defined in the search request take precedence over
	//  fields defined with the same name in the index mappings. (Optional)
	RuntimeMappingsValue RuntimeMappings `bson:"runtime_mappings,omitempty" json:"runtime_mappings,omitempty"`

	// If true, returns sequence number and primary term of the last
	// modification of each hit. See Optimistic concurrency control. (Optional)
	SeqNoPrimaryTermValue *bool `bson:"seq_no_primary_term,omitempty" json:"seq_no_primary_term,omitempty"`
	// The number of hits to return. Defaults to 10. (Optional)
	//
	// By default, you cannot page through more than 10,000 hits using the from and size parameters. To page through more hits, use the search_after parameter.
	SizeValue *int `bson:"size,omitempty" json:"size,omitempty"`

	// Indicates which source fields are returned for matching documents. These fields are returned in the hits._source property of the search response. Defaults to true. (Optional)
	SourceValue *Source `bson:"_source,omitempty" json:"_source,omitempty"`

	// Stats groups to associate with the search. Each group maintains a statistics aggregation for its associated searches. You can retrieve these stats using the indices stats API (Optional).
	StatsValue []string `bson:"stats,omitempty" json:"stats,omitempty"`

	TerminateAfterValue *int `bson:"terminate_after,omitempty" json:"terminate_after,omitempty"`

	TimeoutValue *time.Duration `bson:"timeout,omitempty" json:"timeout,omitempty"`

	VersionValue *bool `bson:"version,omitempty" json:"version,omitempty"`
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
	if s.DocValueFieldsValue == nil {
		s.DocValueFieldsValue = DocValueFields{}
	}
	return s.DocValueFieldsValue
}

// SetDocValueFields sets DocValueFieldsValue to v
func (s *Search) SetDocValueFields(v DocValueFields) *Search {
	s.DocValueFieldsValue = v
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
func (s Search) Fields() Fields {
	if s.FieldsValue != nil {
		return s.FieldsValue
	}
	return Fields{}
}

// SetFields sets the FieldsValue to v
func (s *Search) SetFields(v Fields) *Search {
	s.FieldsValue = v
	return s
}

// Explain indicates whether the search returns detailed information about score
// computation as part of a hit. Defaults to false.
func (s Search) Explain() bool {
	if s.ExplainValue == nil {
		return false
	}
	return *s.ExplainValue
}

// SetExplain sets the ExplainValue to v
func (s *Search) SetExplain(v bool) *Search {
	if s.Explain() != v {
		s.ExplainValue = &v
	}
	return s
}

// From sets the starting document offset. Defaults to 0.
//
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters. To page through more hits, use the search_after parameter.
func (s Search) From() int {
	if s.FromValue == nil {
		return 0
	}
	return *s.FromValue
}

// SetFrom sets the FromValue to v
func (s *Search) SetFrom(v int) *Search {
	if s.From() != v {
		s.FromValue = &v
	}
	return s
}

// IndicesBoost buusts the _score of documents from specified indices
func (s Search) IndicesBoost() IndicesBoost {
	if s.IndicesBoostValue == nil {
		s.IndicesBoostValue = IndicesBoost{}
	}
	return s.IndicesBoostValue
}

// SetIndicesBoost sets IndicesBoostValue to v
func (s *Search) SetIndicesBoost(v IndicesBoost) *Search {
	s.IndicesBoostValue = v
	return s
}

// MinScore is the minimum _score for matching documents. Documents with a lower
// _score are not included in the search results.
func (s Search) MinScore() float32 {
	if s.MinScoreValue == nil {
		return 0
	}
	return *s.MinScoreValue
}

// SetMinScore sets the MinScoreValue to v
func (s *Search) SetMinScore(v float32) *Search {
	if s.MinScore() != v {
		s.MinScoreValue = &v
	}
	return s
}

// SetPointInTime sets the PointInTimeValue to v
func (s *Search) SetPointInTime(v *PointInTime) *Search {
	nv := *v
	s.PointInTimeValue = &nv
	return s
}

// PointInTime is a lightweight view into the state of the data as it existed
// when initiated
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
func (s Search) PointInTime() *PointInTime {
	return s.PointInTimeValue
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
	s.QueryValue = v
	return s
}

// Query defines the search definition using the Query DSL.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
func (s Search) Query() *Query {
	return s.QueryValue
}

func (s Search) RuntimeMappings() RuntimeMappings {
	if s.RuntimeMappingsValue == nil {
		s.RuntimeMappingsValue = RuntimeMappings{}
	}
	return s.RuntimeMappingsValue
}

func (s *Search) SetRuntimeMappings(v RuntimeMappings) *Search {

	s.RuntimeMappingsValue = v
	return s
}

// SeqNoPrimaryTerm https://www.elastic.co/guide/en/elasticsearch/reference/current/optimistic-concurrency-control.html
func (s Search) SeqNoPrimaryTerm() bool {
	if s.SeqNoPrimaryTermValue == nil {
		return false
	}
	return *s.SeqNoPrimaryTermValue
}

func (s *Search) SetSeqNoPrimaryTerm(v bool) *Search {
	if s.SeqNoPrimaryTerm() != v {
		s.SeqNoPrimaryTermValue = &v
	}
	return s
}

// Size is number of hits to return. Defaults to 10.
//
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters. To page through more hits, use the search_after parameter.
func (s Search) Size() int {
	if s.SizeValue == nil {
		return 10
	}
	return *s.SizeValue
}

func (s *Search) SetSize(v int) *Search {
	if s.Size() != v {
		s.SizeValue = &v
	}
	return s
}

// Source indicates which source fields are returned for matching documents.
// These fields are returned in the hits._source property of the search
// response. Defaults to true.
func (s Search) Source() *Source {
	return s.SourceValue
}

// SetSource sets the value of Source
//
// The options are:
//  search.Source
//  *search.Source
//  string
//  []string
//  dynamic.StringOrArrayOfStrings
//  *dynamic.StringOrArrayOfStrings
//  search.SourceSpecifications
//  *search.SourceSpecifications
//  bool
//  *bool
//  nil
// Note, "true" || "false" get parsed as boolean
//
// SetSource panics if v is not one of the types listed above. The expectation
// is that this method will be utilized in a Builder
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
		s.SourceValue = &ts
	case Source:
		s.SourceValue = &t
	default:
		s.SourceValue = &Source{}
		err := s.SourceValue.SetValue(v)
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
	return s.StatsValue
}
func (s *Search) SetStats(v []string) *Search {
	s.StatsValue = v
	return s
}

// TerminateAfter is maximum number of documents to collect for each shard, upon
// reaching which the query execution will terminate early.
//
// Defaults to 0, which does not terminate query execution early.
func (s Search) TerminateAfter() int {
	if s.TerminateAfterValue == nil {
		return 0
	}
	return *s.TerminateAfterValue
}

func (s *Search) SetTerminateAfter(v int) *Search {
	if s.TerminateAfter() != v {
		s.TerminateAfterValue = &v
	}
	return s
}

// Timeout specifies the period of time to wait for a  response. If no response
// is received before the timeout expires, the request fails and returns an
// error. Defaults to no timeout.
func (s Search) Timeout() time.Duration {
	if s.TimeoutValue == nil {
		return time.Duration(0)
	}
	return *s.TimeoutValue
}

func (s *Search) SetTimeout(v time.Duration) *Search {
	if s.Timeout() != v {
		s.TimeoutValue = &v
	}
	return s
}

// Version determines whether the document version should be returned as part a
// hit. Default: false
func (s Search) Version() bool {
	if s.VersionValue == nil {
		return false
	}
	return *s.VersionValue
}

func (s *Search) SetVersion(v bool) *Search {
	if s.Version() != v {
		s.VersionValue = &v
	}
	return s
}

// AddMatch adds a match query to the search. It panics if the field already
// exists or other errors arise (like not setting the query). Use SetMatch to
// overwrite the field instead.
//
// AddMatch panics if there is an error. It is intended to be utilized in a
// builder. To avoid panics, use the same function on the Query itself:
//  s := search.NewSearch()
//  s.AddMatch("field", Match{Query: "example"})
//  s.AddMatch("field", Match{Query: "example"}) // This will panic
//  // this will not:
//  err := s.Query().AddMatch(field, match)
//  _ = err // handle error
func (s *Search) AddMatch(field string, match Match) *Search {
	err := s.Query().AddMatch(field, match)
	if err != nil {
		panic(err)
	}
	return s
}

// SetMatch assigns a match query to the search. It overwrites the field if it
// exists. AddMatch will error instead
//
// SetMatch panics if there is an error. It is intended to be utilized in a
// builder. To avoid panics, use the same function on the Query itself:
//  s := search.NewSearch()
//  err := s.Query().SetMatch(field, match)
//  _ = err // handle error
func (s *Search) SetMatch(field string, match Match) *Search {
	err := s.Query().SetMatch(field, match)
	if err != nil {
		panic(err)
	}
	return s
}
func (s *Search) Clone() *Search {
	n := NewSearch()
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
	return n
}
