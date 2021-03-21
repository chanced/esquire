package mapping

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
	Meta() map[string]string
	SetMeta(v map[string]string) error
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
	MetaValue map[string]string `bson:"meta,omitempty" json:"meta,omitempty"`
}

// Meta is metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that work on the
// same indices to share meta information about fields such as units
func (m MetaParam) Meta() map[string]string {
	return m.MetaValue
}

// SetMeta sets the metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that work on the
// same indices to share meta information about fields such as units
func (m *MetaParam) SetMeta(v map[string]string) error {
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
	if m.MetaValue != nil {
		if _, ok := m.MetaValue["unit"]; ok {
			if u == "" {
				delete(m.MetaValue, "unit")
				return nil
			}
			m.MetaValue["unit"] = u
			return nil
		}
		if len(m.MetaValue) > 4 {
			return ErrMetaLimitExceeded
		}
		m.MetaValue["unit"] = u
	}
	if u == "" {
		return nil
	}
	m.MetaValue = map[string]string{
		"unit": u,
	}
	return nil
}

// Unit associated with a numeric field: "percent", "byte" or a time unit.
//
// By default, a field does not have a unit. Only valid for numeric fields.
// The convention for percents is to use value 1 to mean 100%.
func (m MetaParam) Unit() string {
	if m.MetaValue == nil {
		return ""
	}
	return m.MetaValue["unit"]

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
		return ""
	}
	return m.MetaValue["metric_type"]
}

// SetMetricType sets the type of a numeric field: "gauge" || "counter".
//
// A gauge is a single-value measurement that can go up or down over time, such
// as a temperature. A counter is a single-value cumulative counter that only
// goes up, such as the number of requests processed by a web server. By
// default, no metric type is associated with a field. Only valid for numeric
// fields.
func (m *MetaParam) SetMetricType(v string) error {
	if m.MetaValue != nil {
		if _, ok := m.MetaValue["metric_type"]; ok {
			if v == "" {
				delete(m.MetaValue, "metric_type")
				return nil
			}
			m.MetaValue["metric_type"] = v
			return nil
		}
		if len(m.MetaValue) > 4 {
			return ErrMetaLimitExceeded
		}
		m.MetaValue["metric_type"] = v
	}
	if v == "" {
		return nil
	}
	m.MetaValue = map[string]string{
		"metric_type": v,
	}
	return nil
}
