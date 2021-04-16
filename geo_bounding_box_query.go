package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type GeoBoundingBoxer interface {
	GeoBoundingBox() (*GeoBoundingBoxQuery, error)
}
type GeoBoundingBoxQueryParams struct {
	Type        string        `json:"type,omitempty"`
	BoundingBox BoundingBoxer `json:"bounding_box"`
	Field       string        `json:"field"`
	Name        string        `json:"_name,omitempty"`
	completeClause
}

func (GeoBoundingBoxQueryParams) Kind() QueryKind {
	return QueryKindGeoBoundingBox
}
func (p GeoBoundingBoxQueryParams) Clause() (QueryClause, error) {
	return p.GeoBoundingBox()
}
func (p GeoBoundingBoxQueryParams) GeoBoundingBox() (*GeoBoundingBoxQuery, error) {
	q := &GeoBoundingBoxQuery{}
	q.SetBoundingBox(p.BoundingBox)
	q.SetName(p.Name)
	err := q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	return q, nil
}

type GeoBoundingBoxQuery struct {
	typ            string
	boundingBox    interface{}
	boundingBoxRaw dynamic.JSON
	fieldParam
	nameParam
	completeClause
}

func (g GeoBoundingBoxQuery) Type() string {
	return g.typ
}
func (g *GeoBoundingBoxQuery) SetType(typ string) {
	g.typ = typ
}
func (g *GeoBoundingBoxQuery) SetBoundingBox(bb BoundingBoxer) {
	g.boundingBox = bb.BoundingBox()
}
func (g GeoBoundingBoxQuery) BoundingBox() interface{} {
	return g.boundingBox
}

func (g GeoBoundingBoxQuery) DecodeBoundingBox(v interface{}) error {
	if len(g.boundingBoxRaw) == 0 {
		d, err := json.Marshal(g.boundingBox)
		if err != nil {
			return err
		}
		g.boundingBoxRaw = d
	}
	return json.Unmarshal(g.boundingBoxRaw, &v)
}
func (GeoBoundingBoxQuery) Kind() QueryKind {
	return QueryKindGeoBoundingBox
}

func (g *GeoBoundingBoxQuery) Clause() (QueryClause, error) {
	return g, nil
}
func (g *GeoBoundingBoxQuery) GeoBoundingBox() (*GeoBoundingBoxQuery, error) {
	return g, nil
}
func (g *GeoBoundingBoxQuery) UnmarshalBSON(data []byte) error {
	return g.UnmarshalJSON(data)
}

func (g *GeoBoundingBoxQuery) UnmarshalJSON(data []byte) error {
	*g = GeoBoundingBoxQuery{}

	var obj dynamic.JSONObject
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	for field, d := range obj {
		var err error
		switch field {
		case "_name":
			var name string
			err = json.Unmarshal(d, &name)
			g.name = name
		case "type":
			var typ string
			err = json.Unmarshal(d, &typ)
			g.typ = typ
		default:
			g.field = field
			g.boundingBoxRaw = d
			var bb interface{}
			err = json.Unmarshal(d, &bb)
			g.boundingBox = bb
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (g GeoBoundingBoxQuery) MarshalBSON() ([]byte, error) {
	return g.MarshalJSON()
}

func (g GeoBoundingBoxQuery) MarshalJSON() ([]byte, error) {

	bx, err := json.Marshal(g.boundingBox)
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		g.field: bx,
	}
	if len(g.typ) > 0 {
		typ, err := json.Marshal(g.typ)
		if err != nil {
			return nil, err
		}
		obj["type"] = typ
	}
	if len(g.name) > 0 {
		name, err := json.Marshal(g.name)
		if err != nil {
			return nil, err
		}
		obj["_name"] = name
	}
	return obj.MarshalJSON()
}
func (g *GeoBoundingBoxQuery) IsEmpty() bool {
	return g == nil || g.boundingBox == nil || len(g.field) == 0
}

func (g *GeoBoundingBoxQuery) Clear() {
	if g == nil {
		return
	}
	*g = GeoBoundingBoxQuery{}
}
