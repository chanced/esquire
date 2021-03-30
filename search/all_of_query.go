package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type AllOfer interface {
	AllOf() (*AllOfClause, error)
}

type AllOf struct {
	Name string
}

func (ao AllOf) Kind() Kind {
	return KindAllOf
}

func (ao AllOf) AllOf() (*AllOfClause, error) {
	c := &AllOfClause{}
	c.SetName(ao.Name)
	return c, nil
}

func (ao AllOf) Clause() (CompleteClause, error) {
	return ao.AllOf()
}

type AllOfClause struct {
	nameParam
	clause
	cleared bool // not great
}

func (AllOfClause) Kind() Kind {
	return KindAllOf
}

func (ao AllOfClause) MarshalJSON() ([]byte, error) {
	if ao.IsEmpty() {
		return dynamic.Null, nil
	}
	p, err := marshalParams(&ao)
	if err != nil {
		return nil, err
	}
	return json.Marshal(p)
}
func (ao *AllOfClause) UnmarshalJSON(data []byte) error {
	*ao = AllOfClause{}
	_, err := unmarshalParams(data, ao)
	if err != nil {
		return err
	}
	return nil
}

func (ao *AllOfClause) Clear() {
	ao.cleared = true
}

func (ao *AllOfClause) Enable() {
	ao.cleared = false
}
func (ao *AllOfClause) Disable() {
	ao.cleared = true
}
func (ao *AllOfClause) IsEmpty() bool {
	return ao == nil || ao.cleared
}
