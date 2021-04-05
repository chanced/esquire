package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type CompleteTermer interface {
	Termer
	CompleteClauser
}

type Termer interface {
	Term() (*TermQuery, error)
}

type TermQueryParams struct {
	// The field which is being queried against.
	//
	// This will be ignored if set through a mutator
	Field           string
	Value           string
	Boost           interface{}
	CaseInsensitive bool
	Name            string
	completeClause
}

func (t TermQueryParams) Clause() (QueryClause, error) {
	return t.Term()
}

func (t TermQueryParams) Term() (*TermQuery, error) {
	q := &TermQuery{
		field: t.Field,
	}
	err := q.SetValue(t.Value)
	if err != nil {
		return q, newQueryError(err, QueryKindTerm, t.Field)
	}
	err = q.SetBoost(t.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindTerm, t.Field)
	}
	q.SetCaseInsensitive(t.CaseInsensitive)
	q.SetName(t.Name)
	return q, nil
}

func (t TermQueryParams) Kind() QueryKind {
	return QueryKindTerm
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
func NewTermQuery(params TermQueryParams) (*TermQuery, error) {
	q, err := params.Term()
	if err != nil {
		return nil, newQueryError(err, QueryKindTerm, params.Field)
	}
	if len(q.field) == 0 {
		return nil, newQueryError(ErrFieldRequired, QueryKindTerm)
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
	completeClause
}

func (t *TermQuery) Clause() (QueryClause, error) {
	return t, nil
}
func (t *TermQuery) IsEmpty() bool {
	return t == nil || len(t.value) == 0 || len(t.field) == 0
}

func (t TermQuery) Field() string {
	return t.field
}
func (t TermQuery) Kind() QueryKind {
	return QueryKindTerm
}
func (t TermQuery) Value() string {
	return t.value
}

func (t *TermQuery) SetValue(v string) error {
	if len(v) == 0 {
		return newQueryError(ErrValueRequired, QueryKindTerm, t.field)
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
		return newQueryError(ErrFieldRequired, QueryKindTerm)
	}
	q, err := clause.Term()
	if err != nil {
		return newQueryError(err, QueryKindTerm, field)
	}
	*t = *q
	t.field = field
	return nil
}
func (t *TermQuery) Clear() {
	if t == nil {
		return
	}
	*t = TermQuery{}
}

func (t TermQuery) marshalClauseJSON() (dynamic.JSON, error) {

	params, err := marshalClauseParams(&t)
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

func (t *TermQuery) unmarshalClauseJSON(data []byte) error {
	d := dynamic.JSON(data)
	if d.IsString() {
		return t.unmarshalJSONString(data)
	}
	return t.unmarshalJSONObject(data)
}
