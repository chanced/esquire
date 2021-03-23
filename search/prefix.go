package search

type Prefix struct {
	Value             string
	Rewrite           Rewrite
	IsCaseInsensitive bool
}

func (p Prefix) Type() Type {
	return TypePrefix
}

func (p Prefix) Query() PrefixQueryValue {
	q := PrefixQueryValue{
		Value: p.Value,
	}
	q.SetCaseInsensitive(p.IsCaseInsensitive)
	q.SetRewrite(p.Rewrite)
	return q
}

// PrefixQueryValue returns documents that contain a specific prefix in a provided field.
type PrefixQueryValue struct {
	Value                string `json:"value" bson:"value"`
	RewriteParam         `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

func (p PrefixQueryValue) Type() Type {
	return TypePrefix
}

// PrefixQuery returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	PrefixQueryValue `json:"prefix,omitempty" bson:"prefix,omitempty"`
}

func (p PrefixQuery) Type() Type {
	return TypePrefix
}

// SetPrefix returns documents that contain a specific prefix in a provided field.
//
// SetPrefix panics if Value is not set. It is intended to be used inside of a
// builder.
func (p *PrefixQuery) SetPrefix(v Prefix) {
	if v.Value == "" {
		panic(NewQueryError(ErrValueRequired, TypePrefix))
	}
	p.PrefixQueryValue = v.Query()
}
