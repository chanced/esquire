package picker

import (
	"encoding/json"
	"strings"

	"github.com/chanced/dynamic"
)

type SortMode string

const (
	// SortMdoeMin - Pick the lowest value.
	SortModeMin = "min"
	// SortModeMax - Pick the highest value.
	SortModeMax = "max"
	// SortModeSum - Use the sum of all values as sort value. Only applicable for number based array fields.
	SortModeSum = "sum"
	// SortModeAvg - Use the average of all values as sort value. Only applicable for number based array fields.
	SortModeAvg = "avg"
	// Use the median of all values as sort value. Only applicable for number based array fields.
	SortModeMedian = "median"
)

const (
	// Sort in descending order
	SortOrderDescending = "desc"
	// Sort in ascending order
	SortOrderAscending = "asc"
)

type SortOrder string

func (so SortOrder) String() string {
	return string(so)
}

func (o *SortOrder) IsValid() bool {
	if o == nil {
		return false
	}
	return strings.EqualFold(o.String(), SortOrderDescending) || strings.EqualFold(o.String(), SortOrderAscending)
}

type SortNested struct {
	// A filter that the inner objects inside the nested path should match with
	// in order for its field values to be taken into account by sorting. Common
	// case is to repeat the query / filter inside the nested filter or query.
	// By default no nested_filter is active.
	Filter *Query `json:"filter,omitempty"`
	// Defines on which nested object to sort. The actual sort field must be a
	// direct field inside this nested object. When sorting by nested field,
	// this field is mandatory.
	Path string `json:"path,omitempty"`
	// Same as top-level nested but applies to another nested path within the
	// current nested object.
	Nested *SortNested `json:"nested,omitempty"`
	// The maximum number of children to consider per root document when picking
	// the sort value. Defaults to unlimited.
	MaxChildren int64 `json:"max_children,omitempty"`
}

type Sort []SortEntry

type SortEntry struct {
	Field       string      `json:"-"`
	Order       SortOrder   `json:"order,omitempty"`
	Mode        SortMode    `json:"mode,omitempty"`
	NumericType string      `json:"numeric_type,omitempty"`
	Missing     string      `json:"missing,omitempty"`
	Type        string      `json:"type,omitempty"`
	Script      *Script     `json:"script,omitempty"`
	Nested      *SortNested `json:"nested,omitempty"`
}

func (s SortEntry) MarshalJSON() ([]byte, error) {

	b, err := sort{
		Order:       s.Order,
		Mode:        s.Mode,
		NumericType: s.NumericType,
		Missing:     s.Missing,
		Type:        s.Type,
		Script:      s.Script,
		Nested:      s.Nested,
	}.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return dynamic.JSONObject{s.Field: b}.MarshalJSON()
}

func (s *SortEntry) UnmarshalJSON(data []byte) error {
	*s = SortEntry{}
	var obj dynamic.JSONObject
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	for k, v := range obj {
		s.Field = k
		sv := sort{}
		err := sv.UnmarshalJSON(v)
		if err != nil {
			return err
		}
		s.Missing = sv.Missing
		s.Mode = sv.Mode
		s.Nested = sv.Nested
		s.NumericType = sv.NumericType
		s.Order = sv.Order
		s.Script = sv.Script
		s.Type = sv.Type
		return nil
	}
	// should this return an error if there were no keys or let it be an empty obj?
	return nil
}

//easyjson:json
type sort struct {
	Order       SortOrder   `json:"order,omitempty"`
	Mode        SortMode    `json:"mode,omitempty"`
	NumericType string      `json:"numeric_type,omitempty"`
	Missing     string      `json:"missing,omitempty"`
	Type        string      `json:"type,omitempty"`
	Script      *Script     `json:"script,omitempty"`
	Nested      *SortNested `json:"nested,omitempty"`
}
