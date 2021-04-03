package picker

import (
	"strings"

	"encoding/json"

	"github.com/chanced/dynamic"
)

const (
	ScoreModeUnspecified ScoreMode = ""
	// ScoreModeMultiply - scores are multiplied (default)
	ScoreModeMultiply ScoreMode = "multiply"
	// ScoreModeSum - scores are summed
	ScoreModeSum ScoreMode = "sum"
	// ScoreModeAvg scores are averaged
	ScoreModeAvg ScoreMode = "avg"
	// ScoreModeFirst - the first function that has a matching filter is applied
	ScoreModeFirst ScoreMode = "first"
	// ScoreModeMax - maximum score is used
	ScoreModeMax ScoreMode = "max"
	// ScoreModeMin - minimum score is used
	ScoreModeMin ScoreMode = "min"
)

var DefaultScoreMode = ScoreModeMultiply

type ScoreMode string

func (sm ScoreMode) String() string {
	return string(sm)
}
func (sm *ScoreMode) toLower() {
	*sm = ScoreMode(strings.ToLower(sm.String()))
}
func (sm *ScoreMode) IsValid() bool {
	sm.toLower()
	smv := *sm
	for _, v := range scoreModes {
		if v == smv {
			return true
		}
	}
	return false
}

var scoreModes = []ScoreMode{
	ScoreModeUnspecified,
	ScoreModeMultiply,
	ScoreModeSum,
	ScoreModeAvg,
	ScoreModeFirst,
	ScoreModeMax,
	ScoreModeMin,
}

type scoreModeParam struct {
	scoreMode ScoreMode
}

type WithScoreMode interface {
	SetScoreMode(sm ScoreMode) error
	ScoreMode() ScoreMode
}

func (sm *scoreModeParam) SetScoreMode(scoreMode ScoreMode) error {
	if !scoreMode.IsValid() {
		return ErrInvalidScoreMode
	}
	sm.scoreMode = scoreMode
	return nil
}
func (sm *scoreModeParam) ScoreMode() ScoreMode {
	if sm.scoreMode == ScoreModeUnspecified {
		return DefaultScoreMode
	}
	return sm.scoreMode
}

func marshalScoreModeParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithScoreMode); ok {
		if b.ScoreMode() != DefaultScoreMode {
			return json.Marshal(b.ScoreMode().String())
		}
	}
	return nil, nil
}

func unmarshalScoreModeParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithScoreMode); ok {
		var sm ScoreMode
		err := json.Unmarshal(data, &sm)
		if err != nil {
			return err
		}
		return a.SetScoreMode(sm)
	}
	return nil
}
