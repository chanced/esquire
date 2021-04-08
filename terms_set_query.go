package picker

type TermSetter interface {
	TermSet() (*TermSetQuery, error)
}

type TermSetQueryParams struct {
	Name string
	completeClause
}

func (TermSetQueryParams) Kind() QueryKind {
	return QueryKindTermSet
}

func (p TermSetQueryParams) Clause() (QueryClause, error) {
	return p.TermSet()
}
func (p TermSetQueryParams) TermSet() (*TermSetQuery, error) {
	q := &TermSetQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type TermSetQuery struct {
	nameParam
	completeClause
}

func (TermSetQuery) Kind() QueryKind {
	return QueryKindTermSet
}
func (q *TermSetQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *TermSetQuery) TermSet() (*TermSetQuery, error) {
	return q, nil
}
func (q *TermSetQuery) Clear() {
	if q == nil {
		return
	}
	*q = TermSetQuery{}
}
func (q *TermSetQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q TermSetQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *TermSetQuery) IsEmpty() bool {
	panic("not implemented")
}

type termSetQuery struct {
	Name string `json:"_name,omitempty"`
}
