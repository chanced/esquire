package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Exists returns documents that contain an indexed value for a field.
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
type Exists struct {
	Field string
	Name  string
	completeClause
}

func (e Exists) Clause() (QueryClause, error) {
	return e.Exists()
}

func (e Exists) Exists() (*ExistsClause, error) {
	q := &ExistsClause{}
	err := q.SetField(e.Field)
	if err != nil {
		return q, NewQueryError(err, KindExists, e.Field)
	}
	q.SetName(e.Name)
	return q, nil
}

// ExistsClause returns documents that contain an indexed value for a field.
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
type ExistsClause struct {
	field string
	nameParam
	completeClause
}

var _ QueryClause = (*ExistsClause)(nil)

func (e *ExistsClause) Clause() (QueryClause, error) {
	return e, nil
}

func (e *ExistsClause) Field() string {
	if e == nil {
		return ""
	}
	return e.field
}

func (e *ExistsClause) SetField(field string) error {
	e.field = field
	return nil
}

func (e *ExistsClause) Set(field string) error {
	return e.SetField(field)
}

func (e *ExistsClause) IsEmpty() bool {
	return len(e.field) == 0
}

func (e ExistsClause) Kind() QueryKind {
	return KindExists
}

func (e ExistsClause) MarshalJSON() ([]byte, error) {
	if e.IsEmpty() {
		return dynamic.Null, nil
	}
	return json.Marshal(map[string]string{
		"field": e.field,
	})
}

func (e *ExistsClause) UnmarshalJSON(data []byte) error {
	*e = ExistsClause{}
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

func (e *ExistsClause) Clear() {
	*e = ExistsClause{}
}
