package picker

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/chanced/dynamic"
)

// TODO: This needs to be cleaned up

type Fieldset interface {
	Fields() (Fields, error)
	Field(string) (Field, error)
	Has(name string) bool
	Set(name string, params Fielder) (Field, error)
	Remove(name string) Field
	Len() int
}

// Fields are a collection of Field mappings
type Fields map[string]Field

func (f Fields) Len() int {
	return len(f)
}
func (f Fields) Remove(field string) Field {
	v := f[field]

	if v == nil {
		return nil
	}
	// it doesn't matter if an error occurs here as all valid data is persisted with Field()
	// and the field is being deleted anyway. Presumably, this will have already been validated
	// anyway but if it is later added elsewhere, it'll go through validation again.
	fv, _ := v.Field()
	delete(f, field)
	return fv
}

func (f Fields) FieldMap() FieldMap {
	res := make(FieldMap, len(f))
	for k, v := range f {
		res[k] = v
	}
	return res
}

func (f Fields) Fields() (Fields, error) {
	return f, nil
}

func (f Fields) Field(key string) (Field, error) {
	if f[key] == nil {
		return nil, ErrFieldNotFound
	}
	return f[key].Field()
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

func (f *Fields) UnmarshalBSON(data []byte) error {
	return f.UnmarshalJSON(data)
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

type FieldMap map[string]Fielder

func (f FieldMap) Has(key string) bool {
	_, exists := f[key]
	return exists
}
func (f FieldMap) Field(field string) (Field, error) {
	if f[field] == nil {
		return nil, ErrFieldNotFound
	}
	return f[field].Field()
}

func (f FieldMap) Fields() (Fields, error) {
	e := &MappingError{}
	res := make(Fields, len(f))
	for k, v := range f {
		fv, err := v.Field()
		if err != nil {
			e.Append(&FieldError{Field: k, Err: err})
		}
		res[k] = fv
	}
	return res, e.ErrorOrNil()
}
func (f FieldMap) Set(name string, params Fielder) (Field, error) {
	fld, err := params.Field()
	if err != nil {
		return fld, err
	}
	f[name] = fld
	return fld, nil
}
func (f FieldMap) Remove(field string) Field {
	v := f[field]

	if v == nil {
		return nil
	}
	// it doesn't matter if an error occurs here as all valid data is persisted with Field()
	// and the field is being deleted anyway. Presumably, this will have already been validated
	// anyway but if it is later added elsewhere, it'll go through validation again.
	fv, _ := v.Field()
	delete(f, field)
	return fv
}

func (f FieldMap) Len() int {
	return len(f)
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
	SetFields(v Fields) error
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
func (f *fieldsParam) SetFields(fields Fieldset) error {
	fv, err := fields.Fields()
	if err != nil {
		return err
	}
	f.fields = fv
	return nil
}
