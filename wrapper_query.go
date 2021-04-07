package picker

type WrapperQueryParams struct {
	Name string
	completeClause
}

func (WrapperQueryParams) Kind() QueryKind {
	return QueryKindWrapper
}

func (p WrapperQueryParams) Clause() (QueryClause, error) {
	return p.Wrapper()
}
func (p WrapperQueryParams) Wrapper() (*WrapperQuery, error) {
	q := &WrapperQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type WrapperQuery struct {
	nameParam
	completeClause
}

func (WrapperQuery) Kind() QueryKind {
	return QueryKindWrapper
}
func (q *WrapperQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *WrapperQuery) Wrapper() (QueryClause, error) {
	return q, nil
}
func (q *WrapperQuery) Clear() {
	if q == nil {
		return
	}
	*q = WrapperQuery{}
}
func (q *WrapperQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q WrapperQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *WrapperQuery) IsEmpty() bool {
	panic("not implemented")
}

type wrapperQuery struct {
	Name string `json:"_name,omitempty"`
}
