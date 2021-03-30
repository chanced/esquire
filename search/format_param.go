package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultFormat = "strict_date_optional_time||epoch_millis"

// WithFormat is a type with a format parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-date-format.html
type WithFormat interface {
	// The date format(s) that can be parsed. Defaults to
	//   "strict_date_optional_time||epoch_millis."
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
	Format() string
	// SetFormat sets the Format Value to v
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
	SetFormat(v string)
}

type formatParam struct {
	format string // format
}

//Format is the format(s) that the that can be parsed. Defaults to strict_date_optional_time||epoch_millis.
//
// Multiple formats can be seperated by ||
func (f formatParam) Format() string {
	if f.format == "" {
		return DefaultFormat
	}
	return f.format
}

func (f *formatParam) SetFormat(v string) {
	if v != f.Format() {
		f.format = v
	}
}

func unmarshalFormatParam(value dynamic.JSON, target interface{}) error {
	if r, ok := target.(WithFormat); ok {
		if value.IsNull() {
			return nil
		}
		if value.IsString() {
			var str string
			err := json.Unmarshal(value, &str)
			if err != nil {
				return err
			}
			r.SetFormat(str)
		}
	}
	return nil
}

func marshalFormatParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithFormat); ok {
		if b.Format() != DefaultFormat {
			data[paramBoost] = b.Format()
		}
	}
	return data, nil
}
