package mapping

import (
	"fmt"
	"strings"

	"github.com/chanced/dynamic"
)

var (
	dynamicOpts    = []string{"true", "false", "runtime", "strict"}
	dynamicOptsStr = strings.Join(dynamicOpts, ", ")
)

// FieldWithDynamic is a field with a Dynamic param
//
// Dynamic determines whether or not new properties should be added dynamically to
// an existing object. Inner objects inherit the dynamic setting from their parent
//object or from the mapping type.
//
// true - New fields are added to the mapping (default).
//
// false - New fields are ignored. These fields will not be indexed or searchable
// but will still appear in the _source field of returned hits. These fields will not
// be added to the mapping, and new fields must be added explicitly.
//
// strict - If new fields are detected, an exception is thrown and the document is
// rejected. New fields must be explicitly added to the mapping.
//
// runtime - New fields are added to the mapping as runtime fields. These fields are
// not indexed, and are loaded from _source at query time.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dynamic.html
type FieldWithDynamic interface {
	Dynamic() dynamic.BoolOrString
	// SetDynamic either accepts a string or a bool, returning an error otherwise
	SetDynamic(v interface{}) error
}

// DynamicParam is a mixin for mappings with the Dynamic param
//
// Dynamic determines whether or not new properties should be added dynamically to
// an existing object. Inner objects inherit the dynamic setting from their parent
//object or from the mapping type.
//
// true - New fields are added to the mapping (default).
//
// false - New fields are ignored. These fields will not be indexed or searchable
// but will still appear in the _source field of returned hits. These fields will not
// be added to the mapping, and new fields must be added explicitly.
//
// strict - If new fields are detected, an exception is thrown and the document is
// rejected. New fields must be explicitly added to the mapping.
//
// runtime - New fields are added to the mapping as runtime fields. These fields are
// not indexed, and are loaded from _source at query time.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/dynamic.html
type DynamicParam struct {
	DynamicParamValue dynamic.BoolOrString `bson:"dynamic,omitempty" json:"dynamic,omitempty"`
}

// Dynamic determines whether or not new properties should be added dynamically to
// an existing object. Accepts true (default), false and strict.
func (dp DynamicParam) Dynamic() dynamic.BoolOrString {
	if dp.DynamicParamValue.String() == "" {
		return dynamic.BoolOrString("true")
	}
	return dp.DynamicParamValue
}

// SetDynamic sets the value of Dynamic to v.
func (dp *DynamicParam) SetDynamic(v interface{}) error {
	if dyn, ok := v.(*dynamic.BoolOrString); ok {
		v = dyn.String()
	}
	if dyn, ok := v.(dynamic.BoolOrString); ok {
		v = dyn.String()
	}
	if str, ok := v.(string); ok {
		str = strings.ToLower(str)
		for _, o := range dynamicOpts {
			if str == o {
				return dp.DynamicParamValue.Set(str)
			}
		}
		return fmt.Errorf("%w: expected one of: %s; received %s", ErrInvalidDynamicParam, dynamicOptsStr, str)
	}
	return dp.DynamicParamValue.Set(v)
}
