package search

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Term struct {
	Value           string
	Boost           float64
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

func (t *TermRule) UnmarshalJSON(data []byte) error {
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
	TermRule  *TermRule
}

func (t TermQuery) MarshalJSON() ([]byte, error) {
	return sjson.SetBytes([]byte{}, t.TermField, t.TermRule)
}

func (t *TermQuery) UnmarshalJSON(data []byte) error {
	fmt.Println("inside term")
	t.TermField = ""
	t.TermRule = nil
	g := gjson.ParseBytes(data)
	var val gjson.Result
	g.ForEach(func(key, value gjson.Result) bool {
		t.TermField = key.Str
		val = value
		return false
	})
	switch val.Type {
	case gjson.JSON:
		err := unmarshalRule(val, t.TermRule, nil)
		return err
	case gjson.String:
		t.TermRule = &TermRule{
			TermValue: val.Str,
		}
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
	t.TermRule = r
	return nil
}
func (t *TermQuery) RemoveTerm() {
	t.TermField = ""
	t.TermRule = nil
}
