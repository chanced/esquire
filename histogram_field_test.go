package picker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHistogramField(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
      "mappings": {
        "properties": {
            "histogram": {
                "type": "histogram"
           }
        }
     }
  }`)
	_ = assert
	_ = data

	// i, err := picker.NewIndex(picker.IndexParams{Mappings: picker.Mappings{
	// 	Properties: picker.FieldMap{
	// 		"histogram": picker.HistogramFieldParams{}},
	// }})
	// assert.NoError(err)
	// ixd, err := i.MarshalJSON()
	// assert.NoError(err)
	// assert.True(cmpjson.Equal(data, ixd), cmpjson.Diff(data, ixd))
	// i2 := picker.Index{}
	// err = i2.UnmarshalJSON(data)
	// assert.NoError(err)

}
