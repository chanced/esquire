package picker

import "encoding/json"

type BinaryFieldParams struct {
	DocValues interface{} `json:"doc_values,omitempty"`
	Store     interface{} `json:"store,omitempty"`
}

func (b BinaryFieldParams) Binary() (*BinaryField, error) {
	f := &BinaryField{}
	err := f.SetStore(b.Store)
	if err != nil {
		return f, err
	}
	err = f.SetDocValues(b.DocValues)
	if err != nil {
		return f, err
	}
	return f, nil
}

func (b BinaryFieldParams) Field() (Field, error) {
	return b.Binary()
}
func (BinaryFieldParams) Type() FieldType {
	return FieldTypeBinary
}
func NewBinaryField(params BinaryFieldParams) (*BinaryField, error) {
	return params.Binary()
}

// Whether the field value should be stored and retrievable separately from the _source field. Accepts true or false (default).
// Binary

// BinaryField is a value encoded as a Base64 string.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/binary.html
type BinaryField struct {
	docValuesParam
	storeParam
}

func (BinaryField) Type() FieldType {
	return FieldTypeBinary
}
func (b BinaryField) MarshalJSON() ([]byte, error) {
	return json.Marshal(BinaryFieldParams{
		DocValues: b.docValues.Value(),
		Store:     b.store.Value(),
	})
}
func (b *BinaryField) UnmarshalJSON(data []byte) error {
	p := BinaryFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	n, err := p.Binary()
	if err != nil {
		return err
	}
	*b = *n
	return nil
}
