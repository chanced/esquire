package picker

import "encoding/json"

type HistograpmFieldParams struct{}

func (HistograpmFieldParams) Type() FieldType {
	return FieldTypeHistogram
}
func (p HistograpmFieldParams) Field() (Field, error) {
	return p.Histogram()
}

func (p HistograpmFieldParams) Histogram() (*HistogramField, error) {
	f := &HistogramField{}
	e := &MappingError{}

	return f, e.ErrorOrNil()
}

type HistogramField struct{}

func (HistogramField) Type() FieldType {
	return FieldTypeHistogram
}
func (h *HistogramField) Field() (Field, error) {
	return h, nil
}
func (h *HistogramField) UnmarshalJSON(data []byte) error {

	return nil
}

func (h HistogramField) MarshalJSON() ([]byte, error) {
	return json.Marshal(histogramField{
		Type: h.Type(),
	})
}

func NewHistogramField(params HistograpmFieldParams) (*HistogramField, error) {
	return params.Histogram()
}

//easyjson:json
type histogramField struct {
	Type FieldType `json:"type"`
}
