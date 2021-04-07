package picker

type HasChildQueryParams struct {
	Name string
	completeClause
}

func (HasChildQueryParams) Kind() QueryKind {
	return QueryKindHasChild
}

func (p HasChildQueryParams) Clause() (QueryClause, error) {
	return p.HasChild()
}
func (p HasChildQueryParams) HasChild() (*HasChildQuery, error) {
	q := &HasChildQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type HasChildQuery struct {
	nameParam
	completeClause
}

func (HasChildQuery) Kind() QueryKind {
	return QueryKindHasChild
}
func (q *HasChildQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *HasChildQuery) HasChild() (QueryClause, error) {
	return q, nil
}
func (q *HasChildQuery) Clear() {
	if q == nil {
		return
	}
	*q = HasChildQuery{}
}
func (q *HasChildQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q HasChildQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *HasChildQuery) IsEmpty() bool {
	panic("not implemented")
}

type hasChildQuery struct {
	Name string `json:"_name,omitempty"`
}
