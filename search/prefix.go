package search

type Prefixer interface {
	Prefix() (*PrefixQuery, error)
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
}

func (p Prefix) Type() Type {
	return TypePrefix
}

func (p Prefix) Clause() (Clause, error) {
	return p.Prefix()
}
func (p Prefix) Prefix() (*PrefixQuery, error) {
	q := &PrefixQuery{field: p.Field}
	q.SetCaseInsensitive(p.CaseInsensitive)
	err := q.SetRewrite(p.Rewrite)
	if err != nil {
		return q, NewQueryError(err, TypePrefix, p.Field)
	}
	return q, q.setValue(p.Value)
}

// PrefixQuery returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	value string
	field string
	rewriteParam
	caseInsensitiveParam
	nameParam
}

func (p PrefixQuery) Value() string {
	return p.value
}

func (p *PrefixQuery) setValue(value string) error {
	err := checkValue(value, TypePrefix, p.field)
	if err != nil {
		return err
	}
	p.value = value
	return nil
}

func (p PrefixQuery) Type() Type {
	return TypePrefix
}

// Set sets the value of PrefixQuery.
//
// Valid values:
//  - search.Prefix
//  - search.String
//  - nil (clears PrefixQuery)
func (p *PrefixQuery) Set(field string, prefixer Prefixer) error {
	if prefixer == nil {
		p.Clear()
	}
	err := checkField(field, TypePrefix)
	if err != nil {
		return NewQueryError(err, TypePrefix, field)
	}
	q, err := prefixer.Prefix()
	if err != nil {
		return NewQueryError(err, TypePrefix, field)
	}
	*p = *q
	return nil
}

func (p *PrefixQuery) Clear() {
	*p = PrefixQuery{}
}
