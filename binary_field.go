package picker

import "encoding/json"

type BinaryFieldParams struct {
	// Should the field be stored on disk in a column-stride fashion, so that it can later be used for sorting, aggregations, or scripting? Accepts true or false (default).
	DocValues interface{} `json:"doc_values,omitempty"`
	// Whether the field value should be stored and retrievable separately from the _source field. Accepts true or false (default).
	Store interface{} `json:"store,omitempty"`
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

func (b BinaryField) Field() (Field, error) {
	return &b, nil
}
func (BinaryField) Type() FieldType {
	return FieldTypeBinary
}
func (b BinaryField) MarshalBSON() ([]byte, error) {
	return b.MarshalJSON()
}

func (b BinaryField) MarshalJSON() ([]byte, error) {
	return json.Marshal(binaryField{
		DocValues: b.docValues.Value(),
		Store:     b.store.Value(),
		Type:      b.Type(),
	})
}
func (b *BinaryField) UnmarshalBSON(data []byte) error {
	return b.UnmarshalJSON(data)
}

func (b *BinaryField) UnmarshalJSON(data []byte) error {
	mErr := &MappingError{}
	p := BinaryFieldParams{}
	err := json.Unmarshal(data, &p)
	if err != nil {
		mErr.Append(err)
	}
	n, err := p.Binary()
	if err != nil {
		mErr.Append(err)
	}
	*b = *n
	return mErr.ErrorOrNil()
}

//easyjson:json
type binaryField struct {
	DocValues interface{} `json:"doc_values,omitempty"`
	Store     interface{} `json:"store,omitempty"`
	Type      FieldType   `json:"type"`
}
