package picker

import "github.com/chanced/dynamic"

type GeoShaper interface {
	GeoShape() (*GeoShapeQuery, error)
}
type GeoShapeQueryParams struct {
	Field          string
	Name           string
	Relation       SpatialRelation
	Shape          interface{}
	IndexedShape   *IndexedShape
	IgnoreUnmapped bool
	completeClause
}

func (GeoShapeQueryParams) Kind() QueryKind {
	return QueryKindGeoShape
}

func (p GeoShapeQueryParams) Clause() (QueryClause, error) {
	return p.GeoShape()
}
func (p GeoShapeQueryParams) GeoShape() (*GeoShapeQuery, error) {
	q := &GeoShapeQuery{}
	q.SetName(p.Name)
	err := q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoShape)
	}
	q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	err = q.SetIndexedShape(p.IndexedShape)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoShape, p.Field)
	}
	err = q.SetShape(p.Shape)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoShape, p.Field)
	}
	err = q.SetRelation(p.Relation)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoShape, p.Field)
	}
	return q, nil
}

type GeoShapeQuery struct {
	nameParam
	completeClause
	fieldParam
	shape          interface{}
	indexedShape   *IndexedShape
	relation       SpatialRelation
	ignoreUnmapped bool
}

func (q GeoShapeQuery) IgnoreUnmapped() bool {
	return q.ignoreUnmapped
}

func (q GeoShapeQuery) Relation() SpatialRelation {
	if len(q.relation) == 0 {
		return DefaultSpatialRelation
	}
	return q.relation
}
func (q *GeoShapeQuery) SetRelation(v SpatialRelation) error {
	q.relation = v
	return nil
}
func (q *GeoShapeQuery) SetIgnoreUnmapped(v bool) {
	q.ignoreUnmapped = v
}
func (q *GeoShapeQuery) SetShape(shape interface{}) error {
	// not going to error check for now but leaving
	// the error return just incase it makes sense later
	q.shape = shape
	return nil
}
func (q GeoShapeQuery) Shape() interface{} {
	return q.shape
}
func (q *GeoShapeQuery) SetIndexedShape(indexed *IndexedShape) error {
	// leaving the error just incase I decide to validate indexed
	q.indexedShape = indexed
	return nil
}
func (GeoShapeQuery) Kind() QueryKind {
	return QueryKindGeoShape
}
func (q *GeoShapeQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *GeoShapeQuery) GeoShape() (*GeoShapeQuery, error) {
	return q, nil
}
func (q *GeoShapeQuery) Clear() {
	if q == nil {
		return
	}
	*q = GeoShapeQuery{}
}
func (q *GeoShapeQuery) UnmarshalJSON(data []byte) error {
	*q = GeoShapeQuery{}
	field, fd, err := unmarshalField(data)
	if err != nil {
		return err
	}
	if fd == nil {
		return nil
	}
	q.field = field
	v := geoShapeQuery{}
	err = v.UnmarshalJSON(fd)
	if err != nil {
		return err
	}
	q.shape = v.Shape
	q.indexedShape = v.IndexedShape
	q.ignoreUnmapped = v.IgnoreUnmapped
	q.relation = v.Relation
	q.name = v.Name
	return nil
}
func (q GeoShapeQuery) MarshalJSON() ([]byte, error) {
	v := geoShapeQuery{
		Name:           q.name,
		Relation:       q.relation,
		Shape:          q.shape,
		IndexedShape:   q.indexedShape,
		IgnoreUnmapped: q.ignoreUnmapped,
	}
	d, err := v.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return dynamic.JSONObject{q.field: d}.MarshalJSON()

}
func (q *GeoShapeQuery) IsEmpty() bool {
	return q == nil || len(q.field) == 0 || (q.shape == nil && q.indexedShape == nil)
}

//easyjson:json
type geoShapeQuery struct {
	Name           string          `json:"_name,omitempty"`
	Relation       SpatialRelation `json:"relation,omitempty"`
	Shape          interface{}     `json:"shape,omitempty"`
	IndexedShape   *IndexedShape   `json:"indexed_shape,omitempty"`
	IgnoreUnmapped bool            `json:"ignore_unmapped,omitempty"`
}

// IndexedShape - The shape query supports using a shape which has already been
// indexed in another index. This is particularly useful for when you have a
// pre-defined list of shapes and you want to reference the list using a logical
// name (for example New Zealand) rather than having to provide coordinates each
// time
//easyjson:json
type IndexedShape struct {
	// The ID of the document that containing the pre-indexed shape.
	ID string `json:"id,omitempty"`
	// Name of the index where the pre-indexed shape is. Defaults to shapes.
	Index string `json:"index,omitempty"`
	// The field specified as path containing the pre-indexed shape. Defaults to
	// shape.
	Path string `json:"path,omitempty"`
	// The routing of the shape document if required.
	Routing string `json:"routing,omitempty"`
}

type Shape struct {
	Coordinates interface{} `json:"coordinates"`
	Type        string      `json:"type"`
}
