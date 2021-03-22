package search

// Exists returns documents that contain an indexed value for a field.
//
// An indexed value may not exist for a document’s field due to a variety of
// reasons:
//
// - The field in the source JSON is null or []
//
// - The field has "index" : false set in the mapping
//
// - The length of the field value exceeded an ignore_above setting in the
// mapping
//
// - The field value was malformed and ignore_malformed was defined in the
// mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type Exists struct {
	Field string `json:"field" bson:"field"`
}

// ExistsQuery returns documents that contain an indexed value for a field.
//
// An indexed value may not exist for a document’s field due to a variety of
// reasons:
//
// - The field in the source JSON is null or []
//
// - The field has "index" : false set in the mapping
//
// - The length of the field value exceeded an ignore_above setting in the
// mapping
//
// - The field value was malformed and ignore_malformed was defined in the
// mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-exists-query.html
type ExistsQuery struct {
	ExistsValue *Exists `json:"exists,omitempty" bson:"exists,omitempty"`
}

func (e *ExistsQuery) SetExistsField(field string) {
	if e.ExistsValue == nil {
		e.ExistsValue = &Exists{}
	}
	e.ExistsValue.Field = field
}
