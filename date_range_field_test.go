package picker_test

import (
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestDateRangeField(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"mappings": {
		  "properties": {
			"time_frame": {
			  "type": "date_range", 
			  "format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
			}
		  }
		}
	  }`)
	i, err := picker.NewIndex(picker.IndexParams{Mappings: picker.Mappings{
		Properties: picker.FieldMap{
			"time_frame": picker.DateRangeFieldParams{
				Format: "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
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
