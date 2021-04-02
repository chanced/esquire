package picker

import "github.com/chanced/dynamic"

// Relations are used on JoinFields
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type Relations map[string]dynamic.StringOrArrayOfStrings

func (r Relations) Clone() Relations {
	res := Relations{}
	for k, v := range r {
		if v != nil {
			res[k] = append(dynamic.StringOrArrayOfStrings{}, v...)
		}
	}
	return res
}

// WithRelations is a mapping with a relations parameter
//
// The relations section defines a set of possible relations within the
// documents, each relation being a parent name and a child name.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type WithRelations interface {
	Relations() Relations
	SetRelations(v Relations)
	AddRelation(key string, value dynamic.StringOrArrayOfStrings)
	RemoveRelation(v string)
	ClearRelations()
}

// FieldWithRelations is a Field with the relations parameter
//
//https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type FieldWithRelations interface {
	Field
	WithRelations
}

// RelationsParam is a mixin that adds the relations parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type RelationsParam struct {
	RelationsValue Relations `bson:"relations,omitempty" json:"relations,omitempty"`
}

// Relations defines a set of possible relations within the documents, each
// relation being a parent name and a child name. A parent/child relation can be
// defined as follows:
func (r RelationsParam) Relations() Relations {
	return r.RelationsValue
}

// SetRelations sets the value of Relations to v
func (r *RelationsParam) SetRelations(v Relations) {

	r.RelationsValue = v
}

// AddRelation adds the key to the Relations map, set to value
func (r *RelationsParam) AddRelation(key string, value dynamic.StringOrArrayOfStrings) {
	if r.RelationsValue == nil {
		r.RelationsValue = Relations{}
	}
	r.RelationsValue[key] = value
}

// RemoveRelation deletes the key from the Relations map
func (r *RelationsParam) RemoveRelation(key string) {
	if r.RelationsValue == nil {
		return
	}
	delete(r.RelationsValue, key)
}

// ClearRelations sets Value to a new, empty Relations map
func (r *RelationsParam) ClearRelations() {
	r.RelationsValue = Relations{}
}
