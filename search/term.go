package search

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// Term query returns documents that contain an exact term in a provided field.
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
type Term struct {
	// Term you wish to find in the provided <field>. To return a document, the
	// term must exactly match the field value, including whitespace and
	// capitalization. (Required)
	Value                string `json:"value" bson:"value"`
	BoostParam           `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

func (t Term) MarshalJSON() ([]byte, error) {
	if t.BoostParam.BoostValue == nil && t.CaseInsensitiveParam.CaseInsensitiveValue == nil {
		return json.Marshal(t.Value)
	}
	return json.Marshal(t)
}
func (t *Term) UnmarshalJSON(data []byte) error {
	g := gjson.ParseBytes(data)
	if g.Type == gjson.String {
		t.Value = g.String()
		t.BoostParam = BoostParam{}
		t.CaseInsensitiveParam = CaseInsensitiveParam{}
		return nil
	}
	return json.Unmarshal(data, t)
}

func NewTerm() Term {
	return Term{}
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
	Term map[string]Term `json:"term,omitempty" bson:"term,omitempty"`
}

func NewTermQuery() TermQuery {
	return TermQuery{
		Term: map[string]Term{},
	}
}
func (tq *TermQuery) AddTerm(field string, t Term) {
	if tq.Term == nil {
		tq.Term = map[string]Term{}
	}
	tq.Term[field] = t
}

func (tq *TermQuery) RemoveTerm(field string) {
	delete(tq.Term, field)
}
