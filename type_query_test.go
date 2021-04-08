package picker_test

import (
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestType(t *testing.T) {
	assert := require.New(t)
	_ = assert
	data := []byte(``)
	_ = data
	s, err := picker.NewSearch(picker.SearchParams{})
	assert.NoError(err)
	sd, err := s.MarshalJSON()
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))
}
