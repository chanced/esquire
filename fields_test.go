package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/picker/mapping"
	"github.com/stretchr/testify/require"
)

func TestJSON(t *testing.T) {

	var err error
	assert := require.New(t)

	handler := mapping.FieldTypeHandlers[mapping.FieldTypeInteger]
	hv := handler()
	assert.NotNil(handler)
	assert.NotNil(hv)

	kwraw := []byte(`{
        "type" : "keyword",
        "index" : false
        }`)

	var kwf mapping.Field = &mapping.KeywordField{}
	err = json.Unmarshal(kwraw, kwf)
	assert.NoError(err)
	assert.False(kwf.(*mapping.KeywordField).Index())
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
	m1 := mapping.Mappings{}
	err = json.Unmarshal(raw1, &m1)
	assert.NoError(err)

	fmt.Printf("%+v", m1.Properties["employee-id"])
	emplID := m1.Properties.Field("employee-id")
	assert.NotNil(emplID, "employee-id should have been parsed and aded to the Properties Field map")
	emplIDAsKeyword, ok := emplID.(*mapping.KeywordField)
	assert.True(ok, "employee-id should be unmarshaled and a KeywordField")
	assert.False(emplIDAsKeyword.Index(), "index value should be false")
}
