package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type FuzzyRuler interface {
	FuzzyRule() (*FuzzyRule, error)
}

type FuzzyRuleParams struct {
	// (Required) The term to match
	Term string `json:"term"`
	// (Optional, string) analyzer used to analyze terms in the query. Defaults
	// to the top-level <field>'s analyzer.
	Analyzer string `json:"analyzer,omitempty"`
	// (Optional) An optional interval filter.
	Filter RuleFilterer `json:"filter,omitempty"`
	// (Optional) If specified, then match intervals from this field
	// rather than the top-level <field>.
	//
	// The fuzzy is normalized using the search analyzer from this field,
	// unless a separate analyzer is specified.
	UseField string `json:"use_field,omitempty"`
	// (Optional, integer) Number of beginning characters left unchanged when creating expansions. Defaults to 0.
	PrefixLength int `json:"prefix_length,omitempty"`
	// (Optional, string) Maximum edit distance allowed for matching. See
	// Fuzziness for valid values and more information. Defaults to auto.
	Fuzziness string `json:"fuzziness,omitempty"`
	// (Optional, Boolean) Indicates whether edits include transpositions of two
	// adjacent characters (ab → ba). Defaults to true.
	Transpositions interface{} `json:"transpositions,omitempty"`
}

func (FuzzyRuleParams) Type() RuleType {
	return RuleTypeFuzzy
}
func (p FuzzyRuleParams) Rule() (Rule, error) {
	return p.FuzzyRule()
}

func (p FuzzyRuleParams) FuzzyRule() (*FuzzyRule, error) {
	r := &FuzzyRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetTerm(p.Term)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	return r, nil
}

type FuzzyRule struct {
	analyzerParam
	ruleFilterParam
	useFieldParam
	transpositionsParam
	prefixLengthParam
	fuzzinessParam

	term string
}

func (f FuzzyRule) Term() string {
	return f.term
}

func (f *FuzzyRule) SetTerm(term string) error {
	if len(term) == 0 {
		return ErrTermRequired
	}
	f.term = term
	return nil
}
func (FuzzyRule) Type() RuleType {
	return RuleTypeFuzzy
}
func (f *FuzzyRule) Rule() (QueryRule, error) {
	return f, nil
}

func (f *FuzzyRule) UnmarshalBSON(data []byte) error {
	return f.UnmarshalJSON(data)
}

func (f *FuzzyRule) UnmarshalJSON(data []byte) error {
	*f = FuzzyRule{}
	rd := dynamic.JSON(data)
	var r fuzzyRule
	if rd.IsString() {
		var str string
		err := json.Unmarshal(rd, &str)
		if err != nil {
			return err
		}
		r = fuzzyRule{
			Term: str,
		}
	} else {
		err := json.Unmarshal(rd, &r)
		if err != nil {
			return err
		}
	}

	mv, err := r.FuzzyRule()
	*f = *mv
	return err
}

func (f FuzzyRule) MarshalBSON() ([]byte, error) {
	return f.MarshalJSON()
}

func (f FuzzyRule) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(fuzzyRule{
		Term:           f.term,
		Analyzer:       f.analyzer,
		Filter:         f.filter,
		UseField:       f.useField,
		PrefixLength:   f.PrefixLength(),
		Fuzziness:      f.fuzziness,
		Transpositions: f.Transpositions(),
	})
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		f.Type().String(): d,
	}
	return json.Marshal(obj)
}

//easyjson:json
type fuzzyRule struct {
	Term           string      `json:"term"`
	Analyzer       string      `json:"analyzer,omitempty"`
	Filter         *RuleFilter `json:"filter,omitempty"`
	UseField       string      `json:"use_field,omitempty"`
	PrefixLength   int         `json:"prefix_length,omitempty"`
	Fuzziness      string      `json:"fuzziness,omitempty"`
	Transpositions interface{} `json:"transpositions,omitempty"`
}

func (p fuzzyRule) FuzzyRule() (*FuzzyRule, error) {
	r := &FuzzyRule{}
	r.SetUseField(p.UseField)
	r.SetAnalyzer(p.Analyzer)
	err := r.SetTerm(p.Term)
	if err != nil {
		return r, err
	}
	err = r.SetFilter(p.Filter)
	if err != nil {
		return r, err
	}
	return r, nil
}
