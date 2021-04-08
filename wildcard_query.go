package picker

type WildcardQueryParams struct {
	Name string
	completeClause
}

func (WildcardQueryParams) Kind() QueryKind {
	return QueryKindWildcard
}

func (p WildcardQueryParams) Clause() (QueryClause, error) {
	return p.Wildcard()
}
func (p WildcardQueryParams) Wildcard() (*WildcardQuery, error) {
	q := &WildcardQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type WildcardQuery struct {
	nameParam
	completeClause
}

func (WildcardQuery) Kind() QueryKind {
	return QueryKindWildcard
}
func (q *WildcardQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *WildcardQuery) Wildcard() (QueryClause, error) {
	return q, nil
}
func (q *WildcardQuery) Clear() {
	if q == nil {
		return
	}
	*q = WildcardQuery{}
}
func (q *WildcardQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q WildcardQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *WildcardQuery) IsEmpty() bool {
	panic("not implemented")
}

type wildcardQuery struct {
	Name string `json:"_name,omitempty"`
}
