package picker

import (
	"github.com/chanced/dynamic"
)

type TermsSetter interface {
	TermsSet() (*TermsSetQuery, error)
}

type TermsSetQueryParams struct {
	Name string
	// (Required) Field you wish to search.
	Field string
	// (Required, array of strings) Array of terms you wish to find in the
	// provided <field>. To return a document, a required number of terms must
	// exactly match the field values, including whitespace and capitalization.
	//
	// The required number of matching terms is defined in the
	// minimum_should_match_field or minimum_should_match_script parameter.
	Terms []string
	// (Optional, string) Numeric field containing the number of matching terms
	// required to return a document.
	MinimumShouldMatchField string
	//  Custom script containing the number of matching terms required to return
	//  a document.
	MinimumShouldMatchScript *Script
	Boost                    interface{}
	completeClause
}

func (TermsSetQueryParams) Kind() QueryKind {
	return QueryKindTermsSet
}

func (p TermsSetQueryParams) Clause() (QueryClause, error) {
	return p.TermsSet()
}
func (p TermsSetQueryParams) TermsSet() (*TermsSetQuery, error) {
	q := &TermsSetQuery{}
	var err error
	err = q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindTermsSet)
	}
	err = q.SetTerms(p.Terms)
	if err != nil {
		return q, newQueryError(err, QueryKindTermsSet, p.Field)
	}
	q.SetMinimumShouldMatchField(p.MinimumShouldMatchField)
	q.SetMinimumShouldMatchScript(p.MinimumShouldMatchScript)
	q.SetName(p.Name)
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindTermsSet, p.Field)
	}
	return q, nil
}

type TermsSetQuery struct {
	nameParam
	boostParam
	fieldParam
	terms                    []string
	minimumShouldMatchField  string
	minimumShouldMatchScript *Script
	completeClause
}

func (TermsSetQuery) Kind() QueryKind {
	return QueryKindTermsSet
}
func (q *TermsSetQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *TermsSetQuery) TermsSet() (*TermsSetQuery, error) {
	return q, nil
}
func (q *TermsSetQuery) Clear() {
	if q == nil {
		return
	}
	*q = TermsSetQuery{}
}

func (q TermsSetQuery) Terms() []string {
	return q.terms
}
func (q *TermsSetQuery) SetTerms(terms []string) error {
	if len(terms) == 0 {
		return ErrTermsRequired
	}
	q.terms = terms
	return nil
}

func (q TermsSetQuery) MinimumShouldMatchScript() *Script {
	return q.minimumShouldMatchScript
}

func (q *TermsSetQuery) SetMinimumShouldMatchScript(script *Script) {
	q.minimumShouldMatchScript = script
}
func (q TermsSetQuery) MinimumShouldMatchField() string {
	return q.minimumShouldMatchField
}

func (q *TermsSetQuery) SetMinimumShouldMatchField(v string) {
	q.minimumShouldMatchField = v
}
func (q *TermsSetQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *TermsSetQuery) UnmarshalJSON(data []byte) error {
	*q = TermsSetQuery{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for field, d := range obj {
		q.field = field
		p := termsSetQuery{}
		err := p.UnmarshalJSON(d)
		if err != nil {
			return err
		}
		q.terms = p.Terms
		q.SetBoost(p.Boost)
		q.SetName(p.Name)
		q.SetMinimumShouldMatchField(p.MinimumShouldMatchField)
		q.minimumShouldMatchScript = p.MinimumShouldMatchScript
		return nil
	}
	return nil
}
func (q TermsSetQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q TermsSetQuery) MarshalJSON() ([]byte, error) {
	p := termsSetQuery{
		Name:                    q.name,
		Terms:                   q.terms,
		Boost:                   q.boost.Value(),
		MinimumShouldMatchField: q.minimumShouldMatchField,
	}
	if !q.minimumShouldMatchScript.IsEmpty() {
		p.MinimumShouldMatchScript = q.minimumShouldMatchScript
	}
	pd, err := p.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return dynamic.JSONObject{q.field: pd}.MarshalJSON()
}
func (q *TermsSetQuery) IsEmpty() bool {
	return q == nil || len(q.field) == 0 || len(q.terms) == 0
}

//easyjson:json
type termsSetQuery struct {
	Name                     string      `json:"_name,omitempty"`
	Terms                    []string    `json:"terms"`
	MinimumShouldMatchField  string      `json:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript *Script     `json:"minimum_should_match_script,omitempty"`
	Boost                    interface{} `json:"boost,omitempty"`
}
