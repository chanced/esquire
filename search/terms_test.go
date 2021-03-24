package search_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/chanced/esquire/search"
	"github.com/stretchr/testify/require"
)

func TestTerms(t *testing.T) {
	assert := require.New(t)

	j1f, err := os.Open("./testdata/terms_1.json")
	assert.NoError(err)
	json1, err := ioutil.ReadAll(j1f)
	assert.NoError(err)
	fmt.Println(string(json1))
	var q search.Query
	err = json.Unmarshal(json1, &q)
	assert.NoError(err)
	assert.Equal(float64(1), q.Terms.Boost())
	fmt.Println(q.TermsQuery)
}
