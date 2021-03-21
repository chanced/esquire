package mapping

type Mappings struct {
	Properties Fields `json:"properties"  bson:"properties"`
}
