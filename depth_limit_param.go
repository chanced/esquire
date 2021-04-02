package picker

import "github.com/chanced/dynamic"

var DepthLimitDefault = float64(20)

type WithDepthLimit interface {
	DepthLimit() float64
	SetDepthLimit(v interface{}) error
}

type depthLimitParam struct {
	depthLimit dynamic.Number
}

func (dl depthLimitParam) DepthLimit() float64 {
	if f, ok := dl.depthLimit.Float64(); ok {
		return f
	}
	return DepthLimitDefault
}

func (dl depthLimitParam) SetDepthLimit(v interface{}) error {
	return dl.depthLimit.Set(v)
}
