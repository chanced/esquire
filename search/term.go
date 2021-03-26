package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Term struct {
	Value           string
	Boost           dynamic.Number
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
	if b, ok := t.Boost.Float(); ok {
		q.SetBoost(b)
	}
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
	TermValue string
	boostParam
	caseInsensitiveParam
}

func (tr TermRule) HasTermRule() bool {
	return tr.TermValue != ""
}

func (tr TermRule) MarshalJSON() ([]byte, error) {
	if !tr.HasTermRule() {
		return dynamic.Null, nil
	}
	m, err := marshalParams(&tr)
	if err != nil {
		return nil, err
	}
	m["value"] = tr.TermValue
	return json.Marshal(m)
}

func (tr *TermRule) UnmarshalJSON(data []byte) error {
	tr.TermValue = ""
	tr.boostParam = boostParam{}
	tr.caseInsensitiveParam = caseInsensitiveParam{}
	fields, err := unmarshalParams(data, tr)
	if err != nil {
		return err
	}

	if v, ok := fields["value"]; ok {
		tr.TermValue = v.UnquotedString()
	} else {
		tr.TermValue = ""
	}
	return nil
}
func (tr *TermRule) Type() Type {
	return TypeTerm
}

func (tr *TermRule) SetValue(v string) {
	tr.TermValue = v
}

func (tr TermRule) Value() string {
	return tr.TermValue
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
	TermField string
	TermRule
}

func (tq TermQuery) MarshalJSON() ([]byte, error) {
	if !tq.HasTermRule() || tq.TermField == "" {
		return dynamic.Null, nil
	}

	return json.Marshal(dynamic.Map{tq.TermField: tq.TermRule})
}

func (tq *TermQuery) UnmarshalJSON(data []byte) error {
	tq.TermField = ""
	tq.TermRule = TermRule{}

	m := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		tq.TermField = k
		err := json.Unmarshal(v, &tq.TermRule)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (tq *TermQuery) SetTerm(field string, term *Term) error {
	if term == nil {
		tq.RemoveTerm()
		return nil
	}
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeTerm)
	}
	r, err := term.Term()
	if err != nil {
		return err
	}
	tq.TermField = field
	tq.TermRule = *r
	return nil
}
func (tq *TermQuery) RemoveTerm() {
	tq.TermField = ""
	tq.TermRule = TermRule{}
}
