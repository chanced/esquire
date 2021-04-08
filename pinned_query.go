package picker

type PinnedQueryParams struct {
	Name string
	completeClause
}

func (PinnedQueryParams) Kind() QueryKind {
	return QueryKindPinned
}

func (p PinnedQueryParams) Clause() (QueryClause, error) {
	return p.Pinned()
}
func (p PinnedQueryParams) Pinned() (*PinnedQuery, error) {
	q := &PinnedQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type PinnedQuery struct {
	nameParam
	completeClause
}

func (PinnedQuery) Kind() QueryKind {
	return QueryKindPinned
}
func (q *PinnedQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *PinnedQuery) Pinned() (*PinnedQuery, error) {
	return q, nil
}
func (q *PinnedQuery) Clear() {
	if q == nil {
		return
	}
	*q = PinnedQuery{}
}
func (q *PinnedQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q PinnedQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *PinnedQuery) IsEmpty() bool {
	panic("not implemented")
}

type pinnedQuery struct {
	Name string `json:"_name,omitempty"`
}
