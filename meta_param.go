package picker

import (
	"errors"
	"fmt"
)

type Meta map[string]string

// Unit associated with a numeric field: "percent", "byte" or a time unit.
// By default, a field does not have a unit. Only valid for numeric fields.
// The convention for percents is to use value 1 to mean 100%.
func (m *Meta) Unit() string {
	return m.Get("unit")
}
func (m *Meta) Len() int {
	return len(*m)
}
func (m Meta) IsValid() bool {
	return m.Validate() == nil
}

func (m Meta) Validate() error {
	if m.Len() > 5 {
		return fmt.Errorf("%w; meta may have at most 5 entries, has %d", ErrMetaLimitExceeded, m.Len())
	}
	for k, v := range m {
		err := m.validateKV(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m Meta) Get(key string) string {
	return m[key]
}

func (m Meta) Has(key string) bool {
	_, has := m[key]
	return has
}

func (m *Meta) SetUnit(v string) error {
	return m.Set("unit", v)
}
func (m *Meta) SetMetricType(v string) error {
	return m.Set("metric_type", v)
}

// MetricType is the type of a numeric field: "gauge" || "counter". A gauge
// is a single-value measurement that can go up or down over time, such as a
// temperature. A counter is a single-value cumulative counter that only
// goes up, such as the number of requests processed by a web server. By
// default, no metric type is associated with a field. Only valid for
// numeric fields.
func (m *Meta) MetricType() string {
	v := m.Get("metric_type")
	return v

}

func (m *Meta) Value(key string) (string, bool) {
	if *m == nil {
		*m = Meta{}
	}
	v, exists := (*m)[key]
	return v, exists
}
func (m Meta) Exists(key string) bool {
	_, exists := m.Value(key)
	return exists
}
func (m Meta) isValidKeyLen(key string) bool {
	return len(key) >= 0 && len(key) <= 20
}
func (m Meta) isValidValueLen(value string) bool {
	return len(value) > 0 && len(value) <= 50
}
func (m Meta) validateKV(key string, value string) error {
	if !m.isValidKeyLen(key) {
		return fmt.Errorf("%w; keys must be less than 20 characters; %s is %d characters", ErrMetaLimitExceeded, key, len(key))
	}
	if !m.isValidValueLen(value) {
		return fmt.Errorf("%w; values must be less than 50 characters; value for key %s is %d characters", ErrMetaLimitExceeded, key, len(value))
	}
	return nil
}
func (m *Meta) Set(key string, value string) error {
	if *m == nil {
		*m = Meta{}
	}
	mv := *m
	if key == "" {
		return errors.New("picker: invalid meta key; key can not be empty")
	}
	err := m.validateKV(key, value)
	if err != nil {
		return err
	}
	if value == "" {
		if mv.Exists(key) {
			delete(mv, key)
			return nil
		}
		return nil
	}

	mv[key] = value
	return nil
}

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
}

// metaParam is a mixin that adds the meta field
//
// Metadata attached to the field. This metadata is opaque to Elasticsearch, it
// is only useful for multiple applications that work on the same indices to
// share meta information about fields such as units
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-field-meta.html#mapping-field-meta
type metaParam struct {
	meta Meta
}

// Meta is metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that work on the
// same indices to share meta information about fields such as units
func (m metaParam) Meta() Meta {
	return m.meta
}

// SetMeta sets the metadata attached to the field. This metadata is opaque to
// Elasticsearch, it is only useful for multiple applications that work on the
// same indices to share meta information about fields such as units
func (m *metaParam) SetMeta(v Meta) error {
	if v == nil {
		m.meta = nil
		return nil
	}
	err := v.Validate()
	if err != nil {
		return err
	}
	m.meta = v
	return nil
}
