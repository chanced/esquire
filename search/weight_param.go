package search

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/chanced/dynamic"
)

var DefaultWeight = float64(0)

type WithWeight interface {
	Weight() float64
	SetWeight(v interface{}) error
}

type weightParam struct {
	weightValue *float64
}

func (b weightParam) Weight() float64 {
	if b.weightValue == nil {
		return 0
	}
	return *b.weightValue
}

// SetWeight sets Weight to v
func (b *weightParam) SetWeight(v interface{}) error {
	n, err := dynamic.NewNumber(v)
	if err != nil {
		return err
	}
	if f, ok := n.Float(); ok {
		b.weightValue = &f
	} else if n.IsNil() {
		b.weightValue = nil
	} else {
		return fmt.Errorf("%w <%s>", ErrWeightRequired, v)
	}
	return nil
}

func marshalWeightParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithWeight); ok {
		if b.Weight() != DefaultWeight {
			data["weight"] = b.Weight()
		}
	}
	return data, nil
}
func unmarshalWeightParam(data dynamic.JSON, target interface{}) error {
	if r, ok := target.(WithWeight); ok {
		if data.IsNumber() {
			n, err := dynamic.NewNumber(string(data))
			if err != nil {
				return err
			}
			f, ok := n.Float()
			if !ok {
				return &json.UnmarshalTypeError{
					Value: string(data),
					Type:  reflect.TypeOf(float64(0)),
				}
			}
			r.SetWeight(f)
			return nil
		}
		if data.IsNull() {
			return nil
		}
		if data.IsString() {
			if len(data.UnquotedString()) == 0 {
				return nil
			}
			var str string
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}
			n, err := dynamic.NewNumber(str)
			if err != nil {
				return err
			}
			f, ok := n.Float()
			if ok {
				r.SetWeight(f)
			}
			return nil
		}
	}
	return nil
}
