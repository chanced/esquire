package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestMatchPhrasePrefixQuery(t *testing.T) {
	assert := require.New(t)
	_ = assert
	data := []byte(`{
		"query": {
		  "match_phrase_prefix": {
			"message": {
			  "query": "quick brown f"
			}
		  }
		}
	  }`)
	_ = data
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			MatchPhrasePrefix: picker.MatchPhrasePrefixQueryParams{
				Query: "quick brown f",
				Field: "message",
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)
	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

}
