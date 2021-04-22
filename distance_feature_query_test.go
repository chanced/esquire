package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestDistanceFeature(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "bool": {
			"must": [{
			  "match": {
				"name": {
					"query": "chocolate"
				}
			  }
			}],
			"should": [{
			  "distance_feature": {
				"field": "production_date",
				"pivot": "7d",
				"origin": "now"
			  }
			}]
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			Bool: picker.BoolQueryParams{
				Must: picker.Clauses{
					picker.MatchQueryParams{
						Field: "name",
						Query: "chocolate",
					},
				},
				Should: picker.Clauses{
					picker.DistanceFeatureQueryParams{
						Field:  "production_date",
						Pivot:  "7d",
						Origin: "now",
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
