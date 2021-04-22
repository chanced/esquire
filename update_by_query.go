package picker

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

func (u *UpdateByQuery) Query() *Query {
	if u.query == nil {
		u.query = &Query{}
	}
	return u.query
}

func NewUpdateByQuery(params UpdateByQueryParams) (*UpdateByQuery, error) {
	return params.UpdateByQuery()
}
