package search

type RuntimeMappingType string

const (
	RMTBoolean  RuntimeMappingType = "boolean"
	RMTDate     RuntimeMappingType = "date"
	RMTDouble   RuntimeMappingType = "double"
	RMTGeoPoint RuntimeMappingType = "geo_point"
	RMTIP       RuntimeMappingType = "ip"
	RMTKeyword  RuntimeMappingType = "keyword"
	RMTLong     RuntimeMappingType = "long"
)

type RuntimeMappingField struct {
	Type   RuntimeMappingType `bson:"type" json:"type"`
	Script string             `bson:"script,omitempty" json:"script,omitempty"`
}

type RuntimeMappings map[string]RuntimeMappingField
