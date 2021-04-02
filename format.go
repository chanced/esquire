package picker

type WithFormat interface {
	//The date format(s) that can be parsed. Defaults to
	//strict_date_optional_time||epoch_millis.
	Format() string
	// SetFormat sets the Format Value to v
	SetFormat(v string)
}

// FieldWithFormat is a Field with a format parameter
type FieldWithFormat interface {
	Field
	WithFormat
}

type FormatParam struct {
	FormatValue string `bson:"format,omitempty" json:"format,omitempty"`
}

//Format is the format(s) that the that can be parsed. Defaults to strict_date_optional_time||epoch_millis.
//
// Multiple formats can be seperated by ||
func (f FormatParam) Format() string {
	if f.FormatValue == "" {
		return "strict_date_optional_time||epoch_millis"
	}
	return f.FormatValue
}

func (f *FormatParam) SetFormat(v string) {
	if v != f.Format() {
		f.FormatValue = v
	}
}
