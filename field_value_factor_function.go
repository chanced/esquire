package picker

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

type FieldValueFactorFunc struct {
	// 	Field to be extracted from the document.
	Field string
	// 	Optional factor to multiply the field value with, defaults to 1.
	Factor interface{}
	// Modifier to apply to the field value, can be one of: none, log, log1p, log2p, ln, ln1p, ln2p, square, sqrt, or reciprocal. Defaults to none.
	Modifier Modifier
	// Value used if the document doesn’t have that field. The modifier and factor are still applied to it as though it were read from the document.
	Missing interface{}
	// (optional) Query to filter (such as picker.TermQuery, picker.MatchQuery, and so on)
	Filter CompleteClauser
	// Float
	Weight interface{}
}

func (FieldValueFactorFunc) FuncKind() FuncKind {
	return FuncKindRandomScore
}

func (fvf FieldValueFactorFunc) Function() (Function, error) {
	return fvf.FieldValueFactorFunction()
}
func (fvf FieldValueFactorFunc) FieldValueFactorFunction() (*FieldValueFactorFunction, error) {
	f := &FieldValueFactorFunction{field: fvf.Field}
	err := f.SetWeight(fvf.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetModifier(fvf.Modifier)
	if err != nil {
		return f, err
	}

	err = f.SetFilter(fvf.Filter)
	if err != nil {
		return f, err
	}
	f.SetMissing(fvf.Missing)
	err = f.SetFactor(fvf.Factor)
	if err != nil {
		return f, err
	}
	return f, nil
}

type FieldValueFactorFunction struct {
	field   string
	missing interface{}
	factor  dynamic.Number
	modifierParam
	weightParam
	filter QueryClause
}

func (fvf FieldValueFactorFunction) Field() string {
	return fvf.field
}
func (fvf *FieldValueFactorFunction) SetField(field string) {
	if fvf == nil {
		*fvf = FieldValueFactorFunction{}
	}
	fvf.field = field
}

func (fvf FieldValueFactorFunction) Filter() QueryClause {
	return fvf.filter
}

func (fvf *FieldValueFactorFunction) SetFilter(filter CompleteClauser) error {
	if fvf == nil {
		*fvf = FieldValueFactorFunction{}
	}
	c, err := filter.Clause()
	if err != nil {
		return err
	}
	fvf.filter = c
	return nil
}
func (FieldValueFactorFunction) FuncKind() FuncKind {
	return FuncKindFieldValueFactor
}
func (fvf FieldValueFactorFunction) SetFactor(v interface{}) error {

	fvf.factor = dynamic.Number{}
	err := fvf.factor.Set(v)
	if err != nil {
		return err
	}
	if _, ok := fvf.factor.Float(); !ok {
		return fmt.Errorf("invalid Factor value for %s: <%d>", fvf.field, fvf.factor.Value())
	}
	return nil
}

func (fvf FieldValueFactorFunction) Factor() float64 {
	if fvf.factor.IsNil() {
		return float64(1)
	}
	f, _ := fvf.factor.Float()
	return f
}

func (fvf *FieldValueFactorFunction) SetMissing(missing interface{}) {
	if fvf == nil {
		*fvf = FieldValueFactorFunction{}
	}
	fvf.missing = missing
}

func (fvf FieldValueFactorFunction) Missing() interface{} {
	return fvf.missing
}

type fieldValueFactorParams struct {
	Modifier Modifier    `json:"modifier,omitempty"`
	Field    string      `json:"field"`
	Missing  interface{} `json:"missing,omitempty"`
	Factor   *float64    `json:"factor,omitempty"`
}

func (fvf FieldValueFactorFunction) MarshalJSON() ([]byte, error) {
	return marshalFunction(&fvf)
}

func (fvf *FieldValueFactorFunction) marshalParams(data dynamic.JSONObject) error {
	params := fieldValueFactorParams{
		Field:    fvf.field,
		Missing:  fvf.missing,
		Modifier: fvf.modifier,
	}
	if f, ok := fvf.factor.Float(); ok {
		params.Factor = &f
	}
	fd, err := json.Marshal(params)
	if err != nil {
		return err
	}
	data["field_value_factor"] = fd
	return nil
}
func (fvf *FieldValueFactorFunction) unmarshalParams(data dynamic.JSON) error {
	params := fieldValueFactorParams{}
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	err = fvf.factor.Set(params.Factor)
	if err != nil {
		return err
	}
	fvf.SetModifier(params.Modifier)
	fvf.SetMissing(params.Missing)
	return nil
}
