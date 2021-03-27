package search

import (
	"errors"
	"strings"
)

type QueryError struct {
	Field string
	Err   error
	Type  Type
}

type RuleError struct {
	*QueryError
	Rule Clause
}

func NewRuleError(err error, queryType Type, rule Clause, field ...string) *RuleError {
	return &RuleError{
		Rule:       rule,
		QueryError: NewQueryError(err, queryType, field...),
	}
}
func NewQueryError(err error, queryType Type, field ...string) *QueryError {
	var f string
	if len(field) > 0 {
		f = field[0]
	}
	return &QueryError{
		Err:   err,
		Type:  queryType,
		Field: f,
	}
}

func (s QueryError) Error() string {
	// TODO: clean this error message up
	b := strings.Builder{}
	b.WriteString(s.Err.Error())
	b.WriteString(" for ")
	b.WriteString(s.Type.String())
	if s.Field != "" {
		b.WriteString(" <")
		b.WriteString(s.Field)
		b.WriteRune('>')
	}
	return b.String()
}
func (s QueryError) Unwrap() error {
	return s.Err
}

var (
	ErrFieldRequired     = errors.New("error: field is required")
	ErrValueRequired     = errors.New("error: value is required")
	ErrQueryRequired     = errors.New("error: query is required")
	ErrInvalidSourceType = errors.New("error: invalid source type")
	ErrInvalidRewrite    = errors.New("error: invalid rewrite value")
	ErrFieldExists       = errors.New("error: field exists")
	ErrTypeRequired      = errors.New("error: rule type is required")
	ErrUnsupportedType   = errors.New("error: unsupported rule type")
)
