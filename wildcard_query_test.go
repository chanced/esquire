package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

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
	sdi, err := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(sdi))
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))

}
