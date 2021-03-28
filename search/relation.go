package search

import (
	"github.com/chanced/dynamic"
)

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

// relationParam is a mixin that adds the relation parameter
type relationParam struct {
	relation Relation
}

type WithRelation interface {
	Relation() Relation
	SetRelation(v Relation)
}

// Relation indicates how the range query matches values for range fields.
func (r relationParam) Relation() Relation {
	if r.relation == "" {
		return DefaultRelation
	}
	return r.relation
}

// SetRelation sets Relation to v
func (r *relationParam) SetRelation(v Relation) {
	if r.Relation() != v {
		r.relation = v
	}
}
func unmarshalRelationParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithRelation); ok {
		a.SetRelation(Relation(data.UnquotedString()))
	}
	return nil
}

func marshalRelationParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithRelation); ok {
		if b.Relation() != DefaultRelation {
			data[paramRelation] = b.Relation()
		}
	}
	return data, nil
}
