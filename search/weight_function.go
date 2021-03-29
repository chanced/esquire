package search

type WeightFunc struct {
	// Weight
	Weight float64
	Filter Clause
}

func (WeightFunc) FuncKind() FunctionKind {
	return FuncKindWeight
}
func (w WeightFunc) Function() (Function, error) {
	f := &WeightFunction{}
	if w.Weight == 0 {
		return f, ErrWeightRequired
	}
	err := f.SetWeight(w.Weight)
	return f, err
}

type WeightFunction struct {
	weightParam
	filter Clause
}

func (WeightFunction) FunctionKind() FunctionKind {
	return FuncKindWeight
}
func (w WeightFunction) Filter() Clause {
	return w.filter
}
