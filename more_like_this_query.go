package picker

type MoreLikeThisQueryParams struct {
	Name string
	completeClause
}

func (MoreLikeThisQueryParams) Kind() QueryKind {
	return QueryKindMoreLikeThis
}

func (p MoreLikeThisQueryParams) Clause() (QueryClause, error) {
	return p.MoreLikeThis()
}
func (p MoreLikeThisQueryParams) MoreLikeThis() (*MoreLikeThisQuery, error) {
	q := &MoreLikeThisQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type MoreLikeThisQuery struct {
	nameParam
	completeClause
}

func (MoreLikeThisQuery) Kind() QueryKind {
	return QueryKindMoreLikeThis
}
func (q *MoreLikeThisQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *MoreLikeThisQuery) MoreLikeThis() (QueryClause, error) {
	return q, nil
}
func (q *MoreLikeThisQuery) Clear() {
	if q == nil {
		return
	}
	*q = MoreLikeThisQuery{}
}
func (q *MoreLikeThisQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q MoreLikeThisQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *MoreLikeThisQuery) IsEmpty() bool {
	panic("not implemented")
}

type moreLikeThisQuery struct {
	Name string `json:"_name,omitempty"`
}
