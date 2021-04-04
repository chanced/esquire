package picker

import (
	"fmt"
	"strings"

	"github.com/chanced/dynamic"
)

// IndexOptions is an option to the the index_options parameter controls what
// information is added to the inverted index for search and highlighting
// purposes.
type IndexOptions string

const (
	// IndexOptionsDocs - Only the doc number is indexed. Can answer the
	// question Does this term exist in this field?
	IndexOptionsDocs IndexOptions = "docs"
	// IndexOptionsFreqs - Doc number and term frequencies are indexed. Term
	// frequencies are used to score repeated terms higher than single terms.
	IndexOptionsFreqs IndexOptions = "freqs"
	// IndexOptionsPositions - Doc number, term frequencies, and term positions
	// (or order) are indexed. Positions can be used for proximity or phrase
	// queries.
	IndexOptionsPositions IndexOptions = "positions"
	// IndexOptionsOffsets - Doc number, term frequencies, positions, and start and end character offsets
	// (which map the term back to the original string) are indexed. Offsets are
	// used by the unified highlighter to speed up highlighting.
	IndexOptionsOffsets IndexOptions = "offsets"
)

func (io IndexOptions) String() string {
	return string(io)
}

var allIndexOptions = []IndexOptions{
	IndexOptionsDocs,
	IndexOptionsFreqs,
	IndexOptionsPositions,
	IndexOptionsOffsets,
}
var allIndexOptionsStr = strings.Join([]string{
	IndexOptionsDocs.String(),
	IndexOptionsFreqs.String(),
	IndexOptionsPositions.String(),
	IndexOptionsOffsets.String(),
}, ", ")

// WithIndexOptions is a mapping wit the index_options parameter
//
// The index_options parameter controls what information is added to the
// inverted index for search and highlighting purposes.
//
//  "docs"
// Only the doc number is indexed. Can answer the question Does this term exist
// in this field?
//
//  "freqs"
// Doc number and term frequencies are indexed. Term frequencies are used to
// score repeated terms higher than single terms.
//  "positions" (default)
// Doc number, term frequencies, and term positions (or order) are indexed.
// Positions can be used for proximity or phrase queries.
//
//  "offsets"
// Doc number, term frequencies, positions, and start and end character offsets
// (which map the term back to the original string) are indexed. Offsets are
// used by the unified highlighter to speed up highlighting.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-options.html
type WithIndexOptions interface {
	//IndexOptions parameter controls what information is added to the inverted
	//index for search and highlighting purposes.
	IndexOptions() IndexOptions
	// SetIndexOptions sets IndexOptions value to v
	SetIndexOptions(v IndexOptions) error
}

// FieldWithIndexOptions is a Field mapping with the index_options parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-options.html
type FieldWithIndexOptions interface {
	Field
	WithIndexOptions
}

// indexOptionsParam is a mixin that adds the index_options param to mappings
//
// The index_options parameter controls what information is added to the
// inverted index for search and highlighting purposes.
//
//  "docs"
// Only the doc number is indexed. Can answer the question Does this term exist
// in this field?
//  "freqs"
// Doc number and term frequencies are indexed. Term frequencies are used to
// score repeated terms higher than single terms.
//  "positions" (default)
// Doc number, term frequencies, and term positions (or order) are indexed.
// Positions can be used for proximity or phrase queries.
//  "offsets"
// Doc number, term frequencies, positions, and start and end character offsets
// (which map the term back to the original string) are indexed. Offsets are
// used by the unified highlighter to speed up highlighting.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-options.html
type indexOptionsParam struct {
	indexOptions IndexOptions
}

//IndexOptions parameter controls what information is added to the inverted
//index for search and highlighting purposes.
func (io indexOptionsParam) IndexOptions() IndexOptions {
	if len(io.indexOptions) == 0 {
		return IndexOptionsPositions
	}
	return io.indexOptions
}

// SetIndexOptions sets IndexOptions value to v
func (io *indexOptionsParam) SetIndexOptions(v IndexOptions) error {
	if io.IndexOptions() == v {
		return nil
	}
	if len(io.IndexOptions()) == 0 {
		io.indexOptions = ""
	}
	str := dynamic.NewString(v)
	iov := IndexOptions(str.ToLower().String())
	for _, x := range allIndexOptions {
		if x == iov {
			io.indexOptions = v
			return nil
		}
	}
	return fmt.Errorf("%w: expected one of: [%s]; received: %s",
		ErrInvalidIndexOptions, allIndexOptionsStr, v.String(),
	)
}
