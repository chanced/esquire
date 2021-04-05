package picker

import "github.com/chanced/dynamic"

const DefaultIncludeInParent = false

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
	SetIncludeInParent(v interface{}) error
}

// includeInParentParam is a mixin that adds the include_in_parent param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type includeInParentParam struct {
	includeInParent dynamic.Bool
}

// IncludeInParent (Optional, Boolean) If true, all fields in the nested
// object are also added to the parent document as standard (flat) fields.
// Defaults to false.
func (iip includeInParentParam) IncludeInParent() bool {
	if b, ok := iip.includeInParent.Bool(); ok {
		return b
	}
	return DefaultIncludeInParent
}

// SetIncludeInParent sets the IncludeInParent Value to v
func (iip *includeInParentParam) SetIncludeInParent(v interface{}) error {
	return iip.includeInParent.Set(v)
}
