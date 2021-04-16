package picker

import "encoding/json"

type AnyOfRuleParams struct {
	// (Required, array of rule objects) An array of rules to match.
	Intervals Ruleset `json:"intervals"`
	// (Optional, interval filter rule object) Rule used to filter returned intervals.
	Filter RuleFilterer `json:"filter,omitempty"`
}

func (p AnyOfRuleParams) Rule() (QueryRule, error) {
	return p.AnyOfRule()
}
func (p AnyOfRuleParams) AnyOfRule() (*AnyOfRule, error) {
	r := &AnyOfRule{}
	err := r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	err = r.SetIntervals(p.Intervals)
	if err != nil {
		return r, err
	}
	return r, nil
}
func (AnyOfRuleParams) Type() RuleType {
	return RuleTypeAnyOf
}

type AnyOfRule struct {
	intervals Rules
	ruleFilterParam
}

func (a AnyOfRule) Intervals() Rules {
	return a.intervals
}
func (a *AnyOfRule) SetIntervals(intervals Ruleset) error {

	rl, err := intervals.Rules()
	if err != nil {
		return err
	}
	if len(rl) == 0 {
		return ErrIntervalsRequired
	}
	a.intervals = rl
	return nil
}

func (a *AnyOfRule) UnmarshalBSON(data []byte) error {
	return a.UnmarshalJSON(data)
}

func (a *AnyOfRule) UnmarshalJSON(data []byte) error {
	*a = AnyOfRule{}
	var rv anyOfRule
	err := json.Unmarshal(data, &rv)
	if err != nil {
		return err
	}
	a.filter = rv.Filter
	a.intervals = rv.Intervals

	return nil
}

func (a AnyOfRule) MarshalBSON() ([]byte, error) {
	return a.MarshalJSON()
}

func (a AnyOfRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(anyOfRule{
		Intervals: a.intervals,
		Filter:    a.filter,
	})
}
func (a *AnyOfRule) Rule() (QueryRule, error) {
	if a == nil {
		return nil, nil
	}
	return a.AnyOfRule()
}
func (a *AnyOfRule) AnyOfRule() (*AnyOfRule, error) {
	return a, nil
}
func (AnyOfRule) Type() RuleType {
	return RuleTypeAnyOf
}

//easyjson:json
type anyOfRule struct {
	Intervals Rules       `json:"intervals"`
	Filter    *RuleFilter `json:"filter,omitempty"`
}
