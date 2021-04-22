package picker_test

import (
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestScriptQuery(t *testing.T) {
	assert := require.New(t)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			Script: &picker.ScriptQueryParams{
				Lang:   "painless",
				Name:   "script_query_name",
				Source: "doc['num1'].value > 1",
				Params: MyStruct{Color: "blue", Size: 34},
			},
		},
	})
	assert.NoError(err)
	assert.NotNil(s.Query().Script())
	assert.Equal("painless", s.Query().Script().Lang())
	assert.Equal("doc['num1'].value > 1", s.Query().Script().Source())

}
