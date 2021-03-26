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

func TestTerms(t *testing.T) {
	assert := require.New(t)

	j1, err := os.Open("./testdata/terms_1.json")
	assert.NoError(err)
	defer j1.Close()
	json1, err := ioutil.ReadAll(j1)
	assert.NoError(err)
	var q1 search.Query

	err = json.Unmarshal(json1, &q1)
	assert.NoError(err)
	assert.Equal(float64(1.2), q1.TermsQuery.Boost())
	assert.Equal([]string{"chanced", "kimchy", "elkbee"}, q1.TermsQuery.TermsValue)
	assert.Equal("user.id", q1.TermsField)
	assert.True(q1.TermsQuery.CaseInsensitive())
	json1Res, err := json.MarshalIndent(q1, "", "  ")
	fmt.Println(string(json1Res))
	assert.NoError(err)
	var res1 search.Query
	err = json.Unmarshal(json1Res, &res1)
	assert.NoError(err)
	assert.Equal(float64(1.2), res1.TermsQuery.Boost(), "res1")
	assert.True(res1.TermsQuery.CaseInsensitive(), "res1")
	assert.Equal([]string{"chanced", "kimchy", "elkbee"}, res1.TermsValue, "res1")
	assert.Equal("user.id", res1.TermsField, "res1")

	j2, err := os.Open("./testdata/terms_2.json")
	assert.NoError(err)
	defer j2.Close()
	json2, err := ioutil.ReadAll(j2)
	assert.NoError(err)
	var q2 search.Query

	fmt.Println("json2:\n", string(json2))
	err = json.Unmarshal(json2, &q2)
	assert.NoError(err)
	assert.Equal("2", q2.TermsQuery.TermsID)
	assert.Equal("my-index-000001", q2.TermsQuery.TermsIndex)
	assert.Equal("color", q2.TermsPath)
	assert.Equal("color", q2.TermsField)

	json2Res, err := json.MarshalIndent(q2.TermsQuery, "", "  ")
	assert.NoError(err)
	var res2 search.TermsRule

	fmt.Println(string(json2Res))
	err = json.Unmarshal(json2Res, &res2)
	assert.NoError(err)
	assert.Equal("2", res2.TermsID)
	assert.Equal("my-index-000001", res2.TermsIndex)
	assert.Equal("color", res2.TermsPath)
	assert.Equal("color", res2.TermsField)

}
