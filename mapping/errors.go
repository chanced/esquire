package mapping

import "errors"

// Errors
var (
	ErrMetaLimitExceeded          = errors.New("limit of 5 metadata entries exceeded")
	ErrInvalidDynamicParam        = errors.New("invalid dynamic parameter value")
	ErrInvalidIndexOptionsParam   = errors.New("invalid index options parameter value")
	ErrInvalidIndexPrefixMaxChars = errors.New("invalid index prefix max chars")
	ErrInvalidIndexPrefixMinChars = errors.New("invalid index prefix min chars")
)
