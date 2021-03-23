package search

type Relation string

const (
	//RelationIntersects matches documents with a range field value that
	//intersects the query’s range.
	RelationIntersects Relation = "INTERSECTS"
	// RelationContains matches documents with a range field value that entirely
	// contains the query’s range.
	RelationContains Relation = "CONTAINS"
	// RelationWithin matches documents with a range field value entirely within
	// the query’s range.
	RelationWithin Relation = "WITHIN"
)

// RelationParam is a mixin that adds the relation parameter
type RelationParam struct {
	RelationValue Relation `json:"relation,omitempty" bson:"relation,omitempty"`
}

// Relation indicates how the range query matches values for range fields.
func (r RelationParam) Relation() Relation {
	if r.RelationValue == "" {
		return RelationIntersects
	}
	return r.RelationValue
}

// SetRelation sets Relation to v
func (r *RelationParam) SetRelation(v Relation) {
	if r.Relation() != v {
		r.RelationValue = v
	}
}
