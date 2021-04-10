package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestGeoBoundingBoxQuery(t *testing.T) {
	assert := require.New(t)

	data := []byte(`{
		"query": {
		  "geo_bounding_box": {
			"pin.location": {
			  "top_left": "dr5r9ydj2y73",
			  "bottom_right": "drj7teegpus6"
			}
		  }
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			GeoBoundingBox: picker.GeoBoundingBoxQueryParams{
				Field: "pin.location",
				BoundingBox: picker.BoundingBox{
					TopLeft:     "dr5r9ydj2y73",
					BottomRight: "drj7teegpus6",
				},
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	_ = sd
	_ = data
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)

}
