package picker

import (
	"strings"

	"github.com/chanced/dynamic"
)

type String string

func (s String) String() string {
	return string(s)
}

func (s String) Match() (MatchClause, error) {
	q := MatchClause{}
	err := q.setQuery(s.String())
	return q, err
}

func (s String) Term() (TermClause, error) {
	q := TermClause{}
	err := q.SetValue(s.String())
	return q, err
}

func (s String) Terms() (TermsClause, error) {
	q := TermsClause{}
	strs := strings.Split(s.String(), ",")
	for i, str := range strs {
		str = strings.TrimSpace(str)
		if str != "" {
			strs[i] = str
		}
	}
	err := q.setValue(strs)
	return q, err
}

type Strings []string

func (s Strings) Terms() (TermsClause, error) {
	q := TermsClause{}
	err := q.setValue(s)
	return q, err
}

type number dynamic.Number

// Number returns a new DynamicNumber It panics if v can not be set to a dynamic.Number
//
// see https://github.com/chanced/dynamic/blob/main/number.go
func Number(n dynamic.Number) *number {
	dn := number(n)
	return &dn
}
