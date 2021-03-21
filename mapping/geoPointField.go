package mapping

// A GeoPointField accepts latitude-longitude pairs, which can be used:
//
// - to find geo-points within a bounding box, within a certain distance of a
// central point, or within a polygon or within a geo_shape query.
//
// - to aggregate documents geographically or by distance from a central point.
//
// - to integrate distance into a document’s relevance score.
//
// - to sort documents by distance.
//
//
// Geo-points expressed as an array or string
//
// Please note that string geo-points are ordered as lat,lon, while array
// geo-points are ordered as the reverse: lon,lat.
//
// Originally, lat,lon was used for both array and string, but the array format
// was changed early on to conform to the format used by GeoJSON.
//
// A point can be expressed as a geohash. Geohashes are base32 encoded strings
// of the bits of the latitude and longitude interleaved. Each character in a
// geohash adds additional 5 bits to the precision. So the longer the hash, the
// more precise it is. For the indexing purposed geohashs are translated into
// latitude-longitude pairs. During this process only first 12 characters are
// used, so specifying more than 12 characters in a geohash doesn’t increase the
// precision. The 12 characters provide 60 bits, which should reduce a possible
// error to less than 2cm.
//
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html
type GeoPointField struct {
	BaseField            `bson:",inline" json:",inline"`
	IgnoreMalformedParam `bson:",inline" json:",inline"`
	IgnoreZValueParam    `bson:",inline" json:",inline"`
	CoerceParam          `bson:",inline" json:",inline"`
}

func NewGeoPointField() *GeoPointField {
	return &GeoPointField{BaseField: BaseField{MappingType: TypeGeoPoint}}
}
