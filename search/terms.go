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
type Termser interface {
	Terms() (*TermsQuery, error)
}

// NewTermsQuery returns a new *TermsQuery
//
// Valid options are:
//  - search.Terms
//  - search.Lookup
//  - any type which satisfies Termser that sets the Field value
func NewTermsQuery(params Termser) (*TermsQuery, error) {
	q, err := params.Terms()
	if err != nil {
		return q, NewQueryError(err, KindTerms, q.field)
	}
	err = checkField(q.field, KindTerms)
	if err != nil {
		return q, err
	}
	return q, nil
}

// Terms returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type Terms struct {
	Field           string
	Value           []string
	Boost           interface{}
	CaseInsensitive bool
}

func (t Terms) Clause() (Clause, error) {
	return t.Terms()
}
func (t Terms) Terms() (*TermsQuery, error) {
	q := &TermsQuery{
		field: t.Field,
	}
	err := q.SetBoost(t.Boost)
	if err != nil {
		return q, NewQueryError(err, KindTerms, t.Field)
	}
	err = q.setValue(t.Value)
	if err != nil {
		return q, NewQueryError(err, KindTerms, t.Field)
	}
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t Terms) Kind() Kind {
	return KindTerms
}

func (t *TermsQuery) SetField(field string) {
	t.field = field
}

func (t *TermsQuery) setValue(value []string) error {
	err := checkValues(value, KindTerms, t.field)
	if err != nil {
		return err
	}
	t.lookup = LookupValues{}
	t.value = value
	return nil
}
func (t *TermsQuery) Set(field string, clause Termser) error {
	q, err := clause.Terms()
	if err != nil {
		return NewQueryError(err, KindTerms, field)
	}
	*t = *q
	return nil
}

func (t TermsQuery) Kind() Kind {
	return KindTerms
}

func (t TermsQuery) IsEmpty() bool {
	return (len(t.value) == 0 && t.lookup.IsEmpty())
}

func (t TermsQuery) Value() []string {
	return t.value[:]
}

func (t TermsQuery) Lookup() *TermsLookup {
	return &t.lookup
}

func (t TermsQuery) setLookup(value LookupValues) error {
	t.value = []string{}
	t.lookup = value
	return nil
}

func (t TermsQuery) MarshalJSON() ([]byte, error) {
	if len(t.field) == 0 || t.IsEmpty() {
		return dynamic.Null, nil
	}
	v, err := marshalParams(&t)
	if err != nil {
		return nil, err
	}
	if !t.lookup.IsEmpty() {
		v[t.field] = t.lookup
	} else {
		v[t.field] = t.value
	}
	return json.Marshal(v)
}

func (t *TermsQuery) unmarshalValueJSON(data dynamic.JSON) error {
	var val []string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	t.value = val
	return nil
}
func (t *TermsQuery) unmarshalLookupJSON(data []byte) error {
	var tl TermsLookup
	err := json.Unmarshal(data, &tl)
	if err != nil {
		return err
	}
	t.lookup = tl
	return nil
}
func (t *TermsQuery) UnmarshalJSON(data []byte) error {
	*t = TermsQuery{}
	d := dynamic.JSON(data)
	if d.IsNull() {
		return nil
	}
	fields, err := unmarshalParams(data, t)
	if err != nil {
		return err
	}
	for f, fd := range fields {
		t.field = f
		if fd.IsArray() {
			return t.unmarshalValueJSON(fd)
		}
		return t.unmarshalLookupJSON(fd)
	}
	return err
}

func (t *TermsQuery) UnmarshalBSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

type TermsQuery struct {
	lookup LookupValues
	value  []string
	field  string
	boostParam
	caseInsensitiveParam
	nameParam
}

func (t TermsQuery) Field() string {
	return t.field
}

func (t TermsQuery) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}
