package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
	"github.com/chanced/picker/internal/jsonutil"
	"github.com/tidwall/gjson"
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

func (t TermRule) HasTermRule() bool {
	return t.TermValue != ""
}

func (t *TermRule) UnmarshalJSON(data []byte) error {
	t.TermValue = ""
	t.boostParam = boostParam{}
	t.caseInsensitiveParam = caseInsensitiveParam{}

	r := dynamic.RawJSON(data)
	if r.IsString() {
		t.TermValue = r.String()
		return nil
	}
	unmarshalRule(g, t, func(key, value gjson.Result) error {
		if key.Str == "value" {
			t.TermValue = value.Str
		}
		return nil
	})
	return nil
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

func (t TermQuery) MarshalJSON() ([]byte, error) {
	if !t.HasTermRule() {
		return jsonutil.Nil, nil
	}
	m := M{}
	m, err := marshalParams(m, &t)
	if err != nil {
		return nil, err
	}
	m["value"] = t.TermValue
	return json.Marshal(m)
}

func (t *TermQuery) UnmarshalJSON(data []byte) error {
	t.TermField = ""
	t.TermRule = TermRule{}

	m := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		t.TermField = k
		err := json.Unmarshal(v, &t.TermRule)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (t *TermQuery) SetTerm(field string, term *Term) error {
	if term == nil {
		t.RemoveTerm()
		return nil
	}
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeTerm)
	}
	r, err := term.Term()
	if err != nil {
		return err
	}
	t.TermField = field
	t.TermRule = *r
	return nil
}
func (t *TermQuery) RemoveTerm() {
	t.TermField = ""
	t.TermRule = TermRule{}
}
