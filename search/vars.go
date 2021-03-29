package search

import "reflect"

var (
	trueBytes     = []byte("true")
	falseBytes    = []byte("false")
	typeString    = reflect.TypeOf("")
	typeByteSlice = reflect.TypeOf([]byte(""))
	typeInt64     = reflect.TypeOf(int(0))
	typeFloat64   = reflect.TypeOf(float64(0))
	typeBool      = reflect.TypeOf(true)
)
