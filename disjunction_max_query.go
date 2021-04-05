package picker

import "encoding/json"

type DisjunctionMaxer interface {
	DisjunctionMax() (*DisjunctionMaxQuery, error)
}

type DisjunctionMaxQueryParams struct {
	Queries    Queriers
	TieBreaker float64
	Name       string
}

func (DisjunctionMaxQueryParams) Kind() QueryKind {
	return QueryKindDisjunctionMax
}
func (p DisjunctionMaxQueryParams) Clause() (QueryClause, error) {
	return p.DisjunctionMax()
}
func (p DisjunctionMaxQueryParams) DisjunctionMax() (*DisjunctionMaxQuery, error) {
	q := &DisjunctionMaxQuery{}
	err := q.SetQueries(p.Queries)
	if err != nil {
		return q, err
	}
	err = q.SetTieBreaker(p.TieBreaker)
	if err != nil {
		return q, err
	}
	return q, nil
}

type DisjunctionMaxQuery struct {
	queries    Queries
	tieBreaker float64
	nameParam
	completeClause
}

func (dm DisjunctionMaxQuery) Queries() Queries {
	return dm.queries
}
func (dm DisjunctionMaxQuery) TieBreaker() float64 {
	return dm.tieBreaker
}
func (dm *DisjunctionMaxQuery) SetTieBreaker(v float64) error {
	dm.tieBreaker = v
	// incase I need to validate it in the future
	return nil
}

func (dm *DisjunctionMaxQuery) SetQueries(queries Queryset) error {
	if queries == nil {
		return ErrQueriesRequired
	}
	q, err := queries.Queries()
	if err != nil {
		return err
	}
	if q.IsEmpty() {
		return ErrQueriesRequired
	}
	dm.queries = q
	return nil
}
func (DisjunctionMaxQuery) Kind() QueryKind {
	return QueryKindDisjunctionMax
}

func (dm *DisjunctionMaxQuery) Clear() {
	*dm = DisjunctionMaxQuery{}
}
func (dm *DisjunctionMaxQuery) IsEmpty() bool {
	return dm == nil || dm.queries.IsEmpty()
}

func (dm *DisjunctionMaxQuery) Clause() (QueryClause, error) {
	return dm, nil
}
func (dm *DisjunctionMaxQuery) DisjunectionMax() (*DisjunctionMaxQuery, error) {
	return dm, nil
}

func (dm DisjunctionMaxQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(disjunctionMaxQuery{
		Queries:    dm.queries,
		TieBreaker: dm.tieBreaker,
		Name:       dm.name,
	})
}

func (dm *DisjunctionMaxQuery) UnmarshalJSON(data []byte) error {
	*dm = DisjunctionMaxQuery{}
	var v disjunctionMaxQuery
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	dm.queries = v.Queries
	dm.tieBreaker = v.TieBreaker
	dm.name = v.Name
	return nil
}

type disjunctionMaxQuery struct {
	Queries    Queries `json:"queries"`
	TieBreaker float64 `json:"tie_breaker,omitempty"`
	Name       string  `json:"_name,omitempty"`
}
