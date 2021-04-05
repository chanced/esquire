package picker

type AllOfRuleParams struct {

	// (Required, array of rule objects) An array of rules to combine. All rules
	// must produce a match in a document for the overall source to match.
	Intervals Ruleset `json:"intervals"`
	// (Optional, Boolean) If true, intervals produced by the rules should
	// appear in the order in which they are specified. Defaults to false.
	Ordered interface{} `json:"ordered,omitempty"`
	// (Optional) Rule used to filter returned intervals.
	Filter  RuleFilterer `json:"filter,omitempty"`
	MaxGaps interface{}  `json:"max_gaps,omitempty"`
}

func (p AllOfRuleParams) Rule() (QueryRule, error) {
	return p.AllOfRule()
}
func (p AllOfRuleParams) AllOfRule() (*AllOfRule, error) {
	r := &AllOfRule{}
	err := r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	err = r.SetMaxGaps(p.MaxGaps)
	if err != nil {
		return r, err
	}
	err = r.SetOrdered(p.Ordered)
	if err != nil {
		return r, err
	}
	err = r.SetIntervals(p.Intervals)
	if err != nil {
		return r, err
	}
	return r, nil
}
func (AllOfRuleParams) Type() RuleType {
	return RuleTypeAllOf
}

type AllOfRule struct {
	intervals Rules
	maxGapsParam
	ruleFilterParam
	orderedParam
}

func (a AllOfRule) Intervals() Rules {
	return a.intervals
}
func (a *AllOfRule) SetIntervals(intervals Ruleset) error {

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

func (a *AllOfRule) UnmarshalJSON(data []byte) error {
	*a = AllOfRule{}
	rv := allOfRule{}
	err := rv.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	a.filter = rv.Filter
	a.intervals = rv.Intervals
	err = a.maxGaps.Set(rv.MaxGaps)
	if err != nil {
		return err
	}
	err = a.ordered.Set(rv.Ordered)
	if err != nil {
		return err
	}
	return nil
}

func (a AllOfRule) MarshalJSON() ([]byte, error) {
	return allOfRule{
		Intervals: a.intervals,
		Ordered:   a.ordered,
		Filter:    a.filter,
		MaxGaps:   a.maxGaps.Value(),
	}.MarshalJSON()
}
func (a *AllOfRule) Rule() (QueryRule, error) {
	if a == nil {
		return nil, nil
	}
	return a.AllOfRule()
}
func (a *AllOfRule) AllOfRule() (*AllOfRule, error) {
	return a, nil
}
func (AllOfRule) Type() RuleType {
	return RuleTypeAllOf
}

//easyjson:json
type allOfRule struct {
	Intervals Rules       `json:"intervals"`
	Ordered   interface{} `json:"ordered,omitempty"`
	Filter    *RuleFilter `json:"filter,omitempty"`
	MaxGaps   interface{} `json:"max_gaps,omitempty"`
}
