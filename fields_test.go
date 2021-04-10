package picker_test

import (
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

	data := []byte(`{
        "type" : "keyword",
        "index" : false
        }`)

	var keyword picker.Field = &picker.KeywordField{}
	err = json.Unmarshal(data, keyword)
	assert.NoError(err)
	assert.False(keyword.(*picker.KeywordField).Index())

	data = []byte(`{
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

	m1 := picker.Mappings{}
	err = json.Unmarshal(data, &m1)
	assert.NoError(err)
	emplID, err := m1.Properties.Field("employee-id")
	assert.NoError(err)
	assert.NotNil(emplID, "employee-id should have been parsed and aded to the Properties Field map")
	emplIDAsKeyword, ok := emplID.(*picker.KeywordField)
	assert.True(ok, "employee-id should be unmarshaled and a KeywordField")
	assert.False(emplIDAsKeyword.Index(), "index value should be false")

	email, err := m1.Properties.Field("email")
	assert.NoError(err)
	assert.NotNil(email)
	assert.Equal(picker.FieldTypeKeyword, email.Type())
	name, err := m1.Properties.Field("name")
	assert.NoError(err)
	assert.NotNil(name)
	assert.Equal(picker.FieldTypeText, name.Type())

	_, err = m1.Properties.Field("not_exist")
	assert.ErrorIs(err, picker.ErrFieldNotFound)

}
