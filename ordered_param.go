package picker

import "github.com/chanced/dynamic"

const DefaultOrdered = false

type WithOrdered interface {
	Ordered() bool
	SetOrdered(ordered interface{}) error
}

type orderedParam struct {
	ordered dynamic.Bool
}

func (o orderedParam) Ordered() bool {
	if b, ok := o.ordered.Bool(); ok {
		return b
	}
	return DefaultOrdered
}

func (o *orderedParam) SetOrdered(ordered interface{}) error {
	return o.ordered.Set(ordered)
}
