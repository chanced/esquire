package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Ranger interface {
	Range() (*RangeQuery, error)
}

// Range returns documents that contain terms within a provided range.
type Range struct {
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
}

func (r Range) field() string {
	return r.Field
}

func (r Range) Clause() (Clause, error) {
	return r.Range()
}
func (r Range) Range() (*RangeQuery, error) {
	q := &RangeQuery{field: r.Field}
	err := q.setGreaterThan(r.GreaterThan)
	if err != nil {
		return q, NewQueryError(err, TypeRange, r.Field)
	}
	err = q.setGreaterThan(r.GreaterThanOrEqualTo)
	if err != nil {
		return q, NewQueryError(err, TypeRange, r.Field)
	}
	err = q.setLessThan(r.LessThan)
	if err != nil {
		return q, NewQueryError(err, TypeRange, r.Field)
	}
	err = q.setLessThanOrEqualTo(r.LessThanOrEqualTo)
	if err != nil {
		return q, NewQueryError(err, TypeRange, r.Field)
	}
	err = q.SetRelation(r.Relation)
	if err != nil {
		return q, NewQueryError(err, TypeRange, r.Field)
	}
	q.SetFormat(r.Format)
	err = q.SetBoost(r.Boost)
	if err != nil {
		return q, NewQueryError(err, TypeRange, r.Field)
	}
	q.SetTimeZone(r.TimeZone)
	return q, nil
}

func (r Range) Type() Type {
	return TypeBoolean
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
}

func (RangeQuery) Type() Type {
	return TypeRange
}

func (r *RangeQuery) Set(field string, ranger Ranger) error {
	q, err := ranger.Range()
	if err != nil {
		return NewQueryError(err, TypeRange, field)
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
		return NewQueryError(err, TypeRange, r.field)
	}
	return nil
}

func (r RangeQuery) GreaterThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.greaterThan
}

func (r *RangeQuery) setGreaterThanOrEqualTo(value interface{}) error {
	err := r.greaterThanOrEqualTo.Set(value)
	if err != nil {
		return NewQueryError(err, TypeRange, r.field)
	}
	return nil
}
func (r RangeQuery) LessThan() dynamic.StringNumberOrTime {
	return r.lessThan
}

func (r *RangeQuery) setLessThan(value interface{}) error {
	err := r.lessThan.Set(value)
	if err != nil {
		return NewQueryError(err, TypeRange, r.field)
	}
	return nil
}

func (r RangeQuery) LessThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.lessThanOrEqualTo
}

func (r *RangeQuery) setLessThanOrEqualTo(value interface{}) error {
	err := r.lessThanOrEqualTo.Set(value)
	if err != nil {

		return NewQueryError(err, TypeRange, r.field)
	}
	return nil
}

func (r RangeQuery) marshalClauseJSON() (dynamic.JSON, error) {
	data, err := marshalParams(&r)
	if err != nil {
		return nil, err
	}
	if !r.greaterThan.IsNilOrEmpty() {
		data["gt"] = r.greaterThan
	}
	if !r.greaterThanOrEqualTo.IsNilOrEmpty() {
		data["gte"] = r.greaterThanOrEqualTo
	}
	if !r.lessThan.IsNilOrEmpty() {
		data["lt"] = r.lessThan
	}
	if !r.lessThanOrEqualTo.IsNilOrEmpty() {
		data["lte"] = r.lessThanOrEqualTo
	}
	return json.Marshal(data)
}

func (r *RangeQuery) Clear() {
	*r = RangeQuery{}
}
