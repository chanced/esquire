package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Termer interface {
	Term() (*TermQuery, error)
}

type Term struct {
	// The field which is being queried against.
	//
	// This will be ignored if set through a mutator
	Field           string
	Value           string
	Boost           interface{}
	CaseInsensitive bool
	Name            string
	clause
}

func (t Term) name() string {
	return t.Name
}
func (t Term) field() string {
	return t.Field
}

func (t Term) Clause() (Clause, error) {
	return t.Term()
}

func (t Term) Term() (*TermQuery, error) {
	q := &TermQuery{
		field: t.Field,
	}
	err := q.setValue(t.Value)
	if err != nil {
		return q, NewQueryError(err, KindTerm, t.Field)
	}
	err = q.SetBoost(t.Boost)
	if err != nil {
		return q, NewQueryError(err, KindTerm, t.Field)
	}
	q.SetCaseInsensitive(t.CaseInsensitive)
	q.SetName(t.Name)
	return q, nil
}

func (t Term) Kind() Kind {
	return KindTerm
}

// NewTermQuery creates a new TermQuery
//
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
func NewTermQuery(params Term) (*TermQuery, error) {
	q, err := params.Term()
	if err != nil {
		return nil, NewQueryError(err, KindTerm, params.Field)
	}
	if len(q.field) == 0 {
		return nil, NewQueryError(ErrFieldRequired, KindTerm)
	}
	return q, nil
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
	field string
	value string
	boostParam
	caseInsensitiveParam
	nameParam
	clause
}

func (t *TermQuery) IsEmpty() bool {
	return t == nil || len(t.value) == 0 || len(t.field) == 0
}

func (t TermQuery) Field() string {
	return t.field
}
func (t TermQuery) Kind() Kind {
	return KindTerm
}
func (t TermQuery) Value() string {
	return t.value
}

func (t *TermQuery) setValue(v string) error {
	if len(v) == 0 {
		return NewQueryError(ErrValueRequired, KindTerm, t.field)
	}
	t.value = v
	return nil
}

func (t TermQuery) MarshalJSON() ([]byte, error) {
	if t.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := t.marshalClauseJSON()

	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{t.field: data})
}

func (t *TermQuery) UnmarshalJSON(data []byte) error {
	t.Clear()
	m := map[string]dynamic.JSON{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		t.field = k
		return t.unmarshalClauseJSON(v)
	}
	return nil
}
func (t *TermQuery) Set(field string, clause Termer) error {
	if clause == nil {
		t.Clear()
		return nil
	}
	if len(field) == 0 {
		return NewQueryError(ErrFieldRequired, KindTerm)
	}
	q, err := clause.Term()
	if err != nil {
		return NewQueryError(err, KindTerm, field)
	}
	*t = *q
	t.field = field
	return nil
}
func (t *TermQuery) Clear() {
	*t = TermQuery{}
}

func (t TermQuery) marshalClauseJSON() (dynamic.JSON, error) {

	params, err := marshalParams(&t)
	if err != nil {
		return nil, err
	}
	params["value"] = t.value
	return json.Marshal(params)
}

func (t *TermQuery) unmarshalJSONString(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	t.value = str
	return nil
}

func (t *TermQuery) unmarshalJSONObject(data []byte) error {
	fields, err := unmarshalParams(data, t)
	if err != nil {
		return err
	}
	if v, ok := fields["value"]; ok {
		var str string
		err := json.Unmarshal(v, &str)
		if err != nil {
			return err
		}
		t.value = str
	}
	return nil
}

func (t *TermQuery) unmarshalClauseJSON(data []byte) error {
	d := dynamic.JSON(data)
	if d.IsString() {
		return t.unmarshalJSONString(data)
	}
	return t.unmarshalJSONObject(data)
}
