package picker_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestCompareJSONObject(t *testing.T) {
	assert := require.New(t)
	a := []byte(`{ "v": "v"}`)
	b := []byte(`{ "v": "v"}`)
	assert.NoError(compareJSONObject(a, b))
	a = []byte(`{ "v": "v"}`)
	b = []byte(`{ "b": "v"}`)
	assert.Error(compareJSONObject(a, b))
	a = []byte(`{ "t": true, "v": 
	{"s": 
	["a"]}}`)
	b = []byte(`{ "t": true, 
		"v": {"s": ["a"]}}`)
	assert.NoError(compareJSONObject(a, b))
	a = []byte(`{ "t": true, "v": 
	{"s": 
	["a"]}}`)
	b = []byte(`{ "t": true, 
		"v": {"x": ["a"]}}`)
	assert.Error(compareJSONObject(a, b))
}

func compareJSONObject(a, b []byte) error {
	var am map[string]interface{}
	err := json.Unmarshal(a, &am)
	if err != nil {
		return err
	}
	var bm map[string]interface{}
	err = json.Unmarshal(b, &bm)
	if err != nil {
		return err
	}
	if !cmp.Equal(am, bm) {
		diff := cmp.Diff(am, bm)
		return errors.New(diff)
	}
	return nil
}
