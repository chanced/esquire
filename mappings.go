package picker

import "encoding/json"

type mappings struct {
	Properties Fields `json:"properties"  bson:"properties"`
}

type Mappings struct {
	Properties Fieldset `json:"properties"  bson:"properties"`
}

func (m *Mappings) MarshalJSON() ([]byte, error) {
	f, err := m.Properties.Fields()
	if err != nil {
		return nil, err
	}
	return json.Marshal(f)
}
