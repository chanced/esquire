package search

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Field struct {
	// (Required, string) Wildcard pattern. The request returns doc values for
	// field names matching this pattern.
	Field string `bson:"field" json:"field"`

	// (Optional, string) Format in which the doc values are returned.
	//
	// For date fields, you can specify a date date format. For numeric fields
	// fields, you can specify a DecimalFormat pattern.
	//
	// For other field data types, this parameter is not supported.
	Format string `bson:"format,omitempty" json:"format,omitempty"`
}

type field Field

func (f Field) MarshalJSON() ([]byte, error) {
	if f.Format != "" {
		return json.Marshal(f.Field)
	}
	return json.Marshal(field(f))
}

func (f *Field) UnmarshalJSON(data []byte) error {
	res := gjson.ParseBytes(data)
	if res.Type == gjson.String {
		f.Field = res.Str
		f.Format = ""
		return nil
	}
	f.Field = res.Get("field").Str
	f.Format = res.Get("format").Str
	return nil
}

// Fields allows for retrieving a list of document fields in the search response. It consults both the document _source and the index mappings to return each value in a standardized way that matches its mapping type. By default, date fields are formatted according to the date format parameter in their mappings. You can also use the fields parameter to retrieve runtime field values.
//
// (Optional, array of strings and objects) Array of wildcard (*) patterns. The request returns values for field names matching these patterns in the hits.fields property of the response.
//
// You can specify items in the array as a string or object. See Fields for more details.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#search-fields-param
type Fields []Field
