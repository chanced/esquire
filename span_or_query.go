package picker

type SpanOrQueryParams struct {
	Name string
	completeClause
}

func (SpanOrQueryParams) Kind() QueryKind {
	return QueryKindSpanOr
}

func (p SpanOrQueryParams) Clause() (QueryClause, error) {
	return p.SpanOr()
}
func (p SpanOrQueryParams) SpanOr() (*SpanOrQuery, error) {
	q := &SpanOrQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanOrQuery struct {
	nameParam
	completeClause
}

func (SpanOrQuery) Kind() QueryKind {
	return QueryKindSpanOr
}
func (q *SpanOrQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanOrQuery) SpanOr() (QueryClause, error) {
	return q, nil
}
func (q *SpanOrQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanOrQuery{}
}
func (q *SpanOrQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanOrQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanOrQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanOrQuery struct {
	Name string `json:"_name,omitempty"`
}
