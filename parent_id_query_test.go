package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestParentID(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
			"parent_id": {
				"type": "my-child",
				"id": "1"
			}
		}
	  }`)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			ParentID: picker.ParentIDQueryParams{
				ID:   "1",
				Type: "my-child",
			},
		},
	})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
}
