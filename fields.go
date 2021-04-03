package picker

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/chanced/dynamic"
)

type FieldMappings map[string]Fielder

func (fm FieldMappings) Fields() (Fields, error) {
	res := make(Fields, len(fm))
	for k, v := range fm {
		f, err := v.Field()
		if err != nil {
			return nil, err
		}
		res[k] = f
	}
	return res, nil
}

// Fields are a collection of Field mappings
type Fields map[string]Field

func (f Fields) Field(key string) Field {
	return f[key]
}

func (f Fields) Get(key string) (Field, bool) {
	v, exists := f[key]
	return v, exists
}

func (f Fields) Has(key string) bool {
	_, exists := f[key]
	return exists
}

func (f Fields) Set(key string, field Fielder) (Field, error) {
	fld, err := field.Field()
	if err != nil {
		return fld, err
	}
	f[key] = fld
	return fld, nil
}

func (f *Fields) UnmarshalJSON(data []byte) error {
	var m map[string]dynamic.JSON
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	*f = make(Fields, len(m))

	for fld, fd := range m {
		var props dynamic.JSONObject
		err := json.Unmarshal(fd, &props)
		if err != nil {
			return err
		}
		typ, ok := props["type"]
		if !ok {
			return errors.New("mapping type is missing for " + fld)
		}
		handler, ok := FieldTypeHandlers[FieldType(typ.UnquotedString())]
		if !ok {
			return fmt.Errorf("%w <%s> for field %s", ErrInvalidType, typ, fld)
		}
		nf := handler()
		err = json.Unmarshal(fd, &nf)
		if err != nil {
			return err
		}
		(*f)[fld] = nf
	}
	return nil
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
	// SetFields sets the Fields value to v
	SetFields(v Fields)
}

// FieldWithFields is a Field with the fields paramater
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type FieldWithFields interface {
	Field
	WithFields
}

// fieldsParam is a mixin for mappings that adds the fields param
//
// It is often useful to index the same field in different ways for different
// purposes. This is the purpose of multi-fields. For instance, a string field
// could be mapped as a text field for full-text search, and as a keyword field
// for sorting or aggregations
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type fieldsParam struct {
	fields Fields
}

// Fields (Multi-fields) allow the same string value to be indexed in multiple
// ways for different purposes, such as one field for search and a multi-field
// for sorting and aggregations, or the same string value analyzed by different
// analyzers.
func (f fieldsParam) Fields() Fields {
	if f.fields == nil {
		f.fields = Fields{}
	}
	return f.fields
}

// SetFields sets the Fields value to v
func (f *fieldsParam) SetFields(v Fields) {
	f.fields = v
}
