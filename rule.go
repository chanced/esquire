package picker

import (
	"encoding/json"
	"errors"

	"github.com/chanced/dynamic"
)

type Ruleset []Ruler

func (rs Ruleset) MarshalJSON() ([]byte, error) {
	rl, err := rs.Rules()
	if err != nil {
		return nil, err
	}
	return json.Marshal(rl)
}
func (rs *Ruleset) UnmarshalJSON(data []byte) error {
	var rl Rules
	err := json.Unmarshal(data, &rl)
	if err != nil {
		return err
	}
	*rs = make(Ruleset, len(rl))
	for i, r := range rl {
		(*rs)[i] = r
	}
	return nil
}
func (rs Ruleset) Rules() (Rules, error) {
	res := make(Rules, len(rs))
	for i, v := range rs {
		r, err := v.Rule()
		if err != nil {
			return res, err
		}
		res[i] = r
	}
	return res, nil
}

type Rules []QueryRule

func (rs Rules) MarshalJSON() ([]byte, error) {
	res := make([]dynamic.JSONObject, len(rs))
	for i, r := range rs {
		rd, err := r.MarshalJSON()
		if err != nil {
			return nil, err
		}
		res[i] = dynamic.JSONObject{
			r.Type().String(): rd,
		}
	}
	return json.Marshal(res)
}

func (rs *Rules) UnmarshalJSON(data []byte) error {
	var array []dynamic.JSON
	err := json.Unmarshal(data, &array)
	if err != nil {
		return err
	}
	*rs = make(Rules, len(array))
	for i, rd := range array {
		rv, err := UnmarshalRule(rd)
		if err != nil {
			return err
		}
		(*rs)[i] = rv
	}
	return nil
}

func (rs *Ruleset) Add(rule Ruler) (QueryRule, error) {
	qr, err := rule.Rule()
	if err != nil {
		return qr, err
	}
	*rs = append(*rs, qr)
	return qr, nil
}

type Rule interface {
	Type() RuleType
}

type QueryRule interface {
	Rule
	Ruler
	json.Marshaler
	json.Unmarshaler
}

type Ruler interface {
	Rule() (QueryRule, error)
}

type RuleType string

func (r RuleType) String() string {
	return string(r)
}

const (
	RuleTypeMatch    RuleType = "match"
	RuleTypePrefix   RuleType = "prefix"
	RuleTypeWildcard RuleType = "wildcard"
	RuleTypeFuzzy    RuleType = "fuzzy"
	RuleTypeAllOf    RuleType = "all_of"
	RuleTypeAnyOf    RuleType = "any_of"
)

func MarshalRule(rule QueryRule) ([]byte, error) {
	data, err := rule.MarshalJSON()
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		rule.Type().String(): data,
	}
	return obj.MarshalJSON()
}
func UnmarshalRule(data []byte) (QueryRule, error) {
	var d dynamic.JSONObject
	err := json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}
	var qr QueryRule
	for k, rd := range d {
		handler, ok := ruleTypeHandlers[RuleType(k)]
		if !ok {
			return nil, errors.New("picker: unsupported rule type")
		}
		qr = handler()
		err = qr.UnmarshalJSON(rd)
		return qr, err
	}
	return nil, errors.New("picker: empty rule")
}

var ruleTypeHandlers = map[RuleType]func() QueryRule{
	RuleTypeMatch:    func() QueryRule { return &MatchRule{} },
	RuleTypeFuzzy:    func() QueryRule { return &FuzzyRule{} },
	RuleTypeWildcard: func() QueryRule { return &WildcardRule{} },
	RuleTypePrefix:   func() QueryRule { return &PrefixRule{} },
	// RuleTypeAllOf:    func() QueryRule { return &AllOfRule{} },
	// RuleTypeAnyOf:    func() QueryRule { return &AnyOfRule{} },
}
