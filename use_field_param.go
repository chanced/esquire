package picker

type useFieldParam struct {
	useField string
}

func (uf useFieldParam) UseField() string {
	return uf.useField
}
func (uf *useFieldParam) SetUseField(field string) {
	uf.useField = field
}
