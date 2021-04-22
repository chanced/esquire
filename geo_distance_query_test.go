package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestGeoDistance(t *testing.T) {
	data := []byte(`{
		"query": {
		  "bool": {
			"must": [{
			  "match_all": {}
			}],
			"filter": [{
			  "geo_distance": {
				"distance": "200km",
				"pin.location": {
				  "lat": 40,
				  "lon": -70
				}
			  }
			}]
		  }
		}
	  }`)
	assert := require.New(t)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			Bool: picker.BoolQueryParams{
				Must: picker.Clauses{
					picker.MatchAllQueryParams{},
				},
				Filter: picker.Clauses{
					picker.GeoDistanceQueryParams{
						Distance: "200km",
						Field:    "pin.location",
						GeoPoint: picker.LatLon{
							Lat: 40,
							Lon: -70,
						},
					},
				},
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(sd, data))
}
