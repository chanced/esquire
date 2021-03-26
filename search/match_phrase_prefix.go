package search

type MatchPhrasePrefix struct {
	Query              string `json:"query" bson:"query"`
	AnalyzerParam      `json:",inline" bson:",inline"`
	maxExpansionsParam `json:",inline" bson:",inline"`
	slopParam          `json:",inline" bson:",inline"`
	ZeroTermsQuery     `json:",inline" bson:",inline"`
}

func NewMatchPhrasePrefix(query string) MatchPhrasePrefix {
	return MatchPhrasePrefix{Query: query}
}

type MatchPhrasePrefixQuery struct {
	MatchPhrasePrefix map[string]MatchBooleanPrefix `json:"match_phrase_prefix,omitempty" bson:"match_phrase_prefix,omitempty"`
}
