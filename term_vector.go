package picker

import (
	"fmt"
	"strings"
)

const DefaultTermVector = TermVectorNo

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

func (tv *TermVector) Validate() error {
	if tv.IsValid() {
		return nil
	}
	strs := make([]string, len(termVectorValues))
	for i, k := range termVectorValues {
		strs[i] = k.String()
	}
	return fmt.Errorf("%w; expected one of [%s]", ErrInvalidTermVector, strings.Join(strs, ", "))
}
func (tv *TermVector) IsValid() bool {
	if len(*tv) == 0 {
		return true
	}

	tvv := tv.toLower()
	for _, v := range termVectorValues {
		if tvv == v {
			return true
		}
	}
	return false
}
func (tv *TermVector) toLower() TermVector {
	*tv = TermVector(strings.ToLower(string(*tv)))
	return *tv
}

var termVectorValues = []TermVector{
	TermVectorNo,
	TermVectorYes,
	TermVectorWithPositions,
	TermVectorWithOffsets,
	TermVectorWithPositionsOffsets,
	TermVectorWithPoisitionsPayloads,
	TermVectorPositionsOffsetsPayloads,
}

const (
	// TermVectorNone - No term vectors are stored. (default)
	TermVectorNo TermVector = "no"

	// TermVectorYes - Just the terms in the field are stored.
	TermVectorYes TermVector = "yes"

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
	SetTermVector(v TermVector) error
}

// FieldWithTermVector is a Field with the term_vector paramater
type FieldWithTermVector interface {
	Field
	WithTermVector
}

// termVectorParam is a mixin that adds the term_vector parameter
type termVectorParam struct {
	termVector TermVector
}

// TermVector determines whether term vectors should be stored for the field.
// Defaults to "no" / TermVectorNone.
func (tv termVectorParam) TermVector() TermVector {
	if len(tv.termVector) > 0 {
		return tv.termVector
	}
	return DefaultTermVector
}

// SetTermVector sets the TermVector Value to v
func (tv *termVectorParam) SetTermVector(v TermVector) error {
	err := v.Validate()
	if err != nil {
		return fmt.Errorf("%w; received %s", v)
	}
	tv.termVector = v
	return nil
}
