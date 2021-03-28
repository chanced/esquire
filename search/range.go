package search

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// Range returns documents that contain terms within a provided range.
type Range struct {
	Field                string
	GreaterThan          dynamic.StringNumberOrTime
	GreaterThanOrEqualTo dynamic.StringNumberOrTime
	LessThan             dynamic.StringNumberOrTime
	LessThanOrEqualTo    dynamic.StringNumberOrTime
	Format               string
	TimeZone             string
	Boost                dynamic.Number
	QueryName            string
}

func (r Range) Name() string {
	return r.QueryName
}
func (r Range) Clause() (Clause, error) {
	return r.Range()
}
func (r Range) Range() (*rangeClause, error) {
	q := &rangeClause{}
	q.SetGreaterThan(r.GreaterThan)
	q.SetGreaterThan(r.GreaterThanOrEqualTo)
	q.SetLessThan(r.LessThan)
	q.SetLessThanOrEqualTo(r.LessThanOrEqualTo)
	q.SetFormat(r.Format)
	if b, ok := r.Boost.Float(); ok {
		q.SetBoost(b)
	}
	q.SetTimeZone(r.TimeZone)
	return q, nil
}

func (r Range) Type() Type {
	return TypeBoolean
}

type rangeClause struct {
	greaterThan          dynamic.StringNumberOrTime
	greaterThanOrEqualTo dynamic.StringNumberOrTime
	lessThan             dynamic.StringNumberOrTime
	lessThanOrEqualTo    dynamic.StringNumberOrTime
	formatParam
	timeZoneParam
	boostParam
	nameParam
}

func (r rangeClause) MarshalJSON() ([]byte, error) {
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

func (r *rangeClause) GreaterThan() dynamic.StringNumberOrTime {
	return r.greaterThan
}
func (r *rangeClause) GreaterThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.greaterThan
}

func (r *rangeClause) LessThan() dynamic.StringNumberOrTime {
	return r.lessThan
}

func (r *rangeClause) LessThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.lessThanOrEqualTo
}

func (r *rangeClause) SetGreaterThan(value dynamic.StringNumberOrTime) {
	r.greaterThan = value
}

func (r *rangeClause) SetGreaterThanOrEqualTo(value dynamic.StringNumberOrTime) {
	r.greaterThanOrEqualTo = value
}

func (r *rangeClause) SetLessThan(value dynamic.StringNumberOrTime) {
	r.lessThan = value
}

func (r *rangeClause) SetLessThanOrEqualTo(value dynamic.StringNumberOrTime) {
	r.lessThanOrEqualTo = value
}

func (r rangeClause) Type() Type {
	return TypeBoolean
}

type RangeQuery struct {
	RangeField string
	rangeClause
}
