package picker

type ShapeQueryParams struct {
	Name string
	completeClause
}

func (ShapeQueryParams) Kind() QueryKind {
	return QueryKindShape
}

func (p ShapeQueryParams) Clause() (QueryClause, error) {
	return p.Shape()
}
func (p ShapeQueryParams) Shape() (*ShapeQuery, error) {
	q := &ShapeQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type ShapeQuery struct {
	nameParam
	completeClause
}

func (ShapeQuery) Kind() QueryKind {
	return QueryKindShape
}
func (q *ShapeQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *ShapeQuery) Shape() (QueryClause, error) {
	return q, nil
}
func (q *ShapeQuery) Clear() {
	if q == nil {
		return
	}
	*q = ShapeQuery{}
}
func (q *ShapeQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q ShapeQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *ShapeQuery) IsEmpty() bool {
	panic("not implemented")
}

type shapeQuery struct {
	Name string `json:"_name,omitempty"`
}
