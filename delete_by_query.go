package picker

import (
	"bytes"
	"encoding/json"
)

type DeleteByQuerier interface {
	DeleteByQuery() (*DeleteByQuery, error)
}
type DeleteByQueryParams struct {
	Query Querier
}

func (p DeleteByQueryParams) DeleteByQuery() (*DeleteByQuery, error) {
	d := &DeleteByQuery{}
	if p.Query == nil {
		return d, nil
	}
	q, err := p.Query.Query()
	if err != nil {
		return d, err
	}
	d.query = q
	return d, nil
}

type DeleteByQuery struct {
	query *Query
}

func (d DeleteByQuery) Encode() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(d)
	return buf, err
}

func (d DeleteByQuery) MarshalJSON() ([]byte, error) {
	return deleteByQuery{
		Query: d.query,
	}.MarshalJSON()
}

func (d *DeleteByQuery) UnmarshalJSON(data []byte) error {
	*d = DeleteByQuery{}
	p := deleteByQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	d.query = p.Query
	return nil
}

func (d DeleteByQuery) MarshalBSON() ([]byte, error) {
	return d.MarshalJSON()
}

func (d *DeleteByQuery) UnmarshalBSON(data []byte) error {
	return d.UnmarshalJSON(data)
}

//easyjson:json
type deleteByQuery struct {
	Query *Query `json:"query,omitempty"`
}

func (d *DeleteByQuery) Query() *Query {
	if d.query == nil {
		d.query = &Query{}
	}
	return d.query
}

func NewDeleteByQuery(params DeleteByQueryParams) (*DeleteByQuery, error) {
	return params.DeleteByQuery()
}
