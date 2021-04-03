package picker

import "github.com/chanced/dynamic"

const DefaultStore = false

// WithStore is a mapping with a store paramter.
//
// By default, field values are indexed to make them searchable, but they are
// not stored. This means that the field can be queried, but the original field
// value cannot be retrieved.
//
// Usually this doesn’t matter. The field value is already part of the _source
// field, which is stored by default. If you only want to retrieve the value of
// a single field or of a few fields, instead of the whole _source, then this
// can be achieved with source filtering.
//
// In certain situations it can make sense to store a field. For instance, if
// you have a document with a title, a date, and a very large content field, you
// may want to retrieve just the title and the date without having to extract
// those fields from a large _source field
//
// Stored fields returned as arrays
//
// For consistency, stored fields are always returned as an array because there
// is no way of knowing if the original field value was a single value, multiple
// values, or an empty array.
//
// The original value can be retrieved from the _source field instead.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-store.html
type WithStore interface {
	Store() bool
	SetStore(v interface{}) error
}

// FieldWithStore is a Field with a Store attribute
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-store.html
type FieldWithStore interface {
	Field
	WithStore
}

type storeParam struct {
	store dynamic.Bool
}

// Store returns the StoreAttr Value or false
func (sa storeParam) Store() bool {
	if v, ok := sa.store.Bool(); ok {
		return v
	}
	return DefaultStore
}

// SetStore sets StoreAttr Value to v
func (sa *storeParam) SetStore(v interface{}) error {
	return sa.store.Set(v)
}
