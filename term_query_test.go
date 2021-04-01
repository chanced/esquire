package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/dynamic"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestTerm(t *testing.T) {
	assert := require.New(t)

	json1 := dynamic.JSON(`{
        "term": {
          "user.id": {
            "value": "chanced",
            "boost": 0.2,
            "case_insensitive": true
          }
        }
      }
      `)
	var q1 picker.Query
	err := json.Unmarshal(json1, &q1)
	assert.NoError(err)

	assert.Equal(float64(0.2), q1.Term().Boost())
	assert.Equal("chanced", q1.Term().Value())
	assert.Equal("user.id", q1.Term().Field())
	assert.True(q1.Term().CaseInsensitive())
	json1Res, err := json.MarshalIndent(q1, "", "  ")
	assert.NoError(err)
	fmt.Printf("%+v\n", q1.Term())
	fmt.Println(string(json1Res))

	var rq1 picker.Query
	err = json.Unmarshal(json1Res, &rq1)
	assert.NoError(err)
	assert.Equal("chanced", rq1.Term().Value(), "value should be chanced")
	assert.Equal(float64(0.2), rq1.Term().Boost(), "boost should be 0.2")
	assert.True(rq1.Term().CaseInsensitive(), "case_insensitive should be true")

	assert.Equal("user.id", rq1.Term().Field(), "field should be user.id")

}
