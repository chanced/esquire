package search

// FunctionScore  allows you to modify the score of documents that are retrieved
// by a query. This can be useful if, for example, a score function is
// computationally expensive and it is sufficient to compute the score on a
// filtered set of documents.
//
// To use function_score, the user has to define a query and one or more
// functions, that compute a new score for each document returned by the query.
type FunctionScore struct {
	Query     Query
	Boost     interface{}
	Functions Functions
}

// FunctionScoreQuery allows you to modify the score of documents that are retrieved
// by a query. This can be useful if, for example, a score function is
// computationally expensive and it is sufficient to compute the score on a
// filtered set of documents.
//
// To use function_score, the user has to define a query and one or more
// functions, that compute a new score for each document returned by the query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html
type FunctionScoreQuery struct {
	query QueryValues
}
