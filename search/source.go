package search

import "github.com/chanced/dynamic"

type SourceSpecifications struct {
	Includes dynamic.StringOrArrayOfStrings `bson:"includes,omitempty" json:"includes,omitempty"`
	Excludes dynamic.StringOrArrayOfStrings `bson:"excludes,omitempty" json:"excludes,omitempty"`
}

type Source struct {
	BoolValue       *bool                          `bson:"bool,omitempty" json:"-"`
	WildcardPattern dynamic.StringOrArrayOfStrings `bson:"wildcard,omitempty" json:"-"`
	Specifications  *SourceSpecifications
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
// SetValue panics if v is not one of the types listed above.
func (s *Source) SetValue(v interface{}) {

	switch t := v.(type) {
	case *string:
		if *t == "true" {
			s.SetValue(true)
			return
		}
		if *t == "false" {
			s.SetValue(false)
			return
		}
		s.WildcardPattern = dynamic.StringOrArrayOfStrings{*t}
		s.BoolValue = nil
		s.Specifications = nil
	case string:
		if t == "true" {
			s.SetValue(true)
			return
		}
		if t == "false" {
			s.SetValue(false)
			return
		}
		s.WildcardPattern = dynamic.StringOrArrayOfStrings{t}
		s.BoolValue = nil
		s.Specifications = nil
	case []string:
		s.WildcardPattern = dynamic.StringOrArrayOfStrings{}
		s.WildcardPattern.Set(t)
		s.BoolValue = nil
		s.Specifications = nil
	case dynamic.StringOrArrayOfStrings:
		s.BoolValue = nil
		s.Specifications = nil
		s.WildcardPattern = t
	case *dynamic.StringOrArrayOfStrings:
		s.BoolValue = nil
		s.Specifications = nil
		s.WildcardPattern = *t
	case SourceSpecifications:
		s.BoolValue = nil
		s.WildcardPattern = nil
		s.Specifications = &t
	case *SourceSpecifications:
		s.BoolValue = nil
		s.WildcardPattern = nil
		s.Specifications = t
	case bool:
		s.BoolValue = &t
		s.WildcardPattern = nil
		s.Specifications = nil
	case *bool:
		s.BoolValue = t
		s.WildcardPattern = nil
		s.Specifications = nil
	case nil:
		s.BoolValue = nil
		s.WildcardPattern = nil
		s.Specifications = nil
	default:
		panic("unknown type")
	}

}

// Value indicates which source fields are returned for matching documents.
// These fields are returned in the hits._source property of the search
// response. Defaults to true.
func (s Source) Value() interface{} {
	if s.BoolValue != nil {
		return s.BoolValue
	}
	if s.Specifications != nil {
		return s.Specifications
	}
	if s.WildcardPattern != nil {
		return s.WildcardPattern
	}
	return true
}
