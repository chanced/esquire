package picker

import "encoding/json"

type IndexParams struct {
	Mappings Mappings `json:"mappings"`
}

func (p IndexParams) Index() (*Index, error) {
	fm, err := p.Mappings.FieldMappings()
	i := &Index{}
	if err != nil {
		return i, err
	}
	i.Mappings = fm
	return i, nil
}

type Index struct {
	Mappings FieldMappings `json:"mappings"`
}
type index struct {
	Mappings FieldMappings `json:"mappings"`
}

func (i Index) MarshalBSON() ([]byte, error) {
	return i.MarshalJSON()
}

func (i Index) MarshalJSON() ([]byte, error) {
	return json.Marshal(index{Mappings: i.Mappings})
}

func (i *Index) UnmarshalBSON(data []byte) error {
	return i.UnmarshalJSON(data)
}

func (i *Index) UnmarshalJSON(data []byte) error {
	var idx index
	err := json.Unmarshal(data, &idx)
	if err != nil {
		return err
	}
	i.Mappings = idx.Mappings
	return nil
}

func NewIndex(params IndexParams) (*Index, error) {
	return params.Index()
}
