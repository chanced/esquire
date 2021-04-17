package picker

type DistanceFeaturer interface {
	DistanceFeature() (*DistanceFeatureQuery, error)
}

type DistanceFeatureQueryParams struct {
	Name string
	// (Required, string) Name of the field used to calculate distances.
	//
	// This field must meet the following criteria:
	//
	// - Be a date, date_nanos or geo_point field
	//
	// - Have an index mapping parameter value of true, which is the default
	//
	// - Have an doc_values mapping parameter value of true, which is the
	// default
	Field string
	// (Required, string) Date or point of origin used to calculate distances.
	//
	// If the field value is a date or date_nanos field, the origin value must
	// be a date. Date Math, such as now-1h, is supported.
	//
	// If the field value is a geo_point field, the origin value must be a
	// geopoint.
	Origin string

	// (Required, time unit or distance unit) Distance from the origin at which
	// relevance scores receive half of the boost value.
	//
	// If the field value is a date or date_nanos field, the pivot value must be a
	// time unit, such as 1h or 10d.
	//
	// If the field value is a geo_point field, the pivot value must be a distance
	// unit, such as 1km or 12m.
	Pivot string

	// (Optional, float) Floating point number used to multiply the relevance
	// score of matching documents. This value cannot be negative. Defaults to
	// 1.0.
	Boost interface{}

	completeClause
}

func (DistanceFeatureQueryParams) Kind() QueryKind {
	return QueryKindDistanceFeature
}

func (p DistanceFeatureQueryParams) Clause() (QueryClause, error) {
	return p.DistanceFeature()
}
func (p DistanceFeatureQueryParams) DistanceFeature() (*DistanceFeatureQuery, error) {
	q := &DistanceFeatureQuery{}
	err := q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindDistanceFeature)
	}
	err = q.SetOrigin(p.Origin)
	if err != nil {
		return q, newQueryError(err, QueryKindDistanceFeature, q.field)
	}

	err = q.SetPivot(p.Pivot)
	if err != nil {
		return q, newQueryError(err, QueryKindDistanceFeature, q.field)
	}
	err = q.SetOrigin(p.Origin)
	if err != nil {
		return q, newQueryError(err, QueryKindDistanceFeature, q.field)
	}
	err = q.SetPivot(p.Pivot)
	if err != nil {
		return q, newQueryError(err, QueryKindDistanceFeature, q.field)
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindDistanceFeature, q.field)
	}
	q.SetName(p.Name)
	return q, nil
}

type DistanceFeatureQuery struct {
	nameParam
	boostParam
	fieldParam
	pivot  string
	origin string
	completeClause
}

func (q DistanceFeatureQuery) Pivot() string {
	return q.pivot
}

func (q *DistanceFeatureQuery) SetPivot(pivot string) error {
	if len(pivot) == 0 {
		return newQueryError(ErrPivotRequired, QueryKindDistanceFeature, q.field)
	}
	q.pivot = pivot
	return nil
}

func (q DistanceFeatureQuery) Origin() string {
	return q.origin
}

func (q *DistanceFeatureQuery) SetOrigin(origin string) error {
	if len(origin) == 0 {
		return newQueryError(ErrOriginRequired, QueryKindDistanceFeature, q.field)
	}
	q.origin = origin
	return nil
}

func (DistanceFeatureQuery) Kind() QueryKind {
	return QueryKindDistanceFeature
}
func (q *DistanceFeatureQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *DistanceFeatureQuery) DistanceFeature() (*DistanceFeatureQuery, error) {
	return q, nil
}
func (q *DistanceFeatureQuery) Clear() {
	if q == nil {
		return
	}
	*q = DistanceFeatureQuery{}
}

func (q *DistanceFeatureQuery) UnmarshalJSON(data []byte) error {
	*q = DistanceFeatureQuery{}
	p := distanceFeatureQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	q.field = p.Field
	q.name = p.Name
	q.origin = p.Origin
	q.pivot = p.Pivot
	err = q.boost.Set(p.Boost)
	if err != nil {
		return err
	}
	return nil
}

func (q DistanceFeatureQuery) MarshalJSON() ([]byte, error) {
	return distanceFeatureQuery{
		Name:   q.name,
		Field:  q.field,
		Origin: q.origin,
		Pivot:  q.pivot,
		Boost:  q.boost.Value(),
	}.MarshalJSON()
}
func (q *DistanceFeatureQuery) IsEmpty() bool {
	return q == nil || len(q.field) == 0
}
func (q *DistanceFeatureQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}
func (q DistanceFeatureQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

//easyjson:json
type distanceFeatureQuery struct {
	Name   string      `json:"_name,omitempty"`
	Field  string      `json:"field"`
	Origin string      `json:"origin"`
	Pivot  string      `json:"pivot"`
	Boost  interface{} `json:"boost,omitempty"`
}
