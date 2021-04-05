package picker

import "github.com/chanced/dynamic"

const DefaultMaxGaps = -1

type WithMaxGaps interface {
	MaxGaps() int
	SetMaxGaps(maxGaps interface{}) error
}

type maxGapsParam struct {
	maxGaps dynamic.Number
}

func (mg maxGapsParam) MaxGaps() int {
	if mg.maxGaps.HasValue() {
		if i, ok := mg.maxGaps.Int(); ok {
			return i
		}
		if f, ok := mg.maxGaps.Float64(); ok {
			return int(f)
		}
	}
	return DefaultMaxGaps
}
func (mg *maxGapsParam) SetMaxGaps(maxGaps interface{}) error {
	return mg.maxGaps.Set(maxGaps)
}
