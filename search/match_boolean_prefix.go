package search

type MatchBooleanPrefix struct {
	Query                    string `json:"query" bson:"query"`
	minimumShouldMatchParam  `json:",inline" bson:",inline"`
	operatorParam            `json:",inline" bson:",inline"`
	analyzerParam            `json:",inline" bson:",inline"`
	fuzzinessParam           `json:",inline" bson:",inline"`
	prefixLengthParam        `json:",inline" bson:",inline"`
	fuzzyTranspositionsParam `json:",inline" bson:",inline"`
}

type MatchBooleanPrefixQuery struct {
	MatchBooleanPrefix map[string]MatchBooleanPrefix `json:"match_boolean_prefix,omitempty" bson:"match_boolean_prefix,omitempty"`
}

func NewMatchBooleanPrefixQuery() MatchBooleanPrefixQuery {
	return MatchBooleanPrefixQuery{
		MatchBooleanPrefix: map[string]MatchBooleanPrefix{},
	}
}
