package search_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/chanced/picker/search"
	"github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	assert := require.New(t)

	// {
	//   "match": {
	//     "message": {
	//       "query": "this is a test"
	//     }
	//   }
	// }

	j1, err := os.Open("./testdata/match_1.json")
	assert.NoError(err)
	defer j1.Close()

	json1, err := ioutil.ReadAll(j1)
	assert.NoError(err)
	var q1 search.Query
	err = json.Unmarshal(json1, &q1)
	assert.NoError(err)
	assert.Equal("this is a test", q1.MatchQueryValue.String())
	assert.Equal("message", q1.MatchField)

	rjson1, err := json.Marshal(q1)
	assert.NoError(err)
	fmt.Println(string(rjson1))

	var rq1 search.Query
	err = json.Unmarshal(rjson1, &rq1)
	assert.NoError(err)
	assert.Equal("this is a test", rq1.MatchQueryValue.String())
	assert.Equal("message", rq1.MatchField)

	// {
	//   "match": {
	//     "message": {
	//       "query": 34.78,
	//       "operator": "and",
	//       "fuzziness": "AUTO",
	//       "zero_terms_query": "all",
	//       "cutoff_frequency": 0.001,
	//       "auto_generate_synonyms_phrase_query": false,
	//       "minimum_should_match": "75%",
	//       "fuzzy_transpositions": false,
	//       "prefix_length": 1,
	//       "lenient": true,
	//       "max_expansions": 25,
	//       "analyzer": "test-analyzer"
	//     }
	//   }
	// }

	j2, err := os.Open("./testdata/match_2.json")
	assert.NoError(err)
	defer j2.Close()

	json2, err := ioutil.ReadAll(j2)
	assert.NoError(err)
	var q2 search.Query
	err = json.Unmarshal(json2, &q2)
	assert.NoError(err)

	assert.Equal(search.OperatorAnd, q2.MatchQuery.Operator())
	assert.Equal(search.And, q2.MatchQuery.Operator())
	assert.Equal("AUTO", q2.MatchQuery.Fuzziness())
	assert.Equal(search.ZeroTermsQueryAll, q2.MatchQuery.ZeroTermsQuery())
	assert.Equal(false, q2.MatchQuery.AutoGenerateSynonymsPhraseQuery())
	cutoff, ok := q2.MatchQuery.CutoffFrequency().Float()
	assert.True(ok, "should have a cutoff")
	assert.Equal(float64(0.001), cutoff)
	assert.Equal("75%", q2.MatchQuery.MinimumShouldMatch(), "minimum_should_match should be set")
	assert.Equal(false, q2.MatchQuery.FuzzyTranspositions(), "fuzzy_transpositions should be set")
	assert.Equal("test-analyzer", q2.MatchQuery.Analyzer(), "analyzer should be test-analyzer")
	assert.Equal(int64(1), q2.MatchQuery.PrefixLength(), "prefix_length should be 1")
	assert.Equal(int64(25), q2.MatchQuery.MaxExpansions(), "max_expansions should be set to 25")
	assert.Equal(true, q2.MatchQuery.Lenient())
	rjson2, err := json.Marshal(q2)
	assert.NoError(err)
	fmt.Println(string(rjson2))

	var rq2 search.Query
	err = json.Unmarshal(rjson2, &rq2)
	assert.NoError(err)
	f, ok := rq2.MatchQueryValue.Float()
	assert.True(ok, "q2 match query should be a float")
	assert.Equal(float64(34.78), f, "q2 match query should be 34.78")
	assert.Equal(search.OperatorAnd, rq2.MatchQuery.Operator())
	assert.Equal(search.And, rq2.MatchQuery.Operator())
	assert.Equal("AUTO", rq2.MatchQuery.Fuzziness())
	assert.Equal(search.ZeroTermsQueryAll, rq2.MatchQuery.ZeroTermsQuery())
	assert.Equal(false, rq2.MatchQuery.AutoGenerateSynonymsPhraseQuery())
	assert.Equal("75%", rq2.MatchQuery.MinimumShouldMatch(), "minimum_should_match should be set")
	assert.Equal(false, rq2.MatchQuery.FuzzyTranspositions(), "fuzzy_transpositions should be set")
	assert.Equal("test-analyzer", rq2.MatchQuery.Analyzer(), "analyzer should be test-analyzer")
	assert.Equal(int64(1), rq2.MatchQuery.PrefixLength(), "prefix_length should be 1")
	assert.Equal(int64(25), rq2.MatchQuery.MaxExpansions(), "max_expansions should be set to 25")
	assert.Equal(true, rq2.MatchQuery.Lenient())

	assert.True(ok, "should have a cutoff")
	cutoff2, ok := q2.MatchQuery.CutoffFrequency().Float()
	assert.True(ok, "should have a cutoff")
	// assert.Equal(float64(0.001), cutoff)
	assert.Equal(float64(0.001), cutoff2)
}
