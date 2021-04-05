package picker

import (
	"encoding/json"
)

type IDser interface {
	IDs() (*IDsQuery, error)
}

type IDsQueryParams struct {
	// (Required) An array of document IDs.
	Values []string `json:"values"`
	Name   string   `json:"_name,omitempty"`
}

func (p IDsQueryParams) Clause() (QueryClause, error) {
	return p.IDs()
}
func (p IDsQueryParams) IDs() (*IDsQuery, error) {
	q := &IDsQuery{}
	err := q.SetValues(p.Values)
	if err != nil {
		return q, err
	}
	q.SetName(p.Name)
	return q, nil
}

func (IDsQueryParams) Kind() QueryKind {
	return QueryKindIDs
}

type IDsQuery struct {
	ids []string
	nameParam
	completeClause
}

func (id IDsQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(idsQuery{Values: id.ids, Name: id.name})
}

func (id *IDsQuery) UnmarshalJSON(data []byte) error {
	*id = IDsQuery{}
	var q idsQuery
	err := json.Unmarshal(data, &q)
	if err != nil {
		return err
	}
	id.name = q.Name
	id.ids = q.Values
	return nil
}

func (id *IDsQuery) Clause() (QueryClause, error) {
	return id, nil
}
func (id *IDsQuery) IDs() (*IDsQuery, error) {
	return id, nil
}
func (id IDsQuery) Values() []string {
	return id.ids
}
func (id *IDsQuery) SetValues(ids []string) error {
	if len(ids) == 0 {
		return ErrValuesRequired
	}
	id.ids = ids
	return nil
}

func (id *IDsQuery) Clear() {
	if id == nil || len(id.ids) == 0 {
		return
	}
	*id = IDsQuery{}

}
func (id *IDsQuery) IsEmpty() bool {
	return id == nil || len(id.ids) == 0
}

func (IDsQuery) Kind() QueryKind {
	return QueryKindIDs
}

//easyjson:json
type idsQuery struct {
	Values []string `json:"values"`
	Name   string   `json:"_name,omitempty"`
}
