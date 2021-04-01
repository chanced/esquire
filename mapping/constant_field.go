package mapping

import "encoding/json"

type Constanter interface {
	ConstantField() (*ConstantField, error)
}

// ConstantField is a specialization of the Keyword field for the case
// that all documentsin the index have the same value.
//
// ! X-Pack
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-field-type
type ConstantFieldParams struct {
	Value interface{} `json:"value"`
}

func (c ConstantFieldParams) Field() (Field, error) {
	return c.ConstantField()
}
func (c ConstantFieldParams) ConstantField() (*ConstantField, error) {
	return &ConstantField{
		value: c.Value,
	}, nil
}

// ConstantField is a specialization of the Keyword field for the case
// that all documentsin the index have the same value.
//
// ! X-Pack
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-field-type
type ConstantField struct {
	value interface{}
}

func (c ConstantField) MarshalJSON() ([]byte, error) {
	params := ConstantFieldParams{Value: c.value}
	return json.Marshal(params)
}

func (c *ConstantField) UnmarshalJSON(data []byte) error {
	*c = ConstantField{}
	params := ConstantFieldParams{}
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	c.value = params.Value
	return nil
}

func (ConstantField) Type() FieldType {
	return FieldTypeConstant
}
func (c *ConstantField) Value() interface{} {
	if c == nil {
		return nil
	}
	return c.value
}

func (c ConstantField) SetValue(v interface{}) {
	c.value = v
}

func NewConstantField(params ConstantFieldParams) (*ConstantField, error) {
	return params.ConstantField()
}
