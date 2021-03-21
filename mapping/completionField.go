package mapping

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
	BaseField                       `json:",inline" bson:",inline"`
	AnalyzerParam                   `json:",inline" bson:",inline"`
	PreserveSeperatorsParam         `json:",inline" bson:",inline"`
	PreservePositionIncrementsParam `json:",inline" bson:",inline"`
	MaxInputLengthParam             `json:",inline" bson:",inline"`
}

func NewCompletionField() *CompletionField {
	return &CompletionField{BaseField: BaseField{MappingType: TypeCompletion}}
}
