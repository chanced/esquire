package picker

import (
	"github.com/chanced/dynamic"
)

type GaussFunc struct {
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

func (GaussFunc) FuncKind() FuncKind {
	return FuncKindGauss
}
func (g GaussFunc) Function() (Function, error) {
	f := &GaussFunction{}
	err := f.SetField(g.Field)
	if err != nil {
		return f, err
	}
	err = f.SetWeight(g.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetOrigin(g.Origin)
	if err != nil {
		return f, err
	}
	err = f.SetFilter(g.Filter)
	if err != nil {
		return f, err
	}
	err = f.SetScale(g.Scale)
	if err != nil {
		return f, err
	}
	err = f.SetOffset(g.Offset)
	if err != nil {
		return f, err
	}
	err = f.SetDecay(g.Decay)
	if err != nil {
		return f, err
	}
	return f, nil
}

type GaussFunction struct {
	weightParam
	field  string
	origin interface{}
	offset dynamic.StringNumberOrTime
	filter QueryClause
	decay  dynamic.Number
	scale  dynamic.StringOrNumber
}

func (g *GaussFunction) Field() string {
	if g == nil {
		return ""
	}
	return g.field
}

func (g *GaussFunction) SetField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	g.field = field
	return nil
}
func (GaussFunction) FuncKind() FuncKind {
	return FuncKindGauss
}
func (g GaussFunction) Filter() QueryClause {
	return g.filter
}
func (g *GaussFunction) Decay() dynamic.Number {
	return g.decay
}
func (g *GaussFunction) SetDecay(value interface{}) error {
	return g.decay.Set(value)
}

func (g *GaussFunction) SetScale(scale interface{}) error {
	if scale == nil {
		return ErrScaleRequired
	}
	err := g.scale.Set(scale)
	if err != nil {
		return err
	}
	if g.scale.IsEmptyString() {
		return ErrScaleRequired
	}
	return nil
}

func (g GaussFunction) Scale() dynamic.StringOrNumber {
	return g.scale
}

func (g GaussFunction) Origin() interface{} {
	return g.origin
}

func (g *GaussFunction) Offset() dynamic.StringNumberOrTime {
	return g.offset
}

func (g *GaussFunction) SetOffset(offset interface{}) error {
	return g.offset.Set(offset)
}

func (g *GaussFunction) SetFilter(c CompleteClauser) error {
	if c == nil {
		g.filter = nil
		return nil
	}
	qc, err := c.Clause()
	if err != nil {
		return err
	}
	g.filter = qc
	return nil
}
func (g *GaussFunction) SetOrigin(origin interface{}) error {
	if origin == nil {
		return ErrOriginRequired
	}
	if s, ok := origin.(string); ok && len(s) == 0 {
		return ErrOriginRequired
	}
	g.origin = origin
	return nil
}

func (g *GaussFunction) unmarshalParams(data dynamic.JSON) error {
	return unmarshalDecayFunction(data, g)
}

func (g GaussFunction) MarshalBSON() ([]byte, error) {
	return g.MarshalJSON()
}

func (g GaussFunction) MarshalJSON() ([]byte, error) {
	return marshalFunction(&g)
}
func (g *GaussFunction) marshalParams(data dynamic.JSONObject) error {
	return marshalDecayFunction(data, g)
}
