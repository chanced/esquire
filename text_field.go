package picker

import "encoding/json"

type textField struct {
	Analyzer                 string                    `json:"analyzer,omitempty"`
	EagerGlobalOrdinals      interface{}               `eager_global_ordinals,omitempty`
	FieldData                interface{}               `json:"fielddata,omitempty"`
	FieldDataFrequencyFilter *FieldDataFrequencyFilter `json:"fielddata_frequency_filter,omitempty"`
	Fields                   Fields                    `json:"fields,omitempty"`
	Index                    interface{}               `json:"index,omitempty"`
	IndexOptions             IndexOptions              `json:"index_options,omitempty"`
	IndexPrefixes            *IndexPrefixes            `json:"index_prefixes,omitempty"`
	IndexPhrases             interface{}               `json:"index_phrases,omitempty"`
	Norms                    interface{}               `json:"norms,omitempty"`
	PositionIncrementGap     interface{}               `json:"position_increment_gap,omitempty"`
	Store                    interface{}               `json:"store,omitempty"`
	SearchAnalyzer           string                    `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer      string                    `json:"search_quote_analyzer,omitempty"`
	Similarity               Similarity                `json:"similarity,omitempty"`
	TermVector               TermVector                `json:"term_vector,omitempty"`
	Meta                     Meta                      `json:"meta,omitempty"`
	Boost                    interface{}               `json:"boost,omitempty"`
	Type                     FieldType                 `json:"type"`
}

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
	// Multi-fields allow the same string value to be indexed in multiple ways for different purposes, such as one field for search and a multi-field for sorting and aggregations, or the same string value analyzed by different analyzers.
	Fields FieldMap `json:"fields,omitempty"`
	// Should the field be searchable? Accepts true (default) or false.
	Index interface{} `json:"index,omitempty"`
	// What information should be stored in the index, for search and
	// highlighting purposes. Defaults to positions
	IndexOptions IndexOptions `json:"index_options,omitempty"`
	// If enabled, term prefixes of between 2 and 5 characters are indexed into
	// a separate field. This allows prefix searches to run more efficiently, at
	// the expense of a larger index.
	IndexPrefixes *IndexPrefixes `json:"index_prefixes,omitempty"`

	// If enabled, two-term word combinations (shingles) are indexed into a separate field. This allows exact phrase queries (no slop) to run more efficiently, at the expense of a larger index. Note that this works best when stopwords are not removed, as phrases containing stopwords will not use the subsidiary field and will fall back to a standard phrase query. Accepts true or false (default).
	IndexPhrases interface{} `json:"index_phrases,omitempty"`

	// Whether field-length should be taken into account when scoring queries.
	// Accepts true (default) or false.
	Norms interface{} `json:"norms,omitempty"`

	// The number of fake term position which should be inserted between each element of an array of strings. Defaults to the position_increment_gap configured on the analyzer which defaults to 100. 100 was chosen because it prevents phrase queries with reasonably large slops (less than 100) from matching terms across field values.
	PositionIncrementGap interface{} `json:"position_increment_gap,omitempty"`

	// Whether the field value should be stored and retrievable separately from
	// the _source field. Accepts true or false (default).
	Store interface{} `json:"store,omitempty"`
	// The analyzer that should be used at search time on the text field. Defaults to the analyzer setting.
	SearchAnalyzer string `json:"search_analyzer,omitempty"`
	// The analyzer that should be used at search time when a phrase is
	// encountered. Defaults to the search_analyzer setting.
	SearchQuoteAnalyzer string `json:"search_quote_analyzer,omitempty"`

	// Which scoring algorithm or similarity should be used. Defaults to BM25.
	Similarity Similarity `json:"similarity,omitempty"`
	// Whether term vectors should be stored for the field. Defaults to no.
	TermVector TermVector `json:"term_vector,omitempty"`
	// Metadata about the field.
	Meta Meta `json:"meta,omitempty"`
	// Deprecated
	//
	// Mapping field-level query time boosting. Accepts a floating point number, defaults to 1.0.
	Boost interface{} `json:"boost,omitempty"`
}

func (TextFieldParams) Type() FieldType {
	return FieldTypeText
}

func (p TextFieldParams) Field() (Field, error) {
	return p.Text()
}

func (p TextFieldParams) Text() (*TextField, error) {
	f := &TextField{}
	e := &MappingError{}
	f.SetAnalyzer(p.Analyzer)
	f.SetSearchAnalyzer(p.SearchAnalyzer)
	f.SetSearchQuoteAnalyzer(p.SearchQuoteAnalyzer)
	err := f.SetBoost(p.Boost)
	if err != nil {
		e.Append(err)
	}
	err = f.SetEagerGlobalOrdinals(p.EagerGlobalOrdinals)
	if err != nil {
		e.Append(err)
	}
	err = f.SetFieldData(p.FieldData)
	if err != nil {
		e.Append(err)
	}
	err = f.SetFields(p.Fields)
	if err != nil {
		e.Append(err)
	}
	err = f.SetFieldDataFrequencyFilter(p.FieldDataFrequencyFilter)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndexOptions(p.IndexOptions)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndexPhrases(p.IndexPhrases)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndexPrefixes(p.IndexPrefixes)
	if err != nil {
		e.Append(err)
	}
	err = f.SetMeta(p.Meta)
	if err != nil {
		e.Append(err)
	}
	err = f.SetNorms(p.Norms)
	if err != nil {
		e.Append(err)
	}
	err = f.SetPositionIncrementGap(p.PositionIncrementGap)
	if err != nil {
		e.Append(err)
	}
	err = f.SetSimilarity(p.Similarity)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	err = f.SetTermVector(p.TermVector)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
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
	indexPrefixesParams
	indexPhrasesParam
	normsParam
	positionIncrementGapParam
	storeParam
	analyzerParam
	searchAnalyzerParam
	searchQuoteAnalyzerParam
	similarityParam
	termVectorParam
	metaParam
	boostParam
}

func (f *TextField) Field() (Field, error) {
	return f, nil
}
func (TextField) Type() FieldType {
	return FieldTypeText
}

func (t *TextField) UnmarshalJSON(data []byte) error {
	var params TextFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Text()
	*t = *v
	return err
}

func (t TextField) MarshalJSON() ([]byte, error) {
	return json.Marshal(textField{
		Analyzer:                 t.analyzer,
		EagerGlobalOrdinals:      t.eagerGlobalOrdinals.Value(),
		FieldData:                t.fieldData.Value(),
		FieldDataFrequencyFilter: t.fieldDataFrequencyFilter,
		Index:                    t.index.Value(),
		Store:                    t.store.Value(),
		Meta:                     t.meta,
		Boost:                    t.boost.Value(),
		Fields:                   t.fields,
		IndexOptions:             t.indexOptions,
		IndexPrefixes:            t.indexPrefixes,
		IndexPhrases:             t.indexPhrases.Value(),
		Norms:                    t.norms.Value(),
		PositionIncrementGap:     t.positionIncrementGap.Value(),
		SearchAnalyzer:           t.searchAnalyzer,
		SearchQuoteAnalyzer:      t.SearchQuoteAnalyzerValue,
		Similarity:               t.similarity,
		TermVector:               t.termVector,
		Type:                     t.Type(),
	})
}
func NewTextField(params TextFieldParams) (*TextField, error) {
	return params.Text()
}
