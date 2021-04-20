package picker

import (
	"bytes"
	"time"

	"encoding/json"

	"github.com/chanced/dynamic"
)

var (
	DefaultExplain = false
	DefaultFrom    = int(0)
	DefaultSize    = int(10)
)

// SearchParams are the initial params passed to NewSearch
type SearchParams struct {
	Query QueryParams

	Aggregations map[string]interface{}
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
	DocValueFields SearchFields

	// Array of wildcard (*) patterns. The request returns values for field
	// names matching these patterns in the hits.fields property of the response
	// (Optional).
	//
	// You can specify items in the array as a string or object.
	Fields SearchFields

	// If true, returns detailed information about score computation as part of
	// a hit. Defaults to false. (Optional)
	Explain bool

	// Starting document offset. Defaults to 0.
	//
	// By default, you cannot page through more than 10,000 hits using the from
	// and size parameters. To page through more hits, use the search_after
	// parameter. (Optional)
	From int

	// Boosts the _score of documents from specified indices (Optional).
	IndicesBoost map[string]float64

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
	Size int

	// Indicates which source fields are returned for matching documents. These
	// fields are returned in the hits._source property of the search response.
	// Defaults to true. (Optional)
	Source interface{}

	// Stats groups to associate with the picker. Each group maintains a
	// statistics aggregation for its associated searches. You can retrieve
	// these stats using the indices stats API (Optional).
	Stats []string

	// The maximum number of documents to collect for each shard, upon reaching
	// which the query execution will terminate early. (Optional)
	//
	// Defaults to 0, which does not terminate query execution early.
	TerminateAfter int

	// Specifies the period of time to wait for a response. If no response is
	// received before the timeout expires, the request fails and returns an
	// error. Defaults to no timeout. (Optional)
	Timeout time.Duration

	// If true, returns document version as part of a hit. Defaults to false. (Optional)
	Version bool

	Highlight interface{}
}

func NewSearch(p SearchParams) (*Search, error) {

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
		stats:            p.Stats,
		terminateAfter:   p.TerminateAfter,
		timeout:          p.Timeout,
		version:          p.Version,
		source:           p.Source,
	}
	if p.Size != 0 {
		err := s.SetSize(p.Size)
		if err != nil {
			return s, err
		}
	}

	q, err := p.Query.Query()
	if err != nil {
		return nil, err
	}
	s.query = q
	return s, nil
}

type Search struct {
	// Defines the search definition using the Query DSL. (Optional)
	query            *Query // query
	aggregations     map[string]interface{}
	docValueFields   SearchFields       // docvalue_fields
	fields           SearchFields       // fields
	explain          bool               // explain
	from             int                // from
	indicesBoost     map[string]float64 // indices_boost
	minScore         float64            // min_score
	pointInTime      *PointInTime       // pit
	runtimeMappings  RuntimeMappings    // runtime_mappings
	seqNoPrimaryTerm bool               // seq_no_primary_term
	size             dynamic.Number     // size
	source           interface{}        // _source
	stats            []string           // stats
	terminateAfter   int                // terminate_after
	timeout          time.Duration      // timeout
	version          bool               // version
}

var zeroSearch = &Search{}

func (s *Search) UnmarshalBSON(data []byte) error {
	return s.UnmarshalJSON(data)
}

func (s *Search) UnmarshalJSON(data []byte) (err error) {
	*s = *zeroSearch
	var m map[string]dynamic.JSON
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	// for k, d := range m {
	// 	switch k {
	// 	case "query":
	// 		var q Query
	// 		err = json.Unmarshal(d, &q)
	// 		s.query = &q
	// 	case "aggs":
	// 		var a map[string]interface{}
	// 		err = json.Unmarshal(d, &a)
	// 		s.aggregations = a
	// 	case "docvalue_fields":
	// 		var df SearchFields
	// 		err = json.Unmarshal(d, &df)
	// 		s.docValueFields = df
	// 	case "fields":
	// 		var f SearchFields
	// 		err = json.Unmarshal(d, &f)
	// 		s.fields = f
	// 	case "explain":
	// 		var b dynamic.Bool
	// 		b, err = dynamic.NewBool(d.UnquotedString())
	// 		if v, ok := b.Bool(); ok {
	// 			s.explain = v
	// 		}
	// 	case "from":
	// 		var i int
	// 		err = json.Unmarshal(d, &i)
	// 		s.from = i
	// 	case "indices_boost":
	// 		var ib map[string]float64
	// 		err = json.Unmarshal(d, &ib)
	// 		s.indicesBoost = ib
	// 	case "min_score":
	// 		var i float64
	// 		err = json.Unmarshal(d, &i)
	// 		s.minScore = i
	// 	case "pit":

	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	if d, ok := m["query"]; ok {
		q := Query{}
		err = q.UnmarshalJSON(d)
		if err != nil {
			return err
		}
		s.query = &q
	}

	if d, ok := m["aggs"]; ok {
		var a map[string]interface{}
		err := json.Unmarshal(d, &a)
		if err != nil {
			return err
		}
		s.aggregations = a
	}

	if d, ok := m["docvalue_fields"]; ok {
		var df SearchFields
		err = json.Unmarshal(d, &df)
		if err != nil {
			return err
		}
		s.docValueFields = df
	}
	if d, ok := m["fields"]; ok {
		var f SearchFields
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
		var i int
		err := json.Unmarshal(d, &i)
		if err != nil {
			return err
		}
		s.from = i
	}
	if d, ok := m["indices_boost"]; ok {
		var ib map[string]float64
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
		var v interface{}
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
		var i int
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

func (s Search) MarshalBSON() ([]byte, error) {
	return s.MarshalJSON()
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
	if len(s.aggregations) > 0 {
		aggs, err := json.Marshal(s.aggregations)
		if err != nil {
			return nil, err
		}
		data["aggs"] = aggs
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
func (s Search) DocValueFields() SearchFields {
	if s.docValueFields == nil {
		s.docValueFields = SearchFields{}
	}
	return s.docValueFields
}

// SetDocValueFields sets DocValueFieldsValue to v
func (s *Search) SetDocValueFields(v SearchFields) *Search {
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
func (s *Search) Fields() SearchFields {
	if s.fields == nil {
		s.fields = SearchFields{}
	}
	return s.fields

}

// SetFields sets the FieldsValue to v
func (s *Search) SetFields(v SearchFields) {
	s.fields = v
}

// Explain indicates whether the search returns detailed information about score
// computation as part of a hit. Defaults to false.
func (s Search) Explain() bool {
	return s.explain
}

// SetExplain sets the ExplainValue to v
func (s *Search) SetExplain(v bool) {
	s.explain = v
}

// From sets the starting document offset. Defaults to 0.
//
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters. To page through more hits, use the search_after parameter.
func (s Search) From() int {
	return s.from
}

func (s Search) Encode() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(s)
	return buf, err
}

// SetFrom sets the FromValue to v
func (s *Search) SetFrom(v int) {
	s.from = v
}

// IndicesBoost buusts the _score of documents from specified indices
func (s *Search) IndicesBoost() map[string]float64 {
	if s.indicesBoost == nil {
		s.indicesBoost = map[string]float64{}
	}
	return s.indicesBoost
}

// SetIndicesBoost sets IndicesBoostValue to v
func (s *Search) SetIndicesBoost(v map[string]float64) {
	s.indicesBoost = v
}

// MinScore is the minimum _score for matching documents. Documents with a lower
// _score are not included in the search results.
func (s *Search) MinScore() float64 {
	return s.minScore
}

// SetMinScore sets the MinScoreValue to v
func (s *Search) SetMinScore(v float64) {
	s.minScore = v
}

// SetPointInTime sets the PointInTimeValue to v
func (s *Search) SetPointInTime(v *PointInTime) {
	s.pointInTime = v

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

// SetQuery sets QueryValue to v
func (s *Search) SetQuery(v Querier) error {
	if v == nil {
		s.query = &Query{}
	}
	q, err := v.Query()
	if err != nil {
		return err
	}
	s.query = q
	return nil
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

func (s *Search) SetRuntimeMappings(v RuntimeMappings) {
	s.runtimeMappings = v
}

// SeqNoPrimaryTerm https://www.elastic.co/guide/en/elasticsearch/reference/current/optimistic-concurrency-control.html
func (s Search) SeqNoPrimaryTerm() bool {
	return s.seqNoPrimaryTerm
}

func (s *Search) SetSeqNoPrimaryTerm(v bool) {
	s.seqNoPrimaryTerm = v
}

// Size is number of hits to return. Defaults to 10.
//
// By default, you cannot page through more than 10,000 hits using the from and
// size parameters. To page through more hits, use the search_after parameter.
func (s Search) Size() int {
	if i, ok := s.size.Int(); ok {
		return int(i)
	}
	return DefaultSize
}

func (s *Search) SetSize(v int) error {
	return s.size.Set(v)
}

// Source indicates which source fields are returned for matching documents.
// These fields are returned in the hits._source property of the search
// response. Defaults to true.
func (s Search) Source() interface{} {
	return s.source
}

// SetSource sets the value of Source
func (s *Search) SetSource(v interface{}) error {
	//
	// The options are:
	//  picker.Source, *picker.Source,
	//  string, []string,
	//  dynamic.StringOrArrayOfStrings, *dynamic.StringOrArrayOfStrings,
	//  picker.SourceSpecifications
	//  *picker.SourceSpecifications
	//  bool, *bool
	//  nil
	// Note, "true" || "false" get parsed as boolean
	s.source = v
	return nil
	// switch t := v.(type) {
	// case *SearchSource:
	// 	ts := *t
	// 	s.source = &ts
	// 	return nil
	// case SearchSource:
	// 	s.source = &t
	// 	return nil
	// default:
	// 	s.source = &SearchSource{}
	// 	return s.source.SetValue(v)
	// }
}

// Stats groups to associate with the picker. Each group maintains a statistics
// aggregation for its associated searches. You can retrieve these stats using
// the indices stats API (Optional).
func (s Search) Stats() []string {
	return s.stats
}

func (s *Search) SetStats(v []string) {
	s.stats = v

}

// TerminateAfter is maximum number of documents to collect for each shard, upon
// reaching which the query execution will terminate early.
//
// Defaults to 0, which does not terminate query execution early.
func (s Search) TerminateAfter() int {
	return s.terminateAfter
}

func (s *Search) SetTerminateAfter(v int) {
	s.terminateAfter = v

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
