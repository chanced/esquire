package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestMultiMatchQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "multi_match" : {
			"query" : "this is a test",
			"fields" : [ "subject", "message" ] 
		  }
		}
	  }`)

	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			MultiMatch: picker.MultiMatchQueryParams{
				Query:  "this is a test",
				Fields: []string{"subject", "message"},
			},
		},
	})

	assert.Equal("this is a test", s.Query().MultiMatch().Query())
	assert.NoError(err)
	sd, err := json.Marshal(s)
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr picker.Search
	err = json.Unmarshal(data, &sr)
	assert.Equal("this is a test", sr.Query().MultiMatch().Query())

	assert.NoError(err)
	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

}
