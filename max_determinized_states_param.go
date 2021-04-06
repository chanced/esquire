package picker

import "github.com/chanced/dynamic"

const DefaultMaxDeterminizedStates = 10000

type WithMaxDeterminizedStates interface {
	MaxDeterminizedStates() int
	SetMaxDeterminizedStates(v interface{}) error
}

type maxDeterminizedStatesParam struct {
	maxDeterminizedStates dynamic.Number
}

func (mds maxDeterminizedStatesParam) MaxDeterminizedStates() int {
	if i, ok := mds.maxDeterminizedStates.Int(); ok {
		return i
	}
	if f, ok := mds.maxDeterminizedStates.Float64(); ok {
		return int(f)
	}
	return DefaultMaxDeterminizedStates
}
func (mds *maxDeterminizedStatesParam) SetMaxDeterminizedStates(v interface{}) error {
	return mds.maxDeterminizedStates.Set(v)
}
