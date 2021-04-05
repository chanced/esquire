package picker

import "encoding/json"

type nestedField struct {
	Dynamic         Dynamic     `json:"dynamic,omitempty"`
	Properties      Fieldset    `json:"properties,omitempty"`
	IncludeInParent interface{} `json:"include_in_parent,omitempty"`
	IncludeInRoot   interface{} `json:"include_in_root,omitempty"`
	Type            FieldType   `json:"type"`
}
type NestedFieldParams struct {
	// (Optional) Whether or not new properties should be added dynamically to an existing nested object. Accepts true (default), false and strict.
	Dynamic Dynamic `json:"dynamic,omitempty"`
	// (Optional) The fields within the nested object, which can be of
	// any data type, including nested. New properties may be added to an
	// existing nested object.
	Properties Fieldset `json:"properties,omitempty"`
	// (Optional, Boolean) If true, all fields in the nested object are also
	// added to the parent document as standard (flat) fields. Defaults to
	// false.
	IncludeInParent interface{} `json:"include_in_parent,omitempty"`
	// (Optional, Boolean) If true, all fields in the nested object are also
	// added to the root document as standard (flat) fields. Defaults to false.
	IncludeInRoot interface{} `json:"include_in_root,omitempty"`
}

func (p NestedFieldParams) Nested() (*NestedField, error) {
	f := &NestedField{}
	e := &MappingError{}
	err := f.SetDynamic(p.Dynamic)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIncludeInParent(p.IncludeInParent)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIncludeInRoot(p.IncludeInRoot)
	if err != nil {
		e.Append(err)
	}
	err = f.SetProperties(p.Properties)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()
}
func (p NestedFieldParams) Field() (Field, error) {
	return p.Nested()
}

func (NestedFieldParams) Type() FieldType {
	return FieldTypeNested
}

// NestedField is a specialised version of the object data type that allows
// arrays of objects to be indexed in a way that they can be queried
// independently of each other.
//
// When ingesting key-value pairs with a large, arbitrary set of keys, you might
// consider modeling each key-value pair as its own nested document with key and
// value fields. Instead, consider using the flattened data type, which maps an
// entire object as a single field and allows for simple searches over its
// contents. Nested documents and queries are typically expensive, so using the
// flattened data type for this use case is a better option.
//
// Interacting with nested documents
//
// Nested documents can be:
//
// - queried with the nested query.
//
// - analyzed with the nested and reverse_nested aggregations.
//
// - sorted with nested sorting.
//
// - retrieved and highlighted with nested inner hits.
//
// IMPORTANT
//
// Because nested documents are indexed as separate documents, they can only be
// accessed within the scope of the nested query, the nested/reverse_nested
// aggregations, or nested inner hits.
//
// For instance, if a string field within a nested document has index_options
// set to offsets to allow use of the postings during the highlighting, these
// offsets will not be available during the main highlighting phase. Instead,
// highlighting needs to be performed via nested inner hits. The same
// consideration applies when loading fields during a search through
// docvalue_fields or stored_fields.
//
//
// Limits on nested mappings and objects
//
// As described earlier, each nested object is indexed as a separate Lucene
// document. Continuing with the previous example, if we indexed a single
// document containing 100 user objects, then 101 Lucene documents would be
// created: one for the parent document, and one for each nested object. Because
// of the expense associated with nested mappings, Elasticsearch puts settings
// in place to guard against performance problems:
//
//     index.mapping.nested_fields.limit
// The maximum number of distinct nested mappings in an index. The nested type
// should only be used in special cases, when arrays of objects need to be
// queried independently of each other. To safeguard against poorly designed
// mappings, this setting limits the number of unique nested types per index.
// Default is 50.
//
//     index.mapping.nested_objects.limit
// The maximum number of nested JSON objects that a single document can contain
// across all nested types. This limit helps to prevent out of memory errors
// when a document contains too many nested objects. Default is 10000.
//
// To illustrate how this setting works, consider adding another nested type
// called comments to the previous example mapping. For each document, the
// combined number of user and comment objects it contains must be below the
// limit.
//
// See Settings to prevent mapping explosion regarding additional settings for
// preventing mappings explosion.
// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html
//
type NestedField struct {
	dynamicParam
	propertiesParam
	includeInParentParam
	includeInRootParam
}

func (n *NestedField) Field() (Field, error) {
	return n, nil
}

func (NestedField) Type() FieldType {
	return FieldTypeNested
}

func (n *NestedField) UnmarshalJSON(data []byte) error {
	var p NestedFieldParams
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	f, err := p.Nested()
	*n = *f
	return err
}

func (n NestedField) MarshalJSON() ([]byte, error) {
	return json.Marshal(nestedField{
		Dynamic:         n.dynamic,
		Properties:      n.properties,
		IncludeInParent: n.includeInParent,
		IncludeInRoot:   n.includeInRoot,
		Type:            n.Type(),
	})
}

func NewNestedField(params NestedFieldParams) (*NestedField, error) {
	return params.Nested()
}
