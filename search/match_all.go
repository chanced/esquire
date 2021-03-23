package search

type MatchAll struct {
	Boost float32
}

func (ma MatchAll) Rule() Rule {
	return ma.Rule()
}

func (ma MatchAll) MatchAll() *MatchAllRule {
	r := &MatchAllRule{}
	r.SetBoost(ma.Boost)
	return r
}

func (ma MatchAll) Type() Type {
	return TypeMatchAll
}

type MatchAllRule struct {
	BoostParam `json:",inline" bson:",inline"`
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
