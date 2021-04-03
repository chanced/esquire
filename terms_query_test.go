package picker_test

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Println(string(json1Res))
	assert.NoError(err)
	var res1 picker.Query
	err = json.Unmarshal(json1Res, &res1)
	assert.NoError(err)
	assert.Equal(float64(1.2), res1.Terms().Boost(), "res1")
	assert.True(res1.Terms().CaseInsensitive(), "res1")
	assert.Equal([]string{"chanced", "kimchy", "elkbee"}, res1.Terms().Value(), "res1")
	assert.Equal("user.id", res1.Terms().Field(), "res1")

	j2, err := os.Open("./testdata/terms_2.json")
	assert.NoError(err)
	defer j2.Close()
	json2, err := ioutil.ReadAll(j2)
	assert.NoError(err)
	var q2 picker.Query

	fmt.Println("json2:\n", string(json2))
	err = json.Unmarshal(json2, &q2)
	assert.NoError(err)
	lookup := q2.Terms().Lookup()
	assert.Equal("2", lookup.ID())
	assert.Equal("my-index-000001", lookup.Index())
	assert.Equal("color", lookup.Path())
	assert.Equal("color", q2.Terms().Field())

	json2Res, err := json.MarshalIndent(q2.Terms(), "", "  ")
	assert.NoError(err)
	var res2 picker.TermsQuery

	fmt.Println(string(json2Res))
	err = json.Unmarshal(json2Res, &res2)
	assert.NoError(err)
	lookup = q2.Terms().Lookup()
	assert.Equal("2", lookup.ID())
	assert.Equal("my-index-000001", lookup.Index())
	assert.Equal("color", lookup.Path())
	assert.Equal("color", q2.Terms().Field())

}
