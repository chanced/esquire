package picker

import "github.com/chanced/dynamic"

// Relations are used on JoinFields
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
//
// you can treat this like a map[string][]string
//  r := picker.Relations{"myField": []string{"abc"}}
//  r["myOtherField"] = []string{"other", "relation"}
type Relations map[string]dynamic.StringOrArrayOfStrings

// WithRelations is a mapping with a relations parameter
//
// The relations section defines a set of possible relations within the
// documents, each relation being a parent name and a child name.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type WithRelations interface {
	Relations() Relations
	SetRelations(v Relations)
}

// relationsParam is a mixin that adds the relations parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type relationsParam struct {
	relations Relations
}

// Relations defines a set of possible relations within the documents, each
// relation being a parent name and a child name.
func (r relationsParam) Relations() Relations {
	return r.relations
}

// SetRelations sets the value of Relations to v
func (r *relationsParam) SetRelations(v Relations) {
	r.relations = v
}
