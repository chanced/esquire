package search

type Fuzzier interface {
	Fuzzy() (*FuzzyQuery, error)
}

// Fuzzy returns documents that contain terms similar to the search term,
// as measured by a Levenshtein edit distance.
//
// An edit distance is the number of one-character changes needed to turn one
// term into another. These changes can include:
//
//  	- Changing a character (box → fox)
//
//  	- Removing a character (black → lack)
//
//  	- Inserting a character (sic → sick)
//
//  	- Transposing two adjacent characters (act → cat)
//
// To find similar terms, the fuzzy query creates a set of all possible
// variations, or expansions, of the search term within a specified edit
// distance. The query then returns exact matches for each expansion.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type Fuzzy struct {
	// Value or term to find in the provided <field>. (Required)
	Value string
	// Field which is being queried against.
	//
	// This will be ignored if set through a mutator
	Field string
	// Maximum edit distance allowed for matching. See Fuzziness for valid
	// values and more information. (Optional)
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/common-options.html#fuzziness
	Fuzziness string
	// PrefixLength is the number of beginning characters left unchanged when
	// creating expansions. Defaults to 0. (Optional)
	PrefixLength int
	// NoTranspositions indicates whether edits include transpositions of two
	// adjacent characters (ab → ba). (Optional)
	//
	// Setting NoTranspositions to true sets transpositions to false
	NoTranspositions bool
	// Rewrite method used to rewrite the query. (Optional)
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-term-rewrite.html
	Rewrite Rewrite
	// Maximum number of variations created. Defaults to 50. (Optional)
	//
	// Avoid using a high value in the max_expansions parameter, especially if
	// the prefix_length parameter value is 0. High values in the max_expansions
	// parameter can cause poor performance due to the high number of variations
	// examined.
	MaxExpansions interface{}
	// Name of the query (Optional)
	Name string
}

func (f Fuzzy) Fuzzy() (*FuzzyQuery, error) {
	q := &FuzzyQuery{field: f.Field}
	err := q.setValue(f.Value)
	if err != nil {
		return q, NewQueryError(err, TypeFuzzy, f.Field)
	}
	err = q.SetMaxExpansions(f.MaxExpansions)
	if err != nil {
		return q, NewQueryError(err, TypeFuzzy, f.Field)
	}
	err = q.SetRewrite(f.Rewrite)
	if err != nil {
		return q, NewQueryError(err, TypeFuzzy, f.Field)
	}
	q.SetTranspositions(!f.NoTranspositions)
	q.SetName(f.Name)
	q.SetFuzziness(f.Fuzziness)
	q.SetPrefixLength(f.PrefixLength)
	return q, nil
}

func (f Fuzzy) Clause() (Clause, error) {
	return f.Fuzzy()
}

func NewFuzzyQuery(params Fuzzier) (*FuzzyQuery, error) {
	q, err := params.Fuzzy()
	if err != nil {
		return nil, NewQueryError(err, TypeFuzzy, getField(q, nil))
	}
	err = checkField(q.field, TypeFuzzy)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// FuzzyQuery returns documents that contain terms similar to the search term,
// as measured by a Levenshtein edit distance.
//
// An edit distance is the number of one-character changes needed to turn one
// term into another. These changes can include:
//
//  	- Changing a character (box → fox)
//
//  	- Removing a character (black → lack)
//
//  	- Inserting a character (sic → sick)
//
//  	- Transposing two adjacent characters (act → cat)
//
// To find similar terms, the fuzzy query creates a set of all possible
// variations, or expansions, of the search term within a specified edit
// distance. The query then returns exact matches for each expansion.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-fuzzy-query.html
type FuzzyQuery struct {
	field string
	value string
	fuzzinessParam
	maxExpansionsParam
	prefixLengthParam
	transpositionsParam
	rewriteParam
	nameParam
}

func (f *FuzzyQuery) setValue(v string) error {
	if len(v) == 0 {
		return ErrValueRequired
	}
	return nil
}
func (f FuzzyQuery) Field() string {
	return f.field
}
func (f FuzzyQuery) Type() Type {
	return TypeFuzzy
}

func (f *FuzzyQuery) Set(field string, fuzzier Fuzzier) error {

}
