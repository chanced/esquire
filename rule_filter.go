package picker

import "encoding/json"

type RuleFilterer interface {
	RuleFilter() (*RuleFilter, error)
}

type RuleFilterParams struct {
	// (Optional) Query used to return intervals that follow an
	// interval from the filter rule.
	After Querier
	// (Optional) Query used to return intervals that occur before
	// an interval from the filter rule.
	Before Querier
	// (Optional) Query used to return intervals contained by an
	// interval from the filter rule.
	ContainedBy Querier
	// (Optional) Query used to return intervals that contain an
	// interval from the filter rule.
	Containing Querier
	// (Optional) Query used to return intervals that are not
	// contained by an interval from the filter rule.
	NotContainedBy Querier
	// (Optional) Query used to return intervals that do not
	// contain an interval from the filter rule.
	NotContaining Querier
	// (Optional) Query used to return intervals that do not overlap with an interval from the filter rule.
	NotOverlapping Querier
	// (Optional) Query used to return intervals that overlap with an interval from the filter rule.
	Overlapping Querier
	// Optional) Script used to return matching documents. This script must return a boolean value, true or false. See Script filters for an example.
	Script *Script
}

func (p RuleFilterParams) RuleFilter() (*RuleFilter, error) {
	r := &RuleFilter{}
	var err error
	err = r.SetAfter(p.After)
	if err != nil {
		return r, err
	}
	err = r.SetBefore(p.Before)
	if err != nil {
		return r, err
	}
	err = r.SetContainedBy(p.ContainedBy)
	if err != nil {
		return r, err
	}
	err = r.SetContaining(p.Containing)
	if err != nil {
		return r, err
	}
	err = r.SetNotContainedBy(p.NotContainedBy)
	if err != nil {
		return r, err
	}
	err = r.SetNotContaining(p.NotContaining)
	if err != nil {
		return r, err
	}
	err = r.SetNotOverlapping(p.NotOverlapping)
	if err != nil {
		return r, err
	}
	err = r.SetOverlapping(p.Overlapping)
	if err != nil {
		return r, err
	}
	err = r.SetScript(p.Script)
	if err != nil {
		return r, err
	}
	return r, nil
}

type RuleFilter struct {
	after          *Query
	before         *Query
	containedBy    *Query
	containing     *Query
	notContainedBy *Query
	notContaining  *Query
	notOverlapping *Query
	overlapping    *Query
	script         *Script
}

func (r *RuleFilter) RuleFilter() (*RuleFilter, error) {
	return r, nil
}
func (r *RuleFilter) UnmarshalJSON(data []byte) error {
	*r = RuleFilter{}
	var v ruleFilter
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	r.after = v.After
	r.before = v.Before
	r.containedBy = v.ContainedBy
	r.containing = v.Containing
	r.notContainedBy = v.NotContainedBy
	r.notContaining = v.NotContaining
	r.notOverlapping = v.NotOverlapping
	r.overlapping = v.Overlapping
	r.script = v.Script

	return nil
}

func (r RuleFilter) MarshalJSON() ([]byte, error) {
	v := ruleFilter{}
	if !r.after.IsEmpty() {
		v.After = r.after
	}
	if !r.before.IsEmpty() {
		v.Before = r.before
	}
	if !r.containedBy.IsEmpty() {
		v.ContainedBy = r.containedBy
	}
	if !r.containing.IsEmpty() {
		v.Containing = r.containing
	}
	if !r.notContainedBy.IsEmpty() {
		v.NotContainedBy = r.notContainedBy
	}
	if !r.notContaining.IsEmpty() {
		v.NotContaining = r.notContaining
	}
	if !r.notOverlapping.IsEmpty() {
		v.NotOverlapping = r.notOverlapping
	}
	if !r.overlapping.IsEmpty() {
		v.Overlapping = r.overlapping
	}
	if !r.script.IsEmpty() {
		v.Script = r.script
	}
	return json.Marshal(v)
}

// After - Query used to return intervals that follow an
// interval from the filter rule.
func (r *RuleFilter) After() *Query {
	if r.after == nil {
		r.after = &Query{}
	}
	return r.after
}
func (r *RuleFilter) SetAfter(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.after = q
	return nil
}

// Before - Query used to return intervals that occur before
// an interval from the filter rule.
func (r *RuleFilter) Before() *Query {
	if r.before == nil {
		r.before = &Query{}
	}
	return r.before
}
func (r *RuleFilter) SetBefore(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.before = q
	return nil
}

// ContainedBy - Query used to return intervals contained by an
// interval from the filter rule.
func (r *RuleFilter) ContainedBy() *Query {
	if r.containedBy == nil {
		r.containedBy = &Query{}
	}
	return r.containedBy
}
func (r *RuleFilter) SetContainedBy(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.containedBy = q
	return nil
}

// Containing - Query used to return intervals that contain an
// interval from the filter rule.
func (r *RuleFilter) Containing() *Query {
	if r.containing == nil {
		r.containing = &Query{}
	}
	return r.containing
}
func (r *RuleFilter) SetContaining(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.containing = q
	return nil
}

// NotContainedBy - Query used to return intervals that are not
// contained by an interval from the filter rule.
func (r *RuleFilter) NotContainedBy() *Query {
	if r.notContainedBy == nil {
		r.notContainedBy = &Query{}
	}
	return r.notContainedBy
}
func (r *RuleFilter) SetNotContainedBy(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.notContainedBy = q
	return nil
}

// NotContaining - Query used to return intervals that do not
// contain an interval from the filter rule.
func (r *RuleFilter) NotContaining() *Query {
	if r.notContaining == nil {
		r.notContaining = &Query{}
	}
	return r.notContaining
}
func (r *RuleFilter) SetNotContaining(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.notContaining = q
	return nil
}

// NotOverlapping - Query used to return intervals that do not overlap with an
// interval from the filter rule.
func (r *RuleFilter) NotOverlapping() *Query {
	if r.notOverlapping == nil {
		r.notOverlapping = &Query{}
	}
	return r.notOverlapping
}
func (r *RuleFilter) SetNotOverlapping(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.notOverlapping = q
	return nil
}

// Overlapping - Query used to return intervals that overlap with an interval
// from the filter rule.
func (r *RuleFilter) Overlapping() *Query {
	if r.overlapping == nil {
		r.overlapping = &Query{}
	}
	return r.overlapping
}
func (r *RuleFilter) SetOverlapping(query Querier) error {
	q, err := query.Query()
	if err != nil {
		return err
	}
	r.overlapping = q
	return nil
}
func (r *RuleFilter) Script() *Script {
	if r.script == nil {
		r.script = &Script{}
	}
	return r.script
}

func (r *RuleFilter) SetScript(script *Script) error {
	if script == nil {
		r.script = &Script{}
	}
	r.script = script
	return nil
}

type ruleFilter struct {
	After          *Query  `json:"after,omitempty"`
	Before         *Query  `json:"before,omitempty"`
	ContainedBy    *Query  `json:"contained_by,omitempty"`
	Containing     *Query  `json:"containing,omitempty"`
	NotContainedBy *Query  `json:"not_contained_by,omitempty"`
	NotContaining  *Query  `json:"not_containing,omitempty"`
	NotOverlapping *Query  `json:"not_overlapping,omitempty"`
	Overlapping    *Query  `json:"overlapping,omitempty"`
	Script         *Script `json:"script,omitempty"`
}
