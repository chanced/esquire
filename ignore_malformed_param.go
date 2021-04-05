package picker

import "github.com/chanced/dynamic"

const DefaultIgnoreMalformed = false

// WithIgnoreMalformed is a mapping with the ignore_malformed parameter
//
// Sometimes you don’t have much control over the data that you receive. One
// user may send a login field that is a date, and another sends a login field
// that is an email address.
//
// Trying to index the wrong data type into a field throws an exception by
// default, and rejects the whole document. The ignore_malformed parameter, if
// set to true, allows the exception to be ignored. The malformed field is not
// indexed, but other fields in the document are processed normally.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ignore-malformed.html
type WithIgnoreMalformed interface {
	// IgnoreMalformed determines if malformed numbers are ignored. If true,
	// malformed numbers are ignored. If false (default), malformed numbers
	// throw an exception and reject the whole document.
	IgnoreMalformed() bool
	// SetIgnoreMalformed sets IgnoreMalformed to v
	SetIgnoreMalformed(v interface{}) error
}

type ignoreMalformedParam struct {
	ignoreMalformed dynamic.Bool
}

// IgnoreMalformed determines if malformed numbers are ignored. If true,
// malformed numbers are ignored. If false (default), malformed numbers throw an
// exception and reject the whole document.
func (im ignoreMalformedParam) IgnoreMalformed() bool {
	if v, ok := im.ignoreMalformed.Bool(); ok {
		return v
	}
	return DefaultIgnoreMalformed
}

// SetIgnoreMalformed sets ignore_malformed to v
func (im *ignoreMalformedParam) SetIgnoreMalformed(v interface{}) error {
	return im.ignoreMalformed.Set(v)
}
