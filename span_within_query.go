package picker

type SpanWithinQueryParams struct {
	Name string
	completeClause
}

func (SpanWithinQueryParams) Kind() QueryKind {
	return QueryKindSpanWithin
}

func (p SpanWithinQueryParams) Clause() (QueryClause, error) {
	return p.SpanWithin()
}
func (p SpanWithinQueryParams) SpanWithin() (*SpanWithinQuery, error) {
	q := &SpanWithinQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanWithinQuery struct {
	nameParam
	completeClause
}

func (SpanWithinQuery) Kind() QueryKind {
	return QueryKindSpanWithin
}
func (q *SpanWithinQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanWithinQuery) SpanWithin() (QueryClause, error) {
	return q, nil
}
func (q *SpanWithinQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanWithinQuery{}
}
func (q *SpanWithinQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanWithinQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanWithinQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanWithinQuery struct {
	Name string `json:"_name,omitempty"`
}
