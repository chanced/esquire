package picker

import "encoding/json"

type mappings struct {
	Properties Fields `json:"properties"  bson:"properties"`
}

type Mappings struct {
	Properties FieldMap `json:"properties"  bson:"properties"`
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
