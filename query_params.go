package picker

import (
	"github.com/chanced/dynamic"
)

func isKnownParam(v string) bool {
	_, ok := queryParamMarshalers[v]
	return ok
}

var queryParamMarshalers = map[string]func(data dynamic.Map, source interface{}) (dynamic.Map, error){
	"boost":                               marshalBoostParam,
	"analyzer":                            marshalAnalyzerParam,
	"format":                              marshalFormatParam,
	"case_insensitive":                    marshalCaseInsensitiveParam,
	"fuzziness":                           marshalFuzzinessParam,
	"fuzzy_rewrite":                       marshalFuzzyRewriteParam,
	"fuzzy_transpositions":                marshalFuzzyTranspositionsParam,
	"lenient":                             marshalLenientParam,
	"max_boost":                           marshalMaxBoostParam,
	"minimum_should_match":                marshalMinimumShouldMatchParam,
	"_name":                               marshalNameParam,
	"operator":                            marshalOperatorParam,
	"prefix_length":                       marshalPrefixLengthParam,
	"relation":                            marshalRelationParam,
	"rewrite":                             marshalRewriteParam,
	"zero_terms_query":                    marshalZeroTermsQueryParam,
	"transpositions":                      marshalTranspositionsParam,
	"time_zone":                           marshalTimeZoneParam,
	"slop":                                marshalSlopParam,
	"auto_generate_synonyms_phrase_query": marshalAutoGenerateSynonymsPhraseQueryParam,
	"cutoff_frequency":                    marshalCutoffFrequencyParam,
	"max_expansions":                      marshalMaxExpansionsParam,
	"score_mode":                          marshalScoreModeParam,
	"boost_mode":                          marshalBoostModeParam,
	"min_score":                           marshalMinScoreParam,
	"modifier":                            marshalModifierParam,
}

var queryParamUnmarshalers = map[string]func(data dynamic.JSON, target interface{}) error{
	"boost":                               unmarshalBoostParam,
	"analyzer":                            unmarshalAnalyzerParam,
	"format":                              unmarshalFormatParam,
	"case_insensitive":                    unmarshalCaseInsensitiveParam,
	"fuzziness":                           unmarshalFuzzinessParam,
	"fuzzy_rewrite":                       unmarshalFuzzyRewriteParam,
	"fuzzy_transpositions":                unmarshalFuzzyTranspositionsParam,
	"lenient":                             unmarshalLenientParam,
	"max_boost":                           unmarshalMaxBoostParam,
	"minimum_should_match":                unmarshalMinimumShouldMatchParam,
	"_name":                               unmarshalNameParam,
	"operator":                            unmarshalOperatorParam,
	"prefix_length":                       unmarshalPrefixLengthParam,
	"relation":                            unmarshalRelationParam,
	"rewrite":                             unmarshalRewriteParam,
	"zero_terms_query":                    unmarshalZeroTermsQueryParam,
	"transpositions":                      unmarshalTranspositionsParam,
	"time_zone":                           unmarshalTimeZoneParam,
	"slop":                                unmarshalSlopParam,
	"auto_generate_synonyms_phrase_query": unmarshalAutoGenerateSynonymsPhraseQueryParam,
	"cutoff_frequency":                    unmarshalCutoffFrequencyParam,
	"max_expansions":                      unmarshalMaxExpansionsParam,
	"score_mode":                          unmarshalScoreModeParam,
	"boost_mode":                          unmarshalBoostModeParam,
	"min_score":                           unmarshalMinScoreParam,
	"modifier":                            unmarshalModifierParam,
}

func unmarshalClauseParam(param string, data dynamic.JSON, target interface{}) (bool, error) {
	if unmarshal, ok := queryParamUnmarshalers[param]; ok {
		if data.IsNull() {
			return true, nil
		}
		return true, unmarshal(data, target)
	}
	return false, nil
}

func marshalClauseParams(source interface{}) (dynamic.Map, error) {
	data := dynamic.Map{}
	var err error
	for _, marshal := range queryParamMarshalers {
		data, err = marshal(data, source)
		if err != nil {
			return data, err
		}
	}
	return data, err
}
