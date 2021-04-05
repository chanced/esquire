package picker

type fieldParam struct {
	field string
}

func (f fieldParam) Field() string {
	return f.field
}
func (f *fieldParam) SetField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	f.field = field
	return nil
}
