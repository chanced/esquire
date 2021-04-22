package picker

import (
	"bytes"
	"encoding/json"
)

type UpdateByQuerier interface {
	UpdateByQuery() (*UpdateByQuery, error)
}
type UpdateByQueryParams struct {
	Query Querier
}

func (p UpdateByQueryParams) UpdateByQuery() (*UpdateByQuery, error) {
	u := &UpdateByQuery{}
	if p.Query == nil {
		return u, nil
	}
	q, err := p.Query.Query()
	if err != nil {
		return u, err
	}
	u.query = q
	return u, nil
}

type UpdateByQuery struct {
	query *Query
}

func (u UpdateByQuery) Encode() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(u)
	return buf, err
}

func (u UpdateByQuery) MarshalJSON() ([]byte, error) {
	return updateByQuery{
		Query: u.query,
	}.MarshalJSON()
}

func (u *UpdateByQuery) UnmarshalJSON(data []byte) error {
	*u = UpdateByQuery{}
	p := updateByQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	u.query = p.Query
	return nil
}

func (u UpdateByQuery) MarshalBSON() ([]byte, error) {
	return u.MarshalJSON()
}

func (u *UpdateByQuery) UnmarshalBSON(data []byte) error {
	return u.UnmarshalJSON(data)
}

//easyjson:json
type updateByQuery struct {
	Query *Query `json:"query,omitempty"`
}

func (u *UpdateByQuery) Query() *Query {
	if u.query == nil {
		u.query = &Query{}
	}
	return u.query
}

func NewUpdateByQuery(params UpdateByQueryParams) (*UpdateByQuery, error) {
	return params.UpdateByQuery()
}
