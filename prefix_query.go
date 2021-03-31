package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Prefixer interface {
	Prefix() (*PrefixClause, error)
}

// Prefix returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type Prefix struct {
	// Field which the query is being performed. This is not needed if setting explicitly
	// but is required if the query is being added to a compound query.
	Field string
	// Beginning characters of terms you wish to find in the provided <field>. (Required)
	Value string
	// Method used to rewrite the query. For valid values and more information,
	// see the rewrite parameter. (Optional)
	Rewrite Rewrite
	// Allows ASCII case insensitive matching of the value with the indexed
	// field values when set to true. Default is false which means the case
	// sensitivity of matching depends on the underlying field’s mapping. (Optional)
	CaseInsensitive bool
	// Name of the query (Optional)
	Name string
	clause
}

func (p Prefix) Kind() Kind {
	return KindPrefix
}

func (p Prefix) Clause() (Clause, error) {
	return p.Prefix()
}
func (p Prefix) Prefix() (*PrefixClause, error) {
	q := &PrefixClause{field: p.Field}
	q.SetCaseInsensitive(p.CaseInsensitive)
	err := q.SetRewrite(p.Rewrite)
	if err != nil {
		return q, NewQueryError(err, KindPrefix, p.Field)
	}
	return q, q.setValue(p.Value)
}

// PrefixClause returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixClause struct {
	value string
	field string
	rewriteParam
	caseInsensitiveParam
	nameParam
	clause
}

func (p PrefixClause) Value() string {
	return p.value
}

func (p *PrefixClause) setValue(value string) error {
	err := checkValue(value, KindPrefix, p.field)
	if err != nil {
		return err
	}
	p.value = value
	return nil
}

func (p PrefixClause) Kind() Kind {
	return KindPrefix
}

func (p PrefixClause) MarshalJSON() ([]byte, error) {
	if p.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := p.marshalClauseJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{p.field: data})
}

func (p PrefixClause) marshalClauseJSON() (dynamic.JSON, error) {
	params, err := marshalParams(&p)
	if err != nil {
		return nil, err
	}
	params["value"] = p.value
	return json.Marshal(params)
}

func (p *PrefixClause) UnmarshalJSON(data []byte) error {
	*p = PrefixClause{}

	obj := dynamic.JSONObject{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	for k, v := range obj {
		p.field = k
		return p.unmarshalClauseJSON(v)
	}
	return nil
}

func (p *PrefixClause) unmarshalClauseJSON(data dynamic.JSON) error {
	fields, err := unmarshalParams(data, p)
	if err != nil {
		return err
	}
	if v, ok := fields["query"]; ok {
		var s string
		err := json.Unmarshal(v, &s)
		if err != nil {
			return err
		}
		p.value = s
	}
	return nil
}

// Set sets the value of PrefixQuery.
//
// Valid values:
//  - picker.Prefix
//  - picker.String
//  - nil (clears PrefixQuery)
func (p *PrefixClause) Set(field string, prefixer Prefixer) error {
	if prefixer == nil {
		p.Clear()
	}
	err := checkField(field, KindPrefix)
	if err != nil {
		return NewQueryError(err, KindPrefix, field)
	}
	q, err := prefixer.Prefix()
	if err != nil {
		return NewQueryError(err, KindPrefix, field)
	}
	*p = *q
	return nil
}

func (p *PrefixClause) IsEmpty() bool {
	return p == nil || len(p.field) == 0 || len(p.value) == 0
}

func (p *PrefixClause) Clear() {
	*p = PrefixClause{}
}
