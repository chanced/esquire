package search

import (
	"strings"

	"github.com/chanced/dynamic"
)

const DefaultZeroTermsQuery = ZeroTermsQueryNone

type ZeroTermsQuery string

func (ztq ZeroTermsQuery) String() string {
	return string(ztq)
}
func (ztq ZeroTermsQuery) toLower() ZeroTermsQuery {
	return ZeroTermsQuery(strings.ToLower(string(ztq)))
}
func (ztq ZeroTermsQuery) ref() *ZeroTermsQuery {
	return &ztq
}

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

type zeroTermsQueryParam struct {
	zeroTermsQuery *ZeroTermsQuery
}

// ZeroTermsQuery indicates  whether no documents are returned if the
// analyzer removes all tokens, such as when using a stop filter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#query-dsl-match-query-zero
func (ztq zeroTermsQueryParam) ZeroTermsQuery() ZeroTermsQuery {
	if ztq.zeroTermsQuery != nil {
		return *ztq.zeroTermsQuery
	}
	return DefaultZeroTermsQuery
}

func (ztq *zeroTermsQueryParam) SetZeroTermsQuery(v ZeroTermsQuery) {
	ztq.zeroTermsQuery = v.toLower().ref()
}
func unmarshalZeroTermsQueryParam(value dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithZeroTermsQuery); ok {
		a.SetZeroTermsQuery(ZeroTermsQuery(value.UnquotedString()))
	}
	return nil
}
func marshalZeroTermsQueryParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithZeroTermsQuery); ok {
		if b.ZeroTermsQuery() != DefaultZeroTermsQuery {
			data[paramZeroTermsQuery] = b.ZeroTermsQuery()
		}
	}
	return data, nil
}
