package search

// WithMaxExpansions is a query with the max_expansions param
//
// Maximum number of variations created. Defaults to 50.
//
// WARNING
//
// Avoid using a high value in the max_expansions parameter, especially if the
// prefix_length parameter value is 0. High values in the max_expansions
// parameter can cause poor performance due to the high number of variations
// examined.
type WithMaxExpansions interface {
	// MaxExpansions is the maximum number of variations created. Defaults to 50.
	MaxExpansions() int
	SetMaxExpansions(v int)
}

// MaxExpansionsParam is a mixin that adds the max_expansions param to queries
//
// Maximum number of variations created. Defaults to 50.
//
// WARNING
//
// Avoid using a high value in the max_expansions parameter, especially if the
// prefix_length parameter value is 0. High values in the max_expansions
// parameter can cause poor performance due to the high number of variations
// examined.
type MaxExpansionsParam struct {
	MaxExpansionsValue *int `json:"max_expansions,omitempty" bson:"max_expansions,omitempty"`
}

// MaxExpansions is the maximum number of variations created. Defaults to 50.
func (me MaxExpansionsParam) MaxExpansions() int {
	if me.MaxExpansionsValue == nil {
		return 50
	}
	return *me.MaxExpansionsValue
}
func (me *MaxExpansionsParam) SetMaxExpansions(v int) {
	me.MaxExpansionsValue = &v
}
