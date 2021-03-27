package search

type MatchAll struct {
	Boost float64
}

func (ma MatchAll) Rule() (Clause, error) {
	r := &MatchAllRule{}
	r.SetBoost(ma.Boost)
	return r, nil
}

func (ma MatchAll) Type() Type {
	return TypeMatchAll
}
func (ma MatchAll) MatchAll() *MatchAllRule {
	r := &MatchAllRule{}
	r.SetBoost(ma.Boost)
	return r
}

type MatchAllRule struct {
	boostParam `json:",inline" bson:",inline"`
}

func (MatchAllRule) Type() Type {
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
