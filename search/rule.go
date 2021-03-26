package search

import (
	"encoding/json"
	"errors"

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

func marshalRuleParams(data M, source Rule) (M, error) {
	return marshalParams(data, source)
}

func unmarshalRule(data []byte, target Rule, fn func(key, value gjson.Result) error) error {
	var err error

	g.ForEach(func(key, value gjson.Result) bool {
		var isParam bool
		isParam, err = unmarshalParam(key.Str, target, value)
		if err != nil {
			return false
		}
		if isParam {
			return true
		}
		if fn != nil {
			err = fn(key, value)
			if err != nil {
				return false
			}

		}
		return true
	})
	return err
}
