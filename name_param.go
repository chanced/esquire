package picker

import (
	"encoding/json"

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

func marshalNameParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithName); ok {
		if len(b.Name()) > 0 {
			json.Marshal(b.Name())
		}
	}
	return nil, nil
}
