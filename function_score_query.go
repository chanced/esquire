package picker

// FunctionScoreQuery  allows you to modify the score of documents that are retrieved
// by a query. This can be useful if, for example, a score function is
// computationally expensive and it is sufficient to compute the score on a
// filtered set of documents.
//
// To use function_score, the user has to define a query and one or more
// functions, that compute a new score for each document returned by the query.
type FunctionScoreQuery struct {
	Query Query
	Boost interface{}
	// Documents with a score lower than this floating point number are excluded
	// from the search results. (Optional)
	MinScore  float64
	MaxBoost  float64
	BoostMode BoostMode
	ScoreMode ScoreMode
	Functions Funcs
}

func (fs *FunctionScoreQuery) FunctionScore() (*FunctionScoreClause, error) {
	if fs == nil {
		return nil, nil
	}
	c := &FunctionScoreClause{}
	err := c.SetBoostMode(fs.BoostMode)
	if err != nil {
		return c, err
	}
	err = c.SetScoreMode(fs.ScoreMode)
	if err != nil {
		return c, err
	}
	return c, nil
}
func (fs *FunctionScoreQuery) Clause() (QueryClause, error) {
	return fs.FunctionScore()
}

// FunctionScoreClause allows you to modify the score of documents that are retrieved
// by a query. This can be useful if, for example, a score function is
// computationally expensive and it is sufficient to compute the score on a
// filtered set of documents.
//
// To use function_score, the user has to define a query and one or more
// functions, that compute a new score for each document returned by the query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html
type FunctionScoreClause struct {
	query QueryValues
	boostModeParam
	scoreModeParam
	maxBoostParam
	functions Functions
}
