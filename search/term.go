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

func (t Term) Rule() (Rule, error) {
	return t.Term()
}
func (t Term) Term() (*TermRule, error) {
	q := &TermRule{}
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

// TermRule query returns documents that contain an exact term in a provided field.
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
type TermRule struct {
	TermValue            string `json:"value" bson:"value"`
	BoostParam           `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

func (t *TermRule) Type() Type {
	return TypeTerm
}

func (t *TermRule) SetValue(v string) {
	t.TermValue = v
}

func (t TermRule) Value() string {
	return t.TermValue
}

type term TermRule

func (t TermRule) MarshalJSON() ([]byte, error) {

	if t.BoostParam.BoostValue == nil && t.CaseInsensitiveParam.CaseInsensitiveValue == nil {
		return json.Marshal(t.TermValue)
	}
	return json.Marshal(term(t))
}
func (t *TermRule) UnmarshalJSON(data []byte) error {

	// TODO: bson codec

	g := gjson.ParseBytes(data)
	if g.Type == gjson.String {
		t.TermValue = g.String()
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
	t.TermValue = tt.TermValue
	t.CaseInsensitiveParam = tt.CaseInsensitiveParam
	return nil
}

func newTerm() TermRule {
	return TermRule{}
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
	TermValue map[string]*TermRule `json:"term,omitempty" bson:"term,omitempty"`
}

func NewTermQuery() TermQuery {
	return TermQuery{
		TermValue: map[string]*TermRule{},
	}
}

func (t *TermQuery) SetTerm(term map[string]Term) error {
	t.TermValue = map[string]*TermRule{}
	for k, v := range term {
		err := t.AssignTerm(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TermQuery) AddTerm(field string, term Term) error {
	if t.TermValue == nil {
		t.TermValue = map[string]*TermRule{}
	}
	_, exists := t.TermValue[field]
	if exists {
		return QueryError{
			Field: field,
			Err:   ErrFieldExists,
			Type:  TypeTerm,
		}
	}
	return nil
}

func (t *TermQuery) AssignTerm(field string, term Term) error {
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeTerm)
	}

	if t.TermValue == nil {
		t.TermValue = map[string]*TermRule{}
	}
	v, err := term.Term()
	if err != nil {
		return NewQueryError(err, TypeTerm, field)
	}

	t.TermValue[field] = v
	return nil

}

func (t *TermQuery) RemoveTerm(field string) {
	delete(t.TermValue, field)
}
