package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type LinearFunc struct {
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

func (LinearFunc) FuncKind() FuncKind {
	return FuncKindLinear
}
func (l LinearFunc) Function() (Function, error) {
	f := &LinearFunction{}
	err := f.setField(l.Field)
	if err != nil {
		return f, err
	}
	err = f.SetWeight(l.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetOrigin(l.Origin)
	if err != nil {
		return f, err
	}
	err = f.SetFilter(l.Filter)
	if err != nil {
		return f, err
	}
	err = f.SetScale(l.Scale)
	if err != nil {
		return f, err
	}
	return f, nil
}

type LinearFunction struct {
	weightParam
	field  string
	origin interface{}
	offset dynamic.StringNumberOrTime
	filter QueryClause
	decay  dynamic.Number
	scale  dynamic.StringOrNumber
}

func (LinearFunction) FuncKind() FuncKind {
	return FuncKindLinear
}
func (l LinearFunction) Filter() QueryClause {
	return l.filter
}
func (l *LinearFunction) setField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	l.field = field
	return nil
}
func (l *LinearFunction) SetDecay(value interface{}) error {
	return l.decay.Set(value)
}

func (l *LinearFunction) SetScale(scale interface{}) error {
	if scale == nil {
		return ErrScaleRequired
	}
	err := l.scale.Set(scale)
	if err != nil {
		return err
	}
	if l.scale.IsEmptyString() {
		return ErrScaleRequired
	}
	return nil
}

func (l LinearFunction) Scale() dynamic.StringOrNumber {
	return l.scale
}

func (l LinearFunction) Origin() interface{} {
	return l.origin
}

func (l *LinearFunction) Offset() dynamic.StringNumberOrTime {
	return l.offset
}

func (l *LinearFunction) SetFilter(c CompleteClauser) error {
	if c == nil {
		l.filter = nil
		return nil
	}
	qc, err := c.Clause()
	if err != nil {
		return err
	}
	l.filter = qc
	return nil
}
func (l *LinearFunction) SetOrigin(origin interface{}) error {
	if origin == nil {
		return ErrOriginRequired
	}
	if s, ok := origin.(string); ok && len(s) == 0 {
		return ErrOriginRequired
	}
	l.origin = origin
	return nil
}

func (l *LinearFunction) unmarshalDecay(data dynamic.JSONObject) error {
	n := dynamic.Number{}
	err := n.UnmarshalJSON(data["decay"])
	if err != nil {
		return err
	}
	l.decay = n
	return nil
}

func (l *LinearFunction) unmarshalOffset(data dynamic.JSONObject) error {
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data["offset"])
	if err != nil {
		return err
	}
	l.offset = offset
	return nil
}

func (l *LinearFunction) unmarshalScale(data dynamic.JSONObject) error {
	scale := dynamic.StringOrNumber{}
	err := scale.UnmarshalJSON(data["scale"])
	if err != nil {
		return err
	}
	l.scale = scale
	return nil
}

func (l *LinearFunction) unmarsahlOffset(data dynamic.JSONObject) error {
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data["offset"])
	if err != nil {
		return err
	}
	l.offset = offset
	return nil
}

func (l *LinearFunction) unmarshalField(data dynamic.JSONObject) error {
	var field string
	err := json.Unmarshal(data["field"], &field)
	if err != nil {
		return err
	}
	l.field = field
	return nil
}
func (l *LinearFunction) unmarshalWeight(data dynamic.JSONObject) error {
	var weight *float64
	err := json.Unmarshal(data["weight"], &weight)
	if err != nil {
		return err
	}
	l.weight = weight
	return nil
}

func (l *LinearFunction) unmarshalFilter(data dynamic.JSONObject) error {
	filter, err := unmarshalQueryClause(data["filter"])
	if err != nil {
		return err
	}
	l.filter = filter
	return nil
}

func (l *LinearFunction) UnmarshalJSON(data []byte) error {
	*l = LinearFunction{}
	var fn dynamic.JSONObject
	err := json.Unmarshal(data, &fn)
	if err != nil {
		return err
	}
	unmarshalers := []func(data dynamic.JSONObject) error{
		l.unmarshalField,
		l.unmarsahlOffset,
		l.unmarshalDecay,
		l.unmarshalScale,
		l.unmarshalWeight,
		l.unmarshalFilter,
	}
	for field, d := range fn {
		var params dynamic.JSONObject
		err := json.Unmarshal(d, &params)
		if err != nil {
			return err
		}

		l.field = field
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

func (l LinearFunction) MarshalJSON() ([]byte, error) {
	if l.field == "" {
		return dynamic.Null, nil
	}
	marshalers := []func() (string, dynamic.JSON, error){
		l.marshalField,
		l.marshalDecay,
		l.marshalOffset,
		l.marshalScale,
		l.marshalOrigin,
		l.marshalFilter,
		l.marshalWeight,
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
func (l LinearFunction) marshalOrigin() (string, dynamic.JSON, error) {
	data, err := json.Marshal(l.origin)
	return "origin", data, err
}

func (l LinearFunction) marshalFilter() (string, dynamic.JSON, error) {
	if l.filter == nil {
		return "filter", nil, nil
	}
	data, err := l.filter.MarshalJSON()
	return "filter", data, err
}

func (l LinearFunction) marshalWeight() (string, dynamic.JSON, error) {
	data, err := json.Marshal(l.weight)
	return "weight", data, err
}

func (l LinearFunction) marshalDecay() (string, dynamic.JSON, error) {
	data, err := l.decay.MarshalJSON()
	return "decay", data, err
}

func (l LinearFunction) marshalOffset() (string, dynamic.JSON, error) {
	data, err := l.offset.MarshalJSON()
	return "offset", data, err
}

func (l LinearFunction) marshalField() (string, dynamic.JSON, error) {
	data, err := json.Marshal(l.field)
	return "field", data, err
}

func (l *LinearFunction) marshalScale() (string, dynamic.JSON, error) {
	data, err := l.scale.MarshalJSON()
	return "scale", data, err
}

func (l *LinearFunction) marsahlFilter() (string, dynamic.JSON, error) {
	if l.filter == nil {
		return "filter", nil, nil
	}
	data, err := l.filter.MarshalJSON()
	return "filter", data, err
}
