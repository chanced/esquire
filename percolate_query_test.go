package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestPercolate(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "percolate": {
			"field": "query",
			"document": {
			  "message": "A new bonsai tree in the office"
			}
		  }
		}
	  }`)

	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Percolate: picker.PercolateDocumentQueryParams{
				Field: "query",
				Document: map[string]string{
					"message": "A new bonsai tree in the office",
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

	data = []byte(`{
		"query": {
		  "percolate": {
			"field": "query",
			"documents": [ 
			  {
				"message": "bonsai tree"
			  },
			  {
				"message": "new tree"
			  },
			  {
				"message": "the office"
			  },
			  {
				"message": "office tree"
			  }
			]
		  }
		}
	  }`)

	s, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Percolate: picker.PercolateDocumentsQueryParams{
				Field: "query",
				Documents: []map[string]string{
					{"message": "bonsai tree"},
					{"message": "new tree"},
					{"message": "the office"},
					{"message": "office tree"},
				},
			},
		},
	})
	assert.NoError(err)
	sd, err = s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)

	data = []byte(`{
		"query": {
		  "percolate": {
			"field": "query",
			"index": "my-index-00001",
			"id": "2",
			"version": 1 
		  }
		}
	  }`)
	s, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Percolate: picker.PercolateStoredDocumentQuery{
				Field:   "query",
				ID:      "2",
				Index:   "my-index-00001",
				Version: 1,
			},
		},
	})
	assert.NoError(err)
	sd, err = s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)

}
