package search

import "github.com/tidwall/gjson"

type ZeroTermsQuery string

const (
	// ZeroTermsQueryNone - No documents are returned if the analyzer removes all
	// tokens.
	ZeroTermsQueryNone ZeroTermsQuery = "none"
	// ZeroTermsALl - Returns all documents, similar to a match_all query.
	ZeroTermsQueryAll ZeroTermsQuery = "all"
)

type WithZeroTermsQuery interface {
	// ZeroTermsQuery indicates  whether no documents are returned if the
	// analyzer removes all tokens, such as when using a stop filter
	ZeroTermsQuery() ZeroTermsQuery
	SetZeroTermsQuery(v ZeroTermsQuery)
}

type ZeroTermsQueryParam struct {
	ZeroTermsQueryValue *ZeroTermsQuery `json:"zero_terms_query,omitempty" bson:"zero_terms_query,omitempty"`
}

// ZeroTermsQuery indicates  whether no documents are returned if the
// analyzer removes all tokens, such as when using a stop filter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#query-dsl-match-query-zero
func (ztq ZeroTermsQueryParam) ZeroTermsQuery() ZeroTermsQuery {
	if ztq.ZeroTermsQueryValue != nil {
		return *ztq.ZeroTermsQueryValue
	}
	return ZeroTermsQueryNone
}

func (ztq *ZeroTermsQueryParam) SetZeroTermsQuery(v ZeroTermsQuery) {
	ztq.ZeroTermsQueryValue = &v
}
func unmarshalZeroTermsQueryParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithZeroTermsQuery); ok {
		a.SetZeroTermsQuery(ZeroTermsQuery(value.Str))
	}
	return nil
}
