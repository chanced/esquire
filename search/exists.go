package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Exists returns documents that contain an indexed value for a field.
//
// An indexed value may not exist for a document’s field due to a variety of
// reasons:
//
//  	- The field in the source JSON is null or []
//
//  	- The field has "index" : false set in the mapping
//
//  	- The length of the field value exceeded an ignore_above setting in the
// mapping
//
//  	- The field value was malformed and ignore_malformed was defined in the
// mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type Exists struct {
	Field string
}

func (e Exists) Clause() (Clause, error) {
	return e.Exists()
}

func (e Exists) Exists() (*ExistsQuery, error) {
	q := &ExistsQuery{}
	err := q.SetField(e.Field)
	if err != nil {
		return q, NewQueryError(err, TypeExists, e.Field)
	}
	return q, nil
}

// ExistsQuery returns documents that contain an indexed value for a field.
//
// An indexed value may not exist for a document’s field due to a variety of
// reasons:
//
//  	- The field in the source JSON is null or []
//
//  	- The field has "index" : false set in the mapping
//
//  	- The length of the field value exceeded an ignore_above setting in the
// mapping
//
//  	- The field value was malformed and ignore_malformed was defined in the
// mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type ExistsQuery struct {
	field string
}

func (e ExistsQuery) Field() string {
	return e.field
}

func (e *ExistsQuery) SetField(field string) error {
	e.field = field
	return nil
}

func (e *ExistsQuery) Set(field string) error {
	return e.SetField(field)
}

func (e *ExistsQuery) IsEmpty() bool {
	return len(e.field) == 0
}

func (e ExistsQuery) Type() Type {
	return TypeExists
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
