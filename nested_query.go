package picker

type NestedQueryParams struct {
	Name string
	completeClause
}

func (NestedQueryParams) Kind() QueryKind {
	return QueryKindNested
}

func (p NestedQueryParams) Clause() (QueryClause, error) {
	return p.Nested()
}
func (p NestedQueryParams) Nested() (*NestedQuery, error) {
	q := &NestedQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type NestedQuery struct {
	nameParam
	completeClause
}

func (NestedQuery) Kind() QueryKind {
	return QueryKindNested
}
func (q *NestedQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *NestedQuery) Nested() (QueryClause, error) {
	return q, nil
}
func (q *NestedQuery) Clear() {
	if q == nil {
		return
	}
	*q = NestedQuery{}
}
func (q *NestedQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q NestedQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *NestedQuery) IsEmpty() bool {
	panic("not implemented")
}

type nestedQuery struct {
	Name string `json:"_name,omitempty"`
}
