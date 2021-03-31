package picker

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

func unmarshalCaseInsensitiveParam(data dynamic.JSON, target interface{}) error {
	if p, ok := target.(WithCaseInsensitive); ok {
		var b dynamic.Bool
		var err error
		switch {
		case data.IsNull():
		case data.IsBool():
			b, err = dynamic.NewBool(string(data))
		case data.IsString():
			b, err = dynamic.NewBool(data.UnquotedString())
		default:
			err = &json.UnmarshalTypeError{Value: string(data), Type: typeBool}
		}
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			p.SetCaseInsensitive(v)
		}
		return nil

	}
	return nil
}

func marshalCaseInsensitiveParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithCaseInsensitive); ok {
		if b.CaseInsensitive() {
			return json.Marshal(b.CaseInsensitive())
		}
	}
	return nil, nil
}
