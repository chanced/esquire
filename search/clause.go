package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Clause interface {
	Type() Type
	Name() string
	SetName(name string)
}

type Fielder interface {
	FieldName() string
}

type Clauser interface {
	Clause
	Clause() (Clause, error)
}

// type Rules []Clause

// func (r *Rules) Add(rule Clause) error {
// 	var err error
// 	if v, ok := rule.(Clauser); ok {
// 		rule, err = v.Rule()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	*r = append(*r, rule)
// 	return nil
// }

func marshalClauseParams(source Clause) (dynamic.Map, error) {
	return marshalParams(source)
}

func unmarshalParams(data []byte, target Clause) (map[string]dynamic.RawJSON, error) {
	var raw map[string]dynamic.RawJSON
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}
	res := map[string]dynamic.RawJSON{}
	for key, value := range raw {
		isParam, err := unmarshalParam(key, value, target)
		if err != nil {
			return nil, err
		}
		if !isParam {
			res[key] = value
		}
	}
	return res, nil
}
