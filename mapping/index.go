package mapping

// WithIndex is a mapping with index parameter
//
// The index parameter controls whether field values are indexed. It accepts
// true or false and defaults to true. Fields that are not indexed are not
// queryable.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-index.html
type WithIndex interface {
	Index() bool
	SetIndex(v bool)
}

// FieldWithIndex is a Field with the index parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-index.html
type FieldWithIndex interface {
	Field
	WithIndex
}

// IndexParam is a mapping mixin that adds the index parameter
//
// The index parameter controls whether field values are indexed. It accepts
// true or false and defaults to true. Fields that are not indexed are not
// queryable.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-index.html
type IndexParam struct {
	IndexValue *bool `bson:"index,omitempty" json:"index,omitempty"`
}

// Index controls whether field values are indexed. It accepts true or false and
// defaults to true. Fields that are not indexed are not queryable.
func (i IndexParam) Index() bool {
	if i.IndexValue == nil {
		return true
	}
	return *i.IndexValue
}

// SetIndex sets the Index value to v
//
// Index controls whether field values are indexed. It accepts true or false and
// defaults to true. Fields that are not indexed are not queryable.
func (i *IndexParam) SetIndex(v bool) {
	i.IndexValue = &v
}
