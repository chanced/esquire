package search

// Funcs is a slice of Functioners, valid options are:
//
//  - search.WeightFunc,
//  - search.DecayFunc,
//  - search.RandomScoreFunc,
//  - search.ScriptScoreFunc,
//  -
type Funcs []Functioner

type Function interface {
	FunctionKind() FunctionKind
	Weight() float64
	Filter() Clause
}
type Functions []Function

type Functioner interface {
	Function() (Function, error)
}

type FunctionKind string

func (f FunctionKind) String() string {
	return string(f)
}
