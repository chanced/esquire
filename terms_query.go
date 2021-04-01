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
	Terms() (*TermsClause, error)
}

type TermserComplete interface {
	Termser
	CompleteClauser
}

// NewTermsQuery returns a new *TermsQuery
//
// Valid options are:
//  - picker.Terms
//  - picker.Lookup
//  - any type which satisfies Termser that sets the Field value
func NewTermsQuery(params Termser) (*TermsClause, error) {
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

// TermsQuery returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type TermsQuery struct {
	Field           string
	Value           []string
	Boost           interface{}
	CaseInsensitive bool
	completeClause
}

func (t TermsQuery) Clause() (QueryClause, error) {
	return t.Terms()
}
func (t TermsQuery) Terms() (*TermsClause, error) {
	q := &TermsClause{
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

func (t TermsQuery) Kind() Kind {
	return KindTerms
}

func (t *TermsClause) SetField(field string) {
	t.field = field
}

func (t *TermsClause) setValue(value []string) error {
	err := checkValues(value, KindTerms, t.field)
	if err != nil {
		return err
	}
	t.lookup = LookupValues{}
	t.value = value
	return nil
}
func (t *TermsClause) Set(field string, clause Termser) error {
	q, err := clause.Terms()
	if err != nil {
		return NewQueryError(err, KindTerms, field)
	}
	*t = *q
	return nil
}

func (t TermsClause) Kind() Kind {
	return KindTerms
}

func (t *TermsClause) IsEmpty() bool {
	return t == nil || len(t.field) == 0 || (len(t.value) == 0 && t.lookup.IsEmpty())
}

func (t TermsClause) Value() []string {
	return t.value[:]
}

func (t TermsClause) Lookup() *TermsLookup {
	return &t.lookup
}

func (t TermsClause) setLookup(value LookupValues) error {
	t.value = []string{}
	t.lookup = value
	return nil
}

func (t TermsClause) MarshalJSON() ([]byte, error) {
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

func (t *TermsClause) unmarshalValueJSON(data dynamic.JSON) error {
	var val []string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	t.value = val
	return nil
}
func (t *TermsClause) unmarshalLookupJSON(data []byte) error {
	var tl TermsLookup
	err := json.Unmarshal(data, &tl)
	if err != nil {
		return err
	}
	t.lookup = tl
	return nil
}
func (t *TermsClause) UnmarshalJSON(data []byte) error {
	*t = TermsClause{}

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

func (t *TermsClause) UnmarshalBSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

type TermsClause struct {
	lookup LookupValues
	value  []string
	field  string
	boostParam
	caseInsensitiveParam
	nameParam
	completeClause
}

func (f *TermsClause) Clause() (QueryClause, error) {
	return f, nil
}
func (t TermsClause) Field() string {
	return t.field
}

func (t TermsClause) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}
func (t *TermsClause) Clear() {
	*t = TermsClause{}
}
