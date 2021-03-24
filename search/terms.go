package search

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
//  err = s.AddTerms(&Terms{ Field:"", Value: []string{"kimchy", "elkbee"}})
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

func (t TermsLookup) lookupIsEmpty() bool {
	return len(t.TermsID) == 0 && len(t.TermsIndex) == 0 && len(t.TermsPath) == 0 && len(t.TermsRouting) == 0
}

type TermsRule struct {
	TermsLookup          `json:"-" bson:"-"`
	TermsValue           []string `json:"-" bson:"-"`
	TermsField           string   `json:"-" bson:"-"`
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
	data := []byte{}
	var err error
	data, err = marshalBoost(data, t)
	if err != nil {
		return nil, err
	}

	if t.TermsField != "" {
		if !t.TermsLookup.lookupIsEmpty() {
			return sjson.SetBytes(data, t.TermsField, t.TermsLookup)
		}
		return sjson.SetBytes(data, t.TermsField, t.TermsValue)
	}
	return data, nil
}

func (t *TermsRule) UnmarshalJSON(data []byte) error {
	g := gjson.GetBytes(data, "terms")
	if !g.Exists() {
		return nil
	}

	t.TermsValue = []string{}
	t.TermsLookup = TermsLookup{}

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
func (t *TermsRule) UnmarshalBSON(data []byte) error {
	return t.UnmarshalJSON(data)
}

func (t TermsRule) MarshalBSON() ([]byte, error) {
	return t.MarshalJSON()
}

type TermsQuery struct {
	Terms TermsRule `json:",inline" bson:",inline"`
}

func (t *TermsQuery) UnmarshalJSON(data []byte) error {
	fmt.Println("inside terms")
	return nil
}

func (t TermsQuery) SetTerms(field string, value Termser) error {

	return t.Terms.set(value)
}
