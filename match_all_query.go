package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// MatchAllQueryParams matches all documents, giving them all a _score of 1.0.
type MatchAllQueryParams struct {
	Boost interface{}
	Name  string
}

func (ma MatchAllQueryParams) Clause() (QueryClause, error) {
	return ma.MatchAll()

}

func (ma MatchAllQueryParams) Kind() QueryKind {
	return KindMatchAll
}

func (ma MatchAllQueryParams) MatchAll() (*MatchAllQuery, error) {
	c := &MatchAllQuery{}
	err := c.SetBoost(ma.Boost)
	if err != nil {
		return c, err
	}
	return c, nil
}

// MatchAllQuery matches all documents, giving them all a _score of 1.0.
type MatchAllQuery struct {
	boostParam
	disabled bool
	nameParam
	completeClause
}

func (ma *MatchAllQuery) Clause() (QueryClause, error) {
	return ma, nil
}
func (MatchAllQuery) Kind() QueryKind {
	return KindMatchAll
}

func (ma *MatchAllQuery) Clear() {
	ma.disabled = true
}

func (ma *MatchAllQuery) Enable() {
	if ma == nil {
		*ma = MatchAllQuery{}
	}
	ma.disabled = false
}
func (ma *MatchAllQuery) Disable() {
	if ma == nil {
		return
	}
	ma.disabled = true
}
func (ma *MatchAllQuery) IsEmpty() bool {
	return ma == nil || ma.disabled
}

func (ma *MatchAllQuery) UnmarshalJSON(data []byte) error {
	*ma = MatchAllQuery{}
	_, err := unmarshalClauseParams(data, ma)
	if err != nil {
		return err
	}
	return nil
}
func (ma MatchAllQuery) MarshalJSON() ([]byte, error) {
	if ma.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalClauseParams(ma)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}
