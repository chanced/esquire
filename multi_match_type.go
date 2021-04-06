package picker

type MultiMatchType string

const (
	// (default) Finds documents which match any field, but uses the _score from
	// the best field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#type-best-fields
	MultiMatchBestFields MultiMatchType = "best_fields"
	// Finds documents which match any field and combines the _score from each
	// field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#type-most-fields
	MultiMatchMostFields MultiMatchType = "most_fields"
	// Treats fields with the same analyzer as though they were one big field.
	// Looks for each word in any field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#type-cross-fields
	MultiMatchCrossFields MultiMatchType = "cross_fields"
	// Runs a match_phrase query on each field and uses the _score from the
	// best field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#type-phrase
	MultiMatchPhrase MultiMatchType = "phrase"
	// Runs a match_phrase_prefix query on each field and uses the _score from
	// the best field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#type-phrase
	MultiMatchPhrasePrefix MultiMatchType = "phrase_prefix"
	// Creates a match_bool_prefix query on each field and combines the _score from each field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html#type-bool-prefix
	MultiMatchBoolPrefix MultiMatchType = "bool_prefix"
)
