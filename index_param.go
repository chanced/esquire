package picker

import "github.com/chanced/dynamic"

const DefaultIndex = true

// WithIndex is a mapping with index parameter
//
// The index parameter controls whether field values are indexed. It accepts
// true or false and defaults to true. Fields that are not indexed are not
// queryable.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-index.html
type WithIndex interface {
	Index() bool
	SetIndex(v interface{}) error
}

// FieldWithIndex is a Field with the index parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-index.html
type FieldWithIndex interface {
	Field
	WithIndex
}

type indexParam struct {
	index dynamic.Bool
}

// Index controls whether field values are indexed. It accepts true or false and
// defaults to true. Fields that are not indexed are not queryable.
func (i indexParam) Index() bool {
	if v, ok := i.index.Bool(); ok {
		return v
	}
	return DefaultIndex
}

// SetIndex sets the Index value to v
//
// Index controls whether field values are indexed. It accepts true or false and
// defaults to true. Fields that are not indexed are not queryable.
func (i *indexParam) SetIndex(v interface{}) error {
	return i.index.Set(v)
}
