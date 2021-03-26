package search

import (
	"github.com/chanced/dynamic"
)

type Operator string

func (o Operator) String() string { return string(o) }

const DefaultOperator = Or

const (
	// Or Operator
	//
	// For example, a query value of capital of Hungary is interpreted as
	// capital OR of OR Hungary.
	Or Operator = "OR"
	// And Operator
	//
	// For example, a query value of capital of Hungary is interpreted as capital
	// AND of AND Hungary.
	And Operator = "AND"
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
	o.operator = &v
}
func unmarshalOperatorParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithOperator); ok {
		a.SetOperator(Operator(data.UnquotedString()))
	}
	return nil
}
func marshalOperatorParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithOperator); ok {
		if b.Operator() != DefaultOperator {
			data[paramOperator] = b.Operator()
		}
	}
	return data, nil
}
