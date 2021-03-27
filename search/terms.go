package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Termser can be:
//  *search.Terms
//  *search.Lookup
//  search.String
//  search.Strings
//
// Example:
//  impor
//  s := search.NewSearch()
//  err := s.AddTerms(&Lookup{ID: "1", Index:"my-index-100", Path:"color"})
//  _ = err // handle err
//  err = s.AddTerms(&Terms{ Field:"", Value: []string{"kimchy", "elkbee"}})
//  _ = err // handle err
type Termser interface {
	Terms() (*termsClause, error)
}

// Terms returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type Terms struct {
	Field           string
	Values          []string
	Boost           float64
	CaseInsensitive bool
}

func (t Terms) Rule() (Clause, error) {
	return t.Terms()
}
func (t Terms) Terms() (*termsClause, error) {
	q := &termsClause{
		TermsValue: t.Values,
		TermsField: t.Field,
	}

	q.SetBoost(t.Boost)
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t Terms) Type() Type {
	return TypeTerms
}

type TermsLookup struct {
	TermsID      string `json:"id,omitempty" bson:"id,omitempty"`
	TermsIndex   string `json:"index,omitempty" bson:"index,omitempty"`
	TermsPath    string `json:"path,omitempty" bson:"path,omitempty"`
	TermsRouting string `json:"routing,omitempty" bson:"routing,omitempty"`
}

func (t TermsLookup) lookupIsEmpty() bool {
	return len(t.TermsID) == 0 && len(t.TermsIndex) == 0 && len(t.TermsPath) == 0 && len(t.TermsRouting) == 0
}

type termsClause struct {
	TermsLookup
	TermsValue []string
	TermsField string
	boostParam
	caseInsensitiveParam
}

func (t termsClause) Field() string {
	return t.TermsField
}

func (t termsClause) Type() Type {
	return TypeTerms
}

func (t *termsClause) SetValue(value []string) {
	t.TermsLookup = TermsLookup{}
	if value == nil {
		value = []string{}
	}
	t.TermsValue = value
}
func (t *termsClause) SetField(field string) {
	t.TermsField = field
}
func (t *termsClause) SetLookup(lookup *TermsLookup) {
	t.SetValue([]string{})
	if lookup == nil {
		lookup = &TermsLookup{}
	}
	t.TermsLookup = *lookup

}
func (t *termsClause) set(v Termser) error {
	tv, err := v.Terms()
	if err != nil {
		return err
	}
	t.SetBoost(tv.Boost())
	t.SetCaseInsensitive(tv.CaseInsensitive())
	t.TermsLookup = tv.TermsLookup
	t.TermsValue = tv.TermsValue
	return nil
}

func (t termsClause) Value() []string {
	return t.TermsValue
}

func (t termsClause) selfIdentifying() {}

func (t termsClause) MarshalJSON() ([]byte, error) {
	var v map[string]interface{}
	v, err := marshalClauseParams(&t)
	if err != nil {
		return nil, err
	}
	if t.TermsField == "" {
		return dynamic.Null, nil
	}
	var q interface{}
	if !t.TermsLookup.lookupIsEmpty() {
		q = t.TermsLookup
	} else {
		q = t.TermsValue
	}
	v[t.TermsField] = q
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

func (t *termsClause) UnmarshalJSON(data []byte) error {
	g := dynamic.RawJSON(data)
	if g.IsNull() {
		return nil
	}
	t.TermsValue = []string{}
	t.TermsLookup = TermsLookup{}
	fields, err := unmarshalParams(data, t)
	if err != nil {
		return err
	}
	for fld, val := range fields {
		t.TermsField = fld
		if val.IsArray() {
			var sl []string
			err := json.Unmarshal(val, &sl)
			if err != nil {
				return err
			}
			t.TermsValue = sl
			return nil
		}
		var tl TermsLookup
		err := json.Unmarshal(val, &tl)
		if err != nil {
			return err
		}
		t.TermsLookup = tl
		return nil
	}

	return err
}
func (t *termsClause) UnmarshalBSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

func (t termsClause) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}

type TermsQuery struct {
	termsClause `json:",inline" bson:",inline"`
}

func (t TermsQuery) SetTerms(field string, value Termser) error {
	return t.set(value)
}
