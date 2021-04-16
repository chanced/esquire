package picker

type FieldMaskingSpanQueryParams struct {
	Name string
	completeClause
}

func (FieldMaskingSpanQueryParams) Kind() QueryKind {
	return QueryKindFieldMaskingSpan
}

func (p FieldMaskingSpanQueryParams) Clause() (QueryClause, error) {
	return p.FieldMaskingSpan()
}
func (p FieldMaskingSpanQueryParams) FieldMaskingSpan() (*FieldMaskingSpanQuery, error) {
	q := &FieldMaskingSpanQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type FieldMaskingSpanQuery struct {
	nameParam
	completeClause
}

func (FieldMaskingSpanQuery) Kind() QueryKind {
	return QueryKindFieldMaskingSpan
}
func (q *FieldMaskingSpanQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *FieldMaskingSpanQuery) FieldMaskingSpan() (*FieldMaskingSpanQuery, error) {
	return q, nil
}
func (q *FieldMaskingSpanQuery) Clear() {
	if q == nil {
		return
	}
	*q = FieldMaskingSpanQuery{}
}
func (q *FieldMaskingSpanQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *FieldMaskingSpanQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q FieldMaskingSpanQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q FieldMaskingSpanQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *FieldMaskingSpanQuery) IsEmpty() bool {
	panic("not implemented")
}

type fieldMaskingSpanQuery struct {
	Name string `json:"_name,omitempty"`
}
