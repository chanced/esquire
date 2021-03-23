package search

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// Terms returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type Terms struct {
	Value           []string
	Boost           float32
	CaseInsensitive bool
}

// Termser can either be a pointer to a Terms or a pointer to a Lookup
//
// Example:
//  err := s.AddTerms("color", &Lookup{ID: "1", Index:"my-index-100", Path:"color"})
//  _ = err // handle err
//  err = s.AddTerms("user.id", &Terms{ Value: []string{}})
//  _ = err // handle err
type Termser interface {
	Terms() (*TermsRule, error)
}

func (t Terms) Rule() (Rule, error) {
	return t.Terms()
}
func (t Terms) Terms() (*TermsRule, error) {
	q := &TermsRule{}

	q.SetValue(t.Value)
	q.SetBoost(t.Boost)
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t Terms) Type() Type {
	return TypeTerm
}

type TermsRule struct {
	TermsValue           []string
	Field                string
	BoostParam           `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

func (t *TermsRule) Type() Type {
	return TypeTerms
}

func (t *TermsRule) SetValue(v []string) {
	t.TermsValue = v
}

func (t TermsRule) Value() []string {
	return t.TermsValue
}

func (t TermsRule) MarshalJSON() ([]byte, error) {
	if t.BoostParam.BoostValue == nil && t.CaseInsensitiveParam.CaseInsensitiveValue == nil {
		return json.Marshal(t.TermsValue)
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

func newTerms() TermsRule {
	return TermsRule{
		TermsValue: []string{},
	}
}

type TermsQuery struct {
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
