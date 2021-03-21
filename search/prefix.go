package search

// Prefix returns documents that contain a specific prefix in a provided field.
type Prefix struct {
	Value                string `json:"value" bson:"value"`
	RewriteParam         `json:",inline" bson:",inline"`
	CaseInsensitiveParam `json:",inline" bson:",inline"`
}

// PrefixQuery returns documents that contain a specific prefix in a provided field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-prefix-query.html
type PrefixQuery struct {
	Prefix `json:"prefix,omitempty" bson:"prefix,omitempty"`
}
