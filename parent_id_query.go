package picker

type ParentIDQueryParams struct {
	Name string
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
	_ = q
	panic("not implemented")
	// return q, nil
}

type ParentIDQuery struct {
	nameParam
	completeClause
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
func (q *ParentIDQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q ParentIDQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *ParentIDQuery) IsEmpty() bool {
	panic("not implemented")
}

type parentIDQuery struct {
	Name string `json:"_name,omitempty"`
}
