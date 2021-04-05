package picker

type WithRuleFilter interface {
	Filter() *RuleFilter
	SetFilter(rf RuleFilterer) error
}

type ruleFilterParam struct {
	filter *RuleFilter
}

func (r *ruleFilterParam) SetFilter(filter RuleFilterer) error {
	if filter == nil {
		r.filter = nil
		return nil
	}
	f, err := filter.RuleFilter()
	if err != nil {
		return err
	}
	r.filter = f
	return nil
}
func (r ruleFilterParam) Filter() *RuleFilter {
	return r.filter
}
