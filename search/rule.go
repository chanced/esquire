package search

import (
	"encoding/json"
	"errors"

	"github.com/chanced/dynamic"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Rule interface {
	Type() Type
}

type Ruler interface {
	Rule() (Rule, error)
}

type Clause struct {
	Type Type
	Rule Rule
}

func (c Clause) MarshalJSON() ([]byte, error) {
	return sjson.SetBytes([]byte{}, c.Type.String(), c.Rule)
}

func (c *Clause) UnmarshalJSON(data []byte) error {

	// TODO: really need to do some serious type checking here

	g := gjson.ParseBytes(data)
	c.Rule = nil
	c.Type = ""
	var typ Type
	var rule Rule
	var err error
	g.ForEach(func(key, value gjson.Result) bool {
		c.Type = Type(key.Str)
		if typ == "" {
			return true
		}
		rule = TypeHandlers[typ]()
		err = json.Unmarshal([]byte(value.Raw), rule)
		return false
	})
	if typ == "" {
		return errors.New("missing type on clause")
	}
	c.Type = typ
	c.Rule = rule
	return err
}
func (c *Clause) UnmarshalBSON(data []byte) error {
	return c.UnmarshalJSON(data)
}

func (c Clause) MarshalBSON() ([]byte, error) {
	return c.MarshalJSON()
}

type Rules []Clause

func (r *Rules) Add(typ Type, rule Rule) error {
	var err error
	if v, ok := rule.(Ruler); ok {
		rule, err = v.Rule()
		if err != nil {
			return err
		}
	}
	*r = append(*r, Clause{Type: typ, Rule: rule})
	return nil
}

func marshalRuleParams(source Rule) (dynamic.Map, error) {
	return marshalParams(source)
}

func unmarshalParams(data []byte, target Rule) (map[string]dynamic.RawJSON, error) {
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
