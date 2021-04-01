package mapping

// A TokenCountField is really an integer field which accepts string values,
// analyzes them, then indexes the number of tokens in the string.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/token-count.html
type TokenCountField struct {
	BaseField                     `json:",inline" bson:",inline"`
	analyzerParam                 `json:",inline" bson:",inline"`
	EnablePositionIncrementsParam `json:",inline" bson:",inline"`
	docValuesParam                `json:",inline" bson:",inline"`
	indexParam                    `json:",inline" bson:",inline"`
	nullValueParam                `json:",inline" bson:",inline"`
	storeParam                    `json:",inline" bson:",inline"`
}

func (f TokenCountField) Clone() Field {
	n := NewTokenCountField()
	n.SetEnablePositionIncrements(f.EnablePositionIncrements())
	n.SetDocValues(f.DocValues())
	n.SetNullValue(f.NullValue())
	n.SetAnalyzer(f.Analyzer())
	n.SetSearchQuoteAnalyzer(f.SearchQuoteAnalyzer())
	n.SetSearchAnalyzer(f.SearchAnalyzer())
	n.SetStore(f.Store())
	return n
}

func NewTokenCountField() *TokenCountField {
	return &TokenCountField{BaseField: BaseField{MappingType: FieldTypeTokenCount}}
}
