package search

type AllOf struct {
	QueryName string
}

func (ao AllOf) Name() string {
	return ao.QueryName
}
func (ao AllOf) SetName(name string) {
	ao.QueryName = name
}

func (ao AllOf) Kind() Kind {
	return KindAllOf
}

func (ao AllOf) Clause() (Clause, error) {
	return &AllOf{}, nil
}
