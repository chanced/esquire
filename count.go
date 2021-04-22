package picker

import (
	"bytes"
	"encoding/json"
)

type Counter interface {
	Count() (*Count, error)
}
type CountParams struct {
	Query Querier
}

func (p CountParams) Count() (*Count, error) {
	c := &Count{}
	if p.Query == nil {
		return c, nil
	}
	q, err := p.Query.Query()
	if err != nil {
		return c, err
	}
	c.query = q
	return c, nil
}

type Count struct {
	query *Query
}

func (c *Count) Query() *Query {
	if c.query == nil {
		c.query = &Query{}
	}
	return c.query
}

func NewCount(params CountParams) (*Count, error) {
	return params.Count()
}

//easyjson:json
type count struct {
	Query *Query `json:"query,omitempty"`
}

func (c Count) Encode() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(c)
	return buf, err
}

func (c Count) MarshalJSON() ([]byte, error) {
	return count{
		Query: c.query,
	}.MarshalJSON()
}

func (c *Count) UnmarshalJSON(data []byte) error {
	*c = Count{}
	p := count{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	c.query = p.Query
	return nil
}

func (c Count) MarshalBSON() ([]byte, error) {
	return c.MarshalJSON()
}

func (c *Count) UnmarshalBSON(data []byte) error {
	return c.UnmarshalJSON(data)
}
