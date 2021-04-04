package picker

import "github.com/chanced/dynamic"

const DefaultIncludeInRoot = false

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
	SetIncludeInRoot(v interface{}) error
}

// includeInRootParam is a mixin that adds the include_in_root param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type includeInRootParam struct {
	includeInRoot dynamic.Bool
}

// IncludeInRoot deteremines if all fields in the nested object are also
// added to the root document as standard (flat) fields. Defaults to false
func (iir includeInRootParam) IncludeInRoot() bool {
	if b, ok := iir.includeInRoot.Bool(); ok {
		return b
	}
	return DefaultIncludeInRoot
}

// SetIncludeInRoot sets the IncludeInRoot Value to v
func (iir *includeInRootParam) SetIncludeInRoot(v interface{}) error {
	return iir.includeInRoot.Set(v)
}
