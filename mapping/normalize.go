package mapping

// WithNormalizer is a mixin that adds the normalizer parameter
//
// The normalizer property of keyword fields is similar to analyzer except that
// it guarantees that the analysis chain produces a single token.
//
// The normalizer is applied prior to indexing the keyword, as well as at
// search-time when the keyword field is searched via a query parser such as the
// match query or via a term-level query such as the term query.
//
// A simple normalizer called lowercase ships with elasticsearch and can be
// used. Custom normalizers can be defined as part of analysis settings
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/normalizer.html
type WithNormalizer interface {
	// Normalizer property of keyword fields is similar to analyzer except that it
	// guarantees that the analysis chain produces a single token.
	Normalizer() string
	// SetNormalizer sets the Normalizer value to v
	//
	// The normalizer property of keyword fields is similar to analyzer except that
	// it guarantees that the analysis chain produces a single token.
	SetNormalizer(v string)
}

// FieldWithNormalizer is a Field with the normalizer parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/normalizer.html
type FieldWithNormalizer interface {
	Field
	WithNormalizer
}

type NormalizerParam struct {
	NormalizerValue string `bson:"normalizer,omitempty" json:"normalizer,omitempty"`
}

// Normalizer property of keyword fields is similar to analyzer except that it
// guarantees that the analysis chain produces a single token.
func (n NormalizerParam) Normalizer() string {
	return n.NormalizerValue
}

// SetNormalizer sets the Normalizer value to v
//
// The normalizer property of keyword fields is similar to analyzer except that
// it guarantees that the analysis chain produces a single token.
func (n *NormalizerParam) SetNormalizer(v string) {
	if n.Normalizer() != v {
		n.NormalizerValue = v
	}
}
