package picker

import "github.com/chanced/dynamic"

const DefaultMaxInputLength = 50

// WithMaxInputLength is a mapping with the max_input_length parameter
//
// Limits the length of a single input, defaults to 50 UTF-16 code points. This
// limit is only used at index time to reduce the total number of characters per
// input string in order to prevent massive inputs from bloating the underlying
// datastructure. Most use cases won’t be influenced by the default value since
// prefix completions seldom grow beyond prefixes longer than a handful of
// characters.
type WithMaxInputLength interface {
	// MaxInputLength limits the length of a single input, defaults to 50 UTF-16
	// code points. This limit is only used at index time to reduce the total
	// number of characters per input string in order to prevent massive inputs
	// from bloating the underlying datastructure. Most use cases won’t be
	// influenced by the default value since prefix completions seldom grow
	// beyond prefixes longer than a handful of characters.
	MaxInputLength() int

	// SetMaxInputLength sets the MaxInputLength value to v
	SetMaxInputLength(v interface{}) error
}

type maxInputLengthParam struct {
	maxInputLength dynamic.Number
}

// MaxInputLength limits the length of a single input, defaults to 50 UTF-16
// code points. This limit is only used at index time to reduce the total
// number of characters per input string in order to prevent massive inputs
// from bloating the underlying datastructure. Most use cases won’t be
// influenced by the default value since prefix completions seldom grow
// beyond prefixes longer than a handful of characters.
func (mil maxInputLengthParam) MaxInputLength() int {
	if i, ok := mil.maxInputLength.Int(); ok {
		return int(i)
	}
	return DefaultMaxInputLength
}

// SetMaxInputLength sets the MaxInputLength value to v
func (mil *maxInputLengthParam) SetMaxInputLength(v interface{}) error {
	return mil.maxInputLength.Set(v)
}
