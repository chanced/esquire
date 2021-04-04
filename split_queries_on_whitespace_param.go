package picker

import "github.com/chanced/dynamic"

const DefaultSplitQueriesOnWhitespace = false

// WithSplitQueriesOnWhitespace is mapping with the split_queries_on_whitespace
// parameter
//
// split_queries_on_whitespace determines whether full text queries should split
// the input on whitespace when building a query for this field. Accepts true or
// false (default).
type WithSplitQueriesOnWhitespace interface {
	// SplitQueriesOnWhitespace determines whether full text queries should
	// split the input on whitespace when building a query for this field.
	// Accepts true or false (default).
	SplitQueriesOnWhitespace() bool
	// SetSplitQueriesOnWhitespace sets the SplitQueriesOnWhitespace Value to v
	SetSplitQueriesOnWhitespace(v interface{}) error
}

type splitQueriesOnWhitespaceParam struct {
	splitQueriesOnWhitespace dynamic.Bool
}

// SplitQueriesOnWhitespace determines whether full text queries should split
// the input on whitespace when building a query for this field. Accepts true or
// false (default).
func (sq splitQueriesOnWhitespaceParam) SplitQueriesOnWhitespace() bool {
	if b, ok := sq.splitQueriesOnWhitespace.Bool(); ok {
		return b
	}
	return DefaultSplitQueriesOnWhitespace
}

// SetSplitQueriesOnWhitespace sets the SplitQueriesOnWhitespace Value to v
func (sq *splitQueriesOnWhitespaceParam) SetSplitQueriesOnWhitespace(v interface{}) error {
	return sq.splitQueriesOnWhitespace.Set(v)
}
