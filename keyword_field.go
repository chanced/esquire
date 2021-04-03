package picker

import "encoding/json"

type keywordField struct {
	EagerGlobalOrdinals      interface{}  `eager_global_ordinals,omitempty`
	Fields                   Fields       `json:"fields,omitempty"`
	Index                    interface{}  `json:"index,omitempty"`
	IndexOptions             IndexOptions `json:"index_options,omitempty"`
	Norms                    interface{}  `json:"norms,omitempty"`
	IgnoreAbove              interface{}  `json:"ignore_above,omitempty"`
	NullValue                interface{}  `json:"null_value,omitempty"`
	Store                    interface{}  `json:"store,omitempty"`
	Similarity               Similarity   `json:"similarity,omitempty"`
	Meta                     Meta         `json:"meta,omitempty"`
	Normalizer               interface{}  `json:"normalizer,omitempty"`
	SplitQueriesOnWhitespace interface{}  `json:"split_queries_on_whitespace,omitempty"`
	Boost                    interface{}  `json:"boost,omitempty"`
	Type                     FieldType    `json:"type"`
}

type KeywordFieldParams struct {
	// Should the field be stored on disk in a column-stride fashion, so that it
	// can later be used for sorting, aggregations, or scripting? Accepts true
	// (default) or false.
	DocValues interface{} `json:"doc_values,omitempty"`
	// Should global ordinals be loaded eagerly on refresh? Accepts true or false
	// (default). Enabling this is a good idea on fields that are frequently used
	// for (significant) terms aggregations.
	EagerGlobalOrdinals interface{} `eager_global_ordinals,omitempty`
	// Multi-fields allow the same string value to be indexed in multiple ways
	// for different purposes, such as one field for search and a multi-field
	// for sorting and aggregations, or the same string value analyzed by
	// different analyzers.
	Fields FieldMap `json:"fields,omitempty"`
	// Should the field be searchable? Accepts true (default) or false.
	Index interface{} `json:"index,omitempty"`
	// What information should be stored in the index, for search and
	// highlighting purposes. Defaults to positions
	IndexOptions IndexOptions `json:"index_options,omitempty"`
	// Whether field-length should be taken into account when scoring queries.
	// Accepts true (default) or false.
	Norms interface{} `json:"norms,omitempty"`
	// Do not index any string longer than this value. Defaults to 2147483647 so
	// that all values would be accepted. Please however note that default
	// dynamic mapping rules create a sub keyword field that overrides this
	// default by setting ignore_above: 256.
	IgnoreAbove interface{} `json:"ignore_above,omitempty"`
	// Accepts a string value which is substituted for any explicit null values.
	// Defaults to null, which means the field is treated as missing.
	NullValue interface{} `json:"null_value,omitempty"`
	// Whether the field value should be stored and retrievable separately from
	// the _source field. Accepts true or false (default).
	Store interface{} `json:"store,omitempty"`
	// Which scoring algorithm or similarity should be used. Defaults to BM25.
	Similarity Similarity `json:"similarity,omitempty"`
	// Metadata about the field.
	Meta Meta `json:"meta,omitempty"`
	// How to pre-process the keyword prior to indexing. Defaults to null,
	// meaning the keyword is kept as-is.
	Normalizer string `json:"normalizer,omitempty"`
	// Whether full text queries should split the input on whitespace when
	// building a query for this field. Accepts true or false (default).
	SplitQueriesOnWhitespace interface{} `json:"split_queries_on_whitespace,omitempty"`

	// Deprecated
	//
	// Mapping field-level query time boosting. Accepts a floating point number,
	// defaults to 1.0.
	Boost interface{} `json:"boost,omitempty"`
}

func (KeywordFieldParams) Type() FieldType {
	return FieldTypeKeyword
}
func (p KeywordFieldParams) Field() (Field, error) {
	return p.Keyword()
}
func (p KeywordFieldParams) Keyword() (*KeywordField, error) {
	f := &KeywordField{}
	e := &MappingError{}
	err := f.SetBoost(p.Boost)
	if err != nil {
		e.Append(err)
	}
	err = f.SetDocValues(p.DocValues)
	if err != nil {
		e.Append(err)
	}
	err = f.SetEagerGlobalOrdinals(p.EagerGlobalOrdinals)
	if err != nil {
		e.Append(err)
	}
	err = f.SetFields(p.Fields)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIgnoreAbove(p.IgnoreAbove)
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
	err = f.SetMeta(p.Meta)
	if err != nil {
		e.Append(err)
	}
	err = f.SetNorms(p.Norms)
	if err != nil {
		e.Append(err)
	}
	err = f.SetSimilarity(p.Similarity)
	if err != nil {
		e.Append(err)
	}
	err = f.SetSplitQueriesOnWhitespace(p.SplitQueriesOnWhitespace)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	f.SetNormalizer(p.Normalizer)
	f.SetNullValue(p.NullValue)
	return f, e.ErrorOrNil()
}

// KeywordField keyword, which is used for structured content such as IDs, email
// addresses, hostnames, status codes, zip codes, or tags.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#keyword-field-type
type KeywordField struct {
	docValuesParam
	eagerGlobalOrdinalsParam
	fieldsParam
	ignoreAboveParam
	indexParam
	indexOptionsParam
	normsParam
	nullValueParam
	storeParam
	similarityParam
	normalizerParam
	splitQueriesOnWhitespaceParam
	metaParam
	boostParam
}

func (f *KeywordField) Field() (Field, error) {
	return f, nil
}
func (KeywordField) Type() FieldType {
	return FieldTypeKeyword
}

func (t *KeywordField) UnmarshalJSON(data []byte) error {
	var params KeywordFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Keyword()
	*t = *v
	return err
}

func (t KeywordField) MarshalJSON() ([]byte, error) {
	return json.Marshal(keywordField{

		Meta:                     t.meta,
		Fields:                   t.fields,
		IndexOptions:             t.indexOptions,
		Similarity:               t.similarity,
		NullValue:                t.nullValue,
		Normalizer:               t.normalizer,
		Index:                    t.index.Value(),
		Store:                    t.store.Value(),
		Boost:                    t.boost.Value(),
		Norms:                    t.norms.Value(),
		IgnoreAbove:              t.ignoreAbove.Value(),
		EagerGlobalOrdinals:      t.eagerGlobalOrdinals.Value(),
		SplitQueriesOnWhitespace: t.splitQueriesOnWhitespace.Value(),
		Type:                     t.Type(),
	})
}
func NewKeywordField(params KeywordFieldParams) (*KeywordField, error) {
	return params.Keyword()
}
