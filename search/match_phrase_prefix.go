package search

type MatchPhrasePrefix struct {
	Query string `json:"query" bson:"query"`
	analyzerParam
	maxExpansionsParam
	slopParam
	ZeroTermsQuery
}

func NewMatchPhrasePrefix(query string) MatchPhrasePrefix {
	return MatchPhrasePrefix{Query: query}
}

type MatchPhrasePrefixQuery struct {
	MatchPhrasePrefix map[string]MatchBooleanPrefix `json:"match_phrase_prefix,omitempty" bson:"match_phrase_prefix,omitempty"`
}
