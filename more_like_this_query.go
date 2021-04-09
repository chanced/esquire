package picker

import "github.com/chanced/dynamic"

const (
	DefaultMLTMaxQueryTerms          = 25
	DefaultMLTMinTermFreq            = 2
	DefaultMLTMinDocFreq             = 5
	DefaultMLTMinWordLen             = 0
	DefaultMLTMinimumShouldMatch     = "30%"
	DefaultMLTMaxDocFreq             = 2147483647
	DefaultMLTFailOnUnsupportedField = true
	DefaultMLTBoostTerms             = 0
	DefaultMLTInclude                = false
	DefaultMLTMinWordLength          = 0
	DefaultMLTMaxWordLength          = 0
)

type MoreLikeThiser interface {
	MoreLikeThis() (*MoreLikeThisQuery, error)
}

type MoreLikeThisQueryParams struct {
	Name string
	completeClause
	// The only required parameter of the MLT query is like and follows a
	// versatile syntax, in which the user can specify free form text and/or a
	// single or multiple documents (see examples above). The syntax to specify
	// documents is similar to the one used by the Multi GET API. When
	// specifying documents, the text is fetched from fields unless overridden
	// in each document request. The text is analyzed by the analyzer at the
	// field, but could also be overridden. The syntax to override the analyzer
	// at the field follows a similar syntax to the per_field_analyzer parameter
	// of the Term Vectors API. Additionally, to provide documents not
	// necessarily present in the index, artificial documents are also
	// supported.
	Like interface{}
	// The unlike parameter is used in conjunction with like in order not to
	// select terms found in a chosen set of documents. In other words, we could
	// ask for documents like: "Apple", but unlike: "cake crumble tree". The
	// syntax is the same as like.
	Unlike interface{}
	// A list of fields to fetch and analyze the text from. Defaults to the
	// index.query.default_field index setting, which has a default value of *.
	// The * value matches all fields eligible for term-level queries, excluding
	// metadata fields.
	Fields []string
	// The maximum number of query terms that will be selected. Increasing this
	// value gives greater accuracy at the expense of query execution speed.
	// Defaults to 25.
	MaxQueryTerms interface{}
	// The minimum term frequency below which the terms will be ignored from the
	// input document. Defaults to 2.
	MinTermFrequency interface{}
	// The minimum document frequency below which the terms will be ignored from
	// the input document. Defaults to 5.
	MinDocFrequency interface{}
	// The maximum document frequency above which the terms will be ignored from
	// the input document. This could be useful in order to ignore highly
	// frequent words such as stop words. Defaults to unbounded
	// (Integer.MAX_VALUE, which is 2^31-1 or 2147483647).
	MaxDocFrequency interface{}
	// The minimum word length below which the terms will be ignored. Defaults
	// to 0.
	MinWordLength interface{}
	// The maximum word length above which the terms will be ignored. Defaults
	// to unbounded (0).
	MaxWordLength interface{}
	// An array of stop words. Any word in this set is considered
	// "uninteresting" and ignored. If the analyzer allows for stop words, you
	// might want to tell MLT to explicitly ignore them, as for the purposes of
	// document similarity it seems reasonable to assume that "a stop word is
	// never interesting".
	StopWords []string
	// The analyzer that is used to analyze the free form text. Defaults to the
	// analyzer associated with the first field in fields.
	Analyzer string

	// After the disjunctive query has been formed, this parameter controls the
	// number of terms that must match. The syntax is the same as the minimum
	// should match. (Defaults to "30%").
	MinimumShouldMatch string
	// Controls whether the query should fail (throw an exception) if any of the
	// specified fields are not of the supported types (text or keyword). Set
	// this to false to ignore the field and continue processing. Defaults to
	// true.
	FailOnUnsupportedField interface{}
	// Each term in the formed query could be further boosted by their tf-idf
	// score. This sets the boost factor to use when using this feature.
	// Defaults to deactivated (0). Any other positive value activates terms
	// boosting with the given boost factor.
	BoostTerms interface{}
	// Specifies whether the input documents should also be included in the
	// search results returned. Defaults to false.
	Include interface{}
	// Sets the boost value of the whole query. Defaults to 1.0.
	Boost interface{}
}

func (MoreLikeThisQueryParams) Kind() QueryKind {
	return QueryKindMoreLikeThis
}

func (p MoreLikeThisQueryParams) Clause() (QueryClause, error) {
	return p.MoreLikeThis()
}
func (p MoreLikeThisQueryParams) MoreLikeThis() (*MoreLikeThisQuery, error) {
	q := &MoreLikeThisQuery{}
	q.SetFields(p.Fields)
	err := q.SetLike(p.Like)
	if err != nil {
		return q, err
	}
	q.SetAnalyzer(p.Analyzer)
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	err = q.SetBoostTerms(p.BoostTerms)
	if err != nil {
		return q, err
	}
	err = q.SetFailOnUnsupportedField(p.FailOnUnsupportedField)
	if err != nil {
		return q, err
	}
	err = q.SetInclude(p.Include)
	if err != nil {
		return q, err
	}
	q.SetUnlike(p.Unlike)
	err = q.SetMaxDocFrequency(p.MaxDocFrequency)
	if err != nil {
		return q, err
	}
	err = q.SetMaxQueryTerms(p.MaxQueryTerms)
	if err != nil {
		return q, err
	}
	err = q.SetMaxWordLength(p.MaxWordLength)
	if err != nil {
		return q, err
	}
	err = q.SetMinDocFrequency(p.MinDocFrequency)
	if err != nil {
		return q, err
	}
	err = q.SetMinTermFrequency(p.MinTermFrequency)
	if err != nil {
		return q, err
	}
	err = q.SetMinWordLength(p.MinWordLength)
	if err != nil {
		return q, err
	}
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetName(p.Name)
	q.SetStopWords(p.StopWords)

	return q, nil
}

type MoreLikeThisQuery struct {
	nameParam
	like      interface{}
	unlike    interface{}
	fields    []string
	stopWords []string
	completeClause
	analyzerParam
	boostParam
	minimumShouldMatchParam
	maxWordLength          dynamic.Number
	minWordLength          dynamic.Number
	maxDocFrequency        dynamic.Number
	minDocFrequency        dynamic.Number
	minTermFrequency       dynamic.Number
	boostTerms             dynamic.Number
	maxQueryTerms          dynamic.Number
	include                dynamic.Bool
	failOnUnsupportedField dynamic.Bool
}

func (q MoreLikeThisQuery) MaxWordLength() int {
	if i, ok := q.maxWordLength.Int(); ok {
		return i
	}
	if f, ok := q.maxWordLength.Float64(); ok {
		return int(f)
	}
	return DefaultMLTMaxWordLength
}
func (q *MoreLikeThisQuery) SetStopWords(stopWords []string) {
	q.stopWords = stopWords
}
func (q MoreLikeThisQuery) StopWords() []string {
	return q.stopWords
}
func (q *MoreLikeThisQuery) SetMaxWordLength(v interface{}) error {
	return q.maxWordLength.Set(v)
}
func (q MoreLikeThisQuery) MinWordLength() int {
	if i, ok := q.minWordLength.Int(); ok {
		return i
	}
	if f, ok := q.minWordLength.Float64(); ok {
		return int(f)
	}
	return DefaultMLTMinWordLength
}

func (q *MoreLikeThisQuery) SetMinWordLength(v interface{}) error {
	return q.minWordLength.Set(v)
}

func (q MoreLikeThisQuery) MaxDocFrequency() float64 {
	if f, ok := q.maxDocFrequency.Float64(); ok {
		return f
	}
	return DefaultMLTMaxDocFreq
}
func (q *MoreLikeThisQuery) SetMaxDocFrequency(v interface{}) error {
	return q.maxDocFrequency.Set(v)
}
func (q MoreLikeThisQuery) MinDocFrequency() float64 {
	if f, ok := q.minDocFrequency.Float64(); ok {
		return f
	}
	return DefaultMLTMinDocFreq
}
func (q *MoreLikeThisQuery) SetMinDocFrequency(v interface{}) error {
	return q.minDocFrequency.Set(v)
}
func (q MoreLikeThisQuery) Include() bool {
	if b, ok := q.include.Bool(); ok {
		return b
	}
	return DefaultMLTInclude
}
func (q *MoreLikeThisQuery) SetInclude(v interface{}) error {
	return q.include.Set(v)
}
func (q MoreLikeThisQuery) FailOnUnsupportedField() bool {
	if b, ok := q.failOnUnsupportedField.Bool(); ok {
		return b
	}
	return DefaultMLTFailOnUnsupportedField
}
func (q *MoreLikeThisQuery) SetFailOnUnsupportedField(v interface{}) error {
	return q.failOnUnsupportedField.Set(v)
}
func (q MoreLikeThisQuery) BoostTerms() float64 {
	if f, ok := q.boostTerms.Float64(); ok {
		return f
	}
	return DefaultMLTBoostTerms
}
func (q *MoreLikeThisQuery) SetBoostTerms(v interface{}) error {
	return q.boostTerms.Set(v)
}
func (q MoreLikeThisQuery) MinTermFrequency() float64 {
	if f, ok := q.minTermFrequency.Float64(); ok {
		return f
	}
	return DefaultMLTMinTermFreq
}
func (q *MoreLikeThisQuery) SetMinTermFrequency(v interface{}) error {
	return q.minTermFrequency.Set(v)
}
func (q MoreLikeThisQuery) MaxQueryTerms() int {
	if f, ok := q.maxQueryTerms.Int(); ok {
		return f
	}
	return DefaultMLTMinTermFreq
}
func (q *MoreLikeThisQuery) SetMaxQueryTerms(v interface{}) error {
	return q.maxQueryTerms.Set(v)
}

func (q MoreLikeThisQuery) Like() interface{} {
	return q.like
}
func (q *MoreLikeThisQuery) SetLike(like interface{}) error {
	if like == nil {
		return ErrLikeRequired
	}
	q.like = like
	return nil
}
func (q MoreLikeThisQuery) Unlike() interface{} {
	return q.unlike
}
func (q *MoreLikeThisQuery) SetUnlike(unlike interface{}) {
	q.unlike = unlike
}

func (q MoreLikeThisQuery) Fields() []string {
	return q.fields
}
func (q *MoreLikeThisQuery) SetFields(fields []string) {
	q.fields = fields
}

func (MoreLikeThisQuery) Kind() QueryKind {
	return QueryKindMoreLikeThis
}
func (q *MoreLikeThisQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *MoreLikeThisQuery) MoreLikeThis() (*MoreLikeThisQuery, error) {
	return q, nil
}
func (q *MoreLikeThisQuery) Clear() {
	if q == nil {
		return
	}
	*q = MoreLikeThisQuery{}
}
func (q *MoreLikeThisQuery) UnmarshalJSON(data []byte) error {
	*q = MoreLikeThisQuery{}
	var p moreLikeThisQuery

	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	qv, err := p.MoreLikeThisQuery()
	if err != nil {
		return err
	}
	*q = *qv
	return nil
}
func (q MoreLikeThisQuery) MarshalJSON() ([]byte, error) {
	return moreLikeThisQuery{
		Name:                   q.name,
		Like:                   q.like,
		Unlike:                 q.unlike,
		Fields:                 q.fields,
		MaxQueryTerms:          q.maxQueryTerms.Value(),
		MinTermFrequency:       q.minTermFrequency.Value(),
		MinDocFrequency:        q.minDocFrequency.Value(),
		MaxDocFrequency:        q.maxDocFrequency.Value(),
		MinWordLength:          q.minWordLength.Value(),
		MaxWordLength:          q.maxWordLength.Value(),
		MinimumShouldMatch:     q.minimumShouldMatch,
		StopWords:              q.stopWords,
		Analyzer:               q.analyzer,
		FailOnUnsupportedField: q.failOnUnsupportedField.Value(),
		BoostTerms:             q.boostTerms.Value(),
		Include:                q.include.Value(),
		Boost:                  q.boost.Value(),
	}.MarshalJSON()
}
func (q *MoreLikeThisQuery) IsEmpty() bool {
	return q == nil || q.like == nil
}

//easyjson:json
type moreLikeThisQuery struct {
	Name                   string      `json:"_name,omitempty"`
	Like                   interface{} `json:"like"`
	Unlike                 interface{} `json:"unlike,omitempty"`
	Fields                 []string    `json:"fields,omitempty"`
	MaxQueryTerms          interface{} `json:"max_query_terms,omitempty"`
	MinTermFrequency       interface{} `json:"min_term_freq,omitempty"`
	MinDocFrequency        interface{} `json:"min_doc_freq,omitempty"`
	MaxDocFrequency        interface{} `json:"max_doc_freq,omitempty"`
	MinWordLength          interface{} `json:"min_word_length,omitempty"`
	MaxWordLength          interface{} `json:"max_word_length,omitempty"`
	MinimumShouldMatch     string      `json:"minimum_should_match,omitempty"`
	StopWords              []string    `json:"stop_words,omitempty"`
	Analyzer               string      `json:"analyzer,omitempty"`
	FailOnUnsupportedField interface{} `json:"fail_on_unsupported_field,omitempty"`
	BoostTerms             interface{} `json:"boost_terms,omitempty"`
	Include                interface{} `json:"include,omitempty"`
	Boost                  interface{} `json:"boost,omitempty"`
}

func (p moreLikeThisQuery) MoreLikeThisQuery() (*MoreLikeThisQuery, error) {
	q := &MoreLikeThisQuery{}
	q.SetFields(p.Fields)
	err := q.SetLike(p.Like)
	if err != nil {
		return q, err
	}
	q.SetAnalyzer(p.Analyzer)
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	err = q.SetBoostTerms(p.BoostTerms)
	if err != nil {
		return q, err
	}
	err = q.SetFailOnUnsupportedField(p.FailOnUnsupportedField)
	if err != nil {
		return q, err
	}
	err = q.SetInclude(p.Include)
	if err != nil {
		return q, err
	}
	q.SetUnlike(p.Unlike)
	err = q.SetMaxDocFrequency(p.MaxDocFrequency)
	if err != nil {
		return q, err
	}
	err = q.SetMaxQueryTerms(p.MaxQueryTerms)
	if err != nil {
		return q, err
	}
	err = q.SetMaxWordLength(p.MaxWordLength)
	if err != nil {
		return q, err
	}
	err = q.SetMinDocFrequency(p.MinDocFrequency)
	if err != nil {
		return q, err
	}
	err = q.SetMinTermFrequency(p.MinTermFrequency)
	if err != nil {
		return q, err
	}
	err = q.SetMinWordLength(p.MinWordLength)
	if err != nil {
		return q, err
	}
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetName(p.Name)
	q.SetStopWords(p.StopWords)
	return q, nil
}
