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

func NewQueryError(err error, queryType Type) *QueryError {
	return &QueryError{
		Err:  err,
		Type: queryType,
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
	ErrMissingValue      = errors.New("error: value is required")
	ErrMissingQuery      = errors.New("error: query is required")
	ErrInvalidSourceType = errors.New("error: invalid source type")
	ErrInvalidRewrite    = errors.New("error: invalid rewrite value")
	ErrFieldExists       = errors.New("error: field exists")
)
