package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestMatchBoolPrefixQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "match_bool_prefix": {
			"message": {
			  "query": "quick brown f",
			  "analyzer": "keyword"
			}
		  }
		}
	  }`)

	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			MatchBoolPrefix: picker.MatchBoolPrefixQueryParams{
				Field:    "message",
				Query:    "quick brown f",
				Analyzer: "keyword",
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
}
