package mapping

// FieldWithAnalyzer is a Field mapping with an analyzer
//
// Analyzer
//
// The analyzer parameter specifies the analyzer used for text analysis when
// indexing or searching a text field.
//
// Only text fields support the analyzer mapping parameter.
//
// Search Quote Analyzer
//
// The search_quote_analyzer setting allows you to specify an analyzer for
// phrases, this is particularly useful when dealing with disabling stop words
// for phrase queries.
//
// To disable stop words for phrases a field utilising three analyzer settings
// will be required:
//
// 1. An analyzer setting for indexing all terms including stop words
//
// 2. A search_analyzer setting for non-phrase queries that will remove stop
// words
//
// 3. A search_quote_analyzer setting for phrase queries that will not remove
// stop words
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analyzer.html
type FieldWithAnalyzer interface {
	// Analyzer parameter specifies the analyzer used for text analysis when
	// indexing or searching a text field.
	Analyzer() string
	// SetAnalyzer sets Analyzer to v
	SetAnalyzer(v string)
	// SearchAnalyzer overrides Analyzer for search analysis
	SearchAnalyzer() string
	// SetSearchAnalyzer sets SearchAnalyzer to v
	SetSearchAnalyzer(v string)
	// SearchQuoteAnalyzer setting allows you to specify an analyzer for
	// phrases, this is particularly useful when dealing with disabling stop
	// words for phrase queries.
	SearchQuoteAnalyzer() string
	// SetSearchQuoteAnalyzer sets SearchQuoteAnalyzer to v
	SetSearchQuoteAnalyzer(v string)
}

// AnalyzerParam adds Analyzer, SearchAnalyzer, and SearchQuoteAnalyzer
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analyzer.html
type AnalyzerParam struct {
	AnalyzerValue            string `json:"analyzer,omitempty" bson:"analyzer,omitempty"`
	SearchAnalyzerValue      string `json:"search_analyzer,omitempty" bson:"search_analyzer,omitempty"`
	SearchQuoteAnalyzerValue string `json:"search_quote_analyzer,omitempty" bson:"search_quote_analyzer,omitempty"`
}

// Analyzer parameter specifies the analyzer used for text analysis when
// indexing or searching a text field.
//
// Unless overridden with the search_analyzer mapping parameter, this
// analyzer is used for both index and search analysis.
func (ap AnalyzerParam) Analyzer() string {
	return ap.AnalyzerValue
}

// SetAnalyzer sets Analyzer to v
func (ap *AnalyzerParam) SetAnalyzer(v string) {
	ap.AnalyzerValue = v
}

// SearchAnalyzer overrides Analyzer for search analysis
func (ap AnalyzerParam) SearchAnalyzer() string {
	return ap.SearchAnalyzerValue
}

// SetSearchAnalyzer sets SearchAnalyzer to v
func (ap *AnalyzerParam) SetSearchAnalyzer(v string) {
	ap.SearchAnalyzerValue = v
}

// SearchQuoteAnalyzer setting allows you to specify an analyzer for
// phrases, this is particularly useful when dealing with disabling
// stop words for phrase queries.
func (ap AnalyzerParam) SearchQuoteAnalyzer() string {
	return ap.SearchQuoteAnalyzerValue
}

// SetSearchQuoteAnalyzer sets SearchQuoteAnalyzer to v
func (ap AnalyzerParam) SetSearchQuoteAnalyzer(v string) {
	ap.SearchQuoteAnalyzerValue = v
}
