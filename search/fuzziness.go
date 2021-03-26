package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultFuzziness = "0"

// WithFuzziness is an interface for queries with fuzziness the parameter
//
// Maximum edit distance allowed for matching. See Fuzziness for valid values and more information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
type WithFuzziness interface {
	// Fuzziness is the maximum edit distance allowed for matching. See
	// Fuzziness for valid values and more information. See Fuzziness in the
	// match query for an example.
	Fuzziness() string
	// SetFuzziness sets the FuzzinessValue to v
	SetFuzziness(str string)
	// FuzzyRewrite is the method used to rewrite the query. See the rewrite
	// parameter for valid values and more information.
	//
	// If the fuzziness parameter is not 0, the match query uses a fuzzy_rewrite
	// method of top_terms_blended_freqs_${max_expansions} by default.
	FuzzyRewrite() Rewrite
	// SetFuzzyRewrite sets the value of FuzzyRewrite to v
	SetFuzzyRewrite(v Rewrite)
	DefaultFuzzyRewrite() Rewrite
}

// fuzzinessParam is a mixin that adds the fuzziness parameter to queries
//
// Maximum edit distance allowed for matching. See Fuzziness for valid values and more information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
type fuzzinessParam struct {
	fuzziness    string  `json:"fuzziness,omitempty" bson:"fuzziness,omitempty"`
	fuzzyRewrite Rewrite `json:"fuzzy_rewrite,omitempty" bson:"fuzzy_rewrite,omitempty"`
}

// Fuzziness is the maximum edit distance allowed for matching. See Fuzziness
// for valid values and more information.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
func (f fuzzinessParam) Fuzziness() string {
	return f.fuzziness
}

// SetFuzziness sets the fuzzinessValue to v
func (f *fuzzinessParam) SetFuzziness(v string) {
	if f.Fuzziness() != v {
		f.fuzziness = v
	}
}

// FuzzyRewrite is the method used to rewrite the query. See the rewrite
// parameter for valid values and more information.
//
// If the fuzziness parameter is not 0, the match query uses a fuzzy_rewrite
// method of top_terms_blended_freqs_${max_expansions} by default.
func (f fuzzinessParam) FuzzyRewrite() Rewrite {
	if f.fuzzyRewrite != "" {
		return f.fuzzyRewrite
	}
	return f.DefaultFuzzyRewrite()
}

func (f fuzzinessParam) DefaultFuzzyRewrite() Rewrite {
	if f.Fuzziness() != DefaultFuzziness {
		return RewriteTopTermsBlendedFreqsN
	}
	return RewriteConstantScore
}

// SetFuzzyRewrite sets the value of FuzzyRewrite to v
func (f *fuzzinessParam) SetFuzzyRewrite(v Rewrite) error {
	if !v.IsValid() {
		return ErrInvalidRewrite
	}
	if f.FuzzyRewrite() != v {
		f.fuzzyRewrite = v
	}
	return nil
}

func unmarshalFuzzinessParam(data dynamic.RawJSON, target interface{}) error {
	if r, ok := target.(WithFuzziness); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			r.SetFuzziness(data.UnquotedString())
			return nil
		}
		return &json.UnmarshalTypeError{Value: data.String()}
	}
	return nil
}

func marshalFuzzinessParam(data M, source interface{}) (M, error) {
	if a, ok := source.(WithFuzziness); ok {
		if a.Fuzziness() != DefaultFuzziness {
			data[paramFuzziness] = a.Fuzziness()
		}
	}
	return data, nil
}
func unmarshalFuzzyRewriteParam(data dynamic.RawJSON, target interface{}) error {
	if r, ok := target.(WithFuzziness); ok {
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			r.SetFuzzyRewrite(Rewrite(data.UnquotedString()))
			return nil
		}
		return &json.UnmarshalTypeError{Value: data.String()}
	}
	return nil
}
func marshalFuzzyRewriteParam(data M, source interface{}) (M, error) {
	if a, ok := source.(WithFuzziness); ok {
		if a.FuzzyRewrite() != a.DefaultFuzzyRewrite() {
			data[paramFuzzyRewrite] = a.FuzzyRewrite()
		}
	}
	return data, nil
}
