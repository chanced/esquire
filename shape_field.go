package picker

import "encoding/json"

type shapeField struct {
	Orientation     Orientation `json:"orientation,omitempty"`
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	IgnoreZValue    interface{} `json:"ignore_z_value,omitempty"`
	NullValue       interface{} `json:"null_value,omitempty"`
	Type            FieldType   `json:"type"`
}

type ShapeFieldParams struct {
	// Optionally define how to interpret vertex order for polygons /
	// multipolygons. This parameter defines one of two coordinate system rules
	// (Right-hand or Left-hand) each of which can be specified in three
	// different ways. 1. Right-hand rule: right, ccw, counterclockwise, 2.
	// Left-hand rule: left, cw, clockwise. The default orientation
	// (counterclockwise) complies with the OGC standard which defines outer
	// ring vertices in counterclockwise order with inner ring(s) vertices
	// (holes) in clockwise order. Setting this parameter in the geo_shape
	// mapping explicitly sets vertex order for the coordinate list of a
	// geo_shape field but can be overridden in each individual GeoJSON or WKT
	// document.
	Orientation Orientation `json:"orientation,omitempty"`

	// If true, malformed GeoJSON or WKT shapes are ignored. If false (default),
	// malformed GeoJSON and WKT shapes throw an exception and reject the entire
	// document.
	IgnoreMalformed interface{} `json:"ignore_malformed,omitempty"`
	// If true (default) three dimension points will be accepted (stored in
	// source) but only latitude and longitude values will be indexed; the third
	// dimension is ignored. If false, geo-points containing any more than
	// latitude and longitude (two dimensions) values throw an exception and
	// reject the whole document.
	IgnoreZValue interface{} `json:"ignore_z_value,omitempty"`
}

func (ShapeFieldParams) Type() FieldType {
	return FieldTypeShape
}
func (p ShapeFieldParams) Field() (Field, error) {
	return p.Shape()
}

func (p ShapeFieldParams) Shape() (*ShapeField, error) {
	f := &ShapeField{}
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

// ShapeField  facilitates the indexing of and searching with arbitrary geo
// shapes such as rectangles and polygons. It should be used when either the
// data being indexed or the queries being executed contain shapes other than
// just points.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html
type ShapeField struct {
	orientationParam
	ignoreMalformedParam
	ignoreZValueParam
	nullValueParam
}

func (ShapeField) Type() FieldType {
	return FieldTypeShape
}
func (gs *ShapeField) Field() (Field, error) {
	return gs, nil
}
func (gs *ShapeField) UnmarshalJSON(data []byte) error {
	var params ShapeFieldParams
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	v, err := params.Shape()
	*gs = *v
	return err
}

func (gs ShapeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(shapeField{
		IgnoreMalformed: gs.ignoreMalformed.Value(),
		NullValue:       gs.nullValue,
		IgnoreZValue:    gs.ignoreZ,
		Orientation:     gs.orientation,
		Type:            gs.Type(),
	})
}

func NewShapeField(params ShapeFieldParams) (*ShapeField, error) {
	return params.Shape()
}
