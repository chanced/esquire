package picker

import (
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
	err := f.SetField(l.Field)
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
	err = f.SetOffset(l.Offset)
	if err != nil {
		return f, err
	}
	err = f.SetDecay(l.Decay)
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

func (l *LinearFunction) Field() string {
	if l == nil {
		return ""
	}
	return l.field
}

func (l *LinearFunction) SetField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	l.field = field
	return nil
}
func (LinearFunction) FuncKind() FuncKind {
	return FuncKindLinear
}
func (l LinearFunction) Filter() QueryClause {
	return l.filter
}
func (l *LinearFunction) Decay() dynamic.Number {
	return l.decay
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

func (l *LinearFunction) SetOffset(offset interface{}) error {
	return l.offset.Set(offset)
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

func (l *LinearFunction) unmarshalParams(data dynamic.JSON) error {
	return unmarshalDecayFunction(data, l)
}

func (l LinearFunction) MarshalBSON() ([]byte, error) {
	return l.MarshalJSON()
}

func (l LinearFunction) MarshalJSON() ([]byte, error) {
	return marshalFunction(&l)
}
func (l *LinearFunction) marshalParams(data dynamic.JSONObject) error {
	return marshalDecayFunction(data, l)
}
