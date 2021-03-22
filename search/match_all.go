package search

type MatchAll struct{}

func (ma MatchAll) Type() Type {
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
