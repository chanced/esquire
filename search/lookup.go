package search

type Lookup struct {
	// Index (Required, string) Name of the index from which to fetch field
	// values.
	Index string
	// ID (Required, string) ID of the document from which to fetch field
	// values.
	ID string
	// Path (Required, string) Name of the field from which to fetch field
	// values. Elasticsearch uses these values as search terms for the query.
	//
	// If the field values include an array of nested inner objects, you can
	// access those objects using dot notation syntax.
	Path string
	// Routing value of the document from which to fetch term values. If a
	// custom routing value was provided when the document was indexed, this
	// parameter is required. (Optional)
	Routing string

	Boost           float64
	CaseInsensitive bool
}

func (l Lookup) Rule() (Rule, error) {
	return l.Terms()
}
func (l Lookup) Terms() (*TermsRule, error) {
	q := &TermsRule{}
	q.SetBoost(l.Boost)
	q.SetCaseInsensitive(l.CaseInsensitive)

	return q, nil
}

func (l Lookup) Type() Type {
	return TypeTerms
}
