package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestCommon(t *testing.T) {
	assert := require.New(t)
	_ = assert
	data := []byte(``)
	_ = data
	s, err := picker.NewSearch(picker.SearchParams{})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
}
