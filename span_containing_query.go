package picker

type SpanContainingQueryParams struct {
	Name string
	completeClause
}

func (SpanContainingQueryParams) Kind() QueryKind {
	return QueryKindSpanContaining
}

func (p SpanContainingQueryParams) Clause() (QueryClause, error) {
	return p.SpanContaining()
}
func (p SpanContainingQueryParams) SpanContaining() (*SpanContainingQuery, error) {
	q := &SpanContainingQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanContainingQuery struct {
	nameParam
	completeClause
}

func (SpanContainingQuery) Kind() QueryKind {
	return QueryKindSpanContaining
}
func (q *SpanContainingQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanContainingQuery) SpanContaining() (*SpanContainingQuery, error) {
	return q, nil
}
func (q *SpanContainingQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanContainingQuery{}
}
func (q *SpanContainingQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanContainingQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanContainingQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanContainingQuery struct {
	Name string `json:"_name,omitempty"`
}
