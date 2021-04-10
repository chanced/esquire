package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestIPField(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"mappings": {
		  "properties": {
			"ip_addr": {
			  "type": "ip"
			}
		  }
		}
	  }
	  `)
	i, err := picker.NewIndex(picker.IndexParams{Mappings: picker.Mappings{
		Properties: picker.FieldMap{
			"ip_addr": picker.IPFieldParams{},
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
