package search

type Weight struct{}

func (w Weight) ScoreKind() ScoreKind {
	return ScoreWeight
}
func (w Weight) Function() (Function, error) {
	return &WeightFunction{}, nil
}

type WeightFunction struct{}

func (WeightFunction) FunctionKind() FunctionKind {
	return FunctionKindWeight
}
