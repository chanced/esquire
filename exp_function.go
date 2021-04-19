package picker

import (
	"github.com/chanced/dynamic"
)

type ExpDecayFunctionParams struct {
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

func (ExpDecayFunctionParams) FuncKind() FuncKind {
	return FuncKindExp
}
func (e ExpDecayFunctionParams) Function() (Function, error) {
	f := &ExpDecayFunction{}
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
	err = f.SetDecay(e.Decay)
	if err != nil {
		return f, err
	}
	return f, nil
}

type ExpDecayFunction struct {
	weightParam
	field  string
	origin interface{}
	offset dynamic.StringNumberOrTime
	filter QueryClause
	decay  dynamic.Number
	scale  dynamic.StringOrNumber
}

func (e *ExpDecayFunction) Field() string {
	if e == nil {
		return ""
	}
	return e.field
}

func (e *ExpDecayFunction) SetField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	e.field = field
	return nil
}
func (ExpDecayFunction) FuncKind() FuncKind {
	return FuncKindExp
}
func (e ExpDecayFunction) Filter() QueryClause {
	return e.filter
}
func (e *ExpDecayFunction) Decay() dynamic.Number {
	return e.decay
}
func (e *ExpDecayFunction) SetDecay(value interface{}) error {
	return e.decay.Set(value)
}

func (e *ExpDecayFunction) SetScale(scale interface{}) error {
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

func (e ExpDecayFunction) Scale() dynamic.StringOrNumber {
	return e.scale
}

func (e ExpDecayFunction) Origin() interface{} {
	return e.origin
}

func (e *ExpDecayFunction) Offset() dynamic.StringNumberOrTime {
	return e.offset
}

func (e *ExpDecayFunction) SetOffset(offset interface{}) error {
	return e.offset.Set(offset)
}

func (e *ExpDecayFunction) SetFilter(c CompleteClauser) error {
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
func (e *ExpDecayFunction) SetOrigin(origin interface{}) error {
	if origin == nil {
		return ErrOriginRequired
	}
	if s, ok := origin.(string); ok && len(s) == 0 {
		return ErrOriginRequired
	}
	e.origin = origin
	return nil
}

func (e *ExpDecayFunction) unmarshalParams(data dynamic.JSON) error {
	return unmarshalDecayFunction(data, e)
}

func (e ExpDecayFunction) MarshalBSON() ([]byte, error) {
	return e.MarshalJSON()
}

func (e ExpDecayFunction) MarshalJSON() ([]byte, error) {
	return marshalFunction(&e)
}
func (e *ExpDecayFunction) marshalParams(data dynamic.JSONObject) error {
	return marshalDecayFunction(data, e)
}
