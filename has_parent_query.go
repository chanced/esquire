package picker

import "github.com/chanced/dynamic"

type HasParenter interface {
	HasParent() (*HasParentQuery, error)
}

type HasParentQueryParams struct {
	ParentType string
	// (Required, query object) Query you wish to run on parent documents of the
	// parent_type field. If a parent document matches the search, the query
	// returns its child documents.
	Query Querier
	// (Optional, Boolean) Indicates whether to ignore an unmapped parent_type
	// and not return any documents instead of an error. Defaults to false.
	//
	// If false, Elasticsearch returns an error if the parent_type is unmapped.
	//
	// You can use this parameter to query multiple indices that may not contain
	// the parent_type.
	IgnoreUnmapped interface{}
	// (Optional, Boolean) Indicates whether the relevance score of a matching
	// parent document is aggregated into its child documents. Defaults to
	// false.
	//
	// If false, Elasticsearch ignores the relevance score of the parent
	// document. Elasticsearch also assigns each child document a relevance
	// score equal to the query's boost, which defaults to 1.
	//
	// If true, the relevance score of the matching parent document is
	// aggregated into its child documents' relevance scores.
	Score interface{}
	Name  string
	completeClause
}

func (HasParentQueryParams) Kind() QueryKind {
	return QueryKindHasParent
}

func (p HasParentQueryParams) Clause() (QueryClause, error) {
	return p.HasParent()
}
func (p HasParentQueryParams) HasParent() (*HasParentQuery, error) {
	q := &HasParentQuery{}

	err := q.SetParentType(p.ParentType)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	err = q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	err = q.SetScore(p.Score)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	q.SetName(p.Name)
	return q, nil
}

type HasParentQuery struct {
	nameParam
	score          dynamic.Bool
	query          *Query
	parentTyp      string
	ignoreUnmapped dynamic.Bool
	completeClause
}

func (q HasParentQuery) Score() bool {
	if b, ok := q.score.Bool(); ok {
		return b
	}
	return false // TODO: figure out if this is the default. so far as i can tell, it is.
}

func (q *HasParentQuery) SetScore(score interface{}) error {
	err := q.score.Set(score)
	if err != nil {
		return newQueryError(ErrTypeRequired, QueryKindHasParent)
	}
	return nil
}
func (q HasParentQuery) ParentType() string {
	return q.parentTyp
}
func (q *HasParentQuery) SetParentType(typ string) error {
	if len(typ) == 0 {
		return newQueryError(ErrTypeRequired, QueryKindHasParent)
	}
	q.parentTyp = typ
	return nil
}
func (q HasParentQuery) IgnoreUnmapped() bool {
	if b, ok := q.ignoreUnmapped.Bool(); ok {
		return b
	}
	return false
}
func (q *HasParentQuery) SetIgnoreUnmapped(ignoreUnmapped interface{}) error {
	err := q.ignoreUnmapped.Set(ignoreUnmapped)
	if err != nil {
		return newQueryError(err, QueryKindHasParent)
	}
	return nil
}
func (q HasParentQuery) Query() *Query {
	return q.query
}
func (q *HasParentQuery) SetQuery(query Querier) error {
	qv, err := query.Query()
	if err != nil {
		return newQueryError(err, QueryKindHasParent)
	}
	q.query = qv
	return nil
}

func (HasParentQuery) Kind() QueryKind {
	return QueryKindHasParent
}
func (q *HasParentQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *HasParentQuery) HasParent() (*HasParentQuery, error) {
	return q, nil
}
func (q *HasParentQuery) Clear() {
	if q == nil {
		return
	}
	*q = HasParentQuery{}
}

func (q *HasParentQuery) UnmarshalJSON(data []byte) error {
	*q = HasParentQuery{}
	qv := hasParentQuery{}
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
func (q HasParentQuery) MarshalJSON() ([]byte, error) {
	return hasParentQuery{
		ParentType:     q.parentTyp,
		Query:          q.query,
		Score:          q.score.Value(),
		IgnoreUnmapped: q.ignoreUnmapped.Value(),
		Name:           q.name,
	}.MarshalJSON()
}

func (q HasParentQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}
func (q *HasParentQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}
func (q *HasParentQuery) IsEmpty() bool {
	return q == nil || len(q.parentTyp) == 0
}

//easyjson:json
type hasParentQuery struct {
	ParentType     string      `json:"parent_type"`
	Query          *Query      `json:"query"`
	IgnoreUnmapped interface{} `json:"ignore_unmapped,omitempty"`
	Score          interface{} `json:"score,omitempty"`
	Name           string      `json:"_name,omitempty"`
}

func (p hasParentQuery) data() (*HasParentQuery, error) {
	q := &HasParentQuery{}

	err := q.SetParentType(p.ParentType)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	err = q.SetQuery(p.Query)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	err = q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	err = q.SetScore(p.Score)
	if err != nil {
		return q, newQueryError(err, QueryKindHasParent)
	}
	q.SetName(p.Name)
	return q, nil
}
