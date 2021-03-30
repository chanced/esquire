package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type ExpFunc struct {
	// Field
	Field string
	// Weight (float)
	Weight interface{}
	// The point of origin used for calculating distance. Must be given as a
	// number for numeric field, date for date fields and geo point for geo
	// fields. Required for geo and numeric field. For date fields the default
	// is now. Date math (for example now-1h) is supported for origin.
	Origin interface{}
	// Required for all types. Defines the distance from origin + offset at
	// which the computed score will equal decay parameter. For geo fields: Can
	// be defined as number+unit (1km, 12m,…​). Default unit is meters. For date
	// fields: Can to be defined as a number+unit ("1h", "10d",…​). Default unit
	// is milliseconds. For numeric field: Any number.
	Scale interface{}
	// 	If an offset is defined, the decay function will only compute the decay
	// 	function for documents with a distance greater than the defined offset.
	// 	The default is 0.
	Offset interface{}
	// The decay parameter defines how documents are scored at the distance
	// given at scale. If no decay is defined, documents at the distance scale
	// will be scored 0.5.
	Decay  interface{}
	Filter CompleteClauser
}

func (ExpFunc) FuncKind() FunctionKind {
	return FuncKindExp
}
func (e ExpFunc) Function() (Function, error) {
	f := &ExpFunction{}
	err := f.setField(e.Field)
	if err != nil {
		return f, err
	}
	err = f.SetWeight(e.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetOrigin(e.Origin)
	if err != nil {
		return f, err
	}
	err = f.SetFilter(e.Filter)
	if err != nil {
		return f, err
	}
	err = f.SetScale(e.Scale)
	if err != nil {
		return f, err
	}
	return f, nil
}

type ExpFunction struct {
	weightParam
	field  string
	origin interface{}
	offset dynamic.StringNumberOrTime
	filter QueryClause
	decay  dynamic.Number
	scale  dynamic.StringOrNumber
}

func (ExpFunction) FunctionKind() FunctionKind {
	return FuncKindExp
}
func (e ExpFunction) Filter() QueryClause {
	return e.filter
}
func (e *ExpFunction) setField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	e.field = field
	return nil
}
func (e *ExpFunction) SetDecay(value interface{}) error {
	return e.decay.Set(value)
}

func (e *ExpFunction) SetScale(scale interface{}) error {
	if scale == nil {
		return ErrScaleRequired
	}
	err := e.scale.Set(scale)
	if err != nil {
		return err
	}
	if e.scale.IsEmptyString() {
		return ErrScaleRequired
	}
	return nil
}

func (e ExpFunction) Scale() dynamic.StringOrNumber {
	return e.scale
}

func (e ExpFunction) Origin() interface{} {
	return e.origin
}

func (e *ExpFunction) Offset() dynamic.StringNumberOrTime {
	return e.offset
}

func (e *ExpFunction) SetFilter(c CompleteClauser) error {
	if c == nil {
		e.filter = nil
		return nil
	}
	qc, err := c.Clause()
	if err != nil {
		return err
	}
	e.filter = qc
	return nil
}
func (e *ExpFunction) SetOrigin(origin interface{}) error {
	if origin == nil {
		return ErrOriginRequired
	}
	if s, ok := origin.(string); ok && len(s) == 0 {
		return ErrOriginRequired
	}
	e.origin = origin
	return nil
}

func (e *ExpFunction) unmarshalDecay(data dynamic.JSONObject) error {
	n := dynamic.Number{}
	err := n.UnmarshalJSON(data["decay"])
	if err != nil {
		return err
	}
	e.decay = n
	return nil
}

func (e *ExpFunction) unmarshalOffset(data dynamic.JSONObject) error {
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data["offset"])
	if err != nil {
		return err
	}
	e.offset = offset
	return nil
}

func (e *ExpFunction) unmarshalScale(data dynamic.JSONObject) error {
	scale := dynamic.StringOrNumber{}
	err := scale.UnmarshalJSON(data["scale"])
	if err != nil {
		return err
	}
	e.scale = scale
	return nil
}

func (e *ExpFunction) unmarsahlOffset(data dynamic.JSONObject) error {
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data["offset"])
	if err != nil {
		return err
	}
	e.offset = offset
	return nil
}

func (e *ExpFunction) unmarshalField(data dynamic.JSONObject) error {
	var field string
	err := json.Unmarshal(data["field"], &field)
	if err != nil {
		return err
	}
	e.field = field
	return nil
}
func (e *ExpFunction) unmarshalWeight(data dynamic.JSONObject) error {
	var weight *float64
	err := json.Unmarshal(data["weight"], &weight)
	if err != nil {
		return err
	}
	e.weight = weight
	return nil
}

func (e *ExpFunction) unmarshalFilter(data dynamic.JSONObject) error {
	filter, err := unmarshalQueryClause(data["filter"])
	if err != nil {
		return err
	}
	e.filter = filter
	return nil
}

func (e *ExpFunction) UnmarshalJSON(data []byte) error {
	*e = ExpFunction{}
	var fn dynamic.JSONObject
	err := json.Unmarshal(data, &fn)
	if err != nil {
		return err
	}
	unmarshalers := []func(data dynamic.JSONObject) error{
		e.unmarshalField,
		e.unmarsahlOffset,
		e.unmarshalDecay,
		e.unmarshalScale,
		e.unmarshalWeight,
		e.unmarshalFilter,
	}
	for field, d := range fn {
		var params dynamic.JSONObject
		err := json.Unmarshal(d, &params)
		if err != nil {
			return err
		}

		e.field = field
		for _, unmarshaler := range unmarshalers {
			err = unmarshaler(params)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func (e ExpFunction) MarshalJSON() ([]byte, error) {
	if e.field == "" {
		return dynamic.Null, nil
	}
	marshalers := []func() (string, dynamic.JSON, error){
		e.marshalField,
		e.marshalDecay,
		e.marshalOffset,
		e.marshalScale,
		e.marshalOrigin,
		e.marshalFilter,
		e.marshalWeight,
	}
	obj := dynamic.JSONObject{}

	for _, marshaler := range marshalers {
		param, data, err := marshaler()
		if err != nil {
			return nil, err
		}
		if data == nil || len(data) == 0 || (data.IsString() && len(data) == 2) {
			continue
		}
		obj[param] = data
	}
	return json.Marshal(obj)
}
func (e ExpFunction) marshalOrigin() (string, dynamic.JSON, error) {
	data, err := json.Marshal(e.origin)
	return "origin", data, err
}

func (e ExpFunction) marshalFilter() (string, dynamic.JSON, error) {
	if e.filter == nil {
		return "filter", nil, nil
	}
	data, err := e.filter.MarshalJSON()
	return "filter", data, err
}

func (e ExpFunction) marshalWeight() (string, dynamic.JSON, error) {
	data, err := json.Marshal(e.weight)
	return "weight", data, err
}

func (e ExpFunction) marshalDecay() (string, dynamic.JSON, error) {
	data, err := e.decay.MarshalJSON()
	return "decay", data, err
}

func (e ExpFunction) marshalOffset() (string, dynamic.JSON, error) {
	data, err := e.offset.MarshalJSON()
	return "offset", data, err
}

func (e ExpFunction) marshalField() (string, dynamic.JSON, error) {
	data, err := json.Marshal(e.field)
	return "field", data, err
}

func (e *ExpFunction) marshalScale() (string, dynamic.JSON, error) {
	data, err := e.scale.MarshalJSON()
	return "scale", data, err
}

func (e *ExpFunction) marsahlFilter() (string, dynamic.JSON, error) {
	if e.filter == nil {
		return "filter", nil, nil
	}
	data, err := e.filter.MarshalJSON()
	return "filter", data, err
}
