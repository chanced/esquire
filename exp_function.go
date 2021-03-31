package picker

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

func (ExpFunc) FuncKind() FuncKind {
	return FuncKindExp
}
func (e ExpFunc) Function() (Function, error) {
	f := &ExpFunction{}
	err := f.SetField(e.Field)
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
	err = f.SetOffset(e.Offset)
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

func (e *ExpFunction) Field() string {
	if e == nil {
		return ""
	}
	return e.field
}

func (e *ExpFunction) SetField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	e.field = field
	return nil
}
func (ExpFunction) FuncKind() FuncKind {
	return FuncKindExp
}
func (e ExpFunction) Filter() QueryClause {
	return e.filter
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

func (e *ExpFunction) SetOffset(offset interface{}) error {
	return e.offset.Set(offset)
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

func (e *ExpFunction) unmarshalParams(data []byte) error {
	return unmarshalDecayFunc(data, e)
}

func (e *ExpFunction) UnmarshalJSON(data []byte) error {
	*e = ExpFunction{}
	unmarshalFunction()
}

func (e ExpFunction) MarshalJSON() ([]byte, error) {
	if e.field == "" {
		return dynamic.Null, nil
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
	mv := map[string]dynamic.JSONObject{e.field: obj}
	return json.Marshal(mv)
}
