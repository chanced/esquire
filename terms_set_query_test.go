package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestTermSet(t *testing.T) {
	assert := require.New(t)
	_ = assert
	data := []byte(`{
		"query": {
		  "terms_set": {
			"programming_languages": {
			  "terms": [ "c++", "java", "php" ],
			  "minimum_should_match_field": "required_matches"
			}
		  }
		}
	  }`)
	_ = data
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			TermsSet: picker.TermsSetQueryParams{
				Field:                   "programming_languages",
				Terms:                   []string{"c++", "java", "php"},
				MinimumShouldMatchField: "required_matches",
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

}
