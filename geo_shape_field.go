package picker

import "encoding/json"

type GeoShapeFieldParams struct {
	// Optional. Vertex order for the shape’s coordinates list.
	//
	//œ This parameter sets and returns only a RIGHT (counterclockwise) or LEFT
	// (clockwise) value. However, you can specify either value in multiple
	// ways.
	// Defaults to RIGHT to comply with OGC standards. OGC standards define outer ring vertices in counterclockwise order with inner ring (hole) vertices in clockwise order.
	//
	// Individual GeoJSON or WKT documents can override this parameter.
	Orientation Orientation `json:"orientation,omitempty"`

	// If true, malformed geo-points are ignored. If false (default), malformed
	// geo-points throw an exception and reject the whole document. A geo-point
	// is considered malformed if its latitude is outside the range -90 ⇐
	// latitude ⇐ 90, or if its longitude is outside the range -180 ⇐ longitude
	// ⇐ 180.
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	// If true (default) three dimension points will be accepted (stored in source)
	// but only latitude and longitude values will be indexed; the third dimension
	// is ignored. If false, geo-points containing any more than latitude and
	// longitude (two dimensions) values throw an exception and reject the whole
	// document.
	IgnoreZValue interface{} `json:"ignore_z_value,omitempty"`
}

func (GeoShapeFieldParams) Type() FieldType {
	return FieldTypeGeoShape
}
func (p GeoShapeFieldParams) Field() (Field, error) {
	return p.GeoShape()
}

func (p GeoShapeFieldParams) GeoShape() (*GeoShapeField, error) {
	f := &GeoShapeField{}
	e := &MappingError{}
	err := f.SetIgnoreMalformed(p.IgnoreMalformed)
	if err != nil {
		e.Append(err)
	}
	err = f.SetIgnoreZValue(p.IgnoreZValue)
	if err != nil {
		e.Append(err)
	}
	err = f.SetOrientation(p.Orientation)
	if err != nil {
		e.Append(err)
	}
	return f, e.ErrorOrNil()

}

// GeoShapeField  facilitates the indexing of and searching with arbitrary geo
// shapes such as rectangles and polygons. It should be used when either the
// data being indexed or the queries being executed contain shapes other than
// just points.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html
type GeoShapeField struct {
	orientationParam
	ignoreMalformedParam
	ignoreZValueParam
	nullValueParam
}

func (GeoShapeField) Type() FieldType {
	return FieldTypeGeoShape
}
func (gs *GeoShapeField) Field() (Field, error) {
	return gs, nil
}
func (gs *GeoShapeField) UnmarshalJSON(data []byte) error {
	var params GeoShapeFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.GeoShape()
	*gs = *v
	return err
}

func (gs GeoShapeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(geoShapeField{
		IgnoreMalformed: gs.ignoreMalformed.Value(),
		IgnoreZValue:    gs.ignoreZ.Value(),
		Orientation:     gs.orientation,
		Type:            gs.Type(),
	})
}

func NewGeoShapeField(params GeoShapeFieldParams) (*GeoShapeField, error) {
	return params.GeoShape()
}

//easyjson:json
type geoShapeField struct {
	Orientation     Orientation `json:"orientation,omitempty"`
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	IgnoreZValue    interface{} `json:"ignore_z_value,omitempty"`
	Type            FieldType   `json:"type"`
}
