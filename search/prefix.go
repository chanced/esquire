package search

type Prefix struct {
	Value           string
	Rewrite         Rewrite
	CaseInsensitive bool
	QueryName       string
}

func (p Prefix) Type() Type {
	return TypePrefix
}
func (p Prefix) Name() string {
	return p.QueryName
}

func (p Prefix) Clause() (Clause, error) {
	return p.prefix()
}
func (p Prefix) prefix() (*prefixClause, error) {
	q := &prefixClause{
		Value: p.Value,
	}
	q.SetCaseInsensitive(p.CaseInsensitive)
	q.SetRewrite(p.Rewrite)
	return q, nil
}

// prefixClause returns documents that contain a specific prefix in a provided field.
type prefixClause struct {
	Value string
	rewriteParam
	caseInsensitiveParam
	nameParam
}

func (p prefixClause) Type() Type {
	return TypePrefix
}

// PrefixQuery returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	*prefixClause
}

func (p PrefixQuery) Type() Type {
	return TypePrefix
}

// SetPrefix returns documents that contain a specific prefix in a provided field.
//
// SetPrefix panics if Value is not set. It is intended to be used inside of a
// builder.
func (p *PrefixQuery) SetPrefix(v Prefix) error {
	if v.Value == "" {
		return NewQueryError(ErrValueRequired, TypePrefix)
	}
	r, err := v.prefix()
	if err != nil {
		return err
	}
	p.prefixClause = r
	return nil
}
