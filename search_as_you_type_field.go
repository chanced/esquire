package picker

import "encoding/json"

type searchAsYouTypeField struct {
	Analyzer            string       `json:"analyzer,omitempty"`
	Index               interface{}  `json:"index,omitempty"`
	IndexOptions        IndexOptions `json:"index_options,omitempty"`
	Norms               interface{}  `json:"norms,omitempty"`
	Store               interface{}  `json:"store,omitempty"`
	SearchAnalyzer      string       `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer string       `json:"search_quote_analyzer,omitempty"`
	Similarity          Similarity   `json:"similarity,omitempty"`
	TermVector          TermVector   `json:"term_vector,omitempty"`
	Type                FieldType    `json:"type"`
	MaxShingleSize      int          `json:"max_shingle_size,omitempty"`
}
type SearchAsYouTypeFieldParams struct {
	// (Optional, integer) Largest shingle size to create. Valid values are 2
	// (inclusive) to 4 (inclusive). Defaults to 3.
	MaxShingleSize int `json:"max_shingle_size,omitempty"`

	// The analyzer which should be used for the text field, both at index-time
	// and at search-time (unless overridden by the search_analyzer). Defaults
	// to the default index analyzer, or the standard analyzer.
	Analyzer string `json:"analyzer,omitempty"`

	// Should the field be searchable? Accepts true (default) or false.
	Index interface{} `json:"index,omitempty"`

	// What information should be stored in the index, for search and
	// highlighting purposes. Defaults to positions
	IndexOptions IndexOptions `json:"index_options,omitempty"`

	// Whether field-length should be taken into account when scoring queries.
	// Accepts true (default) or false.
	Norms interface{} `json:"norms,omitempty"`

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

	// Whether term vectors should be stored for the field. Defaults to no. This
	// option configures the root field and shingle subfields, but not the
	// prefix subfield.
	TermVector TermVector `json:"term_vector,omitempty"`
}

func (SearchAsYouTypeFieldParams) Type() FieldType {
	return FieldTypeSearchAsYouType
}
func (p SearchAsYouTypeFieldParams) Field() (Field, error) {
	return p.SearchAsYouType()
}

func (p SearchAsYouTypeFieldParams) SearchAsYouType() (*SearchAsYouTypeField, error) {
	f := &SearchAsYouTypeField{}
	e := &MappingError{}
	f.SetAnalyzer(p.Analyzer)
	f.SetSearchAnalyzer(p.SearchAnalyzer)
	f.SetSearchQuoteAnalyzer(p.SearchQuoteAnalyzer)
	err := f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndexOptions(p.IndexOptions)
	if err != nil {
		e.Append(err)
	}
	err = f.SetMaxShingleSize(p.MaxShingleSize)
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

func NewSearchAsYouTypeField(params SearchAsYouTypeFieldParams) (*SearchAsYouTypeField, error) {
	return params.SearchAsYouType()
}

// SearchAsYouTypeField is a text-like field that is optimized to provide
// out-of-the-box support for queries that serve an as-you-type completion use
// case. It creates a series of subfields that are analyzed to index terms that
// can be efficiently matched by a query that partially matches the entire
// indexed text value. Both prefix completion (i.e matching terms starting at
// the beginning of the input) and infix completion (i.e. matching terms at any
// position within the input) are supported.
//
// The size of shingles in subfields can be configured with the max_shingle_size
// mapping parameter. The default is 3, and valid values for this parameter are
// integer values 2 - 4 inclusive. Shingle subfields will be created for each
// shingle size from 2 up to and including the max_shingle_size. The
// my_field._index_prefix subfield will always use the analyzer from the shingle
// subfield with the max_shingle_size when constructing its own analyzer.
//
// Increasing the max_shingle_size will improve matches for queries with more
// consecutive terms, at the cost of larger index size. The default
// max_shingle_size should usually be sufficient.
//
// The same input text is indexed into each of these fields automatically, with
// their differing analysis chains, when an indexed document has a value for the
// root field my_field.
//
//
// The most efficient way of querying to serve a search-as-you-type use case is
// usually a multi_match query of type bool_prefix that targets the root
// search_as_you_type field and its shingle subfields. This can match the query
// terms in any order, but will score documents higher if they contain the terms
// in order in a shingle subfield.
//
// To search for documents that strictly match the query terms in order, or to
// search using other properties of phrase queries, use a match_phrase_prefix
// query on the root field. A match_phrase query can also be used if the last
// term should be matched exactly, and not as a prefix. Using phrase queries may
// be less efficient than using the match_bool_prefix query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html
type SearchAsYouTypeField struct {
	maxShingleSizeParam
	analyzerParam
	searchAnalyzerParam
	searchQuoteAnalyzerParam
	indexParam
	indexOptionsParam
	normsParam
	storeParam
	similarityParam
	termVectorParam
}

func (s *SearchAsYouTypeField) Field() (Field, error) {
	return s, nil
}

func (SearchAsYouTypeField) Type() FieldType {
	return FieldTypeSearchAsYouType
}

func (s *SearchAsYouTypeField) UnmarshalBSON(data []byte) error {
	return s.UnmarshalJSON(data)
}

func (s *SearchAsYouTypeField) UnmarshalJSON(data []byte) error {
	var params SearchAsYouTypeFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.SearchAsYouType()
	*s = *v
	return err
}

func (s SearchAsYouTypeField) MarshalBSON() ([]byte, error) {
	return s.MarshalJSON()
}

func (s SearchAsYouTypeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(searchAsYouTypeField{
		Analyzer:            s.analyzer,
		Index:               s.index.Value(),
		IndexOptions:        s.indexOptions,
		Norms:               s.norms.Value(),
		Store:               s.store.Value(),
		SearchAnalyzer:      s.searchAnalyzer,
		SearchQuoteAnalyzer: s.searchQuoteAnalyzer,
		Similarity:          s.similarity,
		TermVector:          s.termVector,
		Type:                s.Type(),
	})
}
