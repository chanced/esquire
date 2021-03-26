package search

const DefaultMaxExpansions = int64(50)

// WithMaxExpansions is a query with the max_expansions param
//
// Maximum number of variations created. Defaults to 50.
//
// Warning
//
// Avoid using a high value in the max_expansions parameter, especially if the
// prefix_length parameter value is 0. High values in the max_expansions
// parameter can cause poor performance due to the high number of variations
// examined.
type WithMaxExpansions interface {
	// MaxExpansions is the maximum number of variations created. Defaults to 50.
	MaxExpansions() int64
	SetMaxExpansions(v int64)
}

// maxExpansionsParam is a mixin that adds the max_expansions param to queries
//
// Maximum number of variations created. Defaults to 50.
type maxExpansionsParam struct {
	maxExpansions *int64
}

// MaxExpansions is the maximum number of variations created. Defaults to 50.
func (me maxExpansionsParam) MaxExpansions() int64 {
	if me.maxExpansions == nil {
		return DefaultMaxExpansions
	}
	return *me.maxExpansions
}
func (me *maxExpansionsParam) SetMaxExpansions(v int64) {
	if me.MaxExpansions() != v {
		me.maxExpansions = &v
	}
}
