package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestIDsQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
	  "query": {
	    "ids" : {
	      "values" : ["1", "4", "100"]
	    }
	  }
	}`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			IDs: picker.IDsQueryParams{
				Values: []string{"1", "4", "100"},
			},
		},
	})
	assert.NoError(err)
	sd, err := json.Marshal(s)
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)
	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

}
