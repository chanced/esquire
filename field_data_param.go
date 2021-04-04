package picker

import "github.com/chanced/dynamic"

const DefaultFieldData = false

// WithFieldData is a mapping with the FieldData param
//
// FieldData determines whether the field can  use in-memory fielddata for
// sorting, aggregations, or scripting? Accepts true or false (default).
//
// Text fields are searchable by default, but by default are not available for
// aggregations, sorting, or scripting. If you try to sort, aggregate, or access
// values from a script on a text field, you will see this exception:
//
// Fielddata is disabled on text fields by default. Set fielddata=true on
// your_field_name in order to load fielddata in memory by uninverting the
// inverted index. Note that this can however use significant memory.
//
// Field data is the only way to access the analyzed tokens from a full text
// field in aggregations, sorting, or scripting. For example, a full text field
// like New York would get analyzed as new and york. To aggregate on these
// tokens requires field data.
//
// IMPORTANT
//
// It usually doesn’t make sense to enable fielddata on text fields. Field data
// is stored in the heap with the field data cache because it is expensive to
// calculate. Calculating the field data can cause latency spikes, and
// increasing heap usage is a cause of cluster performance issues.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html#fielddata-mapping-param
type WithFieldData interface {
	// FieldData determines whether the field use in-memory fielddata for
	// sorting, aggregations, or scripting. Accepts true or false (default).
	FieldData() bool
	// SetFieldData sets FieldData to v
	SetFieldData(v bool) error
}

type fieldDataParam struct {
	fieldData dynamic.Bool
}

// FieldData determines whether the field use in-memory fielddata for sorting,
// aggregations, or scripting. Accepts true or false (default).
func (fd fieldDataParam) FieldData() bool {
	if b, ok := fd.fieldData.Bool(); ok {
		return b
	}
	return DefaultFieldData
}

// SetFieldData sets FieldData to v
func (fd *fieldDataParam) SetFieldData(v interface{}) error {
	return fd.fieldData.Set(v)
}
