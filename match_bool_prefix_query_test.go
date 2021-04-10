package picker_test

import (
	"encoding/json"
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
	assert.Equal("quick brown f", s.Query().MatchBoolPrefix().Query())
	assert.Equal("keyword", s.Query().MatchBoolPrefix().Analyzer())
	var sr picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)
	assert.Equal("quick brown f", sr.Query().MatchBoolPrefix().Query())
	assert.Equal("keyword", sr.Query().MatchBoolPrefix().Analyzer())

	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

}
