package picker

import (
	"encoding/json"
	"strings"

	"github.com/chanced/dynamic"
)

type BoostMode string

const DefaultBoostMode = BoostModeMultiply

func (bm BoostMode) String() string {
	return string(bm)
}

func (bm *BoostMode) toLower() {
	*bm = BoostMode(strings.ToLower(string(*bm)))
}
func (bm *BoostMode) IsValid() bool {
	bm.toLower()
	bmv := *bm
	for _, v := range boostModes {
		if v == bmv {
			return true
		}
	}
	return false
}

const (
	BoostModeUnspecified BoostMode = ""
	// query score and function score is multiplied (default)
	BoostModeMultiply BoostMode = "multiply"
	// only function score is used, the query score is ignored
	BoostModeReplace BoostMode = "replace"
	// query score and function score are added
	BoostModeSum BoostMode = "sum"
	// average
	BoostModeAvg BoostMode = "avg"
	// max of query score and function score
	BoostModeMax BoostMode = "max"
	//min of query score and function score
	BoostModeMin BoostMode = "min"
)

var boostModes = []BoostMode{
	BoostModeUnspecified,
	BoostModeMultiply,
	BoostModeReplace,
	BoostModeSum,
	BoostModeAvg,
	BoostModeMax,
	BoostModeMin,
}

type boostModeParam struct {
	boostMode BoostMode
}

type WithBoostMode interface {
	SetBoostMode(sm BoostMode) error
	BoostMode() BoostMode
}

func (sm *boostModeParam) SetBoostMode(boostMode BoostMode) error {
	if !boostMode.IsValid() {
		return ErrInvalidBoostMode
	}
	sm.boostMode = boostMode
	return nil
}
func (sm *boostModeParam) BoostMode() BoostMode {
	if sm.boostMode == BoostModeUnspecified {
		return DefaultBoostMode
	}
	return sm.boostMode
}

func marshalBoostModeParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithBoostMode); ok {
		if b.BoostMode() != DefaultBoostMode {
			data["boost_mode"] = b.BoostMode()
		}
	}
	return data, nil
}
func unmarshalBoostModeParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithBoostMode); ok {
		var sm BoostMode
		err := json.Unmarshal(data, &sm)
		if err != nil {
			return err
		}
		return a.SetBoostMode(sm)
	}
	return nil
}
