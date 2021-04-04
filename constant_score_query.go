package picker

type ConstantScorer interface {
	ConstantScore() (*ConstantScoreQuery, error)
}
type ConstantScoreQueryParams struct {
}
type ConstantScoreQuery struct {
}
