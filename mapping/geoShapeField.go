package mapping

// GeoShapeField  facilitates the indexing of and searching with arbitrary geo
// shapes such as rectangles and polygons. It should be used when either the
// data being indexed or the queries being executed contain shapes other than
// just points.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html
type GeoShapeField struct {
	BaseField            `bson:",inline" json:",inline"`
	OrientationParam     `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	IgnoreZValueParam    `bson:",inline" json:",inline"`
	CoerceParam          `bson:",inline" json:",inline"`
}

func NewGeoShapeField() *GeoShapeField {
	return &GeoShapeField{BaseField: BaseField{MappingType: TypeGeoShape}}
}
