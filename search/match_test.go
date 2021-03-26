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
	//       "query": "this is a test",
	//       "operator": "and",
	//       "fuzziness": "AUTO",
	//       "zero_terms_query": "all",
	//       "cutoff_frequency": 0.001,
	//       "auto_generate_synonyms_phrase_query": false
	//     }
	//   }
	// }

}
