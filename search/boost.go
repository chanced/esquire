package search

type WithBoost interface {
	Boost() float32
	SetBoost(v float32)
}

type BoostParam struct {
	BoostValue *float32 `bson:"boost,omitempty" json:"boost,omitempty"`
}

func (b BoostParam) Boost() float32 {
	if b.BoostValue == nil {
		return 0
	}
	return *b.BoostValue
}

func (b *BoostParam) SetBoost(v float32) {
	b.BoostValue = &v
}
