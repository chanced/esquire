package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type WithSearchAnalyzer interface {
	// SearchAnalyzer overrides Analyzer for search analysis
	SearchAnalyzer() string
	// SetSearchAnalyzer sets SearchAnalyzer to v
	SetSearchAnalyzer(v string)
}

type searchAnalyzerParam struct {
	searchAnalyzer string
}

// SearchAnalyzer overrides Analyzer for search analysis
func (sa searchAnalyzerParam) SearchAnalyzer() string {
	return sa.searchAnalyzer
}

// SetSearchAnalyzer sets search_analyzer to v
func (sa *searchAnalyzerParam) SetSearchAnalyzer(v string) {
	if sa.SearchAnalyzer() != v {
		sa.searchAnalyzer = v
	}
}

func marshalSearchAnalyzerParam(source interface{}) (dynamic.JSON, error) {
	if a, ok := source.(WithSearchAnalyzer); ok {
		if len(a.SearchAnalyzer()) > 0 {
			return json.Marshal(a.SearchAnalyzer())
		}
	}
	return nil, nil
}
func unmarshalSearchAnalyzerParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithSearchAnalyzer); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			var str string
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}
			a.SetSearchAnalyzer(str)
			return nil
		}
		return &json.UnmarshalTypeError{Value: string(data), Type: typeString}
	}
	return nil
}
