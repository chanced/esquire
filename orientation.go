package picker

type Orientation string

const (
	// OrientationRight - Counter-Clockwise Orientation (RIGHT)
	OrientationRight Orientation = "right"
	// OrientationCounterClockwise - Counter-Cloockwise orientation (RIGHT)
	OrientationCounterClockwise Orientation = "counterclockwise"
	// OrientationCCW - Counter-clockwise orientation (RIGHT)
	OrientationCCW Orientation = "ccw"
	// OrientationLeft - Clockwise Orientation
	OrientationLeft Orientation = "left"
	// OrientationClockwise - Clockwise Oreintation (LEFT)
	OrientationClockwise Orientation = "left"
	// OrientationCW Clockwise Orientation (LEFT)
	OrientationCW Orientation = "cw"
)

var CounterClockwiseOrientations = []Orientation{OrientationRight, OrientationCCW, OrientationCounterClockwise}
var ClockwiseOrientations = []Orientation{OrientationLeft, OrientationCW, OrientationClockwise}

func (o Orientation) String() string {
	return string(o)
}

// IsCounterClockwise indicates whether the Orientation is Counter-Clockwise
func (o Orientation) IsCounterClockwise() bool {
	for _, ccw := range CounterClockwiseOrientations {
		if ccw == o {
			return true
		}
	}
	return false
}

// IsClockwise indicates whether the Orientation is Clockwise
func (o Orientation) IsClockwise() bool {
	for _, cw := range ClockwiseOrientations {
		if cw == o {
			return true
		}
	}
	return false
}

// WithOrientation is a mapping with the orientation parameter
type WithOrientation interface {

	// Orientation is the vertex order for the shape’s coordinates list.
	//
	// Defaults to OrientationRight (RIGHT) to comply with OGC standards. OGC
	// standards define outer ring vertices in counterclockwise order with inner
	// ring (hole) vertices in clockwise order.
	Orientation() Orientation
	// SetOrientation sets the Orientation Value to v
	SetOrientation(v Orientation)
}

// FieldWithOrientation is a Field with the orientation paramater
type FieldWithOrientation interface {
	Field
	WithOrientation
}

// OrientationParam is a mixin that adds the orientation parameter
//
// Vertex order for the shape’s coordinates list.
//
// This parameter sets and returns only a RIGHT (counterclockwise) or LEFT
// (clockwise) value. However, you can specify either value in multiple ways.
//
// To set RIGHT, use one of the following arguments or its uppercase variant:
//
//  OrientationRight
//  OrientationCounterClockwise
//  OrientationCCW
//
// To set LEFT, use one of the following arguments or its uppercase variant:
//
//  OrientationLeft
//  OrientationClockwise
//  OrientationCW
//
// Defaults to OrientationRight (RIGHT) to comply with OGC standards. OGC
// standards define outer ring vertices in counterclockwise order with inner
// ring (hole) vertices in clockwise order.
//
// Individual GeoJSON or WKT documents can override this parameter.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html#geo-shape-mapping-options
type OrientationParam struct {
	OrientationValue *Orientation `bson:"orientation,omitempty" json:"orientation,omitempty"`
}

// Orientation is the vertex order for the shape’s coordinates list.
//
// Defaults to OrientationRight (RIGHT) to comply with OGC standards. OGC
// standards define outer ring vertices in counterclockwise order with inner
// ring (hole) vertices in clockwise order.
func (o OrientationParam) Orientation() Orientation {
	if o.OrientationValue == nil {
		return OrientationRight
	}
	return *o.OrientationValue
}

// SetOrientation sets the Orientation Value to v
func (o *OrientationParam) SetOrientation(v Orientation) {
	if o.Orientation() != v {
		o.OrientationValue = &v
	}
}
