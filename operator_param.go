package picker

import (
	"encoding/json"
	"strings"

	"github.com/chanced/dynamic"
)

type Operator string

func (o Operator) String() string { return string(o) }

func (o Operator) toUpper() Operator {
	return Operator(strings.ToUpper(string(o)))
}
func (o Operator) ref() *Operator {
	return &o
}

const DefaultOperator = OperatorOr

const (
	// OperatorOr Operator
	//
	// For example, a query value of capital of Hungary is interpreted as
	// capital OR of OR Hungary.
	OperatorOr Operator = "OR"
	// OperatorAnd Operator
	//
	// For example, a query value of capital of Hungary is interpreted as capital
	// AND of AND Hungary.
	OperatorAnd Operator = "AND"
	And                  = OperatorAnd
	Or                   = OperatorOr
)

// WithOperator is a query with the operator param
//
// Boolean logic used to interpret text in the query value.
type WithOperator interface {
	// Operator is the boolean logic used to interpret text in the query value.
	// Defaults to Or
	Operator() Operator
	// SetOperator sets the Operator to v
	SetOperator(v Operator)
}

// operatorParam is a query mixin that adds the operator param
type operatorParam struct {
	operator *Operator
}

// Operator is the boolean logic used to interpret text in the query value.
// Defaults to Or
func (o operatorParam) Operator() Operator {
	if o.operator != nil {
		return *o.operator
	}
	return DefaultOperator
}

// SetOperator sets the Operator to v
func (o *operatorParam) SetOperator(v Operator) {
	o.operator = v.toUpper().ref()
}
func unmarshalOperatorParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithOperator); ok {
		var str string
		err := json.Unmarshal(data, &str)
		if err != nil {
			return err
		}
		a.SetOperator(Operator(str))
	}
	return nil
}
func marshalOperatorParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithOperator); ok {
		if b.Operator() != DefaultOperator {
			data[paramOperator] = b.Operator()
		}
	}
	return data, nil
}
