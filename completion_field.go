package picker

import "encoding/json"

type Completioner interface {
	Completion() (*CompletionField, error)
}

// CompletionFieldParams creates a completion_field. A completion_field is a
// completion suggester which provides provides auto-complete/search-as-you-type
// functionality. This is a navigational feature to guide users to relevant
// results as they are typing, improving search precision. It is not meant for
// spell correction or did-you-mean functionality like the term or phrase
// suggesters.
//
// Ideally, auto-complete functionality should be as fast as a user types to
// provide instant feedback relevant to what a user has already typed in. Hence,
// completion suggester is optimized for speed. The suggester uses data
// structures that enable fast lookups, but are costly to build and are stored
// in-memory.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
type CompletionFieldParams struct {
	// Analyzer used to convert the text in the query value into tokens.
	// Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used. (Optional)
	Analyzer string `json:"analyzer,omitempty"`
	// SearchAnalyzer overrides Analyzer for search analysis. (Optional)
	SearchAnalyzer string `json:"search_analyzer,omitempty"`
	// SearchQuoteAnalyzer setting allows you to specify an analyzer for
	// phrases, this is particularly useful when dealing with disabling stop
	// words for phrase queries. (Optional)
	SearchQuoteAnalyzer string `json:"search_quote_analyzer,omitempty"`
	// PreserveSeperators preserves the separators, defaults to true. If
	// disabled, you could find a field starting with Foo Fighters, if you
	// suggest for foof.
	PreserveSeperators interface{} `json:"preserve_separators,omitempty"`

	// Enables position increments, defaults to true. If disabled and using
	// stopwords analyzer, you could get a field starting with The Beatles, if you
	// suggest for b. Note: You could also achieve this by indexing two inputs,
	// Beatles and The Beatles, no need to change a simple analyzer, if you are able
	// to enrich your data.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
	PreservePositionIncrements interface{} `json:"preserve_position_increments,omitempty"`
	// MaxInputLength limits the length of a single input, defaults to 50 UTF-16
	// code points. This limit is only used at index time to reduce the total
	// number of characters per input string in order to prevent massive inputs
	// from bloating the underlying datastructure. Most use cases won’t be
	// influenced by the default value since prefix completions seldom grow
	// beyond prefixes longer than a handful of characters.
	MaxInputLength interface{} `json:"max_input_length,omitempty"`
}

func (CompletionFieldParams) Type() FieldType {
	return FieldTypeCompletion
}

func (p CompletionFieldParams) Field() (Field, error) {
	return p.Completion()
}

func (p CompletionFieldParams) Completion() (*CompletionField, error) {
	f := &CompletionField{}
	e := &MappingError{}
	err := f.SetMaxInputLength(p.MaxInputLength)
	if err != nil {
		e.Append(err)
	}
	err = f.SetPreservePositionIncrements(p.PreservePositionIncrements)
	if err != nil {
		e.Append(err)
	}
	err = f.SetPreserveSeperators(p.PreserveSeperators)
	if err != nil {
		e.Append(err)
	}
	f.SetAnalyzer(p.Analyzer)
	f.SetSearchAnalyzer(p.SearchAnalyzer)
	f.SetSearchQuoteAnalyzer(p.SearchQuoteAnalyzer)
	return f, e.ErrorOrNil()
}

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
	analyzerParam
	searchAnalyzerParam
	searchQuoteAnalyzerParam
	preserveSeperatorsParam
	preservePositionIncrementsParam
	maxInputLengthParam
}

func (c *CompletionField) Field() (Field, error) {
	return c, nil
}
func (CompletionField) Type() FieldType {
	return FieldTypeCompletion
}

func NewCompletionField(params CompletionFieldParams) (*CompletionField, error) {
	return params.Completion()
}

func (c CompletionField) MarshalBSON() ([]byte, error) {
	return c.MarshalJSON()
}

func (c CompletionField) MarshalJSON() ([]byte, error) {
	return json.Marshal(completionField{
		Analyzer:                   c.analyzer,
		SearchAnalyzer:             c.searchAnalyzer,
		SearchQuoteAnalyzer:        c.searchQuoteAnalyzer,
		PreserveSeperators:         c.preserveSeperators.Value(),
		PreservePositionIncrements: c.preservePositionIncrements.Value(),
		MaxInputLength:             c.maxInputLength.Value(),
		Type:                       c.Type(),
	})
}

func (c *CompletionField) UnmarshalBSON(data []byte) error {
	return c.UnmarshalJSON(data)
}

func (c *CompletionField) UnmarshalJSON(data []byte) error {
	var p CompletionFieldParams
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.Completion()
	*c = *n
	return err
}

//easyjson:json
type completionField struct {
	Analyzer                   string      `json:"analyzer,omitempty"`
	SearchAnalyzer             string      `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer        string      `json:"search_quote_analyzer,omitempty"`
	PreserveSeperators         interface{} `json:"preserve_separators,omitempty"`
	PreservePositionIncrements interface{} `json:"preserve_position_increments,omitempty"`
	MaxInputLength             interface{} `json:"max_input_length,omitempty"`
	Type                       FieldType   `json:"type"`
}
