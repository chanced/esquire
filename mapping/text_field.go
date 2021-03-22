package mapping

// A TextField is a field to index full-text values, such as the body of an
// email or the description of a product. These fields are analyzed, that is
// they are passed through an analyzer to convert the string into a list of
// individual terms before being indexed. The analysis process allows
// Elasticsearch to search for individual words within each full text field.
// Text fields are not used for sorting and seldom used for aggregations
// (although the significant text aggregation is a notable exception).
//
// text fields are best suited for unstructured but human-readable content. If
// you need to index unstructured machine-generated content, see Mapping
// unstructured content.
//
// If you need to index structured content such as email addresses, hostnames,
// status codes, or tags, it is likely that you should rather use a keyword
// field.
//
// Use a field as both text and keyword
//
// Sometimes it is useful to have both a full text (text) and a keyword
// (keyword) version of the same field: one for full text search and the other
// for aggregations and sorting. This can be achieved with multi-fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html
type TextField struct {
	BaseField                     `json:",inline" bson:",inline"`
	EagerGlobalOrdinalsParam      `json:",inline" bson:",inline"`
	FieldDataParam                `json:",inline" bson:",inline"`
	FieldDataFrequencyFilterParam `json:",inline" bson:",inline"`
	FieldsParam                   `json:",inline" bson:",inline"`
	IndexParam                    `json:",inline" bson:",inline"`
	IndexOptionsParam             `json:",inline" bson:",inline"`
	IndexPrefixesParams           `json:",inline" bson:",inline"`
	IndexPhrasesParam             `json:",inline" bson:",inline"`
	NormsParam                    `json:",inline" bson:",inline"`
	PositionIncrementGapParam     `json:",inline" bson:",inline"`
	StoreParam                    `json:",inline" bson:",inline"`
	AnalyzerParam                 `json:",inline" bson:",inline"`
	SimilarityParam               `json:",inline" bson:",inline"`
	TermVectorParam               `json:",inline" bson:",inline"`
	MetaParam                     `json:",inline" bson:",inline"`
}

func (f TextField) Clone() Field {
	n := NewTextField()
	n.SetEagerGlobalOrdinals(f.EagerGlobalOrdinals())
	n.SetAnalyzer(f.Analyzer())
	n.SetSearchAnalyzer(f.SearchAnalyzer())
	n.SetSearchQuoteAnalyzer(f.SearchQuoteAnalyzer())
	n.SetFields(f.Fields().Clone())
	n.SetFieldData(f.FieldData())
	n.SetFieldDataFrequencyFilter(f.FieldDataFrequencyFilter().Clone())
	n.SetIndex(f.Index())
	n.SetIndexOptions(f.IndexOptions())
	n.SetIndexPhrases(f.IndexPhrases())
	n.SetIndexPrefixesMaxChars(f.IndexPrefixesMaxChars())
	n.SetIndexPrefixesMinChars(f.IndexPrefixesMinChars())
	n.SetMeta(f.Meta().Clone())
	n.SetNorms(f.Norms())
	n.SetPositionIncrementGap(f.PositionIncrementGap())
	n.SetSimilarity(f.Similarity())
	n.SetStore(f.Store())
	n.SetTermVector(f.TermVector())
	return n
}

func NewTextField() *TextField {
	return &TextField{BaseField: BaseField{MappingType: TypeText}}
}
