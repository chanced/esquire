package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultAnalyzer = ""

// WithAnalyzer is a query with the analyzer param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis.html
type WithAnalyzer interface {
	// Analyzer used to convert the text in the query value into tokens.
	// Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used. (Optional)
	Analyzer() string
	// SetAnalyzer sets the Analyzer value to v
	SetAnalyzer(v string)
}

// analyzerParam is a query mixin that adds the analyzer param
//
// Analyzer used to convert the text in the query value into tokens.
// Defaults to the index-time analyzer mapped for the <field>. If no
// analyzer is mapped, the index’s default analyzer is used. (Optional)
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis.html
type analyzerParam struct {
	analyzer string
}

// Analyzer used to convert the text in the query value into tokens.
// Defaults to the index-time analyzer mapped for the <field>. If no
// analyzer is mapped, the index’s default analyzer is used. (Optional)
func (a analyzerParam) Analyzer() string {
	return a.analyzer
}

// SetAnalyzer sets the Analyzer value to v
func (a *analyzerParam) SetAnalyzer(v string) {
	if a.Analyzer() != v {
		a.analyzer = v
	}
}

func marshalAnalyzerParam(data M, source interface{}) (M, error) {
	if a, ok := source.(WithAnalyzer); ok {
		if a.Analyzer() != DefaultAnalyzer {
			data[paramAnalyzer] = a.Analyzer()
		}
	}
	return data, nil
}
func unmarshalAnalyzerParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithAnalyzer); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			a.SetAnalyzer(data.UnquotedString())
		}
		return &json.UnmarshalTypeError{Value: data.String()}
	}
	return nil
}
