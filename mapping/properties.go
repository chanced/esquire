package mapping

// WithProperties is a mixin that adds the properties param
//
// Type mappings, object fields and nested fields contain sub-fields, called
// properties. These properties may be of any data type, including object and
// nested. Properties can be added:
//
//  	- explicitly by defining them when creating an index. - explicitly by
// defining them when adding or updating a mapping type with the PUT mapping
// API. - dynamically just by indexing documents containing new fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type WithProperties interface {
	// Properties are fields within the object, which can be of any data type,
	// including object. New properties may be added to an existing object.
	Properties() Fields
	// SetProperties sets the Properties value to v
	SetProperties(v Fields)
	// Property returns the field with Key if it is exists, otherwise nil
	Property(key string) Field
	// SetProperty sets or adds the given Field v to the Properties param. It
	// initializes PropertiesParam's Value if it is currently nil.
	SetProperty(key string, v Field)
	// DeleteProperty deletes the Properties entry with the given key
	DeleteProperty(key string)
}

// FieldWithProperties is a Field with the properties paramater
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/properties.html
type FieldWithProperties interface {
	Field
	WithFields
}

// PropertiesParam is a mixin for mappings that adds the properties param
//
// Type mappings, object fields and nested fields contain sub-fields, called
// properties. These properties may be of any data type, including object and
// nested. Properties can be added:
//
//  	- explicitly by defining them when creating an index. - explicitly by
// defining them when adding or updating a mapping type with the PUT mapping
// API. - dynamically just by indexing documents containing new fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html
type PropertiesParam struct {
	PropertiesValue Fields `bson:"properties,omitempty" json:"properties,omitempty"`
}

// Properties are fields within the object, which can be of any data type,
// including object. New properties may be added to an existing object.
func (p PropertiesParam) Properties() Fields {
	return p.PropertiesValue
}

// SetProperties sets the Properties value to v
func (p *PropertiesParam) SetProperties(v Fields) {
	p.PropertiesValue = v
}

// Property returns the field with Key if it is exists, otherwise nil
func (p PropertiesParam) Property(key string) Field {
	if p.PropertiesValue == nil {
		return nil
	}
	return p.PropertiesValue[key]
}

// SetProperty sets or adds the given Field v to the Properties param. It
// initializes PropertiesParam's Value if it is currently nil.
func (p *PropertiesParam) SetProperty(key string, v Field) {
	if p.PropertiesValue == nil {
		p.PropertiesValue = Fields{}
	}

	p.PropertiesValue[key] = v

}

// DeleteProperty deletes the Properties entry with the given key
func (p *PropertiesParam) DeleteProperty(key string) {
	if p.PropertiesValue == nil {
		return
	}
	delete(p.PropertiesValue, key)
}
