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

func (ma MatchAll) Kind() Kind {
	return KindMatchAll
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

func (matchAllClause) Kind() Kind {
	return KindMatchAll
}

type MatchAllQuery struct {
	MatchAllValue *MatchAll `json:"match_all,omitempty" bson:"match_all,omitempty"`
}

func (ma MatchAllQuery) Kind() Kind {
	return KindMatchAll
}

func (ma *MatchAllQuery) SetMatchAll(v *MatchAll) {
	ma.MatchAllValue = v
}
