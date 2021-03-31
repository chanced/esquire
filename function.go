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
}
type function interface {
	Function
	unmarshalParams(data []byte) error
	marshalParams(dynamic.JSONObject) error
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
	Decay() dynamic.Number
	SetDecay(v interface{}) error
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
				return nil, err
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
	return fn, nil
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

	unmarshalers := map[string]func(data dynamic.JSON, fn DecayFunction) error{
		"offset": unmarshalOffsetParam,
		"decay":  unmarshalDecayParam,
		"origin": unmarshalOriginParam,
		"scale":  unmarshalScaleParam,
	}
	for key, unmarshal := range unmarshalers {
		err := unmarshal(obj[key], fn)
		if err != nil {
			return err
		}
	}
	return nil
}

func unmarshalDecayFunction(data dynamic.JSON, fn DecayFunction) error {
	if data.Len() == 0 {
		return nil
	}
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

func unmarshalOriginParam(data dynamic.JSON, fn DecayFunction) error {
	if data.Len() == 0 {
		return nil
	}
	var origin interface{}
	err := json.Unmarshal(data, &origin)
	if err != nil {
		return err
	}
	fn.SetOrigin(origin)
	return nil
}

func unmarshalDecayParam(data dynamic.JSON, fn DecayFunction) error {
	if data.Len() == 0 {
		return nil
	}
	n := dynamic.Number{}
	err := n.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	return fn.SetDecay(n)
}

func unmarshalOffsetParam(data dynamic.JSON, fn DecayFunction) error {
	if data.Len() == 0 {
		return nil
	}
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	return fn.SetOffset(offset)

}

func unmarshalScaleParam(data dynamic.JSON, fn DecayFunction) error {
	if data.Len() == 0 {
		return nil
	}
	scale := dynamic.StringOrNumber{}
	err := scale.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	return fn.SetScale(scale)
}

func marshalFunction(fn function) ([]byte, error) {
	obj := dynamic.JSONObject{}
	weight, err := marshalWeightParam(fn)
	if err != nil {
		return nil, err
	}
	if len(weight) > 0 {
		obj["weight"] = weight
	}
	if fn.Filter() != nil {
		filter, err := fn.Filter().MarshalJSON()
		if err != nil {
			return nil, err
		}
		obj["filter"] = filter
	}
	if fn.FuncKind() != FuncKindWeight {
		err := fn.marshalParams(obj)
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(obj)
}
func marshalDecayFunctionParams(obj dynamic.JSONObject, fn DecayFunction) error {
	if len(fn.Field()) == 0 {
		return nil
	}
	params := dynamic.JSONObject{}
	if !fn.Decay().IsNil() {
		decay, err := json.Marshal(fn.Decay())
		if err != nil {
			return err
		}
		params["decay"] = decay
	}
	if !fn.Offset().IsNil() && fn.Offset().String() != "" {
		offset, err := json.Marshal(fn.Offset())
		if err != nil {
			return err
		}
		params["offset"] = offset
	}
	if fn.Origin() != nil {
		origin, err := json.Marshal(fn.Origin())
		if err != nil {
			return err
		}
		params["origin"] = origin
	}
	if !fn.Scale().IsNil() {
		scale, err := fn.Scale().MarshalJSON()
		if err != nil {
			return err
		}
		params["scale"] = scale
	}
	if len(params) > 0 {
		fnData, err := json.Marshal(params)
		if err != nil {
			return err
		}
		obj[string(fn.FuncKind())] = fnData
	}
	return nil
}
