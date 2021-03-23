package search

type AllOf struct {
}

func (ao AllOf) Type() Type {
	return TypeAllOf
}

func (ao AllOf) Rule() (Rule, error) {
	return &AllOf{}, nil
}
