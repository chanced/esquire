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

func (s String) Terms() (*TermsRule, error) {
	strs := strings.Split(s.String(), ",")
	for i, str := range strs {
		strs[i] = strings.TrimSpace(str)
	}
	q := &TermsRule{
		TermsValue: strs,
	}
	return q, nil
}

type Strings []string

func (s Strings) Terms() (*TermsRule, error) {
	q := &TermsRule{
		TermsValue: s,
	}
	return q, nil
}

type DynamicNumber dynamic.Number

// Number returns a new DynamicNumber It panics if v can not be set to a dynamic.Number
//
// see https://github.com/chanced/dynamic/blob/main/number.go
func Number(v interface{}) *DynamicNumber {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		panic(err)
	}
	dn := DynamicNumber(*n)
	return &dn
}
