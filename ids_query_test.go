package picker_test

import (
	"encoding/json"
	"testing"

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
	assert.NoError(compareJSONObject(data, sd))
}
