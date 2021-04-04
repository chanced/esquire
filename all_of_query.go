package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type AllOfer interface {
	AllOf() (*AllOfQuery, error)
}

type AllOfQueryParams struct {
	Name string
}

func (ao AllOfQueryParams) Kind() QueryKind {
	return QueryKindAllOf
}

func (ao AllOfQueryParams) AllOf() (*AllOfQuery, error) {
	c := &AllOfQuery{}
	c.SetName(ao.Name)
	return c, nil
}

func (ao AllOfQueryParams) Clause() (CompleteClause, error) {
	return ao.AllOf()
}

type AllOfQuery struct {
	nameParam
	completeClause
	cleared bool // not great
}

func (ao *AllOfQuery) Clause() (QueryClause, error) {
	return ao, nil
}

func (AllOfQuery) Kind() QueryKind {
	return QueryKindAllOf
}

func (ao AllOfQuery) MarshalJSON() ([]byte, error) {
	if ao.IsEmpty() {
		return dynamic.Null, nil
	}
	p, err := marshalClauseParams(&ao)
	if err != nil {
		return nil, err
	}
	return json.Marshal(p)
}
func (ao *AllOfQuery) UnmarshalJSON(data []byte) error {
	*ao = AllOfQuery{}
	_, err := unmarshalClauseParams(data, ao)
	if err != nil {
		return err
	}
	return nil
}

func (ao *AllOfQuery) Clear() {
	ao.cleared = true
}

func (ao *AllOfQuery) Enable() {
	ao.cleared = false
}
func (ao *AllOfQuery) Disable() {
	ao.cleared = true
}
func (ao *AllOfQuery) IsEmpty() bool {
	return ao == nil || ao.cleared
}
