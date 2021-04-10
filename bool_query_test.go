package picker_test

import (
	"testing"

	"encoding/json"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestBoolean(t *testing.T) {
	assert := require.New(t)
	json1 := []byte(`{
		"bool" : {
		  "must" : {
			"term" : { "user.id" : "chanced" }
		  },
		  "filter": {
			"term" : { "tags" : "production" }
		  },
		  "should" : [
			{ "term" : { "tags" : "env1" } },
			{ "term" : { "tags" : "deployed" } }
		  ],
		  "minimum_should_match" : 1,
		  "boost" : 1.0
		}
	  }`)

	var q1 picker.Query

	err := json.Unmarshal(json1, &q1)
	assert.NoError(err)
	assert.Equal(2, q1.Bool().Should().Len())

	json2 := []byte(`{
		"query": {
		  "bool": {
			"must_not": {
			  "exists": {
				"field": "user.id"
			  }
			}
		  }
		}
	  }
	  `)
	_ = json2
}
