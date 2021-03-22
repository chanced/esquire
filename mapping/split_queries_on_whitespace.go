package mapping

// WithSplitQueriesOnWhitespace is mapping with the split_queries_on_whitespace
// parameter
//
// split_queries_on_whitespace determines whether full text queries should split
// the input on whitespace when building a query for this field. Accepts true or
// false (default).
type WithSplitQueriesOnWhitespace interface {
	// SplitQueriesOnWhitespace determines whether full text queries should split
	// the input on whitespace when building a query for this field. Accepts true or
	// false (default).
	SplitQueriesOnWhitespace() bool
	// SetSplitQueriesOnWhitespace sets the SplitQueriesOnWhitespace Value to v
	SetSplitQueriesOnWhitespace(v bool)
}

// FieldWithSplitQueriesOnWhitespace is a Field with the
// split_queries_on_whitespace param
type FieldWithSplitQueriesOnWhitespace interface {
	Field
	WithSplitQueriesOnWhitespace
}

// SplitQueriesOnWhitespaceParam is a mixin that adds the
// split_queries_on_whitespace paramete
type SplitQueriesOnWhitespaceParam struct {
	SplitQueriesOnWhitespaceValue *bool `bson:"split_queries_on_whitespace,omitempty" json:"split_queries_on_whitespace,omitempty"`
}

// SplitQueriesOnWhitespace determines whether full text queries should split
// the input on whitespace when building a query for this field. Accepts true or
// false (default).
func (sq SplitQueriesOnWhitespaceParam) SplitQueriesOnWhitespace() bool {
	if sq.SplitQueriesOnWhitespaceValue == nil {
		return false
	}
	return *sq.SplitQueriesOnWhitespaceValue
}

// SetSplitQueriesOnWhitespace sets the SplitQueriesOnWhitespace Value to v
func (sq *SplitQueriesOnWhitespaceParam) SetSplitQueriesOnWhitespace(v bool) {
	if sq.SplitQueriesOnWhitespace() != v {
		sq.SplitQueriesOnWhitespaceValue = &v
	}
}
