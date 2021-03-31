package picker

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

// TODO: accept a subset of params so the handler maps aren't needlessly looped through for types that are not applicable

func isKnownParam(v string) bool {
	_, ok := queryParamMarshalers[v]
	return ok
}

type paramMarshaler func(source interface{}) (dynamic.JSON, error)
type paramMarshalers map[string]paramMarshaler

var queryParamMarshalers = paramMarshalers{
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

type paramUnmarshaler func(data dynamic.JSON, target interface{}) error
type paramUnmarshalers map[string]paramUnmarshaler

var queryParamUnmarshalers = paramUnmarshalers{
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

func marshalClauseParams(source interface{}) (dynamic.Map, error) {
	return marshalParams(source, queryParamMarshalers)
}
func marshalParams(source interface{}, marshalers paramMarshalers) (dynamic.Map, error) {
	data := dynamic.Map{}
	for key, marshal := range marshalers {
		d, err := marshal(source)
		if err != nil {
			return nil, fmt.Errorf("%w for parameter %s", err, key)
		}
		if d != nil && !d.IsNull() {
			if (d.IsString() || d.IsArray() || d.IsObject()) && d.Len() < 3 {
				continue
			}
			data[key] = d
		}
	}
	return data, nil
}

func unmarshalParams(data []byte, target interface{}, unmarshalers paramUnmarshalers) (dynamic.JSONObject, error) {
	var raw map[string]dynamic.JSON
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}
	res := map[string]dynamic.JSON{}
	for key, value := range raw {
		isParam, err := unmarshalParam(key, value, target, unmarshalers)
		if err != nil {
			return nil, err
		}
		if !isParam {
			res[key] = value
		}
	}
	return res, nil

}

func unmarshalParam(param string, data dynamic.JSON, target interface{}, unmarshalers paramUnmarshalers) (bool, error) {
	if unmarshal, ok := unmarshalers[param]; ok {
		if data.IsNull() {
			return true, nil
		}
		return true, unmarshal(data, target)
	}
	return false, nil
}

func unmarshalClauseParams(data []byte, target QueryClause) (dynamic.JSONObject, error) {
	return unmarshalParams(data, target, queryParamUnmarshalers)
}
