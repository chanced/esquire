package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultCaseInsensitive = false

// WithCaseInsensitive is a query mixin that adds the case_insensitive param
//
// (Optional, Boolean) Allows ASCII case insensitive matching of the value with
// the indexed field values when set to true. Default is false which means the
// case sensitivity of matching depends on the underlying field’s mapping.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type WithCaseInsensitive interface {
	CaseInsensitive() bool
	SetCaseInsensitive(v bool)
}

// caseInsensitiveParam is a query mixin that adds the case_insensitive param
//
// (Optional, Boolean) Allows ASCII case insensitive matching of the value with
// the indexed field values when set to true. Default is false which means the
// case sensitivity of matching depends on the underlying field’s mapping.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-term-query.html
type caseInsensitiveParam struct {
	caseInsensitive *bool
}

func (ci caseInsensitiveParam) CaseInsensitive() bool {
	if ci.caseInsensitive == nil {
		return DefaultCaseInsensitive
	}
	return *ci.caseInsensitive
}

func (ci *caseInsensitiveParam) SetCaseInsensitive(v bool) {
	if ci.CaseInsensitive() != v {
		ci.caseInsensitive = &v
	}
}

func unmarshalCaseInsensitiveParam(data dynamic.RawJSON, target interface{}) error {
	if r, ok := target.(WithCaseInsensitive); ok {
		if data.IsNull() {
			return nil
		}
		var b bool
		if data.IsBool() {
			err := json.Unmarshal(data, &b)
			if err != nil {
				return err
			}
		}
		if data.IsString() {
			if data.UnquotedString() == "" {
				return nil
			}
			n := dynamic.NewBool(data.UnquotedString())
			v, ok := n.Bool()
			if !ok {
				return &json.UnmarshalTypeError{Value: data.UnquotedString()}
			}
			r.SetCaseInsensitive(v)
			return nil
		}
		return &json.UnmarshalTypeError{Value: data.UnquotedString()}
	}
	return nil
}

func marshalCaseInsensitiveParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithCaseInsensitive); ok {
		if b.CaseInsensitive() {
			data[paramCaseInsensitive] = b.CaseInsensitive()
		}
	}
	return data, nil
}
