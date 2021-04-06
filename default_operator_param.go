package picker

import (
	"strings"

	"encoding/json"

	"github.com/chanced/dynamic"
)

type WithDefaultOperator interface {
	DefaultOperator() Operator
	SetDefaultOperator(v Operator) error
}

type defaultOperatorParam struct {
	defaultOperator Operator
}

func (o defaultOperatorParam) DefaultOperator() Operator {
	if len(o.defaultOperator) == 0 {
		return DefaultOperator
	}
	return o.defaultOperator
}
func (o *defaultOperatorParam) SetDefaultOperator(v Operator) error {
	v = Operator(strings.TrimSpace(v.toUpper().String()))
	o.defaultOperator = v
	return nil
}
func unmarshalDefaultOperatorParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithDefaultOperator); ok {
		var str string
		err := json.Unmarshal(data, &str)
		if err != nil {
			return err
		}
		return a.SetDefaultOperator(Operator(str))
	}
	return nil
}
func marshalDefaultOperatorParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithDefaultOperator); ok {
		if b.DefaultOperator() != DefaultOperator {
			return json.Marshal(b.DefaultOperator())
		}
	}
	return nil, nil
}
