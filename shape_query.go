package picker

import "github.com/chanced/dynamic"

type ShapeQuerier interface {
	ShapeQuery() (*ShapeQuery, error)
}
type ShapeQueryParams struct {
	Field          string
	Name           string
	Relation       SpatialRelation
	Shape          interface{}
	IndexedShape   *IndexedShape
	IgnoreUnmapped bool
	completeClause
}

func (ShapeQueryParams) Kind() QueryKind {
	return QueryKindShape
}

func (p ShapeQueryParams) Clause() (QueryClause, error) {
	return p.ShapeQuery()
}
func (p ShapeQueryParams) ShapeQuery() (*ShapeQuery, error) {
	q := &ShapeQuery{}
	q.SetName(p.Name)
	err := q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindShape)
	}
	q.SetIgnoreUnmapped(p.IgnoreUnmapped)
	err = q.SetIndexedShape(p.IndexedShape)
	if err != nil {
		return q, newQueryError(err, QueryKindShape, p.Field)
	}
	err = q.SetShape(p.Shape)
	if err != nil {
		return q, newQueryError(err, QueryKindShape, p.Field)
	}
	err = q.SetRelation(p.Relation)
	if err != nil {
		return q, newQueryError(err, QueryKindShape, p.Field)
	}
	return q, nil
}

type ShapeQuery struct {
	nameParam
	completeClause
	fieldParam
	shape          interface{}
	indexedShape   *IndexedShape
	relation       SpatialRelation
	ignoreUnmapped bool
}

func (q ShapeQuery) IgnoreUnmapped() bool {
	return q.ignoreUnmapped
}

func (q ShapeQuery) Relation() SpatialRelation {
	if len(q.relation) == 0 {
		return DefaultSpatialRelation
	}
	return q.relation
}
func (q *ShapeQuery) SetRelation(v SpatialRelation) error {
	q.relation = v
	return nil
}
func (q *ShapeQuery) SetIgnoreUnmapped(v bool) {
	q.ignoreUnmapped = v
}
func (q *ShapeQuery) SetShape(shape interface{}) error {
	// not going to error check for now but leaving
	// the error return just incase it makes sense later
	q.shape = shape
	return nil
}
func (q ShapeQuery) Shape() interface{} {
	return q.shape
}
func (q *ShapeQuery) SetIndexedShape(indexed *IndexedShape) error {
	// leaving the error just incase I decide to validate indexed
	q.indexedShape = indexed
	return nil
}
func (ShapeQuery) Kind() QueryKind {
	return QueryKindShape
}
func (q *ShapeQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *ShapeQuery) ShapeQuery() (*ShapeQuery, error) {
	return q, nil
}
func (q *ShapeQuery) Clear() {
	if q == nil {
		return
	}
	*q = ShapeQuery{}
}
func (q *ShapeQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *ShapeQuery) UnmarshalJSON(data []byte) error {
	*q = ShapeQuery{}
	field, fd, err := unmarshalField(data)
	if err != nil {
		return err
	}
	if fd == nil {
		return nil
	}
	q.field = field
	v := shapeQuery{}
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
func (q ShapeQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q ShapeQuery) MarshalJSON() ([]byte, error) {
	v := shapeQuery{
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
func (q *ShapeQuery) IsEmpty() bool {
	return q == nil || len(q.field) == 0 || (q.shape == nil && q.indexedShape == nil)
}

//easyjson:json
type shapeQuery struct {
	Name           string          `json:"_name,omitempty"`
	Relation       SpatialRelation `json:"relation,omitempty"`
	Shape          interface{}     `json:"shape,omitempty"`
	IndexedShape   *IndexedShape   `json:"indexed_shape,omitempty"`
	IgnoreUnmapped bool            `json:"ignore_unmapped,omitempty"`
}
