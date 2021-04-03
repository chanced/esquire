package picker

import (
	"fmt"
	"strings"
)

const DefaultOrientation = OrientationRight

type Orientation string

const (
	OrientationUnspecified Orientation = ""
	// OrientationRight - Counter-Clockwise Orientation (RIGHT)
	OrientationRight Orientation = "right"
	// OrientationCounterClockwise - Counter-Cloockwise orientation (RIGHT)
	OrientationCounterClockwise Orientation = "counterclockwise"
	// OrientationCCW - Counter-clockwise orientation (RIGHT)
	OrientationCCW Orientation = "ccw"
	// OrientationLeft - Clockwise Orientation
	OrientationLeft Orientation = "left"
	// OrientationClockwise - Clockwise Oreintation (LEFT)
	OrientationClockwise Orientation = "clockwise"
	// OrientationCW Clockwise Orientation (LEFT)
	OrientationCW Orientation = "cw"
)

func (o *Orientation) toLower() Orientation {
	*o = Orientation(strings.ToLower(string(*o)))
	return *o
}
func (o *Orientation) IsValid() bool {
	if len(*o) == 0 {
		return true
	}
	n := o.toLower()

	for _, v := range allOrientations {
		if v == n {
			return true
		}
	}
	return false
}

func (o *Orientation) Validate() error {
	if !o.IsValid() {
		strs := make([]string, len(allOrientations))
		for i, v := range allOrientations {
			strs[i] = `"` + v.String() + `"`
		}
		return fmt.Errorf("%w; expected any of [%s]", ErrInvalidOrientation)
	}
	return nil
}

var CounterClockwiseOrientations = []Orientation{OrientationRight, OrientationCCW, OrientationCounterClockwise}
var ClockwiseOrientations = []Orientation{OrientationLeft, OrientationCW, OrientationClockwise}
var allOrientations = append(append([]Orientation{OrientationUnspecified}, CounterClockwiseOrientations...), ClockwiseOrientations...)

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
	SetOrientation(v Orientation) error
}

// FieldWithOrientation is a Field with the orientation paramater
type FieldWithOrientation interface {
	Field
	WithOrientation
}

type orientationParam struct {
	orientation Orientation `json:"orientation,omitempty"`
}

// Orientation is the vertex order for the shape’s coordinates list.
//
// Defaults to OrientationRight (RIGHT) to comply with OGC standards. OGC
// standards define outer ring vertices in counterclockwise order with inner
// ring (hole) vertices in clockwise order.
func (o orientationParam) Orientation() Orientation {
	if o.orientation == "" {
		return DefaultOrientation
	}
	return o.orientation
}

// SetOrientation sets the Orientation Value to v
func (o *orientationParam) SetOrientation(orientation Orientation) error {
	err := orientation.Validate()
	if err != nil {
		return fmt.Errorf("%w; received \"%s\"", err, orientation)
	}
	o.orientation = orientation
	return nil
}
