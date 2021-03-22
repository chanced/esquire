package search

type Range struct {
}

func (r Range) Type() Type {
	return TypeBoolean
}

type RangeQuery struct {
}
