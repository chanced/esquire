package picker

import "github.com/chanced/dynamic"

const DefaultPreservePositionIncrements = true

// WithPreservePositionIncrements is a mapping with the
// preserve_position_increments parameter
//
// Enables position increments, defaults
// to true. If disabled and using stopwords analyzer, you could get a field
// starting with The Beatles, if you suggest for b. Note: You could also achieve
// this by indexing two inputs, Beatles and The Beatles, no need to change a
// simple analyzer, if you are able to enrich your data.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type WithPreservePositionIncrements interface {
	// PreservePositionIncrements enables position increments, defaults to true.
	// If disabled and using stopwords analyzer, you could get a field starting
	// with The Beatles, if you suggest for b. Note: You could also achieve this
	// by indexing two inputs, Beatles and The Beatles, no need to change a
	// simple analyzer, if you are able to enrich your data.
	PreservePositionIncrements() bool
	// SetPreservePositionIncrements sets the PreservePositionIncrements value
	// to V
	SetPreservePositionIncrements(v bool)
}

type FieldWithPreservePositionIncrements interface {
	Field
	WithPreservePositionIncrements
}

type preservePositionIncrementsParam struct {
	preservePositionIncrements dynamic.Bool
}

// PreservePositionIncrements enables position increments, defaults to true.
// If disabled and using stopwords analyzer, you could get a field starting
// with The Beatles, if you suggest for b. Note: You could also achieve this
// by indexing two inputs, Beatles and The Beatles, no need to change a
// simple analyzer, if you are able to enrich your data.
func (ppi preservePositionIncrementsParam) PreservePositionIncrements() bool {
	if b, ok := ppi.preservePositionIncrements.Bool(); ok {
		return b
	}
	return DefaultPreserveSeperators
}

// SetPreservePositionIncrements sets the PreservePositionIncrements value
// to V
func (ppi *preservePositionIncrementsParam) SetPreservePositionIncrements(v interface{}) error {
	return ppi.preservePositionIncrements.Set(v)
}
