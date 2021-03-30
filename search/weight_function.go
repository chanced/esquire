package search

type WeightFunc struct {
	// Weight
	Weight float64
	Filter CompleteClauser
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
	if err != nil {
		return f, err
	}
	err = f.SetFilter(w.Filter)
	return f, err
}

type WeightFunction struct {
	weightParam
	filter QueryClause
}

func (WeightFunction) FunctionKind() FunctionKind {
	return FuncKindWeight
}
func (w WeightFunction) Filter() QueryClause {
	return w.filter
}
func (w *WeightFunction) SetFilter(c CompleteClauser) error {
	if c == nil {
		w.filter = nil
		return nil
	}
	qc, err := c.Clause()
	if err != nil {
		return err
	}
	w.filter = qc
	return nil
}
