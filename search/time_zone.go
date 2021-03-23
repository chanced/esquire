package search

type WithTimeZone interface {
	TimeZone() string
	SetTimeZone(v string)
}

// TimeZoneParam is a mixin that adds the time_zone param to queries
//
// Coordinated Universal Time (UTC) offset or IANA time zone used to convert date values in the query to UTC.
// Valid values are ISO 8601 UTC offsets, such as +01:00 or -08:00, and IANA time zone IDs, such as America/Los_Angeles.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html#range-query-time-zone
type TimeZoneParam struct {
	TimeZoneValue string `json:"time_zone,omitempty" bson:"time_zone,omitempty"`
}

func (tz TimeZoneParam) TimeZone() string {
	return tz.TimeZoneValue
}
func (tz *TimeZoneParam) SetTimeZone(v string) {
	tz.TimeZoneValue = v
}
