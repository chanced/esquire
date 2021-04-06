package picker

type SimpleQueryStringer interface {
	SimpleQueryString() (*SimpleQueryStringQuery, error)
}

type SimpleQueryStringQueryParams struct {
	// (Required, string) Query string you wish to parse and use for search. See
	// Query string syntax.
	Query string `json:"query"`
	// (Optional, string) List of enabled operators for the simple query string
	// syntax. Defaults to ALL (all operators).
	Flags string `json:"flags,omitempty"`
	// Flags string `json:"default_field,omitempty"`
	// AllowLeadingWildcard interface{} `json:"allow_leading_wildcard,omitempty"`

	// (Optional, Boolean) If true, the query attempts to analyze wildcard terms
	// in the query string. Defaults to false.
	AnalyzeWildcard interface{} `json:"analyze_wildcard,omitempty"`
	// (Optional, string) Analyzer used to convert text in the query string into
	// tokens. Defaults to the index-time analyzer mapped for the default_field.
	// If no analyzer is mapped, the index’s default analyzer is used.
	Analyzer string `json:"analyzer,omitempty"`
	// (Optional, Boolean) If true, match phrase queries are automatically
	// created for multi-term synonyms. Defaults to true.
	AutoGenerateSynonymsPhraseQuery interface{} `json:"auto_generate_synonyms_phrase_query,omitempty"`
	// (Optional, float) Floating point number used to decrease or increase the
	// relevance scores of the query. Defaults to 1.0.
	//
	// Boost values are relative to the default value of 1.0. A boost value
	// between 0 and 1.0 decreases the relevance score. A value greater than 1.0
	// increases the relevance score.
	Boost interface{} `json:"boost,omitempty"`
	// (Optional, string) Default boolean logic used to interpret text in the
	// query string if no operators are specified. Valid values are:
	//
	// - OR (Default)
	//
	// - AND
	DefaultOperator Operator `json:"default_operator,omitempty"`

	// (Optional, array of strings) Array of fields you wish to search.
	//
	// You can use this parameter query to search across multiple fields. See
	// Search multiple fields.
	Fields []string `json:"fields,omitempty"`

	FuzzyPrefixLength interface{} `json:"fuzzy_prefix_length,omitempty"`
	// (Optional, integer) Maximum number of terms to which the query expands
	// for fuzzy matching. Defaults to 50.
	FuzzyMaxExpansions interface{} `json:"fuzzy_max_expansions,omitempty"`
	// (Optional, Boolean) If true, edits for fuzzy matching include
	// transpositions of two adjacent characters (ab → ba). Defaults to true.
	FuzzyTranspositions interface{} `json:"fuzzy_transpositions,omitempty"`
	// (Optional, Boolean) If true, format-based errors, such as providing a
	// text value for a numeric field, are ignored. Defaults to false.
	Lenient bool `json:"lenient,omitempty"`
	// (Optional, integer) Maximum number of automaton states required for the
	// query. Default is 10000.
	//
	// Elasticsearch uses Apache Lucene internally to parse regular expressions.
	// Lucene converts each regular expression to a finite automaton containing
	// a number of determinized states.
	//
	// You can use this parameter to prevent that conversion from
	// unintentionally consuming too many resources. You may need to increase
	// this limit to run complex regular expressions.
	MaxDeterminizedStates interface{} `json:"max_determinized_states,omitempty"`
	// (Optional, string) Minimum number of clauses that must match for a
	// document to be returned. See the minimum_should_match parameter for valid
	// values and more information. See How minimum_should_match works for an
	// example.
	MinimumShouldMatch string `json:"minimum_should_match,omitempty"`
	// (Optional, string) Analyzer used to convert quoted text in the query
	// string into tokens. Defaults to the search_quote_analyzer mapped for the
	// default_field.
	//
	// For quoted text, this parameter overrides the analyzer specified in the
	// analyzer parameter.
	QuoteAnalyzer string `json:"quote_analyzer,omitempty"`
	// (Optional, integer) Maximum number of positions allowed between matching
	// tokens for phrases. Defaults to 0. If 0, exact phrase matches are
	// required. Transposed terms have a slop of 2.
	PhraseSlop interface{} `json:"phrase_slop,omitempty"`
	// (Optional, string) Suffix appended to quoted text in the query string.
	//
	// You can use this suffix to use a different analysis method for exact
	// matches. See Mixing exact search with stemming.
	QuoteFieldSuffix string `json:"quote_field_suffix,omitempty"`
	Name             string `json:"_name,omitempty"`
	completeClause
}

func (SimpleQueryStringQueryParams) Kind() QueryKind {
	return QueryKindSimpleQueryString
}
func (p SimpleQueryStringQueryParams) Clause() (QueryClause, error) {
	return p.SimpleQueryString()
}
func (p SimpleQueryStringQueryParams) SimpleQueryString() (*SimpleQueryStringQuery, error) {
	q := &SimpleQueryStringQuery{}
	var err error
	q.SetAnalyzer(p.Analyzer)
	q.SetLenient(p.Lenient)
	q.SetFlags(p.Flags)
	q.SetFields(p.Fields)
	q.SetName(p.Name)
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetQuoteAnalyzer(p.QuoteAnalyzer)
	q.SetQuoteFieldSuffix(p.QuoteFieldSuffix)
	err = q.SetFuzzyPrefixLength(p.FuzzyPrefixLength)
	if err != nil {
		return q, err
	}
	err = q.SetAnalyzeWildcard(p.AnalyzeWildcard)
	if err != nil {
		return q, err
	}
	err = q.SetAutoGenerateSynonymsPhraseQuery(p.AutoGenerateSynonymsPhraseQuery)
	if err != nil {
		return q, err
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	err = q.SetDefaultOperator(p.DefaultOperator)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyMaxExpansions(p.FuzzyMaxExpansions)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyTranspositions(p.FuzzyTranspositions)
	if err != nil {
		return q, err
	}
	err = q.SetMaxDeterminizedStates(p.MaxDeterminizedStates)
	if err != nil {
		return q, err
	}
	err = q.SetPhraseSlop(p.PhraseSlop)
	if err != nil {
		return q, err
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	return q, nil
}

type SimpleQueryStringQuery struct {
	query  string
	flags  string
	fields []string
	quoteFieldSuffixParam
	quoteAnalyzerParam
	phraseSlopParam
	maxDeterminizedStatesParam
	fuzzyMaxExpansionsParam
	analyzeWildcardParam
	fuzzyPrefixLengthParam
	defaultOperatorParam
	fuzzyTranspositionsParam
	lenientParam
	minimumShouldMatchParam
	analyzerParam
	autoGenerateSynonymsPhraseQueryParam
	boostParam
	completeClause
	nameParam
}

func (qs *SimpleQueryStringQuery) Clause() (QueryClause, error) {
	return qs, nil
}
func (qs *SimpleQueryStringQuery) SimpleQueryString() (*SimpleQueryStringQuery, error) {
	return qs, nil
}
func (qs *SimpleQueryStringQuery) IsEmpty() bool {
	return qs == nil || len(qs.query) == 0
}
func (qs *SimpleQueryStringQuery) Clear() {
	if qs == nil {
		return
	}
	*qs = SimpleQueryStringQuery{}
}

func (SimpleQueryStringQuery) Kind() QueryKind {
	return QueryKindSimpleQueryString
}
func (qs SimpleQueryStringQuery) Query() string {
	return qs.query
}
func (qs *SimpleQueryStringQuery) SetQuery(q string) error {
	if len(q) == 0 {
		return ErrQueryRequired
	}
	qs.query = q
	return nil
}
func (qs SimpleQueryStringQuery) Flags() string {
	return qs.flags
}
func (qs *SimpleQueryStringQuery) SetFlags(field string) {
	qs.flags = field
}
func (qs SimpleQueryStringQuery) Fields() []string {
	return qs.fields
}
func (qs *SimpleQueryStringQuery) SetFields(fields []string) {
	qs.fields = fields
}

func (qs *SimpleQueryStringQuery) UnmarshalJSON(data []byte) error {
	q := simpleQueryStringQuery{}
	err := q.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	v, err := q.SimpleQueryString()
	if err != nil {
		return err
	}
	*qs = v
	return nil
}

func (qs SimpleQueryStringQuery) MarshalJSON() ([]byte, error) {
	return simpleQueryStringQuery{
		Query:                           qs.query,
		Flags:                           qs.flags,
		AnalyzeWildcard:                 qs.analyzeWildcard.Value(),
		Analyzer:                        qs.analyzer,
		AutoGenerateSynonymsPhraseQuery: qs.autoGenerateSynonymsPhraseQuery.Value(),
		Boost:                           qs.boost.Value(),
		DefaultOperator:                 qs.defaultOperator,
		Fields:                          qs.fields,
		FuzzyMaxExpansions:              qs.fuzzyMaxExpansions.Value(),
		FuzzyTranspositions:             qs.fuzzyTranspositions.Value(),
		Lenient:                         qs.lenient,
		MaxDeterminizedStates:           qs.maxDeterminizedStates.Value(),
		MinimumShouldMatch:              qs.minimumShouldMatch,
		QuoteAnalyzer:                   qs.quoteAnalyzer,
		PhraseSlop:                      qs.phraseSlop.Value(),
		QuoteFieldSuffix:                qs.quoteFieldSuffix,
		Name:                            qs.name,
	}.MarshalJSON()
}

//easyjson:json
type simpleQueryStringQuery struct {
	Query                           string      `json:"query"`
	Flags                           string      `json:"flags,omitempty"`
	AnalyzeWildcard                 interface{} `json:"analyze_wildcard,omitempty"`
	Analyzer                        string      `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery interface{} `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           interface{} `json:"boost,omitempty"`
	DefaultOperator                 Operator    `json:"default_operator,omitempty"`
	Fields                          []string    `json:"fields,omitempty"`
	FuzzyPrefixLength               interface{} `json:"fuzzy_prefix_length,omitempty"`
	FuzzyMaxExpansions              interface{} `json:"fuzzy_max_expansions,omitempty"`
	FuzzyTranspositions             interface{} `json:"fuzzy_transpositions,omitempty"`
	Lenient                         bool        `json:"lenient,omitempty"`
	MaxDeterminizedStates           interface{} `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch              string      `json:"minimum_should_match,omitempty"`
	QuoteAnalyzer                   string      `json:"quote_analyzer,omitempty"`
	PhraseSlop                      interface{} `json:"phrase_slop,omitempty"`
	QuoteFieldSuffix                string      `json:"quote_field_suffix,omitempty"`
	Name                            string      `json:"_name,omitempty"`
}

func (p simpleQueryStringQuery) SimpleQueryString() (SimpleQueryStringQuery, error) {
	q := SimpleQueryStringQuery{}
	var err error
	q.SetAnalyzer(p.Analyzer)
	q.SetLenient(p.Lenient)
	q.SetFlags(p.Flags)
	q.SetFields(p.Fields)
	q.SetName(p.Name)
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetQuoteAnalyzer(p.QuoteAnalyzer)
	q.SetQuoteFieldSuffix(p.QuoteFieldSuffix)
	err = q.SetAnalyzeWildcard(p.AnalyzeWildcard)
	if err != nil {
		return q, err
	}
	err = q.SetAutoGenerateSynonymsPhraseQuery(p.AutoGenerateSynonymsPhraseQuery)
	if err != nil {
		return q, err
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	err = q.SetDefaultOperator(p.DefaultOperator)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyMaxExpansions(p.FuzzyMaxExpansions)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyTranspositions(p.FuzzyTranspositions)
	if err != nil {
		return q, err
	}
	err = q.SetMaxDeterminizedStates(p.MaxDeterminizedStates)
	if err != nil {
		return q, err
	}
	err = q.SetPhraseSlop(p.PhraseSlop)
	if err != nil {
		return q, err
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	return q, nil
}
