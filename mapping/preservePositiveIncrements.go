package mapping

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

// PreservePositionIncrementsParam is a mixin that adds the
// preserve_position_increments param
//
// Enables position increments, defaults to true. If disabled and using
// stopwords analyzer, you could get a field starting with The Beatles, if you
// suggest for b. Note: You could also achieve this by indexing two inputs,
// Beatles and The Beatles, no need to change a simple analyzer, if you are able
// to enrich your data.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type PreservePositionIncrementsParam struct {
	PreservePositionIncrementsValue *bool `bson:"preserve_position_increments,omitempty" json:"preserve_position_increments,omitempty"`
}

// PreservePositionIncrements enables position increments, defaults to true.
// If disabled and using stopwords analyzer, you could get a field starting
// with The Beatles, if you suggest for b. Note: You could also achieve this
// by indexing two inputs, Beatles and The Beatles, no need to change a
// simple analyzer, if you are able to enrich your data.
func (ppi PreservePositionIncrementsParam) PreservePositionIncrements() bool {
	if ppi.PreservePositionIncrementsValue == nil {
		return true
	}
	return *ppi.PreservePositionIncrementsValue
}

// SetPreservePositionIncrements sets the PreservePositionIncrements value
// to V
func (ppi *PreservePositionIncrementsParam) SetPreservePositionIncrements(v bool) {
	ppi.PreservePositionIncrementsValue = &v
}
