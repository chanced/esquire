package picker

type Regexper interface {
	Regexp() (*RegexpQuery, error)
}

type RegexpQueryParams struct {
	Name string
	completeClause
}

func (RegexpQueryParams) Kind() QueryKind {
	return QueryKindRegexp
}

func (p RegexpQueryParams) Clause() (QueryClause, error) {
	return p.Regexp()
}
func (p RegexpQueryParams) Regexp() (*RegexpQuery, error) {
	q := &RegexpQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type RegexpQuery struct {
	nameParam
	completeClause
}

func (RegexpQuery) Kind() QueryKind {
	return QueryKindRegexp
}
func (q *RegexpQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *RegexpQuery) Regexp() (*RegexpQuery, error) {
	return q, nil
}
func (q *RegexpQuery) Clear() {
	if q == nil {
		return
	}
	*q = RegexpQuery{}
}
func (q *RegexpQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *RegexpQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q RegexpQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q RegexpQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *RegexpQuery) IsEmpty() bool {
	panic("not implemented")
}

type regexpQuery struct {
	Name string `json:"_name,omitempty"`
}
