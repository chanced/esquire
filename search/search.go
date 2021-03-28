package search

import (
	"encoding/json"
	"time"

	"github.com/chanced/dynamic"
)

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
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#docvalue-fields
	//
	// See also:
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/doc-values.html
	DocValueFields Fields

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

func NewSearch(p Params) (*Search, error) {

	s := &Search{
		docValueFields:   p.DocValueFields,
		fields:           p.Fields,
		explain:          p.Explain,
		from:             p.From,
		indicesBoost:     p.IndicesBoost,
		minScore:         p.MinScore,
		pointInTime:      p.PointInTime,
		runtimeMappings:  p.RuntimeMappings,
		seqNoPrimaryTerm: p.SeqNoPrimaryTerm,
		source:           p.Source,
		stats:            p.Stats,
		terminateAfter:   p.TerminateAfter,
		timeout:          p.Timeout,
		version:          p.Version,
	}
	if p.Size != 0 {
		s.SetSize(p.Size)
	}
	qp := QueryParams{
		Term:    p.Term,
		Terms:   p.Terms,
		Match:   p.Match,
		Boolean: p.Boolean,
	}
	q, err := NewQuery(qp)
	if err != nil {
		return nil, err
	}
	s.query = q
	return s, nil
}

type Search struct {
	// Defines the search definition using the Query DSL. (Optional)
	query            *Query          // query
	docValueFields   Fields          // docvalue_fields
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

var zeroSearch = &Search{}

func (s *Search) UnmarshalJSON(data []byte) (err error) {
	*s = *zeroSearch
	var m map[string]dynamic.JSON
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	if d, ok := m["query"]; ok {
		var q Query
		err = json.Unmarshal(d, &q)
		if err != nil {
			return err
		}
		s.query = &q
	}
	if d, ok := m["docvalue_fields"]; ok {
		var df Fields
		err = json.Unmarshal(d, &df)
		if err != nil {
			return err
		}
		s.docValueFields = df
	}
	if d, ok := m["fields"]; ok {
		var f Fields
		err = json.Unmarshal(d, &f)
		if err != nil {
			return err
		}
		s.fields = f
	}
	if d, ok := m["explain"]; ok {
		b, err := dynamic.NewBool(d.UnquotedString())
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			s.explain = v
		}
	}
	if d, ok := m["from"]; ok {
		var i int64
		err := json.Unmarshal(d, &i)
		if err != nil {
			return err
		}
		s.from = i
	}
	if d, ok := m["indices_boost"]; ok {
		var ib IndicesBoost
		err = json.Unmarshal(d, &ib)
		if err != nil {
			return err
		}
		s.indicesBoost = ib
	}
	if d, ok := m["min_score"]; ok {
		var i float64
		err = json.Unmarshal(d, &i)
		if err != nil {
			return err
		}
		s.minScore = i
	}
	if d, ok := m["pit"]; ok {
		var pit PointInTime
		err = json.Unmarshal(d, &pit)
		if err != nil {
			return err
		}
		s.pointInTime = &pit
	}
	if d, ok := m["runtime_mappings"]; ok {
		var r RuntimeMappings
		err = json.Unmarshal(d, &r)
		if err != nil {
			return err
		}
		s.runtimeMappings = r
	}
	if d, ok := m["seq_no_primary_term"]; ok {
		var b bool
		err = json.Unmarshal(d, &b)
		if err != nil {
			return err
		}
	}
	if d, ok := m["size"]; ok {
		n, err := dynamic.NewNumber(d.UnquotedString())
		if err != nil {
			return err
		}
		s.size = n
	}

	if d, ok := m["_source"]; ok {
		var v Source
		err = json.Unmarshal(d, &v)
		if err != nil {
			return err
		}
		s.source = &v
	}
	if d, ok := m["stats"]; ok {
		var v []string
		err = json.Unmarshal(d, &v)
		if err != nil {
			return err
		}
		s.stats = v
	}
	if d, ok := m["terminate_after"]; ok {
		var i int64
		err = json.Unmarshal(d, &i)
		if err != nil {
			return err
		}
		s.terminateAfter = i
	}
	if d, ok := m["timeout"]; ok {
		var v time.Duration
		err = json.Unmarshal(d, &v)
		if err != nil {
			return err
		}
		s.timeout = v
	}
	if d, ok := m["version"]; ok {
		var v bool
		err = json.Unmarshal(d, &v)
		if err != nil {
			return err
		}
		s.version = v
	}
	return nil
}

func (s Search) MarshalJSON() ([]byte, error) {

	data := map[string]dynamic.JSON{}
	if len(s.docValueFields) > 0 {
		b, err := json.Marshal(s.docValueFields)
		if err != nil {
			return nil, err
		}
		data["docvalue_fields"] = b
	}

	if len(s.fields) > 0 {
		b, err := json.Marshal(s.fields)
		if err != nil {
			return nil, err
		}
		data["fields"] = b
	}
	if s.minScore > 0 {
		n, _ := dynamic.NewNumber(s.minScore)
		data["min_score"] = n.Bytes()
	}
	if s.explain {
		data["explain"] = trueBytes
	}
	if s.from > 0 {
		n, _ := dynamic.NewNumber(s.from)
		data["from"] = n.Bytes()
	}
	if len(s.indicesBoost) > 0 {
		b, err := json.Marshal(s.indicesBoost)
		if err != nil {
			return nil, err
		}
		data["indices_boost"] = b
	}
	if s.pointInTime != nil && len(s.pointInTime.ID) > 0 {
		b, err := json.Marshal(s.pointInTime)
		if err != nil {
			return nil, err
		}
		data["pit"] = b
	}
	if s.query != nil {
		b, err := json.Marshal(s.query)
		if err != nil {
			return nil, err
		}
		data["query"] = b
	}
	if len(s.runtimeMappings) > 0 {
		b, err := json.Marshal(s.runtimeMappings)
		if err != nil {
			return nil, err
		}
		data["runtime_mappings"] = b
	}
	if s.seqNoPrimaryTerm {
		data["seq_no_primary_term"] = trueBytes
	}

	if i, ok := s.size.Int(); ok && i != 10 && i != 0 {
		data["size"] = s.size.Bytes()
	}
	if s.source != nil {
		b, err := json.Marshal(s.source)
		if err != nil {
			return nil, err
		}
		data["_source"] = b
	}
	if len(s.stats) > 0 {
		b, err := json.Marshal(s.stats)
		if err != nil {
			return nil, err
		}
		data["stats"] = b
	}
	if s.terminateAfter > 0 {
		n, _ := dynamic.NewNumber(s.terminateAfter)
		data["terminate_after"] = n.Bytes()
	}
	if s.timeout > 0 {
		data["timeout"] = append([]byte{'"'}, append([]byte(s.timeout.String()), '"')...)
	}
	if s.version {
		data["version"] = trueBytes
	}
	return json.Marshal(data)
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
func (s Search) DocValueFields() Fields {
	if s.docValueFields == nil {
		s.docValueFields = Fields{}
	}
	return s.docValueFields
}

// SetDocValueFields sets DocValueFieldsValue to v
func (s *Search) SetDocValueFields(v Fields) *Search {
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
		s.query = &Query{}
	} else {
		s.query = v
	}
	return s
}

// Query defines the search definition using the Query DSL.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
func (s *Search) Query() *Query {
	if s.query == nil {
		s.query = &Query{}
	}
	return s.query
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
