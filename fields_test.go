package picker_test

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {

	var err error
	assert := require.New(t)

	handler := picker.FieldTypeHandlers[picker.FieldTypeInteger]
	hv := handler()
	assert.NotNil(handler)
	assert.NotNil(hv)

	kwraw := []byte(`{
        "type" : "keyword",
        "index" : false
        }`)

	var kwf picker.Field = &picker.KeywordField{}
	err = json.Unmarshal(kwraw, kwf)
	assert.NoError(err)
	assert.False(kwf.(*picker.KeywordField).Index())
	raw1 := []byte(`{
        "properties" : {
            "age" : {
            "type" : "integer"
            },
            "email" : {
            "type" : "keyword"
            },
            "employee-id" : {
            "type" : "keyword",
            "index" : false
            },
            "name" : {
            "type" : "text"
            }
        }
    }`)
	_ = raw1
	m1 := picker.Mappings{}
	err = json.Unmarshal(raw1, &m1)
	assert.NoError(err)

	fmt.Printf("%+v", m1.Properties["employee-id"])
	emplID := m1.Properties.Field("employee-id")
	assert.NotNil(emplID, "employee-id should have been parsed and aded to the Properties Field map")
	emplIDAsKeyword, ok := emplID.(*picker.KeywordField)
	assert.True(ok, "employee-id should be unmarshaled and a KeywordField")
	assert.False(emplIDAsKeyword.Index(), "index value should be false")
}
