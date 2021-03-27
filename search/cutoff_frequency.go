package search

import (
	"github.com/chanced/dynamic"
)

type WithCutoffFrequency interface {
	CutoffFrequency() dynamic.Number
	SetCutoffFreuency(dynamic.Number)
}

type cutoffFrequencyParam struct {
	cutoffFrequency dynamic.Number
}

func (cf cutoffFrequencyParam) CutoffFrequency() dynamic.Number {
	return cf.cutoffFrequency
}
func (cf *cutoffFrequencyParam) SetCutoffFreuency(n dynamic.Number) {
	cf.cutoffFrequency = n
}

func unmarshalCutoffFrequencyParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithCutoffFrequency); ok {
		n := dynamic.NewNumber(data.UnquotedString())
		a.SetCutoffFreuency(n)
	}
	return nil
}
func marshalCutoffFrequencyParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if p, ok := source.(WithCutoffFrequency); ok {
		if p.CutoffFrequency().HasValue() {
			data[paramCutoffFrequency] = p.CutoffFrequency()
		}
	}
	return data, nil
}
