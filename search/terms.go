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
	Terms() (TermsQuery, error)
}

// Terms returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type Terms struct {
	Field           string
	Values          []string
	Boost           dynamic.Number
	CaseInsensitive bool
}

func (t Terms) Clause() (Clause, error) {
	return t.Terms()
}
func (t Terms) Terms() (TermsQuery, error) {
	q := TermsQuery{
		value: t.Values,
		field: t.Field,
	}
	if f, ok := t.Boost.Float(); ok {
		t.Boost.Set(f)
	}
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t Terms) Type() Type {
	return TypeTerms
}

type TermsLookup struct {
	ID      string `json:"id,omitempty"`
	Index   string `json:"index,omitempty"`
	Path    string `json:"path,omitempty"`
	Routing string `json:"routing,omitempty"`
}

func (t TermsLookup) Validate() error {
	if len(t.ID) == 0 {
		return ErrIDRequired
	}
	if len(t.Index) == 0 {
		return ErrIndexRequired
	}
	if len(t.Path) == 0 {
		return ErrPathRequired
	}
	return nil
}

func (t TermsLookup) IsEmpty() bool {
	return len(t.ID) == 0 && len(t.Index) == 0 && len(t.Path) == 0 && len(t.Routing) == 0
}

func (t *TermsQuery) SetField(field string) {
	t.field = field
}

// SetValues sets the Terms value to v and clears the lookup
// It returns an error if v is empty.
//
// If you need to to clear Terms, use Clear()
func (t *TermsQuery) SetValues(v []string) error {
	t.lookup = TermsLookup{}
	t.value = v
	if len(v) == 0 {
		return ErrValueRequired
	}
	return nil
}

func (t TermsQuery) Value() []string {
	return t.value
}

func (t TermsQuery) Lookup() TermsLookup {
	return t.lookup
}

// SetLookup sets the Terms query's lookup and unsets values.
func (t TermsQuery) SetLookup(v TermsLookup) error {

	t.value = []string{}
	t.lookup = v
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
	lookup TermsLookup
	value  []string
	field  string
	boostParam
	caseInsensitiveParam
	nameParam
}

func (t TermsQuery) Field() string {
	return t.field
}

func (t *TermsQuery) Set(field string, clause Termser) error {
	q, err := clause.Terms()
	if err != nil {
		return err
	}
	*t = q
	return nil
}

func (t TermsQuery) Type() Type {
	return TypeTerms
}

func (t TermsQuery) IsEmpty() bool {
	return (len(t.value) == 0 && t.lookup.IsEmpty())
}

func (t TermsQuery) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}
