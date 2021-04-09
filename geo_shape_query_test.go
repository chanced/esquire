package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestGeoShape(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "bool": {
			"must": [
			  { "match_all": {} }
			],
			"filter": [{
			  "geo_shape": {
				"location": {
				  "shape": {
					"type": "envelope",
					"coordinates": [ [ 13, 53 ], [ 14, 52 ] ]
				  },
				  "relation": "within"
				}
			  }
			}]
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Bool: picker.BoolQueryParams{
				Must: picker.Clauses{
					&picker.MatchAllQuery{},
				},
				Filter: picker.Clauses{
					picker.GeoShapeQueryParams{
						Field: "location",
						Shape: picker.Shape{
							Type:        "envelope",
							Coordinates: [][]float64{{13.0, 53.0}, {14.0, 52.0}},
						},
						Relation: "within",
					},
				},
			},
		},
	})
	assert.NoError(err)

	sdp, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(sdp))

	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
}
