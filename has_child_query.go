package picker

import "github.com/chanced/dynamic"

type HasChilder interface {
	HasChild() (*HasChildQuery, error)
}

type HasChildQueryParams struct {
	// (Required, string) Name of the child relationship mapped for the join
	// field.
	Type string
	// (Required, query object) Query you wish to run on child documents of the
	// type field. If a child document matches the search, the query returns the
	// parent document.
	Query Querier
	// (Optional, Boolean) Indicates whether to ignore an unmapped type and not
	// return any documents instead of an error. Defaults to false. If false,
	// Elasticsearch returns an error if the type is unmapped.
	//
	// You can use this parameter to query multiple indices that may not contain
	// the type.
	IgnoreUnmapped interface{}
	// (Optional, integer) Maximum number of child documents that match the
	// query allowed for a returned parent document. If the parent document
	// exceeds this limit, it is excluded from the search results.
	MaxChildren interface{}
	// (Optional, integer) Minimum number of child documents that match the
	// query required to match the query for a returned parent document. If the
	// parent document does not meet this limit, it is excluded from the search
	// results.
	MinChildren interface{}
	// (Optional, string) Indicates how scores for matching child documents affect the root parent document’s relevance score. Valid values are:
	// none (Default)
	// Do not use the relevance scores of matching child documents. The query assigns parent documents a score of 0.
	// avg
	// Use the mean relevance score of all matching child documents.
	// max
	// Uses the highest relevance score of all matching child documents.
	// min
	// Uses the lowest relevance score of all matching child documents.
	// sum
	// Add together the relevance scores of all matching child documents.
	ScoreMode ScoreMode
	Name      string
	completeClause
}

func (HasChildQueryParams) Kind() QueryKind {
	return QueryKindHasChild
}

func (p HasChildQueryParams) Clause() (QueryClause, error) {
	return p.HasChild()
}
func (p HasChildQueryParams) HasChild() (*HasChildQuery, error) {
	q := &HasChildQuery{}
	err := q.SetType(p.Type)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetMaxChildren(p.MaxChildren)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetMinChildren(p.MinChildren)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetScoreMode(p.ScoreMode)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	q.SetName(p.Name)
	return q, nil
}

type HasChildQuery struct {
	nameParam
	scoreMode      ScoreMode
	query          *Query
	typ            string
	ignoreUnmapped dynamic.Bool
	maxChildren    dynamic.Number
	minChildren    dynamic.Number
	completeClause
}

func (q HasChildQuery) ScoreMode() ScoreMode {
	return q.scoreMode
}

func (q *HasChildQuery) SetScoreMode(sm ScoreMode) error {
	// skipping validation for now
	q.scoreMode = sm
	return nil
}
func (q HasChildQuery) Type() string {
	return q.typ
}
func (q *HasChildQuery) SetType(typ string) error {
	if len(typ) == 0 {
		return newQueryError(ErrTypeRequired, QueryKindHasChild)
	}
	q.typ = typ
	return nil
}
func (q HasChildQuery) IgnoreUnmapped() bool {
	if b, ok := q.ignoreUnmapped.Bool(); ok {
		return b
	}
	return false
}
func (q *HasChildQuery) SetIgnoreUnmapped(ignoreUnmapped interface{}) error {
	err := q.ignoreUnmapped.Set(ignoreUnmapped)
	if err != nil {
		return newQueryError(err, QueryKindHasChild)
	}
	return nil
}
func (q *HasChildQuery) MaxChildren() int {
	if i, ok := q.maxChildren.Int(); ok {
		return i
	}
	if f, ok := q.maxChildren.Float64(); ok {
		_ = q.maxChildren.Set(int(f))
		return int(f)
	}
	return 0 // or should it be max int? ES used to regard 0 as having no upper bounds..?
}
func (q *HasChildQuery) SetMaxChildren(max interface{}) error {
	err := q.maxChildren.Set(max)
	if err != nil {
		return newQueryError(err, QueryKindHasChild)
	}
	if max == nil {
		return nil
	}
	if _, ok := q.maxChildren.Int(); ok {
		return nil
	}
	if f, ok := q.maxChildren.Float64(); ok {
		_ = q.maxChildren.Set(int(f))
		return nil
	}
	return nil
}
func (q *HasChildQuery) MinChildren() int {
	if i, ok := q.minChildren.Int(); ok {
		return i
	}
	if f, ok := q.minChildren.Float64(); ok {
		_ = q.minChildren.Set(int(f))
		return int(f)
	}
	return 0 // or should it be max int? ES used to regard 0 as having no upper bounds..?
}
func (q *HasChildQuery) SetMinChildren(min interface{}) error {
	err := q.minChildren.Set(min)
	if err != nil {
		return newQueryError(err, QueryKindHasChild)
	}
	if min == nil {
		return nil
	}
	if _, ok := q.minChildren.Int(); ok {
		return nil
	}
	if f, ok := q.minChildren.Float64(); ok {
		_ = q.minChildren.Set(int(f))
		return nil
	}
	return nil
}

func (q HasChildQuery) Query() *Query {
	return q.query
}
func (q *HasChildQuery) SetQuery(query Querier) error {
	qv, err := query.Query()
	if err != nil {
		return newQueryError(err, QueryKindHasChild)
	}
	q.query = qv
	return nil
}

func (HasChildQuery) Kind() QueryKind {
	return QueryKindHasChild
}
func (q *HasChildQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *HasChildQuery) HasChild() (*HasChildQuery, error) {
	return q, nil
}
func (q *HasChildQuery) Clear() {
	if q == nil {
		return
	}
	*q = HasChildQuery{}
}

func (q *HasChildQuery) UnmarshalJSON(data []byte) error {
	*q = HasChildQuery{}
	qv := hasChildQuery{}
	err := qv.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	nq, err := qv.data()
	if err != nil {
		return err
	}
	*q = *nq
	return nil
}
func (q HasChildQuery) MarshalJSON() ([]byte, error) {
	return hasChildQuery{
		Type:           q.typ,
		Query:          q.query,
		IgnoreUnmapped: q.ignoreUnmapped.Value(),
		MaxChildren:    q.maxChildren.Value(),
		MinChildren:    q.minChildren.Value(),
		ScoreMode:      q.scoreMode,
		Name:           q.name,
	}.MarshalJSON()
}

func (q HasChildQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}
func (q *HasChildQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}
func (q *HasChildQuery) IsEmpty() bool {
	return q == nil || len(q.typ) == 0
}

//easyjson:json
type hasChildQuery struct {
	Type           string      `json:"type,omitempty"`
	Query          *Query      `json:"query,omitempty"`
	IgnoreUnmapped interface{} `json:"ignore_unmapped,omitempty"`
	MaxChildren    interface{} `json:"max_children,omitempty"`
	MinChildren    interface{} `json:"min_children,omitempty"`
	ScoreMode      ScoreMode   `json:"score_mode,omitempty"`
	Name           string      `json:"_name,omitempty"`
}

func (p hasChildQuery) data() (*HasChildQuery, error) {
	q := &HasChildQuery{}

	err := q.SetType(p.Type)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetMaxChildren(p.MaxChildren)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetMinChildren(p.MinChildren)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	err = q.SetScoreMode(p.ScoreMode)
	if err != nil {
		return q, newQueryError(err, QueryKindHasChild)
	}
	q.SetName(p.Name)
	return q, nil
}
