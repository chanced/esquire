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

// SetAnalyzer sets Analyzer to v
func (c *CompletionField) SetAnalyzer(v string) *CompletionField {
	c.AnalyzerValue = v
	return c
}

// SetSearchAnalyzer sets SearchAnalyzer to v
func (c *CompletionField) SetSearchAnalyzer(v string) *CompletionField {
	c.AnalyzerParam.SetSearchAnalyzer(v)
	return c
}

// SetPreserveSeperators sets the PreserveSeperatorParam value to v
func (c *CompletionField) SetPreserveSeperators(v bool) *CompletionField {
	c.PreserveSeperatorsParam.SetPreserveSeperators(v)
	return c
}

func (c *CompletionField) SetPreservePositionIncrements(v bool) *CompletionField {
	c.PreservePositionIncrementsParam.SetPreservePositionIncrements(v)
	return c
}

func (c *CompletionField) SetMaxInputLength(v int) *CompletionField {
	c.MaxInputLengthParam.SetMaxInputLength(v)
	return c
}
