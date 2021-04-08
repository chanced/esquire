package picker

type Commoner interface {
	Common() (*CommonQuery, error)
}
type CommonQueryParams struct {
	Name string
	completeClause
}

func (CommonQueryParams) Kind() QueryKind {
	return QueryKindCommon
}

func (p CommonQueryParams) Clause() (QueryClause, error) {
	return p.Common()
}
func (p CommonQueryParams) Common() (*CommonQuery, error) {
	q := &CommonQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type CommonQuery struct {
	nameParam
	completeClause
}

func (CommonQuery) Kind() QueryKind {
	return QueryKindCommon
}
func (q *CommonQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *CommonQuery) Common() (*CommonQuery, error) {
	return q, nil
}
func (q *CommonQuery) Clear() {
	if q == nil {
		return
	}
	*q = CommonQuery{}
}
func (q *CommonQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q CommonQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *CommonQuery) IsEmpty() bool {
	panic("not implemented")
}

type commonQuery struct {
	Name string `json:"_name,omitempty"`
}
