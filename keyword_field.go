package picker

// KeywordField keyword, which is used for structured content such as IDs, email
// addresses, hostnames, status codes, zip codes, or tags.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#keyword-field-type
type KeywordField struct {
	docValuesParam                `bson:",inline" json:",inline"`
	eagerGlobalOrdinalsParam      `bson:",inline" json:",inline"`
	FieldsParam                   `bson:",inline" json:",inline"`
	ignoreAboveParam              `bson:",inline" json:",inline"`
	indexParam                    `bson:",inline" json:",inline"`
	IndexOptionsParam             `bson:",inline" json:",inline"`
	NormsParam                    `bson:",inline" json:",inline"`
	nullValueParam                `bson:",inline" json:",inline"`
	storeParam                    `bson:",inline" json:",inline"`
	SimilarityParam               `bson:",inline" json:",inline"`
	NormalizerParam               `bson:",inline" json:",inline"`
	SplitQueriesOnWhitespaceParam `bson:",inline" json:",inline"`
	metaParam                     `bson:",inline" json:",inline"`
}

func NewKeywordField(params KeywordFieldParams) (*KeywordField, error) {
	return &KeywordField{BaseField: BaseField{MappingType: FieldTypeKeyword}}
}
