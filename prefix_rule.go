package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type PrefixRuleParams struct {
	// (Required) Beginning characters of terms you wish to find in the
	// top-level <field>.
	Prefix string `json:"prefix"`
	// (Optional, string) analyzer used to analyze terms in the query. Defaults
	// to the top-level <field>'s analyzer.
	Analyzer string `json:"analyzer,omitempty"`
	// (Optional) An optional interval filter.
	Filter RuleFilterer `json:"filter,omitempty"`
	// (Optional) If specified, then match intervals from this field
	// rather than the top-level <field>.
	//
	// The prefix is normalized using the search analyzer from this field,
	// unless a separate analyzer is specified.
	UseField string `json:"use_field,omitempty"`
}

func (PrefixRuleParams) Type() RuleType {
	return RuleTypePrefix
}
func (p PrefixRuleParams) Rule() (Rule, error) {
	return p.PrefixRule()
}

func (p PrefixRuleParams) PrefixRule() (*PrefixRule, error) {
	r := &PrefixRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetPrefix(p.Prefix)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	return r, nil
}

type PrefixRule struct {
	analyzerParam
	ruleFilterParam
	useFieldParam
	prefix string
}

func (p PrefixRule) Prefix() string {
	return p.prefix
}
func (p *PrefixRule) Rule() (QueryRule, error) {
	return p, nil
}
func (p *PrefixRule) SetPrefix(prefix string) error {
	if len(prefix) == 0 {
		return ErrPrefixRequired
	}
	p.prefix = prefix
	return nil
}
func (PrefixRule) Type() RuleType {
	return RuleTypePrefix
}
func (p *PrefixRule) UnmarshalBSON(data []byte) error {
	return p.UnmarshalJSON(data)
}

func (p *PrefixRule) UnmarshalJSON(data []byte) error {
	*p = PrefixRule{}
	rd := dynamic.JSON(data)
	var r prefixRule
	if rd.IsString() {
		var str string
		err := json.Unmarshal(rd, &str)
		if err != nil {
			return err
		}
		r = prefixRule{
			Prefix: str,
		}
	} else {
		err := json.Unmarshal(rd, &r)
		if err != nil {
			return err
		}
	}

	mv, err := r.PrefixRule()
	*p = *mv
	return err
}

func (p PrefixRule) MarshalBSON() ([]byte, error) {
	return p.MarshalJSON()
}

func (p PrefixRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(prefixRule{
		Prefix:   p.prefix,
		Analyzer: p.analyzer,
		Filter:   p.filter,
		UseField: p.useField,
	})
}

//easyjson:json
type prefixRule struct {
	Prefix   string      `json:"prefix"`
	Analyzer string      `json:"analyzer,omitempty"`
	Filter   *RuleFilter `json:"filter,omitempty"`
	UseField string      `json:"use_field,omitempty"`
}

func (p prefixRule) PrefixRule() (*PrefixRule, error) {
	r := &PrefixRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetPrefix(p.Prefix)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	return r, nil
}
