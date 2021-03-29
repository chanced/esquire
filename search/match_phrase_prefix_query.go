package search

type MatchPhrasePrefix struct {
	Query string `json:"query" bson:"query"`
	analyzerParam
	maxExpansionsParam
	slopParam
	ZeroTerms
}

type MatchPhrasePrefixQuery struct {
}
