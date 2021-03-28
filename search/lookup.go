package search

type Lookup struct {
	Field string
	// Name of the index from which to fetch field
	// values.(Required)

	Index string
	// ID of the document from which to fetch field
	// values. (Required)
	ID string

	// Path of the field from which to fetch field values. Elasticsearch
	// uses these values as search terms for the query.(Required)
	//
	// If the field values include an array of nested inner objects, you can
	// access those objects using dot notation syntax.

	Path string
	// Routing value of the document from which to fetch term values. If a
	// custom routing value was provided when the document was indexed, this
	// parameter is required. (Optional)
	Routing         string
	Boost           float64
	CaseInsensitive bool

	QueryName string
}

func (l Lookup) Name() string {
	return l.QueryName
}

func (l Lookup) Clause() (Clause, error) {
	return l.Terms()
}
func (l Lookup) Terms() (TermQuery, error) {
	q := TermQuery{}
	q.SetBoost(l.Boost)
	q.SetCaseInsensitive(l.CaseInsensitive)
	return q, nil
}

func (l Lookup) Type() Type {
	return TypeTerms
}
