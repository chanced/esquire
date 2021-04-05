package picker

import (
	"reflect"

	"encoding/json"

	"github.com/chanced/dynamic"
)

type Ranger interface {
	Range() (*RangeQuery, error)
}

// RangeQueryParams returns documents that contain terms within a provided range.
type RangeQueryParams struct {
	Field                string
	GreaterThan          interface{}
	GreaterThanOrEqualTo interface{}
	LessThan             interface{}
	LessThanOrEqualTo    interface{}
	Format               string
	TimeZone             string
	Boost                interface{}
	Name                 string
	Relation             Relation
	completeClause
}

func (r RangeQueryParams) Clause() (QueryClause, error) {
	return r.Range()
}
func (r RangeQueryParams) Range() (*RangeQuery, error) {
	q := &RangeQuery{field: r.Field}
	err := q.setGreaterThan(r.GreaterThan)
	if err != nil {
		return q, newQueryError(err, QueryKindRange, r.Field)
	}
	err = q.setGreaterThanOrEqualTo(r.GreaterThanOrEqualTo)
	if err != nil {
		return q, newQueryError(err, QueryKindRange, r.Field)
	}
	err = q.setLessThan(r.LessThan)
	if err != nil {
		return q, newQueryError(err, QueryKindRange, r.Field)
	}
	err = q.setLessThanOrEqualTo(r.LessThanOrEqualTo)
	if err != nil {
		return q, newQueryError(err, QueryKindRange, r.Field)
	}
	err = q.SetRelation(r.Relation)
	if err != nil {
		return q, newQueryError(err, QueryKindRange, r.Field)
	}
	err = q.SetBoost(r.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindRange, r.Field)
	}
	q.SetFormat(r.Format)
	q.SetTimeZone(r.TimeZone)
	return q, nil
}

func (r RangeQueryParams) Kind() QueryKind {
	return QueryKindBoolean
}

type RangeQuery struct {
	field                string
	greaterThan          dynamic.StringNumberOrTime
	greaterThanOrEqualTo dynamic.StringNumberOrTime
	lessThan             dynamic.StringNumberOrTime
	lessThanOrEqualTo    dynamic.StringNumberOrTime
	relationParam
	formatParam
	timeZoneParam
	boostParam
	nameParam
	completeClause
}

func (r *RangeQuery) Clause() (QueryClause, error) {
	return r, nil
}
func (RangeQuery) Kind() QueryKind {
	return QueryKindRange
}

func (r *RangeQuery) Set(field string, ranger Ranger) error {
	q, err := ranger.Range()
	if err != nil {
		return newQueryError(err, QueryKindRange, field)
	}
	*r = *q
	return nil
}

func (r RangeQuery) GreaterThan() dynamic.StringNumberOrTime {
	return r.greaterThan
}

func (r *RangeQuery) setGreaterThan(value interface{}) error {
	err := r.greaterThan.Set(value)
	if err != nil {
		return newQueryError(err, QueryKindRange, r.field)
	}
	return nil
}

func (r RangeQuery) GreaterThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.greaterThan
}

func (r *RangeQuery) setGreaterThanOrEqualTo(value interface{}) error {
	err := r.greaterThanOrEqualTo.Set(value)
	if err != nil {
		return newQueryError(err, QueryKindRange, r.field)
	}
	return nil
}
func (r RangeQuery) LessThan() dynamic.StringNumberOrTime {
	return r.lessThan
}

func (r *RangeQuery) setLessThan(value interface{}) error {
	err := r.lessThan.Set(value)
	if err != nil {
		return newQueryError(err, QueryKindRange, r.field)
	}
	return nil
}

func (r RangeQuery) LessThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.lessThanOrEqualTo
}

func (r *RangeQuery) setLessThanOrEqualTo(value interface{}) error {
	err := r.lessThanOrEqualTo.Set(value)
	if err != nil {

		return newQueryError(err, QueryKindRange, r.field)
	}
	return nil
}
func (r *RangeQuery) IsEmpty() bool {
	return r == nil || !(!r.greaterThan.IsNilOrEmpty() ||
		!r.greaterThanOrEqualTo.IsNilOrEmpty() ||
		!r.lessThanOrEqualTo.IsNilOrEmpty() ||
		!r.lessThan.IsNilOrEmpty())
}
func (r RangeQuery) values() map[string]dynamic.StringNumberOrTime {
	return map[string]dynamic.StringNumberOrTime{
		"gt":  r.greaterThan,
		"gte": r.greaterThanOrEqualTo,
		"lt":  r.lessThanOrEqualTo,
		"lte": r.lessThanOrEqualTo,
	}
}
func (r RangeQuery) MarshalJSON() ([]byte, error) {
	if r.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := r.marshalClauseJSON()

	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{r.field: data})
}

func (r RangeQuery) marshalClauseJSON() (dynamic.JSON, error) {
	data, err := marshalClauseParams(&r)
	if err != nil {
		return nil, err
	}
	for key, value := range r.values() {
		if !value.IsNilOrEmpty() {
			data[key] = r.greaterThan
		}
	}
	return json.Marshal(data)
}

func (r *RangeQuery) UnmarshalJSON(data []byte) error {
	*r = RangeQuery{}
	obj := dynamic.JSONObject{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}
	for k, v := range obj {
		r.field = k
		return r.unmarshalClauseJSON(v)
	}
	return nil
}

func (r *RangeQuery) unmarshalClauseJSON(data dynamic.JSON) error {
	fields, err := unmarshalClauseParams(data, r)
	if err != nil {
		return err
	}
	for k, v := range r.values() {
		if fd, ok := fields[k]; ok {
			var val interface{}
			err := json.Unmarshal(fd, &val)
			if err != nil {
				return err
			}
			err = v.Set(val)
			if err != nil {
				return &json.UnmarshalTypeError{
					Value: string(fd),
					Type:  reflect.TypeOf(dynamic.StringNumberOrTime{}),
				}
			}
		}
	}
	return nil
}
func (r *RangeQuery) Clear() {
	*r = RangeQuery{}
}
