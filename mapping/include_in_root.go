package mapping

// WithIncludeInRoot is a mapping with the include_in_root parameter
//
// (Optional, Boolean) If true, all fields in the nested object are also added
// to the root document as standard (flat) fields. Defaults to false
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type WithIncludeInRoot interface {
	// IncludeInRoot deteremines if all fields in the nested object are also
	// added to the root document as standard (flat) fields. Defaults to false
	IncludeInRoot() bool
	// SetIncldueInRoot sets the IncludeInRoot Value to v
	SetIncludeInRoot(v bool)
}

// FieldWithIncludeInRoot is a Field with the include_in_root parameter
type FieldWithIncludeInRoot interface {
	Field
	WithIncludeInRoot
}

// IncludeInRootParam is a mixin that adds the include_in_root param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type IncludeInRootParam struct {
	IncludeInRootValue *bool `bson:"include_in_root,omitempty" json:"include_in_root,omitempty"`
}

// IncludeInRoot deteremines if all fields in the nested object are also
// added to the root document as standard (flat) fields. Defaults to false
func (iir IncludeInRootParam) IncludeInRoot() bool {
	if iir.IncludeInRootValue == nil {
		return false
	}
	return *iir.IncludeInRootValue
}

// SetIncludeInRoot sets the IncludeInRoot Value to v
func (iir *IncludeInRootParam) SetIncludeInRoot(v bool) {
	if iir.IncludeInRoot() != v {
		iir.IncludeInRootValue = &v
	}
}
