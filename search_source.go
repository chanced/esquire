package picker

import (
	"fmt"

	"github.com/chanced/dynamic"
)

type SourceSpecifications struct {
	Includes dynamic.StringOrArrayOfStrings
	Excludes dynamic.StringOrArrayOfStrings
}

type SearchSourceParams struct{}

type SearchSource struct {
	boolean         *bool
	wildcardPattern dynamic.StringOrArrayOfStrings
	specifications  *SourceSpecifications
}

// SetValue sets the value of Source
//
// The options are:
//  string
//  []string
//  dynamic.StringOrArrayOfStrings
//  *dynamic.StringOrArrayOfStrings
//  SourceSpecifications
//  *SourceSpecifications
//  bool
//  *bool
//  nil
//
// SetValue returns an error if v is not one of the types listed above.
func (s *SearchSource) SetValue(v interface{}) error {

	switch t := v.(type) {
	case *string:
		if *t == "true" {
			return s.SetValue(true)
		}
		if *t == "false" {
			return s.SetValue(false)
		}
		s.wildcardPattern = dynamic.StringOrArrayOfStrings{*t}
		s.boolean = nil
		s.specifications = nil
	case string:
		if t == "true" {
			return s.SetValue(true)
		}
		if t == "false" {
			return s.SetValue(false)
		}
		s.wildcardPattern = dynamic.StringOrArrayOfStrings{t}
		s.boolean = nil
		s.specifications = nil
	case []string:
		s.wildcardPattern = dynamic.StringOrArrayOfStrings{}
		s.boolean = nil
		s.specifications = nil
		return s.wildcardPattern.Set(t)
	case dynamic.StringOrArrayOfStrings:
		s.boolean = nil
		s.specifications = nil
		s.wildcardPattern = t
	case *dynamic.StringOrArrayOfStrings:
		s.boolean = nil
		s.specifications = nil
		s.wildcardPattern = *t
	case SourceSpecifications:
		s.boolean = nil
		s.wildcardPattern = nil
		s.specifications = &t
	case *SourceSpecifications:
		s.boolean = nil
		s.wildcardPattern = nil
		s.specifications = t
	case bool:
		s.boolean = &t
		s.wildcardPattern = nil
		s.specifications = nil
	case *bool:
		s.boolean = t
		s.wildcardPattern = nil
		s.specifications = nil
	case nil:
		s.boolean = nil
		s.wildcardPattern = nil
		s.specifications = nil
	default:
		return fmt.Errorf("%w: %t", ErrInvalidSource, v)
	}
	return nil
}

// Value indicates which source fields are returned for matching documents.
// These fields are returned in the hits._source property of the search
// response. Defaults to true.
func (s SearchSource) Value() interface{} {
	if s.boolean != nil {
		return s.boolean
	}
	if s.specifications != nil {
		return s.specifications
	}
	if s.wildcardPattern != nil {
		return s.wildcardPattern
	}
	return true
}
