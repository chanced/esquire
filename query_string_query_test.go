package picker_test

import (
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestQueryStringQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "query_string": {
			"query": "(new york city) OR (big apple)",
			"default_field": "content"
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			QueryString: picker.QueryStringQueryParams{
				Query:        "(new york city) OR (big apple)",
				DefaultField: "content",
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))

	data = []byte(`{
		"query": {
		  "query_string" : {
			"fields" : ["content", "name^5"],
			"query" : "this AND that OR thus",
			"tie_breaker" : 0
		  }
		}
	  }`)
	s, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			QueryString: picker.QueryStringQueryParams{
				Query:      "this AND that OR thus",
				Fields:     []string{"content", "name^5"},
				TieBreaker: 0,
			},
		},
	})
	assert.NoError(err)
	sd, err = s.MarshalJSON()
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))

}
