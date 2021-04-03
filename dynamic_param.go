package picker

import (
	"fmt"
	"strings"

	"github.com/chanced/dynamic"
)

const DefaultDynamic = DynamicTrue

const (
	DynamicUnspecified Dynamic = ""
	DynamicTrue        Dynamic = "true"
	DynamicFalse       Dynamic = "false"
	DynamicRuntime     Dynamic = "runtime"
	DynamicStrict      Dynamic = "strict"
)

type Dynamic string

func (d Dynamic) String() string {
	return string(d)
}

func (d *Dynamic) Validate() error {
	if !d.IsValid() {
		strs := make([]string, len(dynamicOpts)+1)
		strs[0] = `""`
		for i, v := range dynamicOpts {
			strs[i+1] = `"` + v.String() + `"`
		}
		return fmt.Errorf("%w <%s> expected one of [%s]", ErrInvalidDynamic, *d, strings.Join(strs, ", "))
	}
	return nil
}
func (d *Dynamic) IsValid() bool {
	for _, v := range dynamicOpts {
		if *d == v {
			return true
		}
	}
	*d = Dynamic(strings.ToLower(string(*d)))
	for _, v := range dynamicOpts {
		if *d == v {
			return true
		}
	}
	return false
}

var dynamicOpts = []Dynamic{
	DynamicUnspecified,
	DynamicTrue,
	DynamicFalse,
	DynamicRuntime,
	DynamicStrict,
}

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

// dynamicParam is a mixin for mappings with the Dynamic param
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
type dynamicParam struct {
	dynamic Dynamic
}

// Dynamic determines whether or not new properties should be added dynamically to
// an existing object. Accepts true (default), false and strict.
func (dp dynamicParam) Dynamic() Dynamic {
	if len(dp.dynamic) == 0 {
		return DefaultDynamic
	}
	return dp.dynamic
}

// SetDynamic sets the value of Dynamic to v.
func (dp *dynamicParam) SetDynamic(dynamic Dynamic) error {
	err := dynamic.Validate()
	if err != nil {
		return err
	}
	dp.dynamic = dynamic
	return nil
}
