package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestHasChild(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "has_child": {
			"type": "child",
			"query": {
			  "match": {
				"message": {
					"query": "this is a test"
				}
			  }
			},
			"score_mode": "max"
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			HasChild: picker.HasChildQueryParams{
				Type:      "child",
				ScoreMode: picker.ScoreModeMax,
				Query: &picker.QueryParams{
					Match: picker.MatchQueryParams{
						Field: "message",
						Query: "this is a test",
					},
				},
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
