package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestConstantField(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"mappings": {
		  "properties": {
			"level": {
			  "type": "constant_keyword",
			  "value": "debug"
			}
		  }
		}
	  }`)
	i, err := picker.NewIndex(picker.IndexParams{Mappings: picker.Mappings{
		Properties: picker.FieldMap{
			"level": picker.ConstantFieldParams{
				Value: "debug",
			},
		},
	}})
	assert.NoError(err)
	ixd, err := i.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, ixd), cmpjson.Diff(data, ixd))
	i2 := picker.Index{}
	err = i2.UnmarshalJSON(data)
	assert.NoError(err)

}
