package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Exister interface {
	Exists() (*ExistsQuery, error)
}

// ExistsQueryParams returns documents that contain an indexed value for a field.
//
// An indexed value may not exist for a document’s field due to a variety of
// reasons:
//
// - The field in the source JSON is null or []
//
// - The field has "index" : false set in the mapping
//
// - The length of the field value exceeded an ignore_above setting in the
// mapping
//
// - The field value was malformed and ignore_malformed was defined in the
// mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type ExistsQueryParams struct {
	Field string
	Name  string
	completeClause
}

func (e ExistsQueryParams) Clause() (QueryClause, error) {
	return e.Exists()
}

func (e ExistsQueryParams) Exists() (*ExistsQuery, error) {
	q := &ExistsQuery{}
	err := q.SetField(e.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindExists, e.Field)
	}
	q.SetName(e.Name)
	return q, nil
}

// ExistsQuery returns documents that contain an indexed value for a field.
//
// An indexed value may not exist for a document’s field due to a variety of
// reasons:
//
// - The field in the source JSON is null or []
//
// - The field has "index" : false set in the mapping
//
// - The length of the field value exceeded an ignore_above setting in the
// mapping
//
// - The field value was malformed and ignore_malformed was defined in the
// mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type ExistsQuery struct {
	field string
	nameParam
	completeClause
}

var _ QueryClause = (*ExistsQuery)(nil)

func (e *ExistsQuery) Clause() (QueryClause, error) {
	return e, nil
}

func (e *ExistsQuery) Field() string {
	if e == nil {
		return ""
	}
	return e.field
}

func (e *ExistsQuery) SetField(field string) error {
	e.field = field
	return nil
}

func (e *ExistsQuery) Set(field string, exists Exister) error {
	if field == "" {
		e.Clear()
		return nil
	}
	_ = e.SetField(field)
	if exists != nil {
		ex, _ := exists.Exists()
		e.SetName(ex.Name())
	}
	return nil
}

func (e *ExistsQuery) IsEmpty() bool {
	return e == nil || len(e.field) == 0
}

func (e ExistsQuery) Kind() QueryKind {
	return QueryKindExists
}

func (e ExistsQuery) MarshalJSON() ([]byte, error) {
	if e.IsEmpty() {
		return dynamic.Null, nil
	}
	return json.Marshal(map[string]string{
		"field": e.field,
	})
}

func (e *ExistsQuery) UnmarshalJSON(data []byte) error {
	*e = ExistsQuery{}
	d := dynamic.JSON(data)
	if d.IsNull() {
		return nil
	}
	var m map[string]string
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	e.field = m["field"]
	return nil
}

func (e *ExistsQuery) Clear() {
	*e = ExistsQuery{}
}
