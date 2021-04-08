package picker

import (
	"github.com/chanced/dynamic"
)

func unmarshalField(data []byte) (string, dynamic.JSON, error) {
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return "", nil, err
	}
	for k, v := range obj {
		return k, v, nil
	}
	return "", nil, nil
}
