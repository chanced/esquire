package search

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Term struct {
	Value           string
	Boost           float32
	CaseInsensitive bool
}

func (t Term) Query() (TermQueryValue, error) {
	q := TermQueryValue{}
	if t.Value == "" {
		return q, ErrValueRequired
	}
	q.SetValue(t.Value)
	q.SetBoost(t.Boost)
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t Term) Type() Type {
	return TypeTerm
}

// TermQueryValue query returns documents that contain an exact term in a provided field.
//
// You can use the term query to find documents based on a precise value such as
// a price, a product ID, or a username.
//
// Avoid using the term query for text fields.
//
// By default, Elasticsearch changes the values of text fields as part of
// analysis. This can make finding exact matches for text field values
// difficult.
//
// To search text field values, use the match query instead.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type TermQueryValue struct {
	ValueValue           string `json:"value" bson:"value"`
	BoostParam           `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

func (t *TermQueryValue) SetValue(v string) {
	t.ValueValue = v
}

func (t TermQueryValue) Value() string {
	return t.ValueValue
}

type term TermQueryValue

func (t TermQueryValue) MarshalJSON() ([]byte, error) {

	if t.BoostParam.BoostValue == nil && t.CaseInsensitiveParam.CaseInsensitiveValue == nil {
		return json.Marshal(t.ValueValue)
	}
	return json.Marshal(term(t))
}
func (t *TermQueryValue) UnmarshalJSON(data []byte) error {

	// TODO: bson codec

	g := gjson.ParseBytes(data)
	if g.Type == gjson.String {
		t.ValueValue = g.String()
		t.BoostParam = BoostParam{}
		t.CaseInsensitiveParam = CaseInsensitiveParam{}
		return nil
	}
	var tt term
	err := json.Unmarshal(data, &tt)
	if err != nil {
		return err
	}
	t.BoostParam = tt.BoostParam
	t.ValueValue = tt.ValueValue
	t.CaseInsensitiveParam = tt.CaseInsensitiveParam
	return nil
}

func NewTerm() TermQueryValue {
	return TermQueryValue{}
}

// TermQuery returns documents that contain an exact term in a provided field.
//
// You can use the term query to find documents based on a precise value such as
// a price, a product ID, or a username.
//
// Avoid using the term query for text fields.
//
// By default, Elasticsearch changes the values of text fields as part of
// analysis. This can make finding exact matches for text field values
// difficult.
//
// To search text field values, use the match query instead.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type TermQuery struct {
	TermValue map[string]TermQueryValue `json:"term,omitempty" bson:"term,omitempty"`
}

func NewTermQuery() TermQuery {
	return TermQuery{
		TermValue: map[string]TermQueryValue{},
	}
}
func (tq *TermQuery) AddTerm(field string, t TermQueryValue) {
	if tq.TermValue == nil {
		tq.TermValue = map[string]TermQueryValue{}
	}
	tq.TermValue[field] = t
}

func (tq *TermQuery) RemoveTerm(field string) {
	delete(tq.TermValue, field)
}
