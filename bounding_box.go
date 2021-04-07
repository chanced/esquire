package picker

import (
	"encoding/json"

	jwriter "github.com/mailru/easyjson/jwriter"
)

type BoundingBoxer interface {
	BoundingBox() interface{}
	json.Marshaler
}

var _ BoundingBoxer = (*BoundingBox)(nil)
var _ BoundingBoxer = (*Vertices)(nil)
var _ BoundingBoxer = (*BoundingBox)(nil)

// var _ BoundingBoxer = (*WKT)(nil)

// Vertices for GeoBoundingBoxQuery
//easyjson:json
type Vertices struct {
	Top    interface{} `json:"top"`
	Bottom interface{} `json:"bottom"`
	Left   interface{} `json:"left"`
	Right  interface{} `json:"right"`
}

func (v Vertices) BoundingBox() interface{} {
	return v
}

// BoundingBox for GeoBoundingBoxQuery
//easyjson:json
type BoundingBox struct {
	TopLeft     interface{} `json:"top_left"`
	BottomRight interface{} `json:"bottom_right"`
}

func (bb BoundingBox) BoundingBox() interface{} {
	return bb
}

type WKT string

func (wkt WKT) BoundingBox() interface{} {
	return wkt
}
func (wkt WKT) String() string {
	return string(wkt)
}
func (wkt WKT) MarshalJSON() ([]byte, error) {
	return wktVal{WKT: wkt.String()}.MarshalJSON()
}

func (wkt WKT) MarshalEasyJSON(w *jwriter.Writer) {
	wktVal{WKT: wkt.String()}.MarshalEasyJSON(w)
}

//easyjson:json
type wktVal struct {
	WKT string `json:"wkt"`
}
