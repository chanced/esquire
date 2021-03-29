package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Clause interface {
	Kind() Kind
}

type clause interface {
	Name() string
	Clause()
}

type WithField interface {
	Field() string
}

type withfield interface {
	field() string
}

type clauser interface {
	Clause
	Clause() (Clause, error)
}

// type Rules []Clause

// func (r *Rules) Add(rule Clause) error {
//     var err error
//     if v, ok := rule.(Clauser); ok {
//         rule, err = v.Rule()
//         if err != nil {
//             return err
//         }
//     }
//     *r = append(*r, rule)
//     return nil
// }

func marshalClauseParams(source Clause) (dynamic.Map, error) {
	return marshalParams(source)
}

func unmarshalParams(data []byte, target Clause) (map[string]dynamic.JSON, error) {
	var raw map[string]dynamic.JSON
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}
	res := map[string]dynamic.JSON{}
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
