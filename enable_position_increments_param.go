package picker

import "github.com/chanced/dynamic"

const DefaultEnablePositionIncrements = false

// WithEnablePositionIncrements is a mapping with the enable_position_increments
// parameter
//
// EnablePositionIncrements Indicates if position increments should be counted.
// Set to false if you don’t want to count tokens removed by analyzer filters
// (like stop). Defaults to true.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/token-count.html#token-count-params
type WithEnablePositionIncrements interface {
	//Indicates if position increments should be counted. Set to false if you
	//don’t want to count tokens removed by analyzer filters (like stop).
	//Defaults to true.
	EnablePositionIncrements() bool
	// SetEnablePositionIncrements sets the EnablePositionIncrements Value to v
	SetEnablePositionIncrements(v interface{}) error
}

// enablePositionIncrementsParam is a mixin that adds the enable_position_increments param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type enablePositionIncrementsParam struct {
	enablePositionIncrements dynamic.Bool
}

// EnablePositionIncrements deteremines if all fields in the nested object are also
// added to the root document as standard (flat) fields. Defaults to false
func (epi enablePositionIncrementsParam) EnablePositionIncrements() bool {
	if b, ok := epi.enablePositionIncrements.Bool(); ok {
		return b
	}
	return DefaultEnablePositionIncrements
}

// SetEnablePositionIncrements sets the EnablePositionIncrements Value to v
func (epi *enablePositionIncrementsParam) SetEnablePositionIncrements(v interface{}) error {
	return epi.enablePositionIncrements.Set(v)

}
