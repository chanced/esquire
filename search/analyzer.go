package search

import "github.com/tidwall/gjson"

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

// AnalyzerParam is a query mixin that adds the analyzer param
//
// Analyzer used to convert the text in the query value into tokens.
// Defaults to the index-time analyzer mapped for the <field>. If no
// analyzer is mapped, the index’s default analyzer is used. (Optional)
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis.html
type AnalyzerParam struct {
	AnalyzerValue string `json:"analyzer,omitempty" bson:"analyzer,omitempty"`
}

// Analyzer used to convert the text in the query value into tokens.
// Defaults to the index-time analyzer mapped for the <field>. If no
// analyzer is mapped, the index’s default analyzer is used. (Optional)
func (a AnalyzerParam) Analyzer() string {
	return a.AnalyzerValue
}

// SetAnalyzer sets the Analyzer value to v
func (a *AnalyzerParam) SetAnalyzer(v string) {
	if a.Analyzer() != v {
		a.AnalyzerValue = v
	}
}

func marshalAnalyzerParam(data map[string]interface{}, source interface{}) (map[string]interface{}, error) {
	if a, ok := source.(WithAnalyzer); ok {
		if a.Analyzer() != DefaultAnalyzer {
			data[paramAnalyzer] = a.Analyzer()
		}
	}
	return data, nil
}
func unmarshalAnalyzerParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithAnalyzer); ok {
		a.SetAnalyzer(value.Str)
	}
	return nil
}
