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

func TestTerm(t *testing.T) {
	assert := require.New(t)
	// {
	//   "term": {
	//     "user.id": {
	//       "value": "chanced",
	//       "boost": 0.2,
	//       "case_insensitive": true
	//     }
	//   }
	// }
	j1, err := os.Open("./testdata/term_1.json")
	assert.NoError(err)
	defer j1.Close()
	json1, err := ioutil.ReadAll(j1)
	assert.NoError(err)
	var q1 search.Query

	err = json.Unmarshal(json1, &q1)
	assert.NoError(err)

	assert.Equal(float64(0.2), q1.Term.Boost())
	assert.Equal("chanced", q1.Term.Value())
	assert.Equal("user.id", q1.Term.Field())
	assert.True(q1.Term.CaseInsensitive())
	json1Res, err := json.MarshalIndent(q1, "", "  ")
	assert.NoError(err)
	fmt.Println(string(json1Res))

	var rq1 search.Query
	err = json.Unmarshal(json1Res, &rq1)
	assert.NoError(err)
	assert.Equal(float64(0.2), rq1.Term.Boost())
	assert.True(rq1.Term.CaseInsensitive())
	assert.Equal("chanced", rq1.Term.Value())
	assert.Equal("user.id", rq1.Term.Field())

}
