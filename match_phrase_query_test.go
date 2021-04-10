package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestMatchPhraseQuery(t *testing.T) {
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
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	sr := picker.Search{}
	err = sr.UnmarshalJSON(data)
	assert.NoError(err)

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
	assert.True(cmpjson.Equal(expected, sd), cmpjson.Diff(expected, sd))
}
