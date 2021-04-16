package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type MatchRuler interface {
	MatchRule(*MatchRule, error)
}

type MatchRuleParams struct {
	// (Required) Text you wish to find in the provided <field>.
	Query string `json:"query"`
	// (Optional, integer) Maximum number of positions between the matching
	// terms. Terms further apart than this are not considered matches. Defaults
	// to -1.
	//
	// If unspecified or set to -1, there is no width restriction on the match.
	// If set to 0, the terms must appear next to each other.
	MaxGaps interface{} `json:"max_gaps,omitempty"`
	// (Optional, Boolean) If true, matching terms must appear in their
	// specified order. Defaults to false.
	Ordered interface{} `json:"ordered,omitempty"`
	// (Optional, string) analyzer used to analyze terms in the query. Defaults
	// to the top-level <field>'s analyzer.
	Analyzer string `json:"analyzer,omitempty"`
	// (Optional, interval filter rule object) An optional interval filter.
	Filter RuleFilterer `json:"filter,omitempty"`
	// (Optional, string) If specified, then match intervals from this field
	// rather than the top-level <field>. Terms are analyzed using the search
	// analyzer from this field. This allows you to search across multiple
	// fields as if they were all the same field; for example, you could index
	// the same text into stemmed and unstemmed fields, and search for stemmed
	// tokens near unstemmed ones.
	UseField string `json:"use_field,omitempty"`
}

func (MatchRuleParams) Type() RuleType {
	return RuleTypeMatch
}
func (p MatchRuleParams) Rule() (QueryRule, error) {
	return p.MatchRule()
}

func (p MatchRuleParams) MatchRule() (*MatchRule, error) {
	r := &MatchRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetQuery(p.Query)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
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
	return r, nil
}

type MatchRule struct {
	analyzerParam
	maxGapsParam
	orderedParam
	ruleFilterParam
	useFieldParam
	query string
}

func (m *MatchRule) Rule() (QueryRule, error) {
	return m, nil
}

func (m *MatchRule) Query() string {
	return m.query
}

func (m *MatchRule) SetQuery(query string) error {
	if len(query) == 0 {
		return ErrQueryRequired
	}
	m.query = query
	return nil
}
func (MatchRule) Type() RuleType {
	return RuleTypeMatch
}
func (m *MatchRule) UnmarshalBSON(data []byte) error {
	return m.UnmarshalJSON(data)
}

func (m *MatchRule) UnmarshalJSON(data []byte) error {
	*m = MatchRule{}
	rd := dynamic.JSON(data)
	var r matchRule
	if rd.IsString() {
		var str string
		err := json.Unmarshal(rd, &str)
		if err != nil {
			return err
		}
		r = matchRule{
			Query: str,
		}
	} else {
		r = matchRule{}
		err := r.UnmarshalJSON(rd)
		if err != nil {
			return err
		}
	}

	mv, err := r.MatchRule()
	*m = *mv
	return err
}

func (m MatchRule) MarshalBSON() ([]byte, error) {
	return m.MarshalJSON()
}

func (m MatchRule) MarshalJSON() ([]byte, error) {
	return matchRule{
		Query:    m.query,
		MaxGaps:  m.maxGaps.Value(),
		Ordered:  m.ordered.Value(),
		Analyzer: m.analyzer,
		Filter:   m.filter,
		UseField: m.useField,
	}.MarshalJSON()

}

//easyjson:json
type matchRule struct {
	Query    string      `json:"query"`
	MaxGaps  interface{} `json:"max_gaps,omitempty"`
	Ordered  interface{} `json:"ordered,omitempty"`
	Analyzer string      `json:"analyzer,omitempty"`
	Filter   *RuleFilter `json:"filter,omitempty"`
	UseField string      `json:"use_field,omitempty"`
}

func (p matchRule) MatchRule() (*MatchRule, error) {
	r := &MatchRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetQuery(p.Query)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
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
	return r, nil
}
