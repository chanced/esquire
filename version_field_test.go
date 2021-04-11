package picker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionField(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
      "mappings": {
        "properties": {
            "version": {
                "type": "version"
           }
        }
     }
  }`)
	_ = data
	_ = assert
	// i, err := picker.NewIndex(picker.IndexParams{Mappings: picker.Mappings{
	// 	Properties: picker.FieldMap{
	// 		"version": picker.VersionFieldParams{}},
	// }})
	// assert.NoError(err)
	// ixd, err := i.MarshalJSON()
	// assert.NoError(err)
	// assert.True(cmpjson.Equal(data, ixd), cmpjson.Diff(data, ixd))
	// i2 := picker.Index{}
	// err = i2.UnmarshalJSON(data)
	// assert.NoError(err)

}
