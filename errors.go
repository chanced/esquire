package picker

import (
	"errors"
	"fmt"
	"strings"
)

var (

	// TODO: need to be either more or less specific with these errors
	ErrKindRequired    = errors.New("picker: rule type is required")
	ErrUnsupportedType = errors.New("picker: unsupported rule type")

	ErrFieldNotFound              = errors.New("picker: field not found")
	ErrFieldRequired              = errors.New("picker: field is required")
	ErrValueRequired              = errors.New("picker: value is required")
	ErrQueryRequired              = errors.New("picker: query is required")
	ErrInvalidSource              = errors.New("picker: invalid value for source")
	ErrInvalidRewrite             = errors.New("picker: invalid value for rewrite")
	ErrFieldExists                = errors.New("picker: field exists")
	ErrPathRequired               = errors.New("picker: path is required")
	ErrIDRequired                 = errors.New("picker: id is required")
	ErrIndexRequired              = errors.New("picker: index is required")
	ErrInvalidBoost               = errors.New("picker: invalid value for boost")
	ErrInvalidMaxExpansions       = errors.New("picker: invalid value for max_expansions")
	ErrInvalidPrefixLength        = errors.New("picker: invalidvalue for prefix_length")
	ErrInvalidZeroTermQuery       = errors.New("picker: invalid value for zero_terms_query")
	ErrInvalidRelation            = errors.New("picker: invalid value for relation")
	ErrWeightRequired             = errors.New("picker: weight is required")
	ErrScriptRequired             = errors.New("picker: script is required")
	ErrInvalidParams              = errors.New("picker: params should marshal into a JSON object")
	ErrOriginRequired             = errors.New("picker: origin is required")
	ErrScaleRequired              = errors.New("picker: scale is required")
	ErrInvalidScoreMode           = errors.New("picker: invalid value for score_mode")
	ErrInvalidBoostMode           = errors.New("picker: invalid value for boost_mode")
	ErrInvalidModifier            = errors.New("picker: invalid value for modifier")
	ErrMetaLimitExceeded          = errors.New("picker: meta limit exceeded")
	ErrInvalidDynamic             = errors.New("picker: invalid value for dynamic")
	ErrInvalidIndexOptions        = errors.New("picker: invalid value for index_options")
	ErrInvalidIndexPrefixMaxChars = errors.New("picker: invalid value for index prefix_max_chars")
	ErrInvalidIndexPrefixMinChars = errors.New("picker: invalid value for index_prefix_min_chars")
	ErrScalingFactorRequired      = errors.New("picker: scaling_factor is required")
	ErrDimensionsRequired         = errors.New("picker: dimensions is required")
	ErrInvalidOrientation         = errors.New("picker: invalid orientation")
	ErrInvalidTermVector          = errors.New("picker: invalid TermVector")
	ErrMissingType                = errors.New("picker: missing type")
	ErrInvalidMaxShingleSize      = errors.New("picker: invalid max_shingle_size; valid values are [2,3,4]")
)

type FieldError struct {
	Field string
	Err   error
}

func (e *FieldError) Error() string {
	return e.Err.Error()
}

func (e *FieldError) Unwrap() error {
	return e.Err
}

type QueryError struct {
	Field string
	Err   error
	Kind  QueryKind
}

func newQueryError(err error, queryKind QueryKind, field ...string) *QueryError {
	var f string
	if len(field) > 0 {
		f = field[0]
	}
	var qe *QueryError
	if errors.As(err, &qe) {
		if len(f) != 0 {
			qe.Field = f
		}
		if qe.Kind != queryKind {
			qe.Kind = queryKind
		}
		return qe
	}
	return &QueryError{
		Err:   err,
		Kind:  queryKind,
		Field: f,
	}
}

func (s QueryError) Error() string {
	// TODO: clean this error message up
	b := strings.Builder{}
	b.WriteString(s.Err.Error())
	b.WriteString(" for ")
	b.WriteString(s.Kind.String())
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

// Implementation of hashicorp's multierror.
type MappingError struct {
	Errors []error
}

func (e *MappingError) assignField(field string) {
	if e == nil {
		return
	}
	for i, err := range e.Errors {
		var fe *FieldError
		if errors.As(err, &fe) {
			if len(fe.Field) == 0 {
				fe.Field = field
			} else if fe.Field != field {
				fe.Field = field + "->" + fe.Field
			}
		} else {

			fe = &FieldError{
				Field: field,
				Err:   err,
			}
		}
		e.Errors[i] = fe
	}
}

func (e *MappingError) Append(err error) {
	if err == nil {
		return
	}
	if merr, ok := err.(*MappingError); ok {
		if merr.Errors == nil {
			return
		}
		e.Errors = append(e.Errors, merr.Errors...)
	}
	e.Errors = append(e.Errors, err)

}
func (e *MappingError) Error() string {
	if len(e.Errors) == 1 {
		return fmt.Sprintf("1 error occurred:\n\t* %s\n\n", e.Errors[0])
	}

	points := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf(
		"%d errors occurred:\n\t%s\n\n",
		len(e.Errors), strings.Join(points, "\n\t"))

}

// ErrorOrNil returns an error interface if this Error represents
// a list of errors, or returns nil if the list of errors is empty. This
// function is useful at the end of accumulation to make sure that the value
// returned represents the existence of errors.
func (e *MappingError) ErrorOrNil() error {
	if e == nil {
		return nil
	}
	if len(e.Errors) == 0 {
		return nil
	}

	return e
}

func (e *MappingError) GoString() string {
	return fmt.Sprintf("*%#v", *e)
}

// WrappedErrors returns the list of errors that this Error is wrapping. It is
// an implementation of the errwrap.Wrapper interface so that multierror.Error
// can be used with that library.
//
// This method is not safe to be called concurrently. Unlike accessing the
// Errors field directly, this function also checks if the multierror is nil to
// prevent a null-pointer panic. It satisfies the errwrap.Wrapper interface.
func (e *MappingError) WrappedErrors() []error {
	if e == nil {
		return nil
	}
	return e.Errors
}

// Unwrap returns an error from Error (or nil if there are no errors).
// This error returned will further support Unwrap to get the next error,
// etc. The order will match the order of Errors in the multierror.Error
// at the time of calling.
//
// The resulting error supports errors.As/Is/Unwrap so you can continue
// to use the stdlib errors package to introspect further.
//
// This will perform a shallow copy of the errors slice. Any errors appended
// to this error after calling Unwrap will not be available until a new
// Unwrap is called on the multierror.Error.
func (e *MappingError) Unwrap() error {
	// If we have no errors then we do nothing
	if e == nil || len(e.Errors) == 0 {
		return nil
	}

	// If we have exactly one error, we can just return that directly.
	if len(e.Errors) == 1 {
		return e.Errors[0]
	}

	// Shallow copy the slice
	errs := make([]error, len(e.Errors))
	copy(errs, e.Errors)
	return errorChain(errs)
}

type errorChain []error

// Error implements the error interface
func (e errorChain) Error() string {
	return e[0].Error()
}

// Unwrap implements errors.Unwrap by returning the next error in the
// chain or nil if there are no more errors.
func (e errorChain) Unwrap() error {
	if len(e) == 1 {
		return nil
	}

	return e[1:]
}

// As implements errors.As by attempting to map to the current value.
func (e errorChain) As(target interface{}) bool {
	return errors.As(e[0], target)
}

// Is implements errors.Is by comparing the current value directly.
func (e errorChain) Is(target error) bool {
	return errors.Is(e[0], target)
}
