package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestSimpleQueryStringQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "simple_query_string" : {
			  "query": "\"fried eggs\" +(eggplant | potato) -frittata",
			  "fields": ["title^5", "body"],
			  "default_operator": "AND"
		  }
		}
	  }`)

	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			SimpleQueryString: picker.SimpleQueryStringQueryParams{
				Query:           "\"fried eggs\" +(eggplant | potato) -frittata",
				Fields:          []string{"title^5", "body"},
				DefaultOperator: "and",
			},
		},
	})
	assert.NoError(err)
	// sd, err := s.MarshalJSON()
	sd, err := json.MarshalIndent(s, "", "  ")
	assert.NoError(err)

	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)
	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

}
