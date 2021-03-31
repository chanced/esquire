package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultMinScore = float64(0)

type WithMinScore interface {
	MinScore() float64
	SetMinScore(v float64)
}

type minScoreParam struct {
	minScore *float64
}

func (ms minScoreParam) MinScore() float64 {
	if ms.minScore == nil {
		return DefaultMinScore
	}
	return *ms.minScore
}

func (ms *minScoreParam) SetMinScore(v float64) {
	if v > 0 {
		ms.minScore = &v
	}

}
func unmarshalMinScoreParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithMinScore); ok {
		n, err := dynamic.NewNumber(data.UnquotedString())
		if err != nil {
			return &json.UnmarshalTypeError{Value: string(data), Type: typeFloat64}
		}
		if v, ok := n.Float(); ok {
			a.SetMinScore(v)
		}
	}
	return nil
}

func marshalMinScoreParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithMinScore); ok {
		if b.MinScore() != DefaultMinScore {
			data["min_score"] = b.MinScore()
		}
	}
	return data, nil
}
