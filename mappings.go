package picker

import "encoding/json"

type mappings struct {
	Properties Fields `json:"properties"`
}

type Mappings struct {
	Properties FieldMap `json:"properties"`
}

func (m Mappings) FieldMappings() (FieldMappings, error) {
	merr := MappingError{}
	fm := FieldMappings{
		Properties: make(Fields, len(m.Properties)),
	}
	for field, v := range m.Properties {
		f, err := v.Field()
		if err != nil {
			merr.Append(newFieldError(err, field))
			continue
		}
		fm.Properties[field] = f
	}
	return fm, merr.ErrorOrNil()
}

type FieldMappings struct {
	Properties Fields `json:"properties"`
}

func (m *Mappings) UnmarshalBSON(data []byte) error {
	return m.UnmarshalJSON(data)
}

func (m *Mappings) UnmarshalJSON(data []byte) error {
	*m = Mappings{}
	var v mappings
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	m.Properties = v.Properties.FieldMap()
	return nil
}

func (m *Mappings) MarshalJSON() ([]byte, error) {
	f, err := m.Properties.Fields()
	if err != nil {
		return nil, err
	}
	return json.Marshal(f)
}
