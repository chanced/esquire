package search

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chanced/dynamic"
)

const DefaultZeroTermsQuery = ZeroTermsNone

type ZeroTerms string

func (ztq *ZeroTerms) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*ztq = ZeroTerms(strings.ToLower(str))
	return nil
}

func (ztq ZeroTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToLower(ztq.String()))
}
func (ztq ZeroTerms) IsEmpty() bool {
	return len(ztq) == 0
}
func (ztq ZeroTerms) IsValid() bool {
	for _, v := range zeroTermsQueryVals {
		if string(strings.ToLower(ztq.String())) == string(v) {
			return true
		}
	}
	return false
}
func (ztq ZeroTerms) String() string {
	return string(ztq)
}
func (ztq ZeroTerms) toLower() ZeroTerms {
	return ZeroTerms(strings.ToLower(string(ztq)))
}
func (ztq ZeroTerms) ref() *ZeroTerms {
	return &ztq
}

const (
	ZeroTermsUnspecified ZeroTerms = ""
	// ZeroTermsNone - No documents are returned if the analyzer removes all
	// tokens.
	ZeroTermsNone ZeroTerms = "none"
	// ZeroTermsALl - Returns all documents, similar to a match_all query.
	ZeroTermsAll ZeroTerms = "all"
)

var zeroTermsQueryVals = []ZeroTerms{ZeroTermsAll, ZeroTermsNone, ZeroTermsUnspecified}

type WithZeroTermsQuery interface {
	// ZeroTermsQuery indicates  whether no documents are returned if the
	// analyzer removes all tokens, such as when using a stop filter
	ZeroTermsQuery() ZeroTerms
	SetZeroTermsQuery(v ZeroTerms) error
}

type zeroTermsQueryParam struct {
	zeroTermsQuery ZeroTerms
}

// ZeroTermsQuery indicates  whether no documents are returned if the
// analyzer removes all tokens, such as when using a stop filter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html#query-dsl-match-query-zero
func (ztq zeroTermsQueryParam) ZeroTermsQuery() ZeroTerms {
	if ztq.zeroTermsQuery != "" {
		return ztq.zeroTermsQuery
	}
	return DefaultZeroTermsQuery
}

func (ztq *zeroTermsQueryParam) SetZeroTermsQuery(v ZeroTerms) error {
	z := v.toLower()
	if !z.IsValid() {
		return fmt.Errorf("%w <%s>", ErrInvalidZeroTermQuery, v)
	}
	ztq.zeroTermsQuery = z
	return nil
}
func unmarshalZeroTermsQueryParam(value dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithZeroTermsQuery); ok {
		z := ZeroTerms(strings.ToLower(value.UnquotedString()))
		return a.SetZeroTermsQuery(z)
	}
	return nil
}
func marshalZeroTermsQueryParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithZeroTermsQuery); ok {
		if b.ZeroTermsQuery() != DefaultZeroTermsQuery && b.ZeroTermsQuery() != "" {
			data[paramZeroTermsQuery] = b.ZeroTermsQuery()
		}
	}
	return data, nil
}
