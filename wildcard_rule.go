package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type WildcardRuleParams struct {
	// (Required) Wildcard pattern used to find matching terms.
	Wildcard string `json:"wildcard"`
	// (Optional, string) analyzer used to analyze terms in the query. Defaults
	// to the top-level <field>'s analyzer.
	Analyzer string `json:"analyzer,omitempty"`
	// (Optional) An optional interval filter.
	Filter RuleFilterer `json:"filter,omitempty"`
	// (Optional) If specified, then match intervals from this field
	// rather than the top-level <field>.
	//
	// The wildcard is normalized using the search analyzer from this field,
	// unless a separate analyzer is specified.
	UseField string `json:"use_field,omitempty"`
}

func (WildcardRuleParams) Type() RuleType {
	return RuleTypeWildcard
}
func (p WildcardRuleParams) Rule() (Rule, error) {
	return p.WildcardRule()
}

func (p WildcardRuleParams) WildcardRule() (*WildcardRule, error) {
	r := &WildcardRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetWildcard(p.Wildcard)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	return r, nil
}

type WildcardRule struct {
	analyzerParam
	ruleFilterParam
	useFieldParam
	wildcard string
}

func (w WildcardRule) Wildcard() string {
	return w.wildcard
}
func (w *WildcardRule) Rule() (QueryRule, error) {
	return w, nil
}
func (w *WildcardRule) SetWildcard(wildcard string) error {
	if len(wildcard) == 0 {
		return ErrWildcardRequired
	}
	w.wildcard = wildcard
	return nil
}
func (WildcardRule) Type() RuleType {
	return RuleTypeWildcard
}
func (p *WildcardRule) UnmarshalJSON(data []byte) error {
	*p = WildcardRule{}
	rd := dynamic.JSON(data)
	var r wildcardRule
	if rd.IsString() {
		var str string
		err := json.Unmarshal(rd, &str)
		if err != nil {
			return err
		}
		r = wildcardRule{
			Wildcard: str,
		}
	} else {
		err := json.Unmarshal(rd, &r)
		if err != nil {
			return err
		}
	}
	err := json.Unmarshal(rd, &r)
	if err != nil {
		return err
	}
	mv, err := r.WildcardRule()
	*p = *mv
	return err
}

func (p WildcardRule) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(wildcardRule{
		Wildcard: p.wildcard,
		Analyzer: p.analyzer,
		Filter:   p.filter,
		UseField: p.useField,
	})
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		p.Type().String(): d,
	}
	return json.Marshal(obj)
}

type wildcardRule struct {
	Wildcard string      `json:"wildcard"`
	Analyzer string      `json:"analyzer,omitempty"`
	Filter   *RuleFilter `json:"filter,omitempty"`
	UseField string      `json:"use_field,omitempty"`
}

func (p wildcardRule) WildcardRule() (*WildcardRule, error) {
	r := &WildcardRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetWildcard(p.Wildcard)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	return r, nil
}
