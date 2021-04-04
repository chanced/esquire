package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// MatchNoneQueryParams is the inverse of the match_all query, which matches no documents.
type MatchNoneQueryParams struct {
	Name string
}

func (mn MatchNoneQueryParams) Clause() (QueryClause, error) {
	return mn.MatchNone()

}

func (mn MatchNoneQueryParams) Kind() QueryKind {
	return QueryKindMatchNone
}
func (mn MatchNoneQueryParams) MatchNone() (*MatchNoneQuery, error) {
	c := &MatchNoneQuery{}
	return c, nil
}

// MatchNoneQuery is the inverse of the match_all query, which matches no documents.
type MatchNoneQuery struct {
	boostParam
	disabled bool
	nameParam
	completeClause
}

func (mn *MatchNoneQuery) Clause() (QueryClause, error) {
	return mn, nil
}
func (MatchNoneQuery) Kind() QueryKind {
	return QueryKindMatchNone
}

func (mn *MatchNoneQuery) Clear() {
	*mn = MatchNoneQuery{
		disabled: true,
	}
}

func (mn *MatchNoneQuery) Enable() {
	if mn == nil {
		*mn = MatchNoneQuery{}
	}
	mn.disabled = false
}
func (mn *MatchNoneQuery) Disable() {
	if mn == nil {
		return
	}
	mn.disabled = true
}
func (mn *MatchNoneQuery) IsEmpty() bool {
	return mn == nil || mn.disabled
}

func (mn *MatchNoneQuery) UnmarshalJSON(data []byte) error {
	*mn = MatchNoneQuery{}
	_, err := unmarshalClauseParams(data, mn)
	if err != nil {
		return err
	}
	return nil
}
func (mn MatchNoneQuery) MarshalJSON() ([]byte, error) {
	if mn.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalClauseParams(mn)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}
