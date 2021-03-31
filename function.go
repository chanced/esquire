package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const (
	// ScoreExp is an exponential decay function
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay
	FuncKindExp FuncKind = "exp"
	// ScoreGause is normal decay
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay
	FuncKindGauss FuncKind = "gauss"
	// FuncKindLinearDecay decay
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay
	FuncKindLinear FuncKind = "linear"
	// FuncKindWeight allows you to multiply the score by the provided weight. This
	// can sometimes be desired since boost value set on specific queries gets
	// normalized, while for this score function it does not. The number value
	// is of type float.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-weight
	FuncKindWeight FuncKind = "weight"
	// FuncKindScriptScore allows you to wrap another query and customize the
	// scoring of it optionally with a computation derived from other numeric
	// field values in the doc using a script expression.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-script-score
	FuncKindScriptScore FuncKind = "script_score"
	// FuncKindRandomScore generates scores that are uniformly distributed from 0 up to
	// but not including 1. By default, it uses the internal Lucene doc ids as a
	// source of randomness, which is very efficient but unfortunately not
	// reproducible since documents might be renumbered by merges.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-random
	FuncKindRandomScore FuncKind = "random_score"
	// FuncKindFieldValueFactor function allows you to use a field from a document
	// to influence the score. It’s similar to using the script_score function,
	// however, it avoids the overhead of scripting. If used on a multi-valued
	// field, only the first value of the field is used in calculations
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-field-value-factor
	FuncKindFieldValueFactor FuncKind = "field_value_factor"
)

var scoreFunctionHandlers = map[FuncKind]func() function{
	FuncKindExp:              func() function { return &ExpFunction{} },
	FuncKindGauss:            func() function { return &GaussFunction{} },
	FuncKindLinear:           func() function { return &LinearFunction{} },
	FuncKindScriptScore:      func() function { return &ScriptScoreFunction{} },
	FuncKindRandomScore:      func() function { return &RandomScoreFunction{} },
	FuncKindFieldValueFactor: func() function { return &FieldValueFactorFunction{} },
}

// Funcs is a slice of Functioners, valid options are:
//
//  - picker.WeightFunc,
//  - picker.LinearFunc,
//  - picker.ExpFunc,
//  - picker.GaussFunc,
//  - picker.RandomScoreFunc,
//  - picker.ScriptScoreFunc,
//  -
type Funcs []Functioner

func (f Funcs) functions() (Functions, error) {
	res := make(Functions, len(f))
	for i, v := range f {
		fv, err := v.Function()
		if err != nil {
			return nil, err
		}
		res[i] = fv
	}
	return res, nil
}

type Function interface {
	FuncKind() FuncKind
	Weight() float64
	SetWeight(interface{}) error
	Filter() QueryClause
	SetFilter(CompleteClauser) error
	json.Marshaler
	json.Unmarshaler
}
type function interface {
	Function
	unmarshalParams(data []byte) error
}

type DecayFunction interface {
	Function
	Field() string
	Origin() interface{}
	SetOrigin(interface{}) error
	Offset() dynamic.StringNumberOrTime
	SetOffset(interface{}) error
	Scale() dynamic.StringOrNumber
	SetScale(interface{}) error
	SetField(string) error
}

type Functions []Function

// TODO: This needs refactoring. Funcs are a pain to unmarshal
func (f *Functions) UnmarshalJSON(raw []byte) error {
	*f = Functions{}
	var fds []dynamic.JSON
	data := dynamic.JSON(raw)
	if data.IsNull() || len(data) == 0 {
		return nil
	}
	if data.IsNull() {
		return nil
	}

	if data.IsArray() {
		err := json.Unmarshal(data, &fds)
		if err != nil {
			return err
		}
	} else {

		fds = []dynamic.JSON{data}
	}

	for i, fd := range fds {
		fn, err := unmarshalFunction(fd)
		if err != nil {
			return err
		}
		if fn == nil {
			continue
		}
		(*f)[i] = fn
	}

	return nil
}

type Functioner interface {
	Function() (Function, error)
}

type FuncKind string

func (f FuncKind) String() string {
	return string(f)
}

func unmarshalFunction(data dynamic.JSON) (function, error) {
	var params dynamic.JSONObject
	err := json.Unmarshal(data, &params)
	if err != nil {
		return nil, err
	}
	var handler func() function
	var fn function
	for k, fd := range params {
		fk := FuncKind(k)
		if h, ok := scoreFunctionHandlers[fk]; ok {
			handler = h
			fn = handler()
			err := fn.unmarshalParams(fd)
			if err != nil {
				return err
			}
		}

	}
	if handler == nil {
		fn = &WeightFunction{}
	} else {
		fn = handler()
	}
	err = unmarshalWeightParam(params["weight"], fn)
	if err != nil {
		return nil, err
	}
	filter, err := unmarshalQueryClause(params["filter"])
	if err != nil {
		return nil, err
	}
	err = fn.SetFilter(filter)
	if err != nil {
		return nil, err
	}

}

func unmarshalDecayFuncParams(data dynamic.JSON, fn DecayFunction) error {
	var obj dynamic.JSONObject
	if len(data) == 0 {
		// empty? error?
		return nil
	}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	fn.unmarshalDecay()
}

func unmarshalDecayFunc(data dynamic.JSON, fn DecayFunction) error {
	var obj dynamic.JSONObject
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	for field, d := range obj {
		fn.SetField(field)
		return unmarshalDecayFuncParams(d, fn)
	}
	// empty? error?
	return nil
}
