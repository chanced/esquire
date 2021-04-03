package picker

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
	SetEnablePositionIncrements(v bool)
}

// FieldWithEnablePositionIncrements is a Field with the
// enable_position_increments parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/token-count.html#token-count-params
type FieldWithEnablePositionIncrements interface {
	Field
	WithIncludeInRoot
}

// EnablePositionIncrementsParam is a mixin that adds the enable_position_increments param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html#nested-params
type EnablePositionIncrementsParam struct {
	EnablePositionIncrementsValue *bool `bson:"enable_position_increments,omitempty" json:"enable_position_increments,omitempty"`
}

// EnablePositionIncrements deteremines if all fields in the nested object are also
// added to the root document as standard (flat) fields. Defaults to false
func (epi EnablePositionIncrementsParam) EnablePositionIncrements() bool {
	if epi.EnablePositionIncrementsValue == nil {
		return false
	}
	return *epi.EnablePositionIncrementsValue
}

// SetEnablePositionIncrements sets the EnablePositionIncrements Value to v
func (epi *EnablePositionIncrementsParam) SetEnablePositionIncrements(v bool) {
	if epi.EnablePositionIncrements() != v {
		epi.EnablePositionIncrementsValue = &v
	}
}
