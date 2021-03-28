package search

import (
	"github.com/chanced/dynamic"
)

// WithName is a rule / query that has the _name parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html#named-queries
type WithName interface {
	Name() string
	SetName(string)
}

type withName interface {
	name() string
}

const DefaultName = ""

type nameParam struct {
	name string
}

func (n nameParam) Name() string {
	return n.name
}
func (n *nameParam) SetName(name string) {
	if n.Name() != name {
		n.name = name
	}
}
func unmarshalNameParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithName); ok {
		a.SetName(data.UnquotedString())
	}
	return nil
}

func marshalNameParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithName); ok {
		if b.Name() != DefaultName {
			data[paramName] = b.Name()
		}
	}
	return data, nil
}
