package picker

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
type PercolatorField struct {
	BaseField `json:",inline" bson:",inline"`
}

func (f PercolatorField) Clone() Field {
	n := NewPercolatorField()
	return n
}

func NewPercolatorField() *PercolatorField {
	return &PercolatorField{BaseField: BaseField{MappingType: FieldTypePercolator}}
}
