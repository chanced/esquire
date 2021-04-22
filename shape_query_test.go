package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestShape(t *testing.T) {
	data := []byte(`{
		"query": {
		  "shape": {
			"geometry": {
			  "shape": {
				"type": "envelope",
				"coordinates": [ [ 1355.0, 5355.0 ], [ 1400.0, 5200.0 ] ]
			  },
			  "relation": "WITHIN"
			}
		  }
		}
	  }`)
	assert := require.New(t)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			Shape: picker.ShapeQueryParams{
				Field: "geometry",
				Shape: picker.Shape{
					Coordinates: [][]float64{{1355, 5355}, {1400, 5200}},
					Type:        "envelope",
				},
				Relation: picker.SpatialRelationWithin,
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
