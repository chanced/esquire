package mapping

// WithPreserveSeperators is a mapping with the preserve_seperators parameter
//
// Preserves the separators, defaults to true. If disabled, you could find a
// field starting with Foo Fighters, if you suggest for foof.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type WithPreserveSeperators interface {
	// PreserveSeperators preserves the separators, defaults to true. If
	// disabled, you could find a field starting with Foo Fighters, if you
	// suggest for foof.
	PreserveSeperators() bool
	// SetPreserveSeperators sets the PreserveSeperatorParam value to v
	SetPreserveSeperators(v bool)
}

// FieldWithPreserveSeperators is a Field mapping with the preserve_seperators
// parameter
type FieldWithPreserveSeperators interface {
	Field
	WithPreserveSeperators
}

// PreserveSeperatorsParam is a mixin that adds the preserve_separators param
//
// Preserves the separators, defaults to true. If disabled, you could find a
// field starting with Foo Fighters, if you suggest for foof.
type PreserveSeperatorsParam struct {
	PreserveSeperatorsValue *bool `bson:"preserve_separators,omitempty" json:"preserve_separators,omitempty"`
}

// PreserveSeperators preserves the separators, defaults to true. If
// disabled, you could find a field starting with Foo Fighters, if you
// suggest for foof.
func (ps PreserveSeperatorsParam) PreserveSeperators() bool {
	if ps.PreserveSeperatorsValue == nil {
		return true
	}
	return *ps.PreserveSeperatorsValue
}

// SetPreserveSeperators sets the PreserveSeperatorParam value to v
func (ps *PreserveSeperatorsParam) SetPreserveSeperators(v bool) {
	ps.PreserveSeperatorsValue = &v
}
