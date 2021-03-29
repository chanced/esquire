package search

import (
	"errors"
	"strings"
)

var (
	ErrFieldRequired        = errors.New("picker: field is required")
	ErrValueRequired        = errors.New("picker: value is required")
	ErrQueryRequired        = errors.New("picker: query is required")
	ErrInvalidSourceType    = errors.New("picker: invalid source type")
	ErrInvalidRewrite       = errors.New("picker: invalid rewrite value")
	ErrFieldExists          = errors.New("picker: field exists")
	ErrTypeRequired         = errors.New("picker: rule type is required")
	ErrUnsupportedType      = errors.New("picker: unsupported rule type")
	ErrPathRequired         = errors.New("picker: Path required in lookup")
	ErrIDRequired           = errors.New("picker: ID required for lookup")
	ErrIndexRequired        = errors.New("picker: Index required")
	ErrInvalidBoost         = errors.New("picker: invalid boost value")
	ErrInvalidMaxExpansions = errors.New("picker: invalid max expansions")
	ErrInvalidPrefixLength  = errors.New("picker: invalid prefix length")
	ErrInvalidZeroTermQuery = errors.New("picker: invalid zero terms query")
	ErrInvalidRelation      = errors.New("picker: invalid relation")
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

func NewQueryError(err error, queryType Type, field ...string) *QueryError {
	var f string
	if len(field) > 0 {
		f = field[0]
	}
	var qe *QueryError
	if errors.As(err, &qe) {
		if len(f) != 0 {
			qe.Field = f
		}
		if qe.Type != queryType {
			qe.Type = queryType
		}
		return qe
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
