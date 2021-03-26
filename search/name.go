package search

import (
	"github.com/chanced/dynamic"
)

type WithName interface {
	Name() string
	SetName(string)
}

const DefaultName = ""

// nameParam is a mixin that adds the _name parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html#named-queries
type nameParam struct {
	name string `json:"_name,omitempty" bson:"_name,omitempty"`
}

func (n nameParam) Name() string {
	return n.name
}
func (n *nameParam) SetName(name string) {
	if n.Name() != name {
		n.name = name
	}
}
func unmarshalNameParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithName); ok {
		a.SetName(data.UnquotedString())
	}
	return nil
}

func marshalNameParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithName); ok {
		if b.Name() != DefaultName {
			data[paramName] = b.Name()
		}
	}
	return data, nil
}
