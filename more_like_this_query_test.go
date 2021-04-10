package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestMoreLikeThis(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "more_like_this": {
			"fields": [ "title", "description" ],
			"like": [
			  {
				"_index": "imdb",
				"_id": "1"
			  },
			  {
				"_index": "imdb",
				"_id": "2"
			  },
			  "and potentially some more text here as well"
			],
			"min_term_freq": 1,
			"max_query_terms": 12
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			MoreLikeThis: picker.MoreLikeThisQueryParams{
				Fields: []string{"title", "description"},
				Like: []interface{}{
					map[string]interface{}{
						"_index": "imdb",
						"_id":    "1",
					},
					map[string]interface{}{
						"_index": "imdb",
						"_id":    "2",
					},
					"and potentially some more text here as well",
				},
				MinTermFrequency: 1,
				MaxQueryTerms:    12,
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
