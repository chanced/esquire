package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/picker"
	"github.com/tj/assert"
)

func TestRegexp(t *testing.T) {
	// assert := require.New(t)
	// _ = assert
	// data := []byte(``)
	// _ = data
	// s, err := picker.NewSearch(picker.SearchParams{})
	// assert.NoError(err)
	// sd, err := s.MarshalJSON()
	// assert.NoError(err)
	// 	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)

}
