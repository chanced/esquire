package picker

import "encoding/json"

type ConstantScorer interface {
	ConstantScore() (*ConstantScoreQuery, error)
}

type constantScoreQuery struct {
	Filter *Query      `json:"filter"`
	Boost  interface{} `json:"boost,omitempty"`
	Name   string      `json:"_name,omitempty"`
}
type ConstantScoreQueryParams struct {
	// (Required) Filter query you wish to run. Any returned
	// documents must match this query.
	//
	// Filter queries do not calculate relevance scores. To speed up
	// performance, Elasticsearch automatically caches frequently used filter
	// queries.
	Filter Querier
	// (Optional, float) Floating point number used as the constant relevance score for every document matching the filter query. Defaults to 1.0.
	Boost interface{}
}

func (ConstantScoreQueryParams) Kind() QueryKind {
	return QueryKindConstantScore
}
func (p ConstantScoreQueryParams) Clause() (QueryClause, error) {
	return p.ConstantScore()
}
func (p ConstantScoreQueryParams) ConstantScore() (*ConstantScoreQuery, error) {
	q := &ConstantScoreQuery{}
	err := q.SetFilter(p.Filter)
	if err != nil {
		return q, err
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, err
	}
	return q, nil
}

type ConstantScoreQuery struct {
	filter *Query
	nameParam
	boostParam
	completeClause
}

func (ConstantScoreQuery) Kind() QueryKind {
	return QueryKindConstantScore
}
func (cs ConstantScoreQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(constantScoreQuery{
		Filter: cs.filter,
		Boost:  cs.boost.Value(),
		Name:   cs.Name(),
	})
}

func (cs *ConstantScoreQuery) Filter() *Query {
	if cs.filter == nil {
		cs.filter = &Query{}
	}
	return cs.filter
}
func (cs *ConstantScoreQuery) SetFilter(query Querier) error {
	if cs.filter == nil {
		cs.filter = &Query{}
	}
	return cs.filter.Set(query)
}
func (cs *ConstantScoreQuery) UnmarshalJSON(data []byte) error {
	*cs = ConstantScoreQuery{}
	var v constantScoreQuery
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	cs.filter = v.Filter
	err = cs.boost.Set(v.Boost)
	if err != nil {
		return err
	}
	return nil
}
func (cs *ConstantScoreQuery) Clause() (QueryClause, error) {
	return cs, nil
}
func (cs *ConstantScoreQuery) ConstantScore() (*ConstantScoreQuery, error) {
	return cs, nil
}
func (cs *ConstantScoreQuery) Clear() {
	if cs == nil {
		return
	}
	*cs = ConstantScoreQuery{}
}

func (cs *ConstantScoreQuery) IsEmpty() bool {
	return cs == nil || cs.filter.IsEmpty()
}
