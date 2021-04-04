package picker

import "github.com/chanced/dynamic"

const DefaultPreserveSeperators = true

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
	SetPreserveSeperators(v interface{}) error
}

type preserveSeperatorsParam struct {
	preserveSeperators dynamic.Bool
}

// PreserveSeperators preserves the separators, defaults to true. If
// disabled, you could find a field starting with Foo Fighters, if you
// suggest for foof.
func (ps preserveSeperatorsParam) PreserveSeperators() bool {
	if b, ok := ps.preserveSeperators.Bool(); ok {
		return b
	}
	return DefaultPreserveSeperators
}

// SetPreserveSeperators sets the preserve_seperator to v
func (ps *preserveSeperatorsParam) SetPreserveSeperators(v interface{}) error {
	return ps.preserveSeperators.Set(v)
}
