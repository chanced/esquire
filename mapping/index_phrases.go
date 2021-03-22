package mapping

// WithIndexPhrases is a mapping with the index_phrases parameter
//
// If IndexPhrases parameter is enabled, two-term word combinations (shingles)
// are indexed into a separate field. This allows exact phrase queries (no slop)
// to run more efficiently, at the expense of a larger index. Note that this
// works best when stopwords are not removed, as phrases containing stopwords
// will not use the subsidiary field and will fall back to a standard phrase
// query. Accepts true or false (default).
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-phrases.html#index-phrases
type WithIndexPhrases interface {
	// IndexPhrases is a param that if enabled, two-term word combinations
	// (shingles) are indexed into a separate field. This allows exact phrase
	// queries (no slop) to run more efficiently, at the expense of a larger index.
	//
	// Note that this works best when stopwords are not removed, as phrases
	// containing stopwords will not use the subsidiary field and will fall back to
	// a standard phrase query. Accepts true or false (default).
	IndexPhrases() bool
	// SetIndexPhrases sets the IndexPhrases value to v
	SetIndexPhrases(v bool)
}

// FieldWithIndexPhrases is a Field with the index_phrases parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-phrases.html#index-phrases
type FieldWithIndexPhrases interface {
	Field
	WithIndexPhrases
}

// IndexPhrasesParam is a mixin that adds the index_phrases parameter
//
// If IndexPhrases parameter is enabled, two-term word combinations (shingles)
// are indexed into a separate field. This allows exact phrase queries (no slop)
// to run more efficiently, at the expense of a larger index. Note that this
// works best when stopwords are not removed, as phrases containing stopwords
// will not use the subsidiary field and will fall back to a standard phrase
// query. Accepts true or false (default).
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-phrases.html#index-phrases
type IndexPhrasesParam struct {
	IndexPhrasesValue *bool `bson:"index_phrases,omitempty" json:"index_phrases,omitempty"`
}

// IndexPhrases is a param that if enabled, two-term word combinations
// (shingles) are indexed into a separate field. This allows exact phrase
// queries (no slop) to run more efficiently, at the expense of a larger index.
//
// Note that this works best when stopwords are not removed, as phrases
// containing stopwords will not use the subsidiary field and will fall back to
// a standard phrase query. Accepts true or false (default).
func (cp IndexPhrasesParam) IndexPhrases() bool {
	if cp.IndexPhrasesValue != nil {
		return *cp.IndexPhrasesValue
	}
	return false
}

// SetIndexPhrases sets the IndexPhrases value to v
func (cp *IndexPhrasesParam) SetIndexPhrases(v bool) {
	if cp.IndexPhrases() != v {
		cp.IndexPhrasesValue = &v
	}
}
