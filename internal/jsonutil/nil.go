package jsonutil

import (
	"bytes"
	"encoding/json"
)

var Nil = func() json.RawMessage {
	n, err := json.Marshal(nil)
	if err != nil {
		panic(err)
	}
	return n
}()

func IsNil(data []byte) bool {
	return bytes.Equal(Nil, data)
}

func IsNotNil(data []byte) bool {
	return !IsNil(data)
}
