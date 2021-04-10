package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestWildcard(t *testing.T) {
	assert := require.New(t)
	_ = assert
	data := []byte(`{
		"query": {
		  "wildcard": {
			"user.id": {
			  "value": "ki*y",
			  "boost": 1.1,
			  "rewrite": "constant_score_boolean"
			}
		  }
		}
	  }`)
	_ = data
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Wildcard: picker.WildcardQueryParams{
				Value:   "ki*y",
				Boost:   1.1,
				Rewrite: picker.RewriteConstantScoreBoolean,
				Field:   "user.id",
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
