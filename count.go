package picker

type Counter interface {
	Count() (*Count, error)
}
type CountParams struct {
	Query Querier
}

func (p CountParams) Count() (*Count, error) {
	u := &Count{}
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

type Count struct {
	query *Query
}

func (u *Count) Query() *Query {
	if u.query == nil {
		u.query = &Query{}
	}
	return u.query
}

func NewCount(params CountParams) (*Count, error) {
	return params.Count()
}
