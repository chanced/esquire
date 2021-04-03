package picker

import (
	"strings"

	"encoding/json"

	"github.com/chanced/dynamic"
)

type Relation string

func (r Relation) String() string {
	return string(r)
}

func (r Relation) IsValid() bool {
	for _, v := range relationValues {
		if strings.ToUpper(r.String()) == string(v) {
			return true
		}
	}
	return false
}

func (r *Relation) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	*r = Relation(strings.ToUpper(str))
	return nil
}

func (r Relation) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToUpper(r.String()))
}

const DefaultRelation = RelationIntersects

const (
	RelationUnspecified Relation = ""
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

var relationValues = []Relation{RelationContains, RelationWithin, RelationIntersects, RelationUnspecified}

// relationParam is a mixin that adds the relation parameter
type relationParam struct {
	relation Relation
}

type WithRelation interface {
	Relation() Relation
	SetRelation(v Relation) error
}

// Relation indicates how the range query matches values for range fields.
func (r relationParam) Relation() Relation {
	if r.relation == "" {
		return DefaultRelation
	}
	return r.relation
}

// SetRelation sets Relation to v
func (r *relationParam) SetRelation(v Relation) error {
	if !v.IsValid() {
		return ErrInvalidRelation
	}
	r.relation = v
	return nil
}
func unmarshalRelationParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithRelation); ok {
		r := Relation(strings.ToUpper(data.UnquotedString()))
		return a.SetRelation(r)
	}
	return nil
}

func marshalRelationParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithRelation); ok {
		if b.Relation() != DefaultRelation {
			return json.Marshal(b.Relation())
		}
	}
	return nil, nil
}
