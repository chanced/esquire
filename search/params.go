package search

import "github.com/tidwall/gjson"

type Param String

const (
	ParamBoost                           Param = "boost"
	ParamAnalyzer                        Param = "analyzer"
	ParamFormat                          Param = "format"
	ParamCaseInsensitive                 Param = "case_insensitive"
	ParamFuzziness                       Param = "fuzziness"
	ParamFuzzyRewrite                    Param = "fuzzy_rewrite"
	ParamFuzzyTranspositions             Param = "fuzzy_transpositions"
	ParamAutoGenerateSynonymsPhraseQuery Param = "auto_generate_synonyms_phrase_query"
	ParamLenient                         Param = "lenient"
	ParamMaxBoost                        Param = "max_boost"
	ParamMaxExpansions                   Param = "max_expansions"
	ParamMinShouldMatch                  Param = "minimum_should_match"
	ParamName                            Param = "_name"
	ParamOperator                        Param = "operator"
	ParamPrefixLength                    Param = "prefix_length"
	ParamRelation                        Param = "relation"
	ParamRewrite                         Param = "rewrite"
	ParamZeroTermsQuery                  Param = "zero_terms_query"
	ParamTranspositions                  Param = "transpositions"
	ParamTimeZone                        Param = "time_zone"
	ParamSlop                            Param = "slop"
	// ParamQuery                           Param = "query"
)

type Params []Param

func (p Params) Contains(param string) bool {
	v := Param(param)
	for _, k := range p {
		if k == v {
			return true
		}
	}
	return false
}

var paramUnmarshalers = map[Param]func(data gjson.Result, target interface{}) error{
	ParamBoost:                           unmarshalBoostParam,
	ParamAnalyzer:                        unmarshalAnalyzerParam,
	ParamFormat:                          unmarshalFormatParam,
	ParamCaseInsensitive:                 unmarshalCaseInsensitiveParam,
	ParamFuzziness:                       unmarshalFuzzinessParam,
	ParamFuzzyRewrite:                    unmarshalFuzzyRewriteParam,
	ParamFuzzyTranspositions:             unmarshalFuzzyTranspositionsParam,
	ParamLenient:                         unmarshalLenientParam,
	ParamMaxBoost:                        unmarshalMaxBoostParam,
	ParamMinShouldMatch:                  unmarshalMinShouldMatchParam,
	ParamName:                            unmarshalNameParam,
	ParamOperator:                        unmarshalOperatorParam,
	ParamPrefixLength:                    unmarshalPrefixLengthParam,
	ParamRelation:                        unmarshalRelationParam,
	ParamRewrite:                         unmarshalRewriteParam,
	ParamZeroTermsQuery:                  unmarshalZeroTermsQueryParam,
	ParamTranspositions:                  unmarshalTranspositionsParam,
	ParamTimeZone:                        unmarshalTimeZoneParam,
	ParamSlop:                            unmarshalSlopParam,
	ParamAutoGenerateSynonymsPhraseQuery: unmarshalAutoGenerateSynonymsPhraseQueryParam,
	// ParamQuery:                           unmarshalQueryParam,

}

func unmarshalParam(param Param, target interface{}, value gjson.Result) (bool, error) {
	if unmarshal, ok := paramUnmarshalers[param]; ok {
		return true, unmarshal(value, target)
	}
	return false, nil
}
