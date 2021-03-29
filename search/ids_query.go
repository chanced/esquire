package search

type IDsQuery struct {
	IDs []string `json:"values" bson:"values"`
}
