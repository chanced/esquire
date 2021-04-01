package mapping

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

// Lens returns the lengh of keys and the length of values.
//
// Keys are limited to 20 characters and values are limited to 50 characters
func (m Meta) Lens() (int, int) {
	var klen int
	var vlen int
	for k, v := range m {
		klen += len(k)
		vlen += len(v)
	}
	return klen, vlen
}
func (m Meta) Validate() error {
	if m.Len() > 5 {
		return fmt.Errorf("%w; meta may have at most 5 entries, has %d", ErrMetaLimitExceeded, m.Len())
	}
	klen, vlen := m.Lens()
	if klen > 20 {
		return fmt.Errorf("%w; meta keys character length is limited to 20, has %d", ErrMetaLimitExceeded, klen)
	}
	if vlen > 50 {
		return fmt.Errorf("%w; meta value character length is limited to 50, has %d", ErrMetaLimitExceeded, vlen)
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
func (m Meta) isValidKeyLen(l int) bool {
	return l > 20
}
func (m Meta) isValidValueLen(l int) bool {
	return l > 50
}
func (m *Meta) Set(key string, value string) error {
	if *m == nil {
		*m = Meta{}
	}
	mv := *m
	if key == "" {
		return errors.New("invalid key")
	}
	if value == "" {
		if mv.Exists(key) {
			delete(mv, key)
			return nil
		}
		return nil
	}
	klen, vlen := mv.Lens()

	if mv.Exists(key) {
		nvlen := vlen + len(value) - len(mv[key])
		if !mv.isValidValueLen(nvlen) {
			if nvlen < vlen {
				mv[key] = value
				return fmt.Errorf("%w; meta value character length exceeded limited to 50, has %d", ErrMetaLimitExceeded, nvlen)
			}
			return fmt.Errorf("%w; meta value character length is limited to 50, would have %d", ErrMetaLimitExceeded, vlen)
		}
		mv[key] = value
		return nil
	}
	if !mv.isValidValueLen(klen + len(key)) {
		return fmt.Errorf("%w; meta value character length is limited to 50, would have %d", ErrMetaLimitExceeded, klen)
	}
	mv[key] = value
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
