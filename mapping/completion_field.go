package mapping

type Completioner interface {
	Completion() (*CompletionField, error)
}

// CompletionFieldParams creates a completion_field. A completion_field is a
// completion suggester which provides provides auto-complete/search-as-you-type
// functionality. This is a navigational feature to guide users to relevant
// results as they are typing, improving search precision. It is not meant for
// spell correction or did-you-mean functionality like the term or phrase
// suggesters.
//
// Ideally, auto-complete functionality should be as fast as a user types to
// provide instant feedback relevant to what a user has already typed in. Hence,
// completion suggester is optimized for speed. The suggester uses data
// structures that enable fast lookups, but are costly to build and are stored
// in-memory.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type CompletionFieldParams struct {
	// Analyzer used to convert the text in the query value into tokens.
	// Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used. (Optional)
	Analyzer string `json:"analyzer,omitempty"`
	// SearchAnalyzer overrides Analyzer for search analysis. (Optional)
	SearchAnalyzer string `json:"search_analyzer,omitempty"`
	// SearchQuoteAnalyzer setting allows you to specify an analyzer for
	// phrases, this is particularly useful when dealing with disabling stop
	// words for phrase queries. (Optional)
	SearchQuoteAnalyzer string `json:"search_quote_analyzer,omitempty"`
}

// The CompletionField is a completion suggester which provides provides
// auto-complete/search-as-you-type functionality. This is a navigational
// feature to guide users to relevant results as they are typing, improving
// search precision. It is not meant for spell correction or did-you-mean
// functionality like the term or phrase suggesters.
//
// Ideally, auto-complete functionality should be as fast as a user types to
// provide instant feedback relevant to what a user has already typed in. Hence,
// completion suggester is optimized for speed. The suggester uses data
// structures that enable fast lookups, but are costly to build and are stored
// in-memory.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type CompletionField struct {
	analyzerParam
	searchAnalyzerParam
	searchQuoteAnalyzerParam
	PreserveSeperatorsParam         `json:",inline" bson:",inline"`
	PreservePositionIncrementsParam `json:",inline" bson:",inline"`
	MaxInputLengthParam             `json:",inline" bson:",inline"`
}

func (CompletionField) Type() FieldType {
	return FieldTypeCompletion
}

func NewCompletionField() *CompletionField {
	return &CompletionField{}
}
