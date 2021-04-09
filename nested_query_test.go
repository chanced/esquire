package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestNested(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "nested": {
			"path": "driver",
			"query": {
			  "nested": {
				"path": "driver.vehicle",
				"query": {
				  "bool": {
					"must": [
					  { "match": { "driver.vehicle.make": { "query": "Powell Motors" } } },
					  { "match": { "driver.vehicle.model": { "query": "Canyonero" } } }
					]
				  }
				}
			  }
			}
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Nested: picker.NestedQueryParams{
				Path: "driver",
				Query: &picker.QueryParams{
					Nested: picker.NestedQueryParams{
						Path: "driver.vehicle",
						Query: &picker.QueryParams{
							Bool: picker.BoolQueryParams{
								Must: picker.Clauses{
									picker.MatchQueryParams{
										Field: "driver.vehicle.make",
										Query: "Powell Motors",
									},
									picker.MatchQueryParams{
										Field: "driver.vehicle.model",
										Query: "Canyonero",
									},
								},
							},
						},
					},
				},
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
}
