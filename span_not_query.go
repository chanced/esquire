package picker

type SpanNotQueryParams struct {
	Name string
	completeClause
}

func (SpanNotQueryParams) Kind() QueryKind {
	return QueryKindSpanNot
}

func (p SpanNotQueryParams) Clause() (QueryClause, error) {
	return p.SpanNot()
}
func (p SpanNotQueryParams) SpanNot() (*SpanNotQuery, error) {
	q := &SpanNotQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanNotQuery struct {
	nameParam
	completeClause
}

func (SpanNotQuery) Kind() QueryKind {
	return QueryKindSpanNot
}
func (q *SpanNotQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanNotQuery) SpanNot() (*SpanNotQuery, error) {
	return q, nil
}
func (q *SpanNotQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanNotQuery{}
}
func (q *SpanNotQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *SpanNotQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanNotQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q SpanNotQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanNotQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanNotQuery struct {
	Name string `json:"_name,omitempty"`
}
