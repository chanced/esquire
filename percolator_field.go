package picker

import "encoding/json"

type percolatorField struct {
	Type FieldType `json:"type"`
}

type PercolatorFieldParams struct{}

func (PercolatorFieldParams) Type() FieldType {
	return FieldTypePercolator
}
func (p PercolatorFieldParams) Field() (Field, error) {
	return p.Percolator()
}
func (PercolatorFieldParams) Percolator() (*PercolatorField, error) {
	return &PercolatorField{}, nil
}

// The PercolatorField type parses a json structure into a native query and
// stores that query, so that the percolate query can use it to match provided
// documents.
//
// Any field that contains a json object can be configured to be a percolator
// field. The percolator field type has no settings. Just configuring the
// percolator field type is sufficient to instruct Elasticsearch to treat a
// field as a query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/percolator.html
type PercolatorField struct{}

func (PercolatorField) Type() FieldType {
	return FieldTypePercolator
}
func (p *PercolatorField) Field() (Field, error) {
	return p, nil
}
func NewPercolatorField() (*PercolatorField, error) {
	return &PercolatorField{}, nil
}

func (PercolatorField) MarshalJSON() ([]byte, error) {
	return json.Marshal(percolatorField{
		Type: FieldTypePercolator,
	})
}

func (PercolatorField) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &percolatorField{})
}
