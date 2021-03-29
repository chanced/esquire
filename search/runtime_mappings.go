package search

type RuntimeMappingKind string

const (
	RMTBoolean  RuntimeMappingKind = "boolean"
	RMTDate     RuntimeMappingKind = "date"
	RMTDouble   RuntimeMappingKind = "double"
	RMTGeoPoint RuntimeMappingKind = "geo_point"
	RMTIP       RuntimeMappingKind = "ip"
	RMTKeyword  RuntimeMappingKind = "keyword"
	RMTLong     RuntimeMappingKind = "long"
)

type RuntimeMappingField struct {
	Kind   RuntimeMappingKind `bson:"type" json:"type"`
	Script string             `bson:"script,omitempty" json:"script,omitempty"`
}

type RuntimeMappings map[string]RuntimeMappingField
