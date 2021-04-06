package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultAnalyzeWildcard = false

type WithAnalyzeWildcard interface {
	// Default: false
	AnalyzeWildcard() bool
	// SetAnalyzeWildcard sets analyze_wildcard to v
	SetAnalyzeWildcard(v interface{}) error
}

type analyzeWildcardParam struct {
	analyzeWildcard dynamic.Bool
}

// If true, the query attempts to analyze wildcard terms in the query string. Defaults to false.
func (cp analyzeWildcardParam) AnalyzeWildcard() bool {
	if v, ok := cp.analyzeWildcard.Bool(); ok {
		return v
	}
	return DefaultAnalyzeWildcard
}

// SetAnalyzeWildcard sets analyze_wildcard to v
func (cp *analyzeWildcardParam) SetAnalyzeWildcard(v interface{}) error {
	return cp.analyzeWildcard.Set(v)
}

func unmarshalAnalyzeWildcardParam(value dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithAnalyzeWildcard); ok {
		b, err := dynamic.NewBool(value)
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			return a.SetAnalyzeWildcard(v)
		}
	}
	return nil
}
func marshalAnalyzeWildcardParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithAnalyzeWildcard); ok {
		if !b.AnalyzeWildcard() {
			return json.Marshal(b.AnalyzeWildcard())
		}
	}
	return nil, nil
}
