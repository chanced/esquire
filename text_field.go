package picker

type textField struct{}

type TextFieldParams struct {
	// The analyzer which should be used for the text field, both at index-time
	// and at search-time (unless overridden by the search_analyzer). Defaults
	// to the default index analyzer, or the standard analyzer.
	Analyzer string `json:"analyzer,omitempty"`
	// Should global ordinals be loaded eagerly on refresh? Accepts true or false
	// (default). Enabling this is a good idea on fields that are frequently used
	// for (significant) terms aggregations.
	EagerGlobalOrdinals interface{} `eager_global_ordinals,omitempty`
	// Can the field use in-memory fielddata for sorting, aggregations, or
	// scripting? Accepts true or false (default).
	FieldData interface{} `json:"fielddata,omitempty"`
	// Expert settings which allow to decide which values to load in memory when
	// fielddata is enabled. By default all values are loaded.
	FieldDataFrequencyFilter *FieldDataFrequencyFilter `json:"fielddata_frequency_filter,omitempty"`
	// Whether field-length should be taken into account when scoring queries.
	// Accepts true (default) or false.

	Fields Fields
	Norms  interface{} `json:"norms,omitempty"`
	// Deprecated
	//
	// Mapping field-level query time boosting. Accepts a floating point number, defaults to 1.0.
	Boost interface{} `json`
}

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
	eagerGlobalOrdinalsParam
	fieldDataParam
	fieldDataFrequencyFilterParam
	fieldsParam
	indexParam
	indexOptionsParam
	IndexPrefixesParams
	IndexPhrasesParam
	NormsParam
	positionIncrementGapParam
	storeParam
	analyzerParam
	similarityParam
	TermVectorParam
	metaParam
	boostParam
}

func NewTextField() *TextField {
	return &TextField{BaseField: BaseField{MappingType: FieldTypeText}}
}
