package search

// QueryWithCaseInsensitive is a query mixin that adds the case_insensitive param
//
// (Optional, Boolean) Allows ASCII case insensitive matching of the value with
// the indexed field values when set to true. Default is false which means the
// case sensitivity of matching depends on the underlying field’s mapping.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type QueryWithCaseInsensitive interface {
	CaseInsensitive() bool
	SetCaseInsensitive(v bool)
}

// CaseInsensitiveParam is a query mixin that adds the case_insensitive param
//
// (Optional, Boolean) Allows ASCII case insensitive matching of the value with
// the indexed field values when set to true. Default is false which means the
// case sensitivity of matching depends on the underlying field’s mapping.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type CaseInsensitiveParam struct {
	CaseInsensitiveValue *bool `json:"case_insensitive,omitempty" bson:"case_insensitive,omitempty"`
}

func (ci CaseInsensitiveParam) Default() bool {
	return false
}

func (ci CaseInsensitiveParam) CaseInsensitive() bool {
	if ci.CaseInsensitiveValue == nil {
		return ci.Default()
	}
	return *ci.CaseInsensitiveValue
}

func (ci *CaseInsensitiveParam) SetCaseInsensitive(v bool) {
	if ci.CaseInsensitive() != v {
		ci.CaseInsensitiveValue = &v
	}
}
