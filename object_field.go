package picker

import "encoding/json"

type objectField struct {
	Properties Fields      `json:"properties,omitempty"`
	Enabled    interface{} `json:"enabled,omitempty"`
	Dynamic    Dynamic     `json:"dynamic,omitempty"`
	Type       FieldType   `json:"type"`
}

type ObjectFieldParams struct {
	Properties Fields      `json:"properties,omitempty"`
	Enabled    interface{} `json:"enabled,omitempty"`
	Dynamic    Dynamic     `json:"dynamic,omitempty"`
}

func (p ObjectFieldParams) Field() (Field, error) {
	return p.Object()
}
func (p ObjectFieldParams) Object() (*ObjectField, error) {
	f := &ObjectField{}
	e := &MappingError{}

	return f, e.ErrorOrNil()
}

func (ObjectFieldParams) Type() FieldType {
	return FieldTypeObject
}

func NewObjectField(params ObjectFieldParams) (*ObjectField, error) {
	return params.Object()
}

// An ObjectField is a field type that contains other documents
//
// JSON documents are hierarchical in nature: the document may contain inner
// objects which, in turn, may contain inner objects themselves
//
// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/object.html
type ObjectField struct {
	propertiesParam
	enabledParam
	dynamicParam
}

func (ObjectField) Type() FieldType {
	return FieldTypeObject
}

func (o *ObjectField) Field() (Field, error) {
	return o, nil
}

func (o *ObjectField) UnmarshalJSON(data []byte) error {
	var p ObjectFieldParams
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	f, err := p.Object()
	*o = *f
	return err
}

func (o ObjectField) MarshalJSON() ([]byte, error) {
	return json.Marshal(objectField{
		Dynamic:    o.dynamic,
		Properties: o.properties,
		Enabled:    o.enabled.Value(),
		Type:       o.Type(),
	})
}
