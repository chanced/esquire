package picker

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

//easyjson:json
type Index struct {
	Mappings FieldMappings `json:"mappings"`
}

func NewIndex(params IndexParams) (*Index, error) {
	return params.Index()
}
