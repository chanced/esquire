package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type SearchField struct {
	// Wildcard pattern. The request returns doc values for
	// field names matching this pattern. (Required)
	Field string

	//  Format in which the doc values are returned.
	//
	// For date fields, you can specify a date date format. For numeric fields
	// fields, you can specify a DecimalFormat pattern. (Optional)
	//
	// For other field data types, this parameter is not supported.
	Format string
}

type field SearchField

func (f SearchField) MarshalJSON() ([]byte, error) {
	if f.Format != "" {
		return json.Marshal(f.Field)
	}
	return json.Marshal(field(f))
}

func (f *SearchField) UnmarshalJSON(data []byte) error {
	d := dynamic.JSON(data)
	if d.IsString() {
		var str string
		err := json.Unmarshal(d, &str)
		if err != nil {
			return err
		}
		f.Field = str
		f.Format = ""
		return nil
	}

	var sm map[string]string
	err := json.Unmarshal(d, &sm)
	if err != nil {
		return err
	}
	f.Field = sm["field"]
	f.Format = sm["format"]
	return nil
}

// Fields allows for retrieving a list of document fields in the search response. It consults both the document _source and the index mappings to return each value in a standardized way that matches its mapping type. By default, date fields are formatted according to the date format parameter in their mappings. You can also use the fields parameter to retrieve runtime field values.
//
// (Optional, array of strings and objects) Array of wildcard (*) patterns. The request returns values for field names matching these patterns in the hits.fields property of the response.
//
// You can specify items in the array as a string or object. See Fields for more details.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#search-fields-param
type SearchFields []SearchField
