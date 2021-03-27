package search

type MatchAll struct {
	Boost     float64
	QueryName string
}

func (ma MatchAll) Name() string {
	return ma.QueryName
}

func (ma *MatchAll) SetName(name string) {
	ma.QueryName = name
}

func (ma MatchAll) Clause() (Clause, error) {
	r := &matchAllClause{}
	r.SetBoost(ma.Boost)
	return r, nil
}

func (ma MatchAll) Type() Type {
	return TypeMatchAll
}
func (ma MatchAll) MatchAll() *matchAllClause {
	r := &matchAllClause{}
	r.SetBoost(ma.Boost)
	return r
}

type matchAllClause struct {
	boostParam
	nameParam
}

func (matchAllClause) Type() Type {
	return TypeMatchAll
}

type MatchAllQuery struct {
	MatchAllValue *MatchAll `json:"match_all,omitempty" bson:"match_all,omitempty"`
}

func (ma MatchAllQuery) Type() Type {
	return TypeMatchAll
}

func (ma *MatchAllQuery) SetMatchAll(v *MatchAll) {
	ma.MatchAllValue = v
}
