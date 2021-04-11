package picker

import "encoding/json"

type WildcardFieldParams struct {
	// NullValue parameter allows you to replace explicit null values with the
	// specified value so that it can be indexed and searched
	NullValue interface{} `json:"null_value,omitempty"`
	// IgnoreAbove signiall to not index any string longer than this value.
	// Defaults to 2147483647 so that all values would be accepted. Please
	// however note that default dynamic mapping rules create a sub keyword
	// field that overrides this default by setting ignore_above: 256.
	IgnoreAbove interface{} `json:"ignore_above,omitempty"`
}

func (WildcardFieldParams) Type() FieldType {
	return FieldTypeWildcardKeyword
}

func (p WildcardFieldParams) Field() (Field, error) {
	return p.Wildcard()
}

func (p WildcardFieldParams) Wildcard() (*WildcardField, error) {
	f := &WildcardField{}
	e := &MappingError{}

	f.SetNullValue(p.NullValue)
	err := f.SetIgnoreAbove(p.IgnoreAbove)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}

func NewWildcardField(params WildcardFieldParams) (*WildcardField, error) {
	return params.Wildcard()
}

// A WildcardField stores values optimised for wildcard grep-like queries. Wildcard queries are possible on other field types but suffer from constraints:
//
// - text fields limit matching of any wildcard expressions to individual tokens rather than the original whole value held in a field
//
// - keyword fields are untokenized but slow at performing wildcard queries (especially patterns with leading wildcards).
//
// Internally the wildcard field indexes the whole field value using ngrams and stores the full string. The index is used as a rough filter to cut down the number of values that are then checked by retrieving and checking the full values. This field is especially well suited to run grep-like queries on log lines. Storage costs are typically lower than those of keyword fields but search speeds for exact matches on full terms are slower.
//
// Limitations
//
// wildcard fields are untokenized like keyword fields, so do not support queries that rely on word positions such as phrase queries.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#wildcard-field-type
type WildcardField struct {
	nullValueParam   `bson:",inline" json:",inline"`
	ignoreAboveParam `bson:",inline" json:",inline"`
}

func (w *WildcardField) Field() (Field, error) {
	return w, nil
}

func (WildcardField) Type() FieldType {
	return FieldTypeWildcardKeyword
}

func (w WildcardField) MarshalJSON() ([]byte, error) {
	return wildcardField{
		NullValue:   w.nullValue,
		IgnoreAbove: w.ignoreAbove.Value(),
		Type:        string(w.Type()),
	}.MarshalJSON()
}
func (w *WildcardField) UnmarshalJSON(data []byte) error {
	*w = WildcardField{}
	var p WildcardFieldParams
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.Wildcard()
	*w = *n
	return err
}

//easyjson:json
type wildcardField struct {
	NullValue   interface{} `json:"null_value,omitempty"`
	IgnoreAbove interface{} `json:"ignore_above,omitempty"`
	Type        string      `json:"type"`
}
