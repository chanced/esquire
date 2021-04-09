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
		Query: picker.QueryParams{
			MultiMatch: picker.MultiMatchQueryParams{
				Query:  "this is a test",
				Fields: []string{"subject", "message"},
			},
		},
	})
	assert.NoError(err)
	sd, err := json.Marshal(s)
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
}
