package picker

type PercolateQueryParams struct {
	Name string
	completeClause
}

func (PercolateQueryParams) Kind() QueryKind {
	return QueryKindPercolate
}

func (p PercolateQueryParams) Clause() (QueryClause, error) {
	return p.Percolate()
}
func (p PercolateQueryParams) Percolate() (*PercolateQuery, error) {
	q := &PercolateQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type PercolateQuery struct {
	nameParam
	completeClause
}

func (PercolateQuery) Kind() QueryKind {
	return QueryKindPercolate
}
func (q *PercolateQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *PercolateQuery) Percolate() (QueryClause, error) {
	return q, nil
}
func (q *PercolateQuery) Clear() {
	if q == nil {
		return
	}
	*q = PercolateQuery{}
}
func (q *PercolateQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q PercolateQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *PercolateQuery) IsEmpty() bool {
	panic("not implemented")
}

type percolateQuery struct {
	Name string `json:"_name,omitempty"`
}
