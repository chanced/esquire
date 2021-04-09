package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

const DefaultDistanceType = DistanceTypeArc

type GeoDistancer interface {
	GeoDistance() (*GeoDistanceQuery, error)
}

type DistanceType string

const (
	DistanceTypeArc   = "arc"
	DistanceTypePlane = "plane"
)

const DefaultDistanceValidationMethod = ValidationMethodStrict

type ValidationMethod string

const (
	ValidationMethodStrict          ValidationMethod = "STRICT"
	ValidationMethodIgnoreMalformed ValidationMethod = "IGNORE_MALFORMED"
	ValidationMethodCoerce          ValidationMethod = "COERCE"
)

// type GeoDistancePropertiesQueryParams struct {
// 	Distance     string
// 	DistanceType DistanceType
// 	Name string
// 	completeClause
// }

// LatLon properties
//easyjson:json
type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type GeoDistanceQueryParams struct {
	Distance         string
	Field            string
	DistanceType     DistanceType
	Name             string
	ValidationMethod ValidationMethod
	GeoPoint         interface{}
	completeClause
}

func (GeoDistanceQueryParams) Kind() QueryKind {
	return QueryKindGeoDistance
}

func (p GeoDistanceQueryParams) Clause() (QueryClause, error) {
	return p.GeoDistance()
}
func (p GeoDistanceQueryParams) GeoDistance() (*GeoDistanceQuery, error) {
	q := &GeoDistanceQuery{}
	err := q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoDistance)
	}
	err = q.SetDistance(p.Distance)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoDistance, p.Field)
	}
	err = q.SetDistanceType(p.DistanceType)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoDistance, p.Field)
	}
	err = q.SetGeoPoint(p.GeoPoint)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoDistance, p.Field)
	}
	err = q.SetValidationMethod(p.ValidationMethod)
	if err != nil {
		return q, newQueryError(err, QueryKindGeoDistance, p.Field)
	}
	q.SetName(p.Name)
	return q, nil
}

type GeoDistanceQuery struct {
	nameParam
	distance string
	fieldParam
	validationMethod ValidationMethod
	distanceType     DistanceType
	geoPoint         interface{}
	geoPointRaw      dynamic.JSON
	completeClause
}

func (q GeoDistanceQuery) DistanceType() DistanceType {
	if len(q.distanceType) > 0 {
		return q.distanceType
	}
	return DefaultDistanceType
}
func (q *GeoDistanceQuery) SetDistanceType(typ DistanceType) error {
	// skipping validation at the moment. It may all be removed
	q.distanceType = typ
	return nil
}
func (q GeoDistanceQuery) ValidationMethod() ValidationMethod {
	if len(q.validationMethod) > 0 {
		return q.validationMethod
	}
	return DefaultDistanceValidationMethod
}
func (q *GeoDistanceQuery) SetValidationMethod(method ValidationMethod) error {
	// skipping validation at the moment. It may all be removed
	q.validationMethod = method
	return nil
}

func (q GeoDistanceQuery) Distance() string {
	return q.distance
}
func (q *GeoDistanceQuery) SetDistance(distance string) error {
	if len(distance) == 0 {
		return ErrDistanceRequired
	}
	q.distance = distance
	return nil
}
func (q GeoDistanceQuery) GeoPoint() interface{} {
	return q.geoPoint
}

func (q *GeoDistanceQuery) SetGeoPoint(v interface{}) error {
	if v == nil {
		return ErrGeoPointRequired
	}
	q.geoPoint = v
	return nil
}

func (q GeoDistanceQuery) DecodeGeoPoint(v interface{}) error {
	if len(q.geoPointRaw) == 0 {
		d, err := json.Marshal(q.geoPoint)
		if err != nil {
			return err
		}
		q.geoPointRaw = d
	}
	return json.Unmarshal(q.geoPointRaw, v)
}

func (GeoDistanceQuery) Kind() QueryKind {
	return QueryKindGeoDistance
}
func (q *GeoDistanceQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *GeoDistanceQuery) GeoDistance() (*GeoDistanceQuery, error) {
	return q, nil
}

func (q *GeoDistanceQuery) Clear() {
	if q == nil {
		return
	}
	*q = GeoDistanceQuery{}
}
func (q *GeoDistanceQuery) UnmarshalJSON(data []byte) error {
	*q = GeoDistanceQuery{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for k, d := range obj {
		switch k {
		case "_name":
			err = unmarshalNameParam(d, q)
		case "distance_type":
			var dt DistanceType
			err = json.Unmarshal(d, &dt)
			q.distanceType = dt
		case "validation_method":
			var vm ValidationMethod
			err = json.Unmarshal(d, &vm)
			q.validationMethod = vm
		case "distance":
			var distance string
			err = json.Unmarshal(d, &distance)
			q.distance = distance
		default:
			q.field = k
			q.geoPointRaw = d
			var geoPoint interface{}
			err = json.Unmarshal(d, &geoPoint)
			q.geoPoint = geoPoint
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (q GeoDistanceQuery) MarshalJSON() ([]byte, error) {
	geoPoint, err := json.Marshal(q.geoPoint)
	if err != nil {
		return nil, err
	}
	distance, err := json.Marshal(q.distance)
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{q.field: geoPoint, "distance": distance}
	if len(q.name) > 0 {
		name, err := marshalNameParam(&q)
		if err != nil {
			return nil, err
		}
		obj["_name"] = name
	}
	if len(q.distanceType) > 0 {
		dt, err := json.Marshal(q.distanceType)
		if err != nil {
			return nil, err
		}
		obj["distance_type"] = dt
	}
	if len(q.validationMethod) > 0 {
		vm, err := json.Marshal(q.validationMethod)
		if err != nil {
			return nil, err
		}
		obj["validation_method"] = vm
	}
	return obj.MarshalJSON()
}
func (q *GeoDistanceQuery) IsEmpty() bool {
	return q == nil || len(q.field) == 0
}
