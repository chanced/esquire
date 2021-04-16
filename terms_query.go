package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Termser can be:
//  *picker.Terms
//  *picker.Lookup
//  picker.String
//  picker.Strings
type Termser interface {
	Terms() (*TermsQuery, error)
}

type CompleteTermser interface {
	Termser
	CompleteClauser
}

// TermsQueryParams returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type TermsQueryParams struct {
	Field           string
	Value           []string
	Boost           interface{}
	CaseInsensitive bool
	completeClause
}

func (t TermsQueryParams) Clause() (QueryClause, error) {
	return t.Terms()
}
func (t TermsQueryParams) Terms() (*TermsQuery, error) {
	q := &TermsQuery{
		field: t.Field,
	}
	err := q.SetBoost(t.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindTerms, t.Field)
	}
	err = q.SetValue(t.Value)
	if err != nil {
		return q, newQueryError(err, QueryKindTerms, t.Field)
	}
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t TermsQueryParams) Kind() QueryKind {
	return QueryKindTerms
}

func (t *TermsQuery) SetField(field string) {
	t.field = field
}

func (t *TermsQuery) SetValue(value []string) error {
	err := checkValues(value, QueryKindTerms, t.field)
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
		return newQueryError(err, QueryKindTerms, field)
	}
	*t = *q
	return nil
}

func (t TermsQuery) Kind() QueryKind {
	return QueryKindTerms
}

func (t *TermsQuery) IsEmpty() bool {
	return t == nil || len(t.field) == 0 || (len(t.value) == 0 && t.lookup.IsEmpty())
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
	v, err := marshalClauseParams(&t)
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
func (t *TermsQuery) UnmarshalBSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

func (t *TermsQuery) UnmarshalJSON(data []byte) error {
	*t = TermsQuery{}

	d := dynamic.JSON(data)
	if len(data) == 0 || d.IsEmptyObject() {
		return nil
	}
	if d.IsNull() {
		return nil
	}

	fields, err := unmarshalClauseParams(data, t)
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

type TermsQuery struct {
	lookup LookupValues
	value  []string
	field  string
	boostParam
	caseInsensitiveParam
	nameParam
	completeClause
}

func (t *TermsQuery) Terms() (*TermsQuery, error) {
	return t, nil
}
func (t *TermsQuery) Clause() (QueryClause, error) {
	return t, nil
}
func (t TermsQuery) Field() string {
	return t.field
}

func (t TermsQuery) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}
func (t *TermsQuery) Clear() {
	*t = TermsQuery{}
}
