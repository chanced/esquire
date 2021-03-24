package search

import "github.com/tidwall/gjson"

type WithPrefixLength interface {
	PrefixLength() int
	SetPrefixLength(v int)
}

type PrefixLengthParam struct {
	PrefixLengthValue *int `json:"prefix_length,omitempty" bson:"prefix_length,omitempty"`
}

func (pl PrefixLengthParam) PrefixLength() int {
	if pl.PrefixLengthValue == nil {
		return 0
	}
	return *pl.PrefixLengthValue
}

func (pl *PrefixLengthParam) SetPrefixLength(v int) {
	pl.PrefixLengthValue = &v
}
func unmarshalPrefixLengthParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithPrefixLength); ok {
		a.SetPrefixLength(int(value.Int()))
	}
	return nil
}
