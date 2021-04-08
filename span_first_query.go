package picker

type SpanFirstQueryParams struct {
	Name string
	completeClause
}

func (SpanFirstQueryParams) Kind() QueryKind {
	return QueryKindSpanFirst
}

func (p SpanFirstQueryParams) Clause() (QueryClause, error) {
	return p.SpanFirst()
}
func (p SpanFirstQueryParams) SpanFirst() (*SpanFirstQuery, error) {
	q := &SpanFirstQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanFirstQuery struct {
	nameParam
	completeClause
}

func (SpanFirstQuery) Kind() QueryKind {
	return QueryKindSpanFirst
}
func (q *SpanFirstQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanFirstQuery) SpanFirst() (*SpanFirstQuery, error) {
	return q, nil
}
func (q *SpanFirstQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanFirstQuery{}
}
func (q *SpanFirstQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanFirstQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanFirstQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanFirstQuery struct {
	Name string `json:"_name,omitempty"`
}
