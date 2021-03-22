package search

type WithName interface {
	Name() string
	SetName(string)
}

// NameParam is a mixin that adds the _name parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html#named-queries
type NameParam struct {
	NameValue string `json:"_name,omitempty" bson:"_name,omitempty"`
}

func (n NameParam) Name() string {
	return n.NameValue
}
func (n *NameParam) SetName(name string) {
	if n.Name() != name {
		n.NameValue = name
	}
}
