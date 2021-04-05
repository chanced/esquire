package picker

import "encoding/json"

type FlattenedFieldParams struct {
	// The maximum allowed depth of the flattened object field, in terms of
	// nested inner objects. If a flattened object field exceeds this limit,
	// then an error will be thrown. Defaults to 20. Note that depth_limit can
	// be updated dynamically through the update mapping API.
	DepthLimit interface{} `json:"depth_limit,omitempty"`
	// Should the field be stored on disk in a column-stride fashion, so that it
	// can later be used for sorting, aggregations, or scripting? Accepts true
	// (default) or false.
	DocValues interface{} `json:"doc_values,omitempty"`
	// Should global ordinals be loaded eagerly on refresh? Accepts true or
	// false (default). Enabling this is a good idea on fields that are
	// frequently used for terms aggregations.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eager-global-ordinals.html
	EagerGlobalOrdinals interface{} `json:"eager_global_ordinals,omitempty"`
	// Leaf values longer than this limit will not be indexed. By default, there
	// is no limit and all values will be indexed. Note that this limit applies
	// to the leaf values within the flattened object field, and not the length
	// of the entire field.
	IgnoreAbove interface{} `json:"ignore_above,omitempty"`
	// Index controls whether field values are indexed. It accepts true or false
	// and defaults to true. Fields that are not indexed are not queryable.
	// (Optional, bool or string that can be parsed as a bool)
	Index interface{} `bson:"index,omitempty" json:"index,omitempty"`
	// What information should be stored in the index for scoring purposes.
	// Defaults to docs but can also be set to freqs to take term frequency into
	// account when computing scores.
	IndexOptions IndexOptions `json:"index_options,omitempty"`
	// A string value which is substituted for any explicit null values within
	// the flattened object field. Defaults to null, which means null sields are
	// treated as if it were missing.
	NullValue interface{} `json:"null_value,omitempty"`
	// Which scoring algorithm or similarity should be used. Defaults to BM25.
	Similarity Similarity `json:"similarity,omitempty"`
	// Whether full text queries should split the input on whitespace when
	// building a query for this field. Accepts true or false (default).
	SplitQueriesOnWhitespace interface{} `json:"split_queries_on_whitespace,omitempty"`
}

func (FlattenedFieldParams) Type() FieldType {
	return FieldTypeFlattened
}

func (p FlattenedFieldParams) Field() (Field, error) {
	return p.Flattened()
}

func (p FlattenedFieldParams) Flattened() (*FlattenedField, error) {
	f := &FlattenedField{}
	e := &MappingError{}

	err := f.SetDepthLimit(p.DepthLimit)
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
	err = f.SetSimilarity(p.Similarity)
	if err != nil {
		e.Append(err)
	}
	err = f.SetSplitQueriesOnWhitespace(p.SplitQueriesOnWhitespace)
	if err != nil {
		e.Append(err)
	}
	f.SetNullValue(p.NullValue)
	return f, e.ErrorOrNil()
}

// FlattenedField maps an entire object as a single field.
//
// By default, each subfield in an object is mapped and indexed separately. If
// the names or types of the subfields are not known in advance, then they are
// mapped dynamically.
//
// The flattened type provides an alternative approach, where the entire object
// is mapped as a single field. Given an object, the flattened mapping will
// parse out its leaf values and index them into one field as keywords. The
// object’s contents can then be searched through simple queries and
// aggregations.
//
// This data type can be useful for indexing objects with a large or unknown
// number of unique keys. Only one field mapping is created for the whole JSON
// object, which can help prevent a mappings explosion from having too many
// distinct field mappings.
//
// On the other hand, flattened object fields present a trade-off in terms of search functionality. Only basic queries are allowed, with no support for numeric range queries or highlighting. Further information on the limitations can be found in the Supported operations section.
//
// ! X-Pack
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/flattened.html
type FlattenedField struct {
	depthLimitParam
	docValuesParam
	eagerGlobalOrdinalsParam
	ignoreAboveParam
	indexParam
	indexOptionsParam
	nullValueParam
	similarityParam
	splitQueriesOnWhitespaceParam
}

func (f *FlattenedField) Field() (Field, error) {
	return f, nil
}

func (FlattenedField) Type() FieldType {
	return FieldTypeFlattened
}
func (f FlattenedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(flattenedField{
		DepthLimit:               f.depthLimit.Value(),
		DocValues:                f.docValues.Value(),
		EagerGlobalOrdinals:      f.eagerGlobalOrdinals.Value(),
		IgnoreAbove:              f.ignoreAbove.Value(),
		Index:                    f.index.Value(),
		IndexOptions:             f.indexOptions,
		NullValue:                f.nullValue,
		Similarity:               f.similarity,
		SplitQueriesOnWhitespace: f.splitQueriesOnWhitespace.Value(),
		Type:                     f.Type(),
	})
}
func (f *FlattenedField) UnmarshalJSON(data []byte) error {
	var p FlattenedFieldParams

	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.Flattened()
	*f = *n
	return err

}
func NewFlattenedField(params FlattenedFieldParams) (*FlattenedField, error) {
	return params.Flattened()
}

//easyjson:json
type flattenedField struct {
	DepthLimit               interface{}  `json:"depth_limit,omitempty"`
	DocValues                interface{}  `json:"doc_values,omitempty"`
	EagerGlobalOrdinals      interface{}  `json:"eager_global_ordinals,omitempty"`
	IgnoreAbove              interface{}  `json:"ignore_above,omitempty"`
	Index                    interface{}  `bson:"index,omitempty" json:"index,omitempty"`
	IndexOptions             IndexOptions `json:"index_options,omitempty"`
	NullValue                interface{}  `json:"null_value,omitempty"`
	Similarity               Similarity   `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace interface{}  `json:"split_queries_on_whitespace,omitempty"`
	Type                     FieldType    `json:"type"`
}
