package picker

// WithProperties is a mixin that adds the properties param
//
// Type mappings, object fields and nested fields contain sub-fields, called
// properties. These properties may be of any data type, including object and
// nested. Properties can be added:
//
// - explicitly by defining them when creating an index. - explicitly by
// defining them when adding or updating a mapping type with the PUT mapping
// API. - dynamically just by indexing documents containing new fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type WithProperties interface {
	// Properties are fields within the object, which can be of any data type,
	// including object. New properties may be added to an existing object.
	Properties() Fields
	// SetProperties sets the Properties value to v
	SetProperties(v FieldMap) error
}

// propertiesParam is a mixin for mappings that adds the properties param
//
// Type mappings, object fields and nested fields contain sub-fields, called
// properties. These properties may be of any data type, including object and
// nested. Properties can be added:
//
// - explicitly by defining them when creating an index. - explicitly by
// defining them when adding or updating a mapping type with the PUT mapping
// API. - dynamically just by indexing documents containing new fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type propertiesParam struct {
	properties Fields
}

// Properties are fields within the object, which can be of any data type,
// including object. New properties may be added to an existing object.
func (p propertiesParam) Properties() Fields {
	return p.properties
}

// SetProperties sets the Properties value to v
func (p *propertiesParam) SetProperties(v Fieldset) error {
	if v == nil {
		p.properties = nil
		return nil
	}
	f, err := v.Fields()
	if err != nil {
		return err
	}
	p.properties = f
	return nil
}
