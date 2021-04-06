package picker_test

import (
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestMatchPhrasePrefixQuery(t *testing.T) {
	assert := require.New(t)

	data := []byte(`{
		"query": {
		  "match_phrase": {
			"message": {
			  "query": "this is a test",
			  "analyzer": "my_analyzer"
			}
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			MatchPhrase: picker.MatchPhraseQueryParams{
				Query:    "this is a test",
				Analyzer: "my_analyzer",
				Field:    "message",
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))
	data = []byte(`{
		"query": {
		  "match_phrase": {
			"message": "this is a test"
		  }
		}
	  }`)

	err = s.UnmarshalJSON(data)
	assert.NoError(err)
	sd, err = s.MarshalJSON()
	assert.NoError(err)
	expected := []byte(`{
		"query": {
		  "match_phrase": {
			"message": {
			  "query": "this is a test"
			}
		  }
		}
	  }`)
	assert.NoError(compareJSONObject(expected, sd))
}
