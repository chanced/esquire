package search

import (
	"github.com/chanced/dynamic"
)

const (
	paramBoost                           = "boost"
	paramAnalyzer                        = "analyzer"
	paramFormat                          = "format"
	paramCaseInsensitive                 = "case_insensitive"
	paramFuzziness                       = "fuzziness"
	paramFuzzyRewrite                    = "fuzzy_rewrite"
	paramFuzzyTranspositions             = "fuzzy_transpositions"
	paramAutoGenerateSynonymsPhraseQuery = "auto_generate_synonyms_phrase_query"
	paramLenient                         = "lenient"
	paramMaxBoost                        = "max_boost"
	ParamMaxExpansions                   = "max_expansions"
	paramMinimumShouldMatch              = "minimum_should_match"
	paramName                            = "_name"
	paramOperator                        = "operator"
	paramPrefixLength                    = "prefix_length"
	paramRelation                        = "relation"
	paramRewrite                         = "rewrite"
	paramZeroTermsQuery                  = "zero_terms_query"
	paramTranspositions                  = "transpositions"
	paramTimeZone                        = "time_zone"
	paramSlop                            = "slop"
	// ParamQuery                           Param = "query"
)

var paramMarshalers = map[string]func(data M, source interface{}) (M, error){
	paramBoost:                           marshalBoostParam,
	paramAnalyzer:                        marshalAnalyzerParam,
	paramFormat:                          marshalFormatParam,
	paramCaseInsensitive:                 marshalCaseInsensitiveParam,
	paramFuzziness:                       marshalFuzzinessParam,
	paramFuzzyRewrite:                    marshalFuzzyRewriteParam,
	paramFuzzyTranspositions:             marshalFuzzyTranspositionsParam,
	paramLenient:                         marshalLenientParam,
	paramMaxBoost:                        marshalMaxBoostParam,
	paramMinimumShouldMatch:              marshalMinimumShouldMatchParam,
	paramName:                            marshalNameParam,
	paramOperator:                        marshalOperatorParam,
	paramPrefixLength:                    marshalPrefixLengthParam,
	paramRelation:                        marshalRelationParam,
	paramRewrite:                         marshalRewriteParam,
	paramZeroTermsQuery:                  marshalZeroTermsQueryParam,
	paramTranspositions:                  marshalTranspositionsParam,
	paramTimeZone:                        marshalTimeZoneParam,
	paramSlop:                            marshalSlopParam,
	paramAutoGenerateSynonymsPhraseQuery: marshalAutoGenerateSynonymsPhraseQueryParam,
}

var paramUnmarshalers = map[string]func(data dynamic.RawJSON, target interface{}) error{
	paramBoost:                           unmarshalBoostParam,
	paramAnalyzer:                        unmarshalAnalyzerParam,
	paramFormat:                          unmarshalFormatParam,
	paramCaseInsensitive:                 unmarshalCaseInsensitiveParam,
	paramFuzziness:                       unmarshalFuzzinessParam,
	paramFuzzyRewrite:                    unmarshalFuzzyRewriteParam,
	paramFuzzyTranspositions:             unmarshalFuzzyTranspositionsParam,
	paramLenient:                         unmarshalLenientParam,
	paramMaxBoost:                        unmarshalMaxBoostParam,
	paramMinimumShouldMatch:              unmarshalMinimumShouldMatchParam,
	paramName:                            unmarshalNameParam,
	paramOperator:                        unmarshalOperatorParam,
	paramPrefixLength:                    unmarshalPrefixLengthParam,
	paramRelation:                        unmarshalRelationParam,
	paramRewrite:                         unmarshalRewriteParam,
	paramZeroTermsQuery:                  unmarshalZeroTermsQueryParam,
	paramTranspositions:                  unmarshalTranspositionsParam,
	paramTimeZone:                        unmarshalTimeZoneParam,
	paramSlop:                            unmarshalSlopParam,
	paramAutoGenerateSynonymsPhraseQuery: unmarshalAutoGenerateSynonymsPhraseQueryParam,
}

func unmarshalParam(param string, data dynamic.RawJSON, target interface{}) (bool, error) {

	if unmarshal, ok := paramUnmarshalers[param]; ok {
		if data.IsNull() {
			return true, nil
		}
		return true, unmarshal(data, target)
	}
	return false, nil
}

func marshalParams(data M, source interface{}) (M, error) {
	var err error
	for _, marshal := range paramMarshalers {

		data, err = marshal(data, source)
		if err != nil {
			return data, err
		}
	}
	return data, err
}
