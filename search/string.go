package search

import (
	"strings"

	"github.com/chanced/dynamic"
)

type String string

func (s String) String() string {
	return string(s)
}

// TODO: Split string

func (s String) Terms() (*termsClause, error) {
	strs := strings.Split(s.String(), ",")
	for i, str := range strs {
		strs[i] = strings.TrimSpace(str)
	}
	q := &termsClause{
		TermsValue: strs,
	}
	return q, nil
}

type Strings []string

func (s Strings) Terms() (*termsClause, error) {
	q := &termsClause{
		TermsValue: s,
	}
	return q, nil
}

type number dynamic.Number

// Number returns a new DynamicNumber It panics if v can not be set to a dynamic.Number
//
// see https://github.com/chanced/dynamic/blob/main/number.go
func Number(v interface{}) *number {
	n := dynamic.NewNumber(v)
	dn := number(n)
	return &dn
}
