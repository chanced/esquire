package search

import "github.com/tidwall/gjson"

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
	paramMinShouldMatch                  = "minimum_should_match"
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

var paramMarshalers = map[string]func(data map[string]interface{}, source interface{}) (map[string]interface{}, error){
	paramBoost:                           marshalBoostParam,
	paramAnalyzer:                        marshalAnalyzerParam,
	paramFormat:                          marshalFormatParam,
	paramCaseInsensitive:                 marshalCaseInsensitiveParam,
	paramFuzziness:                       marshalFuzzinessParam,
	paramFuzzyRewrite:                    marshalFuzzyRewriteParam,
	paramFuzzyTranspositions:             unmarshalFuzzyTranspositionsParam,
	paramLenient:                         unmarshalLenientParam,
	paramMaxBoost:                        unmarshalMaxBoostParam,
	paramMinShouldMatch:                  unmarshalMinShouldMatchParam,
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

var paramUnmarshalers = map[string]func(data gjson.Result, target interface{}) error{
	paramBoost:                           unmarshalBoostParam,
	paramAnalyzer:                        unmarshalAnalyzerParam,
	paramFormat:                          unmarshalFormatParam,
	paramCaseInsensitive:                 unmarshalCaseInsensitiveParam,
	paramFuzziness:                       unmarshalFuzzinessParam,
	paramFuzzyRewrite:                    unmarshalFuzzyRewriteParam,
	paramFuzzyTranspositions:             unmarshalFuzzyTranspositionsParam,
	paramLenient:                         unmarshalLenientParam,
	paramMaxBoost:                        unmarshalMaxBoostParam,
	paramMinShouldMatch:                  unmarshalMinShouldMatchParam,
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

func unmarshalParam(param string, target interface{}, value gjson.Result) (bool, error) {
	if unmarshal, ok := paramUnmarshalers[param]; ok {
		return true, unmarshal(value, target)
	}
	return false, nil
}
