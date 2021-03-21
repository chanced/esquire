package mapping

import "errors"

type Meta map[string]string

func (m Meta) Clone() Meta {
	res := Meta{}
	for k, v := range m {
		res[k] = v
	}
	return res
}
func (m Meta) Unit() (string, bool) {
	v, exists := m["unit"]
	return v, exists
}

func (m Meta) SetUnit(v string) error {
	return m.Set("unit", v)
}
func (m Meta) SetMetricType(v string) error {
	return m.Set("metric_type", v)
}
func (m Meta) MetricType() (string, bool) {
	v, exists := m["metric_type"]
	return v, exists
}

func (m Meta) Value(key string) (string, bool) {
	v, exists := m[key]
	return v, exists
}
func (m Meta) Exists(key string) bool {
	_, exists := m.Value(key)
	return exists
}
func (m Meta) Set(key string, value string) error {
	if key == "" {
		return errors.New("invalid key")
	}
	if value == "" {
		if m.Exists(key) {
			delete(m, key)
			return nil
		}
		return nil
	}
	if m.Exists(key) {
		m[key] = value
	}
	if len(m) > 5 {
		return ErrMetaLimitExceeded
	}
	m[key] = value
	return nil
}

// FieldWithMeta is a Field with a meta parameter.

// WithMeta is a mapping with the meta parameter
//
// Meta is metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that
// work on the same indices to share meta information about fields such
// as units
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
type WithMeta interface {
	Meta() Meta
	SetMeta(Meta) error
	// SetUnit associated with a numeric field: "percent", "byte" or a time
	// unit. By default, a field does not have a unit. Only valid for numeric
	// fields. The convention for percents is to use value 1 to mean 100%.
	SetUnit(u string) error
	// Unit associated with a numeric field: "percent", "byte" or a time unit.
	// By default, a field does not have a unit. Only valid for numeric fields.
	// The convention for percents is to use value 1 to mean 100%.
	Unit() string
	// MetricType is the type of a numeric field: "gauge" || "counter". A gauge
	// is a single-value measurement that can go up or down over time, such as a
	// temperature. A counter is a single-value cumulative counter that only
	// goes up, such as the number of requests processed by a web server. By
	// default, no metric type is associated with a field. Only valid for
	// numeric fields.
	MetricType() string
	// SetMetricType sets the type of a numeric field: "gauge" || "counter".
	// A gauge is a single-value measurement that can go up or down over time,
	// such as a temperature. A counter is a single-value cumulative counter that
	// only goes up, such as the number of requests processed by a web server. By
	//default, no metric type is associated with a field. Only valid for numeric
	//fields.
	SetMetricType(v string) error
}

// MetaParam is a mixin that adds the meta field
//
// Metadata attached to the field. This metadata is opaque to Elasticsearch, it
// is only useful for multiple applications that work on the same indices to
// share meta information about fields such as units
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
type MetaParam struct {
	MetaValue Meta `bson:"meta,omitempty" json:"meta,omitempty"`
}

// Meta is metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that work on the
// same indices to share meta information about fields such as units
func (m MetaParam) Meta() Meta {
	return m.MetaValue
}

// SetMeta sets the metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that work on the
// same indices to share meta information about fields such as units
func (m *MetaParam) SetMeta(v Meta) error {
	if len(v) > 5 {
		return ErrMetaLimitExceeded
	}
	m.MetaValue = v
	return nil
}

// SetUnit associated with a numeric field: "percent", "byte" or a time unit.
//
// By default, a field does not have a unit. Only valid for numeric fields. The
// convention for percents is to use value 1 to mean 100%.
func (m *MetaParam) SetUnit(u string) error {
	if m.MetaValue == nil {
		m.MetaValue = Meta{}
	}
	return m.MetaValue.SetUnit(u)
}

// Unit associated with a numeric field: "percent", "byte" or a time unit.
//
// By default, a field does not have a unit. Only valid for numeric fields.
// The convention for percents is to use value 1 to mean 100%.
func (m MetaParam) Unit() string {
	if m.MetaValue == nil {
		m.MetaValue = Meta{}
	}
	v, _ := m.MetaValue.Unit()
	return v
}

// MetricType is the type of a numeric field: "gauge" || "counter".
//
// A gauge is a single-value measurement that can go up or down over time, such
// as a temperature. A counter is a single-value cumulative counter that only
// goes up, such as the number of requests processed by a web server. By
// default, no metric type is associated with a field. Only valid for numeric
// fields.
func (m MetaParam) MetricType() string {
	if m.MetaValue == nil {
		m.MetaValue = Meta{}
	}
	v, _ := m.MetaValue.MetricType()
	return v

}

// SetMetricType sets the type of a numeric field: "gauge" || "counter".
//
// A gauge is a single-value measurement that can go up or down over time, such
// as a temperature. A counter is a single-value cumulative counter that only
// goes up, such as the number of requests processed by a web server. By
// default, no metric type is associated with a field. Only valid for numeric
// fields.
func (m *MetaParam) SetMetricType(v string) error {
	if m.MetaValue == nil {
		m.MetaValue = Meta{}
	}
	return m.MetaValue.SetMetricType(v)
}
