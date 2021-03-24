package search

import "github.com/tidwall/gjson"

type Relation string

func (r Relation) String() string {
	return string(r)
}

const DefaultRelation = RelationIntersects

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

type WithRelation interface {
	Relation() Relation
	SetRelation(v Relation)
}

// Relation indicates how the range query matches values for range fields.
func (r RelationParam) Relation() Relation {
	if r.RelationValue == "" {
		return DefaultRelation
	}
	return r.RelationValue
}

// SetRelation sets Relation to v
func (r *RelationParam) SetRelation(v Relation) {
	if r.Relation() != v {
		r.RelationValue = v
	}
}
func unmarshalRelationParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithRelation); ok {
		a.SetRelation(Relation(value.String()))
	}
	return nil
}

func marshalRelationParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithRelation); ok {
		if b.Relation() != DefaultRelation {
			data[paramRelation] = b.Relation()
		}
	}
	return data, nil
}
