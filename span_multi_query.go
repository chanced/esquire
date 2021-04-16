package picker

type SpanMultiQueryParams struct {
	Name string
	completeClause
}

func (SpanMultiQueryParams) Kind() QueryKind {
	return QueryKindSpanMulti
}

func (p SpanMultiQueryParams) Clause() (QueryClause, error) {
	return p.SpanMulti()
}
func (p SpanMultiQueryParams) SpanMulti() (*SpanMultiQuery, error) {
	q := &SpanMultiQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanMultiQuery struct {
	nameParam
	completeClause
}

func (SpanMultiQuery) Kind() QueryKind {
	return QueryKindSpanMulti
}
func (q *SpanMultiQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanMultiQuery) SpanMulti() (*SpanMultiQuery, error) {
	return q, nil
}
func (q *SpanMultiQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanMultiQuery{}
}
func (q *SpanMultiQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *SpanMultiQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanMultiQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q SpanMultiQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanMultiQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanMultiQuery struct {
	Name string `json:"_name,omitempty"`
}
