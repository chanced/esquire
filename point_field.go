package picker

import "encoding/json"

type pointField struct {
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	IgnoreZValue    interface{} `json:"ignore_z_value,omitempty"`
	NullValue       interface{} `json:"null_value,omitempty"`
	Type            FieldType   `json:"type"`
}

type PointFieldParams struct {

	// If true, malformed points are ignored. If false (default), malformed
	// points throw an exception and reject the whole document.
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	// If true (default) three dimension points will be accepted (stored in
	// source) but only x and y values will be indexed; the third dimension is
	// ignored. If false, points containing any more than x and y (two
	// dimensions) values throw an exception and reject the whole document.
	IgnoreZValue interface{} `json:"ignore_z_value,omitempty"`
	// Accepts an point value which is substituted for any explicit null values.
	// Defaults to null, which means the field is treated as missing.
	NullValue interface{} `json:"null_value,omitempty"`
}

func (PointFieldParams) Type() FieldType {
	return FieldTypePoint
}
func (p PointFieldParams) Field() (Field, error) {
	return p.Point()
}

func (p PointFieldParams) Point() (*PointField, error) {
	f := &PointField{}
	e := &MappingError{}

	err := f.SetIgnoreMalformed(p.IgnoreMalformed)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIgnoreZValue(p.IgnoreZValue)
	if err != nil {
		e.Append(err)
	}
	f.SetNullValue(p.NullValue)
	return f, e.ErrorOrNil()
}

// A PointField accepts latitude-longitude pairs, which can be used:
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
type PointField struct {
	ignoreMalformedParam
	ignoreZValueParam
	nullValueParam
}

func (PointField) Type() FieldType {
	return FieldTypePoint
}
func (gp *PointField) Field() (Field, error) {
	return gp, nil
}
func (gp *PointField) UnmarshalJSON(data []byte) error {
	var params PointFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Point()
	if err != nil {
		return err
	}
	*gp = *v
	return nil
}

func (gp PointField) MarshalJSON() ([]byte, error) {
	return json.Marshal(pointField{
		IgnoreMalformed: gp.ignoreMalformed.Value(),
		NullValue:       gp.nullValue,
		IgnoreZValue:    gp.ignoreZ,
		Type:            gp.Type(),
	})
}

func NewPointField(params PointFieldParams) (*PointField, error) {
	return params.Point()
}
