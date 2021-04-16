package picker

import "encoding/json"

type tokenCountField struct {
	Analyzer                 string      `json:"analyzer,omitempty"`
	EnablePositionIncrements interface{} `json:"enable_position_increments,omitempty"`
	DocValues                interface{} `json:"doc_values,omitempty"`
	Index                    interface{} `json:"index,omitempty"`
	NullValue                interface{} `json:"null_value,omitempty"`
	Boost                    interface{} `json:"boost,omitempty"`
	Store                    interface{} `json:"store,omitempty"`
	Type                     FieldType   `json:"type"`
}

type TokenCountFieldParams struct {
	// The analyzer which should be used to analyze the string value. Required.
	// For best performance, use an analyzer without token filters.
	Analyzer string `json:"analyzer,omitempty"`
	// Indicates if position increments should be counted. Set to false if you
	// don’t want to count tokens removed by analyzer filters (like stop).
	// Defaults to true.
	EnablePositionIncrements interface{} `json:"enable_position_increments,omitempty"`
	// 	Should the field be stored on disk in a column-stride fashion, so that
	// 	it can later be used for sorting, aggregations, or scripting? Accepts
	// 	true (default) or false.
	DocValues interface{} `json:"doc_values,omitempty"`
	// Should the field be searchable? Accepts true (default) and false.
	Index interface{} `json:"index,omitempty"`
	// Accepts a numeric value of the same type as the field which is
	// substituted for any explicit null values. Defaults to null, which means
	// the field is treated as missing.
	NullValue interface{} `json:"null_value,omitempty"`
	// Deprecated
	Boost interface{} `json:"boost,omitempty"`
	// Whether the field value should be stored and retrievable separately from
	// the _source field. Accepts true or false (default).
	Store interface{} `json:"store,omitempty"`
}

func (TokenCountFieldParams) Type() FieldType {
	return FieldTypeTokenCount
}

func (p TokenCountFieldParams) Field() (Field, error) {
	return p.TokenCount()
}
func (p TokenCountFieldParams) TokenCount() (*TokenCountField, error) {
	f := &TokenCountField{}
	e := &MappingError{}

	f.SetAnalyzer(p.Analyzer)
	f.SetNullValue(p.NullValue)
	err := f.SetBoost(p.Boost)
	if err != nil {

		e.Append(err)
	}
	err = f.SetDocValues(p.DocValues)
	if err != nil {
		e.Append(err)
	}
	err = f.SetEnablePositionIncrements(p.EnablePositionIncrements)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIndex(p.Index)
	if err != nil {
		e.Append(err)
	}
	err = f.SetStore(p.Store)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewTokenCountField(params TokenCountFieldParams) (*TokenCountField, error) {
	return params.TokenCount()
}

// A TokenCountField is really an integer field which accepts string values,
// analyzes them, then indexes the number of tokens in the string.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/token-count.html
type TokenCountField struct {
	analyzerParam
	boostParam
	enablePositionIncrementsParam
	docValuesParam
	indexParam
	nullValueParam
	storeParam
}

func (TokenCountField) Type() FieldType {
	return FieldTypeTokenCount
}
func (t *TokenCountField) Field() (Field, error) {
	return t, nil
}
func (t *TokenCountField) UnmarshalBSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

func (t *TokenCountField) UnmarshalJSON(data []byte) error {

	var params TokenCountFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.TokenCount()
	*t = *v
	return err
}

func (t TokenCountField) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}

func (t TokenCountField) MarshalJSON() ([]byte, error) {
	return json.Marshal(tokenCountField{
		Analyzer:                 t.analyzer,
		NullValue:                t.nullValue,
		Index:                    t.index.Value(),
		Store:                    t.store.Value(),
		Boost:                    t.boost.Value(),
		DocValues:                t.docValues.Value(),
		EnablePositionIncrements: t.enablePositionIncrements.Value(),
		Type:                     t.Type(),
	})
}
