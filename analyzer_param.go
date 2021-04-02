package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// WithAnalyzer is a Field mapping with an analyzer
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
type WithAnalyzer interface {
	// Analyzer parameter specifies the analyzer used for text analysis when
	// indexing or searching a text field.
	Analyzer() string
	// SetAnalyzer sets Analyzer to v
	SetAnalyzer(v string)
}

// analyzerParam adds Analyzer, SearchAnalyzer, and SearchQuoteAnalyzer
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analyzer.html
type analyzerParam struct {
	analyzer string
}

// Analyzer parameter specifies the analyzer used for text analysis when
// indexing or searching a text field.
//
// Unless overridden with the search_analyzer mapping parameter, this
// analyzer is used for both index and search analysis.
func (ap analyzerParam) Analyzer() string {
	return ap.analyzer
}

// SetAnalyzer sets Analyzer to v
func (ap *analyzerParam) SetAnalyzer(v string) {
	if ap.Analyzer() != v {
		ap.analyzer = v

	}
}

func marshalAnalyzerParam(source interface{}) (dynamic.JSON, error) {
	if a, ok := source.(WithAnalyzer); ok {
		if len(a.Analyzer()) > 0 {
			return json.Marshal(a.Analyzer())
		}
	}
	return nil, nil
}
func unmarshalAnalyzerParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithAnalyzer); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			var str string
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}
			a.SetAnalyzer(str)
			return nil
		}
		return &json.UnmarshalTypeError{Value: string(data), Type: typeString}
	}
	return nil
}
