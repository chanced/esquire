package picker

type QueryStringQueryParams struct {
	// (Required, string) Query string you wish to parse and use for search. See
	// Query string syntax.
	Query string `json:"query"`
	// (Optional, string) Default field you wish to search if no field is
	// provided in the query string.
	//
	// Defaults to the index.query.default_field index setting, which has a
	// default value of *. The * value extracts all fields that are eligible for
	// term queries and filters the metadata fields. All extracted fields are
	// then combined to build a query if no prefix is specified.
	//
	// Searching across all eligible fields does not include nested documents.
	// Use a nested query to search those documents.
	//
	// For mappings with a large number of fields, searching across all eligible
	// fields could be expensive.
	//
	// There is a limit on the number of fields that can be queried at once. It
	// is defined by the indices.query.bool.max_clause_count search setting,
	// which defaults to 1024.
	DefaultField string `json:"default_field,omitempty"`
	// (Optional, Boolean) If true, the wildcard characters * and ? are allowed
	// as the first character of the query string. Defaults to true.
	AllowLeadingWildcard interface{} `json:"allow_leading_wildcard,omitempty"`
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
	DefaultOperator string `json:"default_operator,omitempty"`
	// (Optional, Boolean) If true, enable position increments in queries
	// constructed from a query_string search. Defaults to true.
	EnablePositionIncrements interface{} `json:"enable_position_increments,omitempty"`

	// (Optional, array of strings) Array of fields you wish to search.
	//
	// You can use this parameter query to search across multiple fields. See
	// Search multiple fields.
	Fields []string `json:"fields,omitempty"`
	// (Optional, string) Maximum edit distance allowed for matching. See
	// Fuzziness for valid values and more information.
	Fuzziness string `json:"fuzziness"`
	// (Optional, integer) Maximum number of terms to which the query expands
	// for fuzzy matching. Defaults to 50.
	FuzzyMaxExpansions interface{} `json:"fuzzy_max_expansions,omitempty"`
	// (Optional, Boolean) If true, edits for fuzzy matching include
	// transpositions of two adjacent characters (ab → ba). Defaults to true.
	FuzzyTranspositions interface{} `json:"fuzzy_transpositions,omitempty"`
	// (Optional, Boolean) If true, format-based errors, such as providing a
	// text value for a numeric field, are ignored. Defaults to false.
	Lenient interface{} `json:"lenient,omitempty"`
	// (Optional, integer) Maximum number of automaton states required for the query. Default is 10000.
	//
	// Elasticsearch uses Apache Lucene internally to parse regular expressions.
	// Lucene converts each regular expression to a finite automaton containing a
	// number of determinized states.
	//
	// You can use this parameter to prevent that conversion from unintentionally
	// consuming too many resources. You may need to increase this limit to run
	// complex regular expressions.
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

}
