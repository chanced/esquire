package picker

// WithIncludeInParent is a mapping with the include_in_parent parameter
//
// (Optional, Boolean) If true, all fields in the nested object are also added
// to the parent document as standard (flat) fields. Defaults to false.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type WithIncludeInParent interface {
	// IncludeInParent (Optional, Boolean) If true, all fields in the nested
	// object are also added to the parent document as standard (flat) fields.
	// Defaults to false.
	IncludeInParent() bool
	// SetIncludeInParent sets the IncludeInParent Value to v
	SetIncludeInParent(v bool)
}

type FieldWithIncludeInParent interface {
	Field
	WithIncludeInParent
}

// IncludeInParentParam is a mixin that adds the include_in_parent param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type IncludeInParentParam struct {
	IncludeInParentValue *bool `bson:"include_in_parent,omitempty" json:"include_in_parent,omitempty"`
}

// IncludeInParent (Optional, Boolean) If true, all fields in the nested
// object are also added to the parent document as standard (flat) fields.
// Defaults to false.
func (iip IncludeInParentParam) IncludeInParent() bool {
	if iip.IncludeInParentValue == nil {
		return false
	}
	return *iip.IncludeInParentValue
}

// SetIncludeInParent sets the IncludeInParent Value to v
func (iip *IncludeInParentParam) SetIncludeInParent(v bool) {
	if iip.IncludeInParent() != v {
		iip.IncludeInParentValue = &v
	}
}
