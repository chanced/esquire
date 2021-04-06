package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type WithQuoteAnalyzer interface {

	// QuoteAnalyzer setting allows you to specify an analyzer for
	// phrases, this is particularly useful when dealing with disabling stop
	// words for phrase queries.
	QuoteAnalyzer() string
	// SetQuoteAnalyzer sets QuoteAnalyzer to v
	SetQuoteAnalyzer(v string)
}

type quoteAnalyzerParam struct {
	quoteAnalyzer string
}

// QuoteAnalyzer setting allows you to specify an analyzer for
// phrases, this is particularly useful when dealing with disabling
// stop words for phrase queries.
func (sq quoteAnalyzerParam) QuoteAnalyzer() string {
	return sq.quoteAnalyzer
}

// SetQuoteAnalyzer sets search_quote_analyzer to v
func (sq quoteAnalyzerParam) SetQuoteAnalyzer(v string) {
	if sq.QuoteAnalyzer() != v {
		sq.quoteAnalyzer = v
	}
}

func marshalQuoteAnalyzerParam(source interface{}) (dynamic.JSON, error) {
	if a, ok := source.(WithQuoteAnalyzer); ok {
		if len(a.QuoteAnalyzer()) > 0 {
			return json.Marshal(a.QuoteAnalyzer())
		}
	}
	return nil, nil
}
func unmarshalQuoteAnalyzerParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithQuoteAnalyzer); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			var str string
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}
			a.SetQuoteAnalyzer(str)
			return nil
		}
		return &json.UnmarshalTypeError{Value: string(data), Type: typeString}
	}
	return nil
}
