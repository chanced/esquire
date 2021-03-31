package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Termer interface {
	Term() (*TermClause, error)
}

type TermQuery struct {
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

func (t TermQuery) name() string {
	return t.Name
}
func (t TermQuery) field() string {
	return t.Field
}

func (t TermQuery) Clause() (Clause, error) {
	return t.Term()
}

func (t TermQuery) Term() (*TermClause, error) {
	q := &TermClause{
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

func (t TermQuery) Kind() Kind {
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
func NewTermQuery(params TermQuery) (*TermClause, error) {
	q, err := params.Term()
	if err != nil {
		return nil, NewQueryError(err, KindTerm, params.Field)
	}
	if len(q.field) == 0 {
		return nil, NewQueryError(ErrFieldRequired, KindTerm)
	}
	return q, nil
}

// TermClause returns documents that contain an exact term in a provided field.
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
type TermClause struct {
	field string
	value string
	boostParam
	caseInsensitiveParam
	nameParam
	clause
}

func (t *TermClause) IsEmpty() bool {
	return t == nil || len(t.value) == 0 || len(t.field) == 0
}

func (t TermClause) Field() string {
	return t.field
}
func (t TermClause) Kind() Kind {
	return KindTerm
}
func (t TermClause) Value() string {
	return t.value
}

func (t *TermClause) setValue(v string) error {
	if len(v) == 0 {
		return NewQueryError(ErrValueRequired, KindTerm, t.field)
	}
	t.value = v
	return nil
}

func (t TermClause) MarshalJSON() ([]byte, error) {
	if t.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := t.marshalClauseJSON()

	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{t.field: data})
}

func (t *TermClause) UnmarshalJSON(data []byte) error {
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
func (t *TermClause) Set(field string, clause Termer) error {
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
func (t *TermClause) Clear() {
	*t = TermClause{}
}

func (t TermClause) marshalClauseJSON() (dynamic.JSON, error) {

	params, err := marshalClauseParams(&t)
	if err != nil {
		return nil, err
	}
	params["value"] = t.value
	return json.Marshal(params)
}

func (t *TermClause) unmarshalJSONString(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	t.value = str
	return nil
}

func (t *TermClause) unmarshalJSONObject(data []byte) error {
	fields, err := unmarshalClauseParams(data, t)
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

func (t *TermClause) unmarshalClauseJSON(data []byte) error {
	d := dynamic.JSON(data)
	if d.IsString() {
		return t.unmarshalJSONString(data)
	}
	return t.unmarshalJSONObject(data)
}
