package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultAllowLeadingWildcard = true

type WithAllowLeadingWildcard interface {
	AllowLeadingWildcard() bool
	SetAllowLeadingWildcard(v interface{}) error
}

type allowLeadingWildcardParam struct {
	allowLeadingWildcard dynamic.Bool
}

func (cp allowLeadingWildcardParam) AllowLeadingWildcard() bool {
	if v, ok := cp.allowLeadingWildcard.Bool(); ok {
		return v
	}
	return DefaultAllowLeadingWildcard
}

func (cp *allowLeadingWildcardParam) SetAllowLeadingWildcard(v interface{}) error {
	return cp.allowLeadingWildcard.Set(v)
}

func unmarshalAllowLeadingWildcardParam(value dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithAllowLeadingWildcard); ok {
		b, err := dynamic.NewBool(value)
		if err != nil {
			return err
		}
		if v, ok := b.Bool(); ok {
			return a.SetAllowLeadingWildcard(v)
		}
	}
	return nil
}
func marshalAllowLeadingWildcardParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithAllowLeadingWildcard); ok {
		if !b.AllowLeadingWildcard() {
			return json.Marshal(b.AllowLeadingWildcard())
		}
	}
	return nil, nil
}
