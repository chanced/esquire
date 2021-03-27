package search

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

type Term struct {
	// The field which is being queried against.
	//
	// This does not need to be set if you are explicitly setting the TermQuery.
	// It is only needed when adding to a set of Clauses:
	//  s := search.NewSearch()
	//  s.SetTerm("my-field", Term{ Value: "term-value" }) // good
	//  s.SetShould(Term{ FieldName: "my-field", Value: "term-value" })
	Field           string
	Value           string
	Boost           dynamic.Number
	CaseInsensitive bool
	QueryName       string
}

func (t Term) FieldName() string {
	return t.Field
}

func (t Term) Name() string {
	return t.QueryName
}

func (t Term) Clause() (Clause, error) {
	return t.Term()
}
func (t Term) Term() (*termClause, error) {
	q := &termClause{}
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

type termClause struct {
	TermValue string
	boostParam
	caseInsensitiveParam
	nameParam
}

func (tr termClause) HasTermRule() bool {
	return tr.TermValue != ""
}

func (tr termClause) MarshalJSON() ([]byte, error) {
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

func (tr *termClause) UnmarshalJSON(data []byte) error {
	tr.TermValue = ""
	tr.boostParam = boostParam{}
	tr.caseInsensitiveParam = caseInsensitiveParam{}

	d := dynamic.RawJSON(data)
	if d.IsString() {
		tr.TermValue = d.UnquotedString()
		return nil
	}
	fields, err := unmarshalParams(data, tr)
	if err != nil {
		return err
	}

	if v, ok := fields["value"]; ok {
		tr.TermValue = v.UnquotedString()
	}
	return nil
}
func (tr termClause) Type() Type {
	return TypeTerm
}

func (tr *termClause) SetValue(v string) {
	tr.TermValue = v
}

func (tr termClause) Value() string {
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
	termClause
}

func (t TermQuery) FieldName() string {
	return t.TermField
}

func (t TermQuery) MarshalJSON() ([]byte, error) {
	if !t.HasTermRule() || t.TermField == "" {
		return dynamic.Null, nil
	}
	return json.Marshal(dynamic.Map{t.TermField: t.termClause})
}

func (t *TermQuery) UnmarshalJSON(data []byte) error {
	t.TermField = ""
	t.termClause = termClause{}
	rd := dynamic.RawJSON(data)
	fmt.Println(rd.String())
	m := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		t.TermField = k
		err := json.Unmarshal(v, &t.termClause)
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
	term.Field = field
	if term.Field == "" {
		term.Field = t.TermField
	}
	if term.Field == "" {
		return NewQueryError(ErrFieldRequired, TypeTerm)
	}
	r, err := term.Term()
	if err != nil {
		return err
	}
	t.termClause = *r
	return nil
}
func (t *TermQuery) RemoveTerm() {
	t.TermField = ""
	t.termClause = termClause{}
}
