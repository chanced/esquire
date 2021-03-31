package picker

import (
	"encoding/json"

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
	err := f.setField(g.Field)
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

func (GaussFunction) FuncKind() FuncKind {
	return FuncKindGauss
}
func (g GaussFunction) Filter() QueryClause {
	return g.filter
}
func (g *GaussFunction) setField(field string) error {
	if len(field) == 0 {
		return ErrFieldRequired
	}
	g.field = field
	return nil
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

func (g *GaussFunction) unmarshalDecay(data dynamic.JSONObject) error {
	n := dynamic.Number{}
	err := n.UnmarshalJSON(data["decay"])
	if err != nil {
		return err
	}
	g.decay = n
	return nil
}

func (g *GaussFunction) unmarshalOffset(data dynamic.JSONObject) error {
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data["offset"])
	if err != nil {
		return err
	}
	g.offset = offset
	return nil
}

func (g *GaussFunction) unmarshalScale(data dynamic.JSONObject) error {
	scale := dynamic.StringOrNumber{}
	err := scale.UnmarshalJSON(data["scale"])
	if err != nil {
		return err
	}
	g.scale = scale
	return nil
}

func (g *GaussFunction) unmarsahlOffset(data dynamic.JSONObject) error {
	offset := dynamic.StringNumberOrTime{}
	err := offset.UnmarshalJSON(data["offset"])
	if err != nil {
		return err
	}
	g.offset = offset
	return nil
}

func (g *GaussFunction) unmarshalField(data dynamic.JSONObject) error {
	var field string
	err := json.Unmarshal(data["field"], &field)
	if err != nil {
		return err
	}
	g.field = field
	return nil
}
func (g *GaussFunction) unmarshalWeight(data dynamic.JSONObject) error {
	var weight *float64
	err := json.Unmarshal(data["weight"], &weight)
	if err != nil {
		return err
	}
	g.weight = weight
	return nil
}

func (g *GaussFunction) unmarshalFilter(data dynamic.JSONObject) error {
	filter, err := unmarshalQueryClause(data["filter"])
	if err != nil {
		return err
	}
	g.filter = filter
	return nil
}

func (g *GaussFunction) UnmarshalJSON(data []byte) error {
	*g = GaussFunction{}
	var fn dynamic.JSONObject
	err := json.Unmarshal(data, &fn)
	if err != nil {
		return err
	}
	unmarshalers := []func(data dynamic.JSONObject) error{
		g.unmarshalField,
		g.unmarsahlOffset,
		g.unmarshalDecay,
		g.unmarshalScale,
		g.unmarshalWeight,
		g.unmarshalFilter,
	}
	for field, d := range fn {
		var params dynamic.JSONObject
		err := json.Unmarshal(d, &params)
		if err != nil {
			return err
		}

		g.field = field
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

func (g GaussFunction) MarshalJSON() ([]byte, error) {
	if g.field == "" {
		return dynamic.Null, nil
	}
	marshalers := []func() (string, dynamic.JSON, error){
		g.marshalField,
		g.marshalDecay,
		g.marshalOffset,
		g.marshalScale,
		g.marshalOrigin,
		g.marshalFilter,
		g.marshalWeight,
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
func (g GaussFunction) marshalOrigin() (string, dynamic.JSON, error) {
	data, err := json.Marshal(g.origin)
	return "origin", data, err
}

func (g GaussFunction) marshalFilter() (string, dynamic.JSON, error) {
	if g.filter == nil {
		return "filter", nil, nil
	}
	data, err := g.filter.MarshalJSON()
	return "filter", data, err
}

func (g GaussFunction) marshalWeight() (string, dynamic.JSON, error) {
	data, err := json.Marshal(g.weight)
	return "weight", data, err
}

func (g GaussFunction) marshalDecay() (string, dynamic.JSON, error) {
	data, err := g.decay.MarshalJSON()
	return "decay", data, err
}

func (g GaussFunction) marshalOffset() (string, dynamic.JSON, error) {
	data, err := g.offset.MarshalJSON()
	return "offset", data, err
}

func (g GaussFunction) marshalField() (string, dynamic.JSON, error) {
	data, err := json.Marshal(g.field)
	return "field", data, err
}

func (g *GaussFunction) marshalScale() (string, dynamic.JSON, error) {
	data, err := g.scale.MarshalJSON()
	return "scale", data, err
}

func (g *GaussFunction) marsahlFilter() (string, dynamic.JSON, error) {
	if g.filter == nil {
		return "filter", nil, nil
	}
	data, err := g.filter.MarshalJSON()
	return "filter", data, err
}
