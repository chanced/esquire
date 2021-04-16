package picker

type SpanNearQueryParams struct {
	Name string
	completeClause
}

func (SpanNearQueryParams) Kind() QueryKind {
	return QueryKindSpanNear
}

func (p SpanNearQueryParams) Clause() (QueryClause, error) {
	return p.SpanNear()
}
func (p SpanNearQueryParams) SpanNear() (*SpanNearQuery, error) {
	q := &SpanNearQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanNearQuery struct {
	nameParam
	completeClause
}

func (SpanNearQuery) Kind() QueryKind {
	return QueryKindSpanNear
}
func (q *SpanNearQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanNearQuery) SpanNear() (*SpanNearQuery, error) {
	return q, nil
}
func (q *SpanNearQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanNearQuery{}
}
func (q *SpanNearQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *SpanNearQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanNearQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q SpanNearQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanNearQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanNearQuery struct {
	Name string `json:"_name,omitempty"`
}
