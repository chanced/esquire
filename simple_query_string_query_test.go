package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

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
	fmt.Println(string(sd))
	assert.NoError(err)

	assert.NoError(compareJSONObject(data, sd))
}
