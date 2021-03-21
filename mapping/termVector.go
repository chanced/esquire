package mapping

// TermVector indicates which Term vectors, which contain information about the
// terms produced by the analysis process, including:
//
// - a list of terms.
//
// - the position (or order) of each term.
//
// - the start and end character offsets mapping the term to its origin in the original string.
//
// - payloads (if they are available) — user-defined binary data associated with each term position.
//
// - These term vectors can be stored so that they can be retrieved for a particular document.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/term-vector.html
type TermVector string

func (tv TermVector) String() string {
	return string(tv)
}

const (
	// TermVectorNone - No term vectors are stored. (default)
	TermVectorNone TermVector = "no"

	// TermVectorTerms - Just the terms in the field are stored.
	TermVectorTerms TermVector = "yes"

	// TermVectorWithPositions  - Terms and positions are stored.
	TermVectorWithPositions TermVector = "with_positions"

	// TermVectorWithOffsets - Terms and character offsets are stored.
	TermVectorWithOffsets TermVector = "with_offsets"

	// TermVectorWithPositionsOffsets  - Terms, positions, and character offsets are stored.
	TermVectorWithPositionsOffsets TermVector = "with_positions_offsets"

	// TermVectorWithPoisitionsPayloads - Terms, positions, and payloads are stored.
	TermVectorWithPoisitionsPayloads TermVector = "with_positions_payloads"

	// TTermVectorPositionsOffsetsPayloads - Terms, positions, offsets and payloads are stored
	TermVectorPositionsOffsetsPayloads TermVector = "with_positions_offsets_payloads"
)

// WithTermVector is a mapping with the term_vector parameter
//
// Term vectors contain information about the terms produced by the analysis
// process, including:
//
//
// - a list of terms.
//
// - the position (or order) of each term.
//
// - the start and end character offsets mapping the term to its origin in the
// original string.
//
// - payloads (if they are available) — user-defined binary data associated with
// each term position.
//
// These term vectors can be stored so that they can be retrieved for a
// particular document.
//
// WARNING
//
// Setting with_positions_offsets will double the size of a field’s index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/term-vector.html
type WithTermVector interface {
	// TermVector determines whether term vectors should be stored for the field.
	// Defaults to "no" / TermVectorNone.
	TermVector() TermVector
	// SetTermVector sets the TermVector Value to v
	SetTermVector(v TermVector)
}

// FieldWIthTermVector is a Field with the term_vector paramater
type FieldWIthTermVector interface {
	Field
	WithTermVector
}

// TermVectorParam is a mixin that adds the term_vector parameter
type TermVectorParam struct {
	TermVectorValue *TermVector `bson:"term_vector,omitempty" json:"term_vector,omitempty"`
}

// TermVector determines whether term vectors should be stored for the field.
// Defaults to "no" / TermVectorNone.
func (tv TermVectorParam) TermVector() TermVector {
	if tv.TermVectorValue == nil {
		return TermVectorNone
	}
	return *tv.TermVectorValue
}

// SetTermVector sets the TermVector Value to v
func (tv *TermVectorParam) SetTermVector(v TermVector) {
	tv.TermVectorValue = &v
}
