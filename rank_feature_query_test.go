package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestRankFeature(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "rank_feature": {
			"field": "pagerank",
			"saturation": {
			  "pivot": 8
			}
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			RankFeature: picker.RankFeatureQueryParams{
				Field: "pagerank",
				Saturation: picker.SaturationFunctionParams{
					Pivot: 8,
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
