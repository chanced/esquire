package mapping

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
	SetMaxInputLength(v int)
}

// FieldWithMaxInputLength is a Field mapping with the max_input_length parameter
type FieldWithMaxInputLength interface {
	Field
	WithMaxInputLength
}

type MaxInputLengthParam struct {
	MaxInputLengthValue *int `bson:"max_input_length,omitempty" json:"max_input_length,omitempty"`
}

// MaxInputLength limits the length of a single input, defaults to 50 UTF-16
// code points. This limit is only used at index time to reduce the total
// number of characters per input string in order to prevent massive inputs
// from bloating the underlying datastructure. Most use cases won’t be
// influenced by the default value since prefix completions seldom grow
// beyond prefixes longer than a handful of characters.
func (mil MaxInputLengthParam) MaxInputLength() int {
	if mil.MaxInputLengthValue == nil {
		return 50
	}
	return *mil.MaxInputLengthValue
}

// SetMaxInputLength sets the MaxInputLength value to v
func (mil *MaxInputLengthParam) SetMaxInputLength(v int) {
	if mil.MaxInputLength() != v {
		mil.MaxInputLengthValue = &v

	}
}
