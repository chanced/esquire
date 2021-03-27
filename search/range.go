package search

import "github.com/chanced/dynamic"

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
	err := q.SetGreaterThan(r.GreaterThan)
	if err != nil {
		return q, err
	}
	err = q.SetGreaterThan(r.GreaterThanOrEqualTo)
	if err != nil {
		return q, err
	}
	err = q.SetLessThan(r.LessThan)
	if err != nil {
		return q, err
	}
	err = q.SetLessThanOrEqualTo(r.LessThanOrEqualTo)
	if err != nil {
		return q, err
	}
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
	GreaterThanValue          dynamic.StringNumberOrTime `json:"gt,omitempty" bson:"gt,omitempty"`
	GreaterThanOrEqualToValue dynamic.StringNumberOrTime `json:"gte,omitempty" bson:"gt,omitempty"`
	LessThanValue             dynamic.StringNumberOrTime `json:"lt,omitempty" bson:"lt,omitempty"`
	LessThanOrEqualToValue    dynamic.StringNumberOrTime `json:"lte,omitempty" bson:"lte,omitempty"`
	FormatParam
	timeZoneParam
	boostParam
	nameParam
}

func (r *rangeClause) GreaterThan() dynamic.StringNumberOrTime {
	return r.GreaterThanValue
}
func (r *rangeClause) GreaterThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.GreaterThanValue
}

func (r *rangeClause) LessThan() dynamic.StringNumberOrTime {
	return r.LessThanValue
}

func (r *rangeClause) LessThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.LessThanOrEqualToValue
}

func (r *rangeClause) SetGreaterThan(value interface{}) error {
	v := dynamic.NewStringNumberOrTime()
	err := v.Set(value)
	if err != nil {
		return err
	}
	r.GreaterThanValue = v
	return nil
}

func (r *rangeClause) SetGreaterThanOrEqualTo(value interface{}) error {
	v := dynamic.NewStringNumberOrTime()
	err := v.Set(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.GreaterThanOrEqualToValue = v
	return nil
}

func (r *rangeClause) SetLessThan(value interface{}) error {
	v := dynamic.NewStringNumberOrTime()
	err := v.Set(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.LessThanValue = v
	return nil
}

func (r *rangeClause) SetLessThanOrEqualTo(value interface{}) error {
	v := dynamic.NewStringNumberOrTime()
	err := v.Set(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.LessThanOrEqualToValue = v
	return nil
}

func (r rangeClause) Type() Type {
	return TypeBoolean
}

type RangeQuery struct {
	RangeField string
	rangeClause
}
