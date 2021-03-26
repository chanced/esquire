package search

type Prefix struct {
	Value           string
	Rewrite         Rewrite
	CaseInsensitive bool
}

func (p Prefix) Type() Type {
	return TypePrefix
}
func (p Prefix) Rule() (Rule, error) {
	return p.Query()
}
func (p Prefix) Query() (*PrefixRule, error) {
	q := &PrefixRule{
		Value: p.Value,
	}
	q.SetCaseInsensitive(p.CaseInsensitive)
	q.SetRewrite(p.Rewrite)
	return q, nil
}

// PrefixRule returns documents that contain a specific prefix in a provided field.
type PrefixRule struct {
	Value string `json:"value" bson:"value"`
	rewriteParam
	caseInsensitiveParam
}

func (p PrefixRule) Type() Type {
	return TypePrefix
}

// PrefixQuery returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	*PrefixRule `json:"prefix,omitempty" bson:"prefix,omitempty"`
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
	r, err := v.Query()
	if err != nil {
		return err
	}
	p.PrefixRule = r
	return nil
}
