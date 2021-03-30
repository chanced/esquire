package search

type GaussFunc struct {
	// Weight
	Weight interface{}
	Filter Clause
}

func (GaussFunc) FuncKind() FunctionKind {
	return FuncKindGauss
}
func (w GaussFunc) Function() (Function, error) {
	f := &GaussFunction{}
	err := f.SetWeight(w.Weight)
	return f, err
}

type GaussFunction struct {
	weightParam
	filter Clause
}

func (GaussFunction) FunctionKind() FunctionKind {
	return FuncKindGauss
}
func (w GaussFunction) Filter() Clause {
	return w.filter
}
