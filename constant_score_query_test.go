package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestConstantScore(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
			"constant_score": {
				"filter": {
					"term": { "user.id": { "value": "kimchy" } }
				},
				"boost": 1.2
			}
		}
	}`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			ConstantScore: picker.ConstantScoreQueryParams{
				Filter: &picker.QueryParams{
					Term: picker.TermQueryParams{
						Field: "user.id",
						Value: "kimchy",
					},
				},
				Boost: 1.2,
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
