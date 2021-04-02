package picker

// WithEnabled is a mapping with an enabled parameter
//
// Enabled determines whether the JSON value given for the object field
// should be parsed and indexed (true, default) or completely ignored (false).
//
// Elasticsearch tries to index all of the fields you give it, but
// sometimes you want to just store the field without indexing it. For instance,
// imagine that you are using Elasticsearch as a web session store. You may
// want to index the session ID and last update time, but you don’t need to
// query or run aggregations on the session data itself.
//
// The enabled setting, which can be applied only to the top-level mapping
// definition and to object fields, causes Elasticsearch to skip parsing of the
// contents of the field entirely. The JSON can still be retrieved from the
// _source field, but it is not searchable or stored in any other way
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/enabled.html
type WithEnabled interface {
	// Enabled determines whether the JSON value given for the object field should be
	// parsed and indexed (true, default) or completely ignored (false).
	Enabled() bool
	// SetEnabled sets Enabled to v
	SetEnabled(v bool)
}

// FieldWithEnabled is a Field with an Enabled param
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/enabled.html
type FieldWithEnabled interface {
	Field
	WithEnabled
}

// EnabledParam is a mixin for mappings that adds the enabled parameter
//
// Enabled determines whether the JSON value given for the object field should
// be parsed and indexed (true, default) or completely ignored (false).
//
// Elasticsearch tries to index all of the fields you give it, but sometimes you
// want to just store the field without indexing it. For instance, imagine that
// you are using Elasticsearch as a web session store. You may want to index the
// session ID and last update time, but you don’t need to query or run
// aggregations on the session data itself.
//
// The enabled setting, which can be applied only to the top-level mapping
// definition and to object fields, causes Elasticsearch to skip parsing of the
// contents of the field entirely. The JSON can still be retrieved from the
// _source field, but it is not searchable or stored in any other way
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/enabled.html
type EnabledParam struct {
	EnabledValue *bool `bson:"enabled,omitempty" json:"enabled,omitempty"`
}

// Enabled determines whether the JSON value given for the object field should be
// parsed and indexed (true, default) or completely ignored (false).
func (ep EnabledParam) Enabled() bool {
	if ep.EnabledValue != nil {
		return *ep.EnabledValue
	}
	return true
}

// SetEnabled sets Enabled to v
func (ep *EnabledParam) SetEnabled(v bool) {
	if ep.Enabled() != v {
		ep.EnabledValue = &v
	}
}
