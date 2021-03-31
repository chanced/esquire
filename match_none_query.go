package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// MatchNoneQuery is the inverse of the match_all query, which matches no documents.
type MatchNoneQuery struct {
	Name string
}

func (mn MatchNoneQuery) Clause() (QueryClause, error) {
	return mn.MatchNone()

}

func (mn MatchNoneQuery) Kind() Kind {
	return KindMatchNone
}
func (mn MatchNoneQuery) MatchNone() (*MatchNoneClause, error) {
	c := &MatchNoneClause{}
	return c, nil
}

// MatchNoneClause is the inverse of the match_all query, which matches no documents.
type MatchNoneClause struct {
	boostParam
	disabled bool
	nameParam
	completeClause
}

func (mn *MatchNoneClause) Clause() (QueryClause, error) {
	return mn, nil
}
func (MatchNoneClause) Kind() Kind {
	return KindMatchNone
}

func (mn *MatchNoneClause) Clear() {
	*mn = MatchNoneClause{
		disabled: true,
	}
}

func (mn *MatchNoneClause) Enable() {
	if mn == nil {
		*mn = MatchNoneClause{}
	}
	mn.disabled = false
}
func (mn *MatchNoneClause) Disable() {
	if mn == nil {
		return
	}
	mn.disabled = true
}
func (mn *MatchNoneClause) IsEmpty() bool {
	return mn == nil || mn.disabled
}

func (mn *MatchNoneClause) UnmarshalJSON(data []byte) error {
	*mn = MatchNoneClause{}
	_, err := unmarshalClauseParams(data, mn)
	if err != nil {
		return err
	}
	return nil
}
func (mn MatchNoneClause) MarshalJSON() ([]byte, error) {
	if mn.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalClauseParams(mn)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}
