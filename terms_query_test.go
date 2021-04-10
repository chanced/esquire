package picker_test

import (
	"testing"

	"encoding/json"

	"github.com/chanced/dynamic"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestTerms(t *testing.T) {
	assert := require.New(t)

	json1 := dynamic.JSON(`
      {
          "terms": {
          "user.id": ["chanced", "kimchy", "elkbee"],
          "boost": 1.2,
          "case_insensitive": true
        }
      }  
    `)
	var q1 picker.Query

	err := json.Unmarshal(json1, &q1)
	assert.NoError(err)
	assert.Equal(float64(1.2), q1.Terms().Boost())
	assert.Equal([]string{"chanced", "kimchy", "elkbee"}, q1.Terms().Value())
	assert.Equal("user.id", q1.Terms().Field())
	assert.True(q1.Terms().CaseInsensitive())
	json1Res, err := json.MarshalIndent(q1, "", "  ")
	assert.NoError(err)
	var res1 picker.Query
	err = json.Unmarshal(json1Res, &res1)
	assert.NoError(err)
	assert.Equal(float64(1.2), res1.Terms().Boost(), "res1")
	assert.True(res1.Terms().CaseInsensitive(), "res1")
	assert.Equal([]string{"chanced", "kimchy", "elkbee"}, res1.Terms().Value(), "res1")
	assert.Equal("user.id", res1.Terms().Field(), "res1")

	json2 := []byte(`{
		"query": {
		  "terms": {
			  "color" : {
				  "index" : "my-index-000001",
				  "id" : "2",
				  "path" : "color"
			  }
		  }
		}
	  }`)
	assert.NoError(err)
	var s2 picker.Search

	err = json.Unmarshal(json2, &s2)
	assert.NoError(err)
	lookup := s2.Query().Terms().Lookup()
	assert.Equal("2", lookup.ID())
	assert.Equal("my-index-000001", lookup.Index())
	assert.Equal("color", lookup.Path())
	assert.Equal("color", s2.Query().Terms().Field())

	json2Res, err := json.MarshalIndent(s2.Query().Terms(), "", "  ")
	assert.NoError(err)
	var res2 picker.TermsQuery

	err = json.Unmarshal(json2Res, &res2)
	assert.NoError(err)
	lookup = s2.Query().Terms().Lookup()
	assert.Equal("2", lookup.ID())
	assert.Equal("my-index-000001", lookup.Index())
	assert.Equal("color", lookup.Path())
	assert.Equal("color", s2.Query().Terms().Field())

}
