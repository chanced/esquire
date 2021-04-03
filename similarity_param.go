package picker

// Similarity is mostly useful for text fields, but can also apply to other field types.
type Similarity string

const (
	// The Okapi BM25 algorithm. The algorithm used by default in Elasticsearch and Lucene.
	//
	// https://en.wikipedia.org/wiki/Okapi_BM25
	SimilarityBM25 Similarity = "BM25"
	// SimilarityClassic is the former default in ElasticSsearch and Lucene that has since been deprecated
	//
	// Deprecated: TF/IDF algorithm, the former default in Elasticsearch and Lucene.
	SimilarityClassic Similarity = "classic"

	// SimilarityBoolean is a simple boolean similarity, which is used when
	// full-text ranking is not needed and the score should only be based on
	// whether the query terms match or not. Boolean similarity gives terms a
	// score equal to their query boost.
	SimilarityBoolean Similarity = "boolean"
)

func (s Similarity) String() string {
	return string(s)
}

// WithSimilarity is a mapping with the similarity parameter
//
// Elasticsearch allows you to configure a scoring algorithm or similarity per
// field. The similarity setting provides a simple way of choosing a similarity
// algorithm other than the default BM25, such as TF/IDF.
//
// Similarities are mostly useful for text fields, but can also apply to other
// field types.
//
// Custom similarities can be configured by tuning the parameters of the
// built-in similarities. For more details about this expert options, see the
// similarity module.
//
// The only similarities which can be used out of the box, without any further
// configuration are:
//
//  "BM25"
// The Okapi BM25 algorithm. The algorithm used by default in Elasticsearch and
// Lucene.
//
//  "classic" (deprecated)
// TF/IDF algorithm, the former default in Elasticsearch and Lucene.
//
//  "boolean"
// A simple boolean similarity, which is used when full-text ranking is not
// needed and the score should only be based on whether the query terms match or
// not. Boolean similarity gives terms a score equal to their query boost.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/similarity.html
type WithSimilarity interface {
	// Similarity is mostly useful for text fields, but can also apply to other
	// field types.
	Similarity() Similarity
	// SetSimilarity sets the similarity param to v
	SetSimilarity(v Similarity) error
}

// FieldWithSimilarity is a Field with the similarity paramater
type FieldWithSimilarity interface {
	Field
	WithSimilarity
}

// similarityParam is a mixin that adds the similarity parameter
//
// Elasticsearch allows you to configure a scoring algorithm or similarity per
// field. The similarity setting provides a simple way of choosing a similarity
// algorithm other than the default BM25, such as TF/IDF.
//
// Similarities are mostly useful for text fields, but can also apply to other
// field types.
//
// Custom similarities can be configured by tuning the parameters of the
// built-in similarities. For more details about this expert options, see the
// similarity module.
//
// The only similarities which can be used out of the box, without any further
// configuration are:
//
//  "BM25"
// The Okapi BM25 algorithm. The algorithm used by default in Elasticsearch and
// Lucene.
//
//  "classic" (deprecated)
// TF/IDF algorithm, the former default in Elasticsearch and Lucene.
//
//  "boolean"
// A simple boolean similarity, which is used when full-text ranking is not
// needed and the score should only be based on whether the query terms match or
// not. Boolean similarity gives terms a score equal to their query boost.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/similarity.html
type similarityParam struct {
	similarity Similarity
}

// Similarity is mostly useful for text fields, but can also apply to other
// field types.
func (s similarityParam) Similarity() Similarity {
	if len(s.similarity) == 0 {
		return SimilarityBM25
	}
	return s.similarity
}

// SetSimilarity sets the Similarity Value to v
func (s *similarityParam) SetSimilarity(v Similarity) error {

	// TODO: Figure out if similarity is case sensitive.
	// TODO: validate v

	s.similarity = v

	return nil
}
