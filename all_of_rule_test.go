package picker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllOfRule(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "intervals" : {
			"my_text" : {
			  "all_of" : {
				"intervals" : [
				  { "match" : { "query" : "the" } },
				  { "any_of" : {
					  "intervals" : [
						  { "match" : { "query" : "big" } },
						  { "match" : { "query" : "big bad" } }
					  ] } },
				  { "match" : { "query" : "wolf" } }
				],
				"max_gaps" : 0,
				"ordered" : true
			  }
			}
		  }
		}
	  }`)
	_ = assert
	_ = data
}
