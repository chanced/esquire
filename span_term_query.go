package picker

type SpanTermQueryParams struct {
	Name string
	completeClause
}

func (SpanTermQueryParams) Kind() QueryKind {
	return QueryKindSpanTerm
}

func (p SpanTermQueryParams) Clause() (QueryClause, error) {
	return p.SpanTerm()
}
func (p SpanTermQueryParams) SpanTerm() (*SpanTermQuery, error) {
	q := &SpanTermQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type SpanTermQuery struct {
	nameParam
	completeClause
}

func (SpanTermQuery) Kind() QueryKind {
	return QueryKindSpanTerm
}
func (q *SpanTermQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *SpanTermQuery) SpanTerm() (*SpanTermQuery, error) {
	return q, nil
}
func (q *SpanTermQuery) Clear() {
	if q == nil {
		return
	}
	*q = SpanTermQuery{}
}
func (q *SpanTermQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q SpanTermQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *SpanTermQuery) IsEmpty() bool {
	panic("not implemented")
}

type spanTermQuery struct {
	Name string `json:"_name,omitempty"`
}
