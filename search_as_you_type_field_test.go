package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestSearchAsYouType(t *testing.T) {
	assert := require.New(t)

	data := []byte(`{
		"mappings": {
		  "properties": {
			"my_field": {
			  "type": "search_as_you_type"
			}
		  }
		}
	  }`)

	var index picker.Index

	err := json.Unmarshal(data, &index)
	assert.NoError(err)

}
