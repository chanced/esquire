package mapping

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type WithSearchQuoteAnalyzer interface {

	// SearchQuoteAnalyzer setting allows you to specify an analyzer for
	// phrases, this is particularly useful when dealing with disabling stop
	// words for phrase queries.
	SearchQuoteAnalyzer() string
	// SetSearchQuoteAnalyzer sets SearchQuoteAnalyzer to v
	SetSearchQuoteAnalyzer(v string)
}

type searchQuoteAnalyzerParam struct {
	SearchQuoteAnalyzerValue string
}

// SearchQuoteAnalyzer setting allows you to specify an analyzer for
// phrases, this is particularly useful when dealing with disabling
// stop words for phrase queries.
func (sq searchQuoteAnalyzerParam) SearchQuoteAnalyzer() string {
	return sq.SearchQuoteAnalyzerValue
}

// SetSearchQuoteAnalyzer sets search_quote_analyzer to v
func (sq searchQuoteAnalyzerParam) SetSearchQuoteAnalyzer(v string) {
	if sq.SearchQuoteAnalyzer() != v {
		sq.SearchQuoteAnalyzerValue = v
	}
}

func marshalSearchQuoteAnalyzerParam(source interface{}) (dynamic.JSON, error) {
	if a, ok := source.(WithSearchQuoteAnalyzer); ok {
		if len(a.SearchQuoteAnalyzer()) > 0 {
			return json.Marshal(a.SearchQuoteAnalyzer())
		}
	}
	return nil, nil
}
func unmarshalSearchQuoteAnalyzerParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithSearchQuoteAnalyzer); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			var str string
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}
			a.SetSearchQuoteAnalyzer(str)
			return nil
		}
		return &json.UnmarshalTypeError{Value: string(data), Type: typeString}
	}
	return nil
}
