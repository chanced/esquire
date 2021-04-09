package picker

import "github.com/chanced/dynamic"

type NestedQueryParams struct {
	Name           string
	Query          Querier
	Path           string
	ScoreMode      ScoreMode
	IgnoreUnmapped interface{}
	completeClause
}

type NestedQuerier interface {
	Nested() (*NestedQuery, error)
}

func (NestedQueryParams) Kind() QueryKind {
	return QueryKindNested
}

func (p NestedQueryParams) Clause() (QueryClause, error) {
	return p.Nested()
}
func (p NestedQueryParams) Nested() (*NestedQuery, error) {
	q := &NestedQuery{}
	q.SetName(p.Name)
	err := q.SetPath(p.Path)
	if err != nil {
		return q, err
	}
	err = q.SetScoreMode(p.ScoreMode)
	if err != nil {
		return q, err
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	err = q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	if err != nil {
		return q, err
	}
	return q, nil
}

type NestedQuery struct {
	nameParam
	query *Query
	path  string
	scoreModeParam
	ignoreUnmapped dynamic.Bool
	completeClause
}

func (NestedQuery) Kind() QueryKind {
	return QueryKindNested
}
func (q *NestedQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *NestedQuery) Nested() (*NestedQuery, error) {
	return q, nil
}
func (q *NestedQuery) Clear() {
	if q == nil {
		return
	}
	*q = NestedQuery{}
}

func (q NestedQuery) IgnoreUnmapped() bool {
	if b, ok := q.ignoreUnmapped.Bool(); ok {
		return b
	}
	return false
}
func (q *NestedQuery) SetIgnoreUnmapped(v interface{}) error {
	return q.ignoreUnmapped.Set(v)
}
func (q NestedQuery) Path() string {
	return q.path
}
func (q *NestedQuery) SetPath(v string) error {
	if len(v) == 0 {
		return ErrPathRequired
	}
	q.path = v
	return nil
}
func (q NestedQuery) Query(query Querier) *Query {
	return q.query
}
func (q *NestedQuery) SetQuery(query Querier) error {
	qv, err := query.Query()
	if err != nil {
		return err
	}
	if qv.IsEmpty() {
		return ErrQueryRequired
	}
	q.query = qv
	return nil
}
func (q *NestedQuery) UnmarshalJSON(data []byte) error {
	*q = NestedQuery{}
	p := nestedQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	q.SetName(p.Name)
	err = q.SetPath(p.Path)
	if err != nil {
		return err
	}
	err = q.SetScoreMode(p.ScoreMode)
	if err != nil {
		return err
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return err
	}
	err = q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	if err != nil {
		return err
	}
	return nil
}
func (q NestedQuery) MarshalJSON() ([]byte, error) {
	return nestedQuery{
		Name:           q.name,
		Query:          q.query,
		Path:           q.path,
		ScoreMode:      q.scoreMode,
		IgnoreUnmapped: q.ignoreUnmapped.Value(),
	}.MarshalJSON()
}
func (q *NestedQuery) IsEmpty() bool {
	return q == nil || len(q.path) == 0 || q.query.IsEmpty()
}

//easyjson:json
type nestedQuery struct {
	Name           string      `json:"_name,omitempty"`
	Query          *Query      `json:"query"`
	Path           string      `json:"path"`
	ScoreMode      ScoreMode   `json:"score_mode,omitempty"`
	IgnoreUnmapped interface{} `json:"ignore_unmapped,omitempty"`
}
