package search

type Prefix struct {
	Value             string
	Rewrite           Rewrite
	IsCaseInsensitive bool
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

// PrefixQuery returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	PrefixQueryValue `json:"prefix,omitempty" bson:"prefix,omitempty"`
}

// Prefix returns documents that contain a specific prefix in a provided field.
//
// Prefix panics if Value is not set. It is intended to be used inside of a builder.
func (p *PrefixQuery) Prefix(v Prefix) {
	if v.Value == "" {
		panic(NewSearchError(ErrMissingValue, QueryTypePrefix))
	}
	p.PrefixQueryValue = v.Query()
}
