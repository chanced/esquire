package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// MatchAll matches all documents, giving them all a _score of 1.0.
type MatchAll struct {
	Boost interface{}
	Name  string
}

func (ma MatchAll) Clause() (QueryClause, error) {
	return ma.MatchAll()

}

func (ma MatchAll) Kind() Kind {
	return KindMatchAll
}

func (ma MatchAll) MatchAll() (*MatchAllClause, error) {
	c := &MatchAllClause{}
	err := c.SetBoost(ma.Boost)
	if err != nil {
		return c, err
	}
	return c, nil
}

// MatchAllClause matches all documents, giving them all a _score of 1.0.
type MatchAllClause struct {
	boostParam
	disabled bool
	nameParam
}

func (MatchAllClause) Kind() Kind {
	return KindMatchAll
}

func (ma *MatchAllClause) Clear() {
	ma.disabled = true
}

func (ma *MatchAllClause) Enable() {
	if ma == nil {
		*ma = MatchAllClause{}
	}
	ma.disabled = false
}
func (ma *MatchAllClause) Disable() {
	if ma == nil {
		return
	}
	ma.disabled = true
}
func (ma *MatchAllClause) IsEmpty() bool {
	return ma == nil || ma.disabled
}

func (ma *MatchAllClause) UnmarshalJSON(data []byte) error {
	*ma = MatchAllClause{}
	_, err := unmarshalParams(data, ma)
	if err != nil {
		return err
	}
	return nil
}
func (ma MatchAllClause) MarshalJSON() ([]byte, error) {
	if ma.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := marshalParams(ma)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}
