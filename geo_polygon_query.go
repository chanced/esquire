package picker

type GeoPolygonQueryParams struct {
	Name string
	completeClause
}

func (GeoPolygonQueryParams) Kind() QueryKind {
	return QueryKindGeoPolygon
}

func (p GeoPolygonQueryParams) Clause() (QueryClause, error) {
	return p.GeoPolygon()
}
func (p GeoPolygonQueryParams) GeoPolygon() (*GeoPolygonQuery, error) {
	q := &GeoPolygonQuery{}
	_ = q
	panic("not implemented")
	// return q, nil
}

type GeoPolygonQuery struct {
	nameParam
	completeClause
}

func (GeoPolygonQuery) Kind() QueryKind {
	return QueryKindGeoPolygon
}
func (q *GeoPolygonQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *GeoPolygonQuery) GeoPolygon() (QueryClause, error) {
	return q, nil
}
func (q *GeoPolygonQuery) Clear() {
	if q == nil {
		return
	}
	*q = GeoPolygonQuery{}
}
func (q *GeoPolygonQuery) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}
func (q GeoPolygonQuery) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}
func (q *GeoPolygonQuery) IsEmpty() bool {
	panic("not implemented")
}

type geoPolygonQuery struct {
	Name string `json:"_name,omitempty"`
}
