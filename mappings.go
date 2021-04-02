package picker

type Mappings struct {
	Properties Fields `json:"properties"  bson:"properties"`
}

func NewMappings() Mappings {
	return Mappings{
		Properties: Fields{},
	}
}
