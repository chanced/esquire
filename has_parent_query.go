package picker

type HasParentQueryParams struct {
	Name string
	completeClause
}

func (HasParentQueryParams) Kind() QueryKind {
	return QueryKindHasParent
}

func (p HasParentQueryParams) Clause() (QueryClause, error) {
	return p.HasParent()
}
func (p HasParentQueryParams) HasParent() (*HasParentQuery, error) {
	q := &HasParentQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type HasParentQuery struct {
	nameParam
	completeClause
}

func (HasParentQuery) Kind() QueryKind {
	return QueryKindHasParent
}
func (q *HasParentQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *HasParentQuery) HasParent() (*HasParentQuery, error) {
	return q, nil
}
func (q *HasParentQuery) Clear() {
	if q == nil {
		return
	}
	*q = HasParentQuery{}
}
func (q *HasParentQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *HasParentQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q HasParentQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q HasParentQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *HasParentQuery) IsEmpty() bool {
	panic("not implemented")
}

type hasParentQuery struct {
	Name string `json:"_name,omitempty"`
}
