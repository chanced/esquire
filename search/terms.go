package search

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// Termser can be:
//  *search.Terms
//  *search.Lookup
//  search.String
//  search.Strings
//
// Example:
//  impor
//  s := search.NewSearch()
//  err := s.AddTerms(&Lookup{ID: "1", Index:"my-index-100", Path:"color"})
//  _ = err // handle err
//  err = s.AddTerms(&Terms{ Field:"", Value: []string{"chanced", "kimchy", "elkbee"}})
//  _ = err // handle err
type Termser interface {
	Terms() (*TermsRule, error)
}

// Terms returns documents that contain one or more exact terms in a provided
// field.
//
// The terms query is the same as the term query, except you can search for
// multiple values.
type Terms struct {
	Values          []string
	Boost           float64
	CaseInsensitive bool
}

func (t Terms) Rule() (Rule, error) {
	return t.Terms()
}
func (t Terms) Terms() (*TermsRule, error) {
	q := &TermsRule{
		TermsValue: t.Values,
	}
	q.SetBoost(t.Boost)
	q.SetCaseInsensitive(t.CaseInsensitive)
	return q, nil
}

func (t Terms) Type() Type {
	return TypeTerms
}

type TermsLookup struct {
	TermsID      string `json:"id,omitempty" bson:"id,omitempty"`
	TermsIndex   string `json:"index,omitempty" bson:"index,omitempty"`
	TermsPath    string `json:"path,omitempty" bson:"path,omitempty"`
	TermsRouting string `json:"routing,omitempty" bson:"routing,omitempty"`
}

type TermsRule struct {
	TermsValue           []string `json:"-"`
	TermsField           string   `json:"-"`
	TermsLookup          `json:",inline" bson:",inline"`
	BoostParam           `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

func (t *TermsRule) Type() Type {
	return TypeTerms
}

func (t *TermsRule) SetValue(value []string) {
	t.TermsLookup = TermsLookup{}
	if value == nil {
		value = []string{}
	}
	t.TermsValue = value
}
func (t *TermsRule) SetField(field string) {
	t.TermsField = field
}
func (t *TermsRule) SetLookup(lookup *TermsLookup) {
	t.SetValue([]string{})
	if lookup == nil {
		lookup = &TermsLookup{}
	}
	t.TermsLookup = *lookup

}
func (t *TermsRule) set(v Termser) error {
	tv, err := v.Terms()
	if err != nil {
		return err
	}
	t.SetBoost(tv.Boost())
	t.SetCaseInsensitive(tv.CaseInsensitive())
	t.TermsLookup = tv.TermsLookup
	t.TermsValue = tv.TermsValue
	return nil
}

func (t TermsRule) Value() []string {
	return t.TermsValue
}

func (t TermsRule) MarshalJSON() ([]byte, error) {
	panic("not imp")
}
func (t *TermsRule) UnmarshalJSON(data []byte) error {
	t.TermsValue = []string{}
	t.TermsLookup = TermsLookup{}

	g := gjson.ParseBytes(data)
	err := unmarshalRule(g, t, func(key, value gjson.Result) error {
		t.TermsField = key.Str
		if value.IsArray() {
			value.ForEach(func(key, value gjson.Result) bool {
				t.TermsValue = append(t.TermsValue, value.String())
				return true
			})
		} else {
			err := json.Unmarshal([]byte(value.Raw), &t.TermsLookup)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

type TermsQuery struct {
	TermsRule `json:",inline" bson:",inline"`
}

func (t *TermsQuery) SetTerms(field string, value Termser) error {
	t.TermsField = field
	return t.set(value)
}
