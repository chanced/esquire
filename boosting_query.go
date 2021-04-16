package picker

import "encoding/json"

type Boostinger interface {
	Boosting() (*BoostingQuery, error)
}

// BoostingQuery returns documents matching a positive query while reducing the
// relevance score of documents that also match a negative query.
//
// You can use the boosting query to demote certain documents without excluding
// them from the search results.
type BoostingQuery struct {
	positive *Query
	negative *Query
	negativeBoostParam
	nameParam
	completeClause
}

func (b *BoostingQuery) Clause() (QueryClause, error) {
	return b, nil
}

func (BoostingQuery) Kind() QueryKind {
	return QueryKindBoosting
}

func (b *BoostingQuery) Clear() {
	*b = BoostingQuery{}
}
func (b *BoostingQuery) IsEmpty() bool {
	return b == nil || b.negative.IsEmpty() || b.positive.IsEmpty()
}

func (b BoostingQuery) Positive() *Query {
	if b.positive == nil {
		b.positive = &Query{}
	}
	return b.positive
}
func (b BoostingQuery) Negative() *Query {
	return b.negative
}

func (b *BoostingQuery) SetNegative(params Querier) error {
	if params == nil {
		return ErrNegativeRequired
	}
	q, err := params.Query()
	if err != nil {
		return err
	}
	b.negative = q
	return nil
}

func (b *BoostingQuery) SetPositive(params Querier) error {
	if params == nil {
		return ErrPositiveRequired
	}
	q, err := params.Query()
	if err != nil {
		return err
	}
	b.positive = q
	return nil
}

func (b BoostingQuery) MarshalBSON() ([]byte, error) {
	return b.MarshalJSON()
}

func (b BoostingQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(boostingQuery{
		Negative:      b.negative,
		Positive:      b.positive,
		NegativeBoost: b.negativeBoost,
	})
}

func (b *BoostingQuery) UnmarshalBSON(data []byte) error {
	return b.UnmarshalJSON(data)
}

func (b *BoostingQuery) UnmarshalJSON(data []byte) error {
	*b = BoostingQuery{}
	var bq boostingQuery

	err := json.Unmarshal(data, &bq)
	if err != nil {
		return err
	}
	b.negative = bq.Negative
	b.positive = bq.Positive
	b.negativeBoost = bq.NegativeBoost
	return nil
}

type boostingQuery struct {
	Negative      *Query  `json:"negative"`
	Positive      *Query  `json:"positive"`
	NegativeBoost float64 `json:"negative_boost"`
	Name          string  `json:"_name,omitempty"`
}

type BoostingQueryParams struct {
	// (Required, query object) Query you wish to run. Any returned documents
	// must match this query.
	Positive Querier
	// (Required, query object) Query used to decrease the relevance score of
	// matching documents.
	//
	// If a returned document matches the positive query and this query, the
	// boosting query calculates the final relevance score for the document as
	// follows:
	//
	// 1. Take the original relevance score from the positive query. 2. Multiply
	// the score by the negative_boost value.
	Negative Querier

	// (Required, float) Floating point number between 0 and 1.0 used to
	// decrease the relevance scores of documents matching the negative query.
	NegativeBoost float64

	Name string
}

func (p BoostingQueryParams) Clause() (QueryClause, error) {
	return p.Boosting()
}

func (p BoostingQueryParams) Boosting() (*BoostingQuery, error) {
	q := &BoostingQuery{}
	q.SetName(p.Name)
	err := q.SetNegative(p.Negative)
	if err != nil {
		return q, err
	}
	err = q.SetPositive(p.Positive)
	if err != nil {
		return q, err
	}
	err = q.SetNegativeBoost(p.NegativeBoost)
	if err != nil {
		return q, err
	}
	return q, nil
}
