package search

import "github.com/tidwall/gjson"

type Operator string

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

// OperatorParam is a query mixin that adds the operator param
type OperatorParam struct {
	OperatorValue *Operator `json:"operator,omitempty" bson:"operator,omitempty"`
}

// Operator is the boolean logic used to interpret text in the query value.
// Defaults to Or
func (o OperatorParam) Operator() Operator {
	if o.OperatorValue != nil {
		return *o.OperatorValue
	}
	return Or
}

// SetOperator sets the Operator to v
func (o *OperatorParam) SetOperator(v Operator) {
	o.OperatorValue = &v
}
func unmarshalOperatorParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithOperator); ok {
		a.SetOperator(Operator(value.Str))
	}
	return nil
}
