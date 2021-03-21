package mapping

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// Fields are a collection of Field mappings
type Fields map[string]Field

func (flds Fields) Field(key string) Field {
	return flds[key]
}

func (flds *Fields) UnmarshalJSON(data []byte) error {
	if flds == nil || *flds == nil {
		*flds = Fields{}
	}
	r := gjson.ParseBytes(data)
	var err error
	r.ForEach(func(key, value gjson.Result) bool {
		vt := Type(value.Get("type").String())
		handler, ok := FieldTypeHandlers[vt]
		if !ok {
			err = ErrInvalidType
			return false
		}
		fld := handler()
		if fld == nil {
			panic("field is nil")
		}

		err = json.Unmarshal([]byte(value.Raw), fld)
		if err != nil {
			return false
		}
		(*flds)[key.Str] = fld
		return true
	})
	return err
}

// WithFields is a mixin that adds the fields param
//
// It is often useful to index the same field in different ways for different purposes.
// This is the purpose of multi-fields. For instance, a string field could be mapped
// as a text field for full-text search, and as a keyword field for sorting or aggregations
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type WithFields interface {
	// Fields, also known as Multi-fields, allow the same string value to be
	// indexed in multiple ways for different purposes, such as one field for
	// search and a multi-field for sorting and aggregations, or the same string
	// value analyzed by different analyzers.
	Fields() Fields
	// Field returns the field with Key if it is exists, otherwise nil
	Field(key string) Field
	// SetFields sets the Fields value to v
	SetFields(v Fields)
	// SetField sets or adds the given Field v to the Fields param. It
	// initializes FieldsParam's Value if it is currently nil.
	SetField(key string, v Field)
	// DeleteField deletes the Fields entry with the given key
	DeleteField(key string)
}

// FieldWithFields is a Field with the fields paramater
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type FieldWithFields interface {
	Field
	WithFields
}

// FieldsParam is a mixin for mappings that adds the fields param
//
// It is often useful to index the same field in different ways for different
// purposes. This is the purpose of multi-fields. For instance, a string field
// could be mapped as a text field for full-text search, and as a keyword field
// for sorting or aggregations
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type FieldsParam struct {
	FieldsValue Fields `bson:"fields,omitempty" json:"fields,omitempty"`
}

// Fields (Multi-fields) allow the same string value to be indexed in multiple
// ways for different purposes, such as one field for search and a multi-field
// for sorting and aggregations, or the same string value analyzed by different
// analyzers.
func (f FieldsParam) Fields() Fields {
	return f.FieldsValue
}

// SetFields sets the Fields value to v
func (f *FieldsParam) SetFields(v Fields) {
	f.FieldsValue = v
}

// Field returns the field with Key if it is exists, otherwise nil
func (f FieldsParam) Field(key string) Field {
	if f.FieldsValue == nil {
		return nil
	}
	return f.FieldsValue[key]
}

// SetField sets or adds the given Field v to the Fields param. It initializes
// FieldsParam's Value if it is currently nil.
func (f *FieldsParam) SetField(key string, v Field) {
	if f.FieldsValue == nil {
		f.FieldsValue = Fields{}
	}

	f.FieldsValue[key] = v

}

// DeleteField deletes the Fields entry with the given key
func (f *FieldsParam) DeleteField(key string) {
	if f.FieldsValue == nil {
		return
	}
	delete(f.FieldsValue, key)
}
