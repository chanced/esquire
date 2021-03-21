package search

// Fuzzy returns documents that contain terms similar to the search term,
// as measured by a Levenshtein edit distance.
//
// An edit distance is the number of one-character changes needed to turn one
// term into another. These changes can include:
//
// - Changing a character (box → fox)
//
// - Removing a character (black → lack)
//
// - Inserting a character (sic → sick)
//
// - Transposing two adjacent characters (act → cat)
//
// To find similar terms, the fuzzy query creates a set of all possible
// variations, or expansions, of the search term within a specified edit
// distance. The query then returns exact matches for each expansion.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type Fuzzy struct {
	Value               string `json:"value" bson:"value"`
	FuzzinessParam      `json:",inline" bson:",inline"`
	MaxExpansionsParam  `json:",inline" bson:",inline"`
	PrefixLengthParam   `json:",inline" bson:",inline"`
	TranspositionsParam `json:",inline" bson:",inline"`
	RewriteParam        `json:",inline" bson:",inline"`
}

func NewFuzzy() Fuzzy {
	return Fuzzy{}
}

// FuzzyQuery returns documents that contain terms similar to the search term,
// as measured by a Levenshtein edit distance.
//
// An edit distance is the number of one-character changes needed to turn one
// term into another. These changes can include:
//
// - Changing a character (box → fox)
//
// - Removing a character (black → lack)
//
// - Inserting a character (sic → sick)
//
// - Transposing two adjacent characters (act → cat)
//
// To find similar terms, the fuzzy query creates a set of all possible
// variations, or expansions, of the search term within a specified edit
// distance. The query then returns exact matches for each expansion.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type FuzzyQuery struct {
	Fuzzy map[string]Fuzzy `json:"fuzzy,omitempty" bson:"fuzzy,omitempty"`
}

func NewFuzzyQuery() FuzzyQuery {
	return FuzzyQuery{
		Fuzzy: map[string]Fuzzy{},
	}
}
