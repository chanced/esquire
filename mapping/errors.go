package mapping

import "errors"

// Errors
var (
	ErrMetaLimitExceeded          = errors.New("picker: metadata limit exceeded")
	ErrInvalidDynamicParam        = errors.New("picker: invalid dynamic parameter value")
	ErrInvalidIndexOptionsParam   = errors.New("picker: invalid index options parameter value")
	ErrInvalidIndexPrefixMaxChars = errors.New("picker: invalid index prefix max chars")
	ErrInvalidIndexPrefixMinChars = errors.New("picker: invalid index prefix min chars")
	ErrPathRequired               = errors.New("picker: Path is required")
)

type FieldError struct {
}
