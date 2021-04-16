package picker

import "github.com/chanced/dynamic"

type ParentIDer interface {
	ParentID() (*ParentIDQuery, error)
}

type ParentIDQueryParams struct {
	Name string
	// (Required) ID of the parent document. The query will return child documents of this parent docume
	ID string
	// (Required, string) Name of the child relationship mapped for the join field.
	Type string
	// (Optional, Boolean) Indicates whether to ignore an unmapped type and not
	// return any documents instead of an error. Defaults to false.
	//
	// If false, Elasticsearch returns an error if the type is unmapped.
	//
	// You can use this parameter to query multiple indices that may not contain the
	// type.
	IgnoreUnmapped interface{}
	completeClause
}

func (ParentIDQueryParams) Kind() QueryKind {
	return QueryKindParentID
}

func (p ParentIDQueryParams) Clause() (QueryClause, error) {
	return p.ParentID()
}
func (p ParentIDQueryParams) ParentID() (*ParentIDQuery, error) {
	q := &ParentIDQuery{}
	err := q.SetID(p.ID)
	if err != nil {
		return q, err
	}
	q.name = p.Name
	err = q.SetType(p.Type)
	if err != nil {
		return q, err
	}
	err = q.ignoreUnmapped.Set(p.IgnoreUnmapped)
	if err != nil {
		return q, err
	}
	return q, nil
}

type ParentIDQuery struct {
	nameParam
	id             string
	typ            string
	ignoreUnmapped dynamic.Bool
	completeClause
}

func (q ParentIDQuery) ID() string {
	return q.id
}
func (q *ParentIDQuery) SetID(id string) error {
	if len(id) == 0 {
		return ErrIDRequired
	}
	q.id = id
	return nil
}
func (q ParentIDQuery) IgnoreUnmapped() bool {
	if b, ok := q.ignoreUnmapped.Bool(); ok {
		return b
	}
	return false
}
func (q *ParentIDQuery) SetIgnoreUnmapped(ignore interface{}) error {
	return q.ignoreUnmapped.Set(ignore)
}

func (q ParentIDQuery) Type() string {
	return q.typ
}
func (q *ParentIDQuery) SetType(typ string) error {
	if len(typ) == 0 {
		return ErrTypeRequired
	}
	q.typ = typ
	return nil
}

func (ParentIDQuery) Kind() QueryKind {
	return QueryKindParentID
}
func (q *ParentIDQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *ParentIDQuery) ParentID() (*ParentIDQuery, error) {
	return q, nil
}
func (q *ParentIDQuery) Clear() {
	if q == nil {
		return
	}
	*q = ParentIDQuery{}
}
func (q *ParentIDQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *ParentIDQuery) UnmarshalJSON(data []byte) error {
	q.Clear()
	p := parentIDQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	q.id = p.ID
	q.name = p.Name
	q.typ = p.Type
	q.ignoreUnmapped.Set(p.IgnoreUnmapped)
	return nil
}
func (q ParentIDQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q ParentIDQuery) MarshalJSON() ([]byte, error) {
	return parentIDQuery{
		Name:           q.name,
		ID:             q.id,
		Type:           q.typ,
		IgnoreUnmapped: q.ignoreUnmapped.Value(),
	}.MarshalJSON()
}
func (q *ParentIDQuery) IsEmpty() bool {
	return q == nil || len(q.typ) == 0 || len(q.id) == 0
}

//easyjson:json
type parentIDQuery struct {
	Name           string      `json:"_name,omitempty"`
	ID             string      `json:"id"`
	Type           string      `json:"type"`
	IgnoreUnmapped interface{} `json:"ignore_unmapped,omitempty"`
}
