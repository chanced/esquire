package search

import (
	"errors"
	"strings"
)

type QueryError struct {
	Field     string
	Err       error
	QueryType QueryType
}

func NewQueryError(err error, queryType QueryType) *QueryError {
	return &QueryError{
		Err:       err,
		QueryType: queryType,
	}
}

func (s QueryError) Error() string {
	// TODO: clean this error message up
	b := strings.Builder{}
	b.WriteString(s.Err.Error())
	b.WriteString(" for ")
	b.WriteString(s.QueryType.String())
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
