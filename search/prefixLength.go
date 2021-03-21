package search

type WithPrefixLength interface {
	PrefixLength() int
	SetPrefixLength(v int)
}

type PrefixLengthParam struct {
	PrefixLengthValue *int `json:"prefix_length,omitempty" bson:"prefix_length,omitempty"`
}

func (pl PrefixLengthParam) PrefixLength() int {
	if pl.PrefixLengthValue == nil {
		return 0
	}
	return *pl.PrefixLengthValue
}

func (pl *PrefixLengthParam) SetPrefixLength(v int) {
	pl.PrefixLengthValue = &v
}
