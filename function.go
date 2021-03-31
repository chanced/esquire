package picker

import (
	"encoding/json"
	"fmt"
	"strings"

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

var scoreFunctionHandlers = map[FuncKind]func() Function{
	FuncKindExp:              func() Function { return &ExpFunction{} },
	FuncKindGauss:            func() Function { return &GaussFunction{} },
	FuncKindLinear:           func() Function { return &LinearFunction{} },
	FuncKindWeight:           func() Function { return &WeightFunction{} },
	FuncKindScriptScore:      func() Function { return &ScriptScoreFunction{} },
	FuncKindRandomScore:      func() Function { return &RandomScoreFunction{} },
	FuncKindFieldValueFactor: func() Function { return &FieldValueFactorFunction{} },
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
	Filter() QueryClause
	json.Marshaler
	json.Unmarshaler
}
type Functions []Function

type Functioner interface {
	Function() (Function, error)
}

type FuncKind string

func (f FuncKind) String() string {
	return string(f)
}
func unmarshalFunctions(data dynamic.JSON) (Functions, error) {
	var fds []dynamic.JSON
	if data.IsNull() || len(data) == 0 {
		return nil, nil
	}
	if !data.IsArray() {
		fds = []dynamic.JSON{data}
	} else {
		err := json.Unmarshal(data, &fds)
		if err != nil {
			return nil, err
		}
	}
	res := make(Functions, len(fds))
	for i, fd := range fds {
		f, err := unmarshalFunction(fd)
		if err != nil {
			return nil, err
		}
		res[i] = f
	}
	return res, nil
}

func unmarshalFunction(data dynamic.JSON) (Function, error) {
	var params dynamic.JSONObject
	err := json.Unmarshal(data, &params)
	if err != nil {
		return nil, err
	}

	handlers := map[FuncKind]func() Function{}
	for k := range params {
		fk := FuncKind(strings.ToLower(k))
		if handler, ok := scoreFunctionHandlers[fk]; ok {
			handlers[fk] = handler
		}
	}
	if len(handlers) > 1 {
		for k, handler := range handlers {
			if k != FuncKindWeight {
				f := handler()
				err := f.UnmarshalJSON(data)
				return f, err
			}
		}
	}
	for _, handler := range handlers {
		f := handler()
		err := f.UnmarshalJSON(data)
		return f, err
	}
	return nil, fmt.Errorf("unknown function %s", data)
}
