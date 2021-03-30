package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// MatchNone is the inverse of the match_all query, which matches no documents.
type MatchNone struct {
	Name string
}

func (ma MatchNone) Clause() (QueryClause, error) {
	return ma.MatchNone()

}

func (ma MatchNone) Kind() Kind {
	return KindMatchNone
}
func (ma MatchNone) MatchNone() (*MatchNoneClause, error) {
	c := &MatchNoneClause{}
	return c, nil
}

// MatchNoneClause is the inverse of the match_all query, which matches no documents.
type MatchNoneClause struct {
	boostParam
	disabled bool
	nameParam
}

func (MatchNoneClause) Kind() Kind {
	return KindMatchNone
}

func (ma *MatchNoneClause) Clear() {
	ma.disabled = true
}

func (ma *MatchNoneClause) Enable() {
	if ma == nil {
		*ma = MatchNoneClause{}
	}
	ma.disabled = false
}
func (ma *MatchNoneClause) Disable() {
	if ma == nil {
		return
	}
	ma.disabled = true
}
func (ma *MatchNoneClause) IsEmpty() bool {
	return ma == nil || ma.disabled
}

func (ma *MatchNoneClause) UnmarshalJSON(data []byte) error {
	*ma = MatchNoneClause{}
	_, err := unmarshalParams(data, ma)
	if err != nil {
		return err
	}
	return nil
}
func (ma MatchNoneClause) MarshalJSON() ([]byte, error) {
	if ma.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalParams(ma)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}
