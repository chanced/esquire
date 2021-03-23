package search

import "github.com/chanced/dynamic"

type Range struct {
	GreaterThan          interface{}
	GreaterThanOrEqualTo interface{}
	LessThan             interface{}
	LessThanOrEqualTo    interface{}
	Format               string
	TimeZone             string
	Boost                float32
}

func (r Range) QueryValue() (RangeQueryValue, error) {
	q := RangeQueryValue{}
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
	q.SetBoost(r.Boost)
	q.SetTimeZone(r.TimeZone)
	return q, nil
}

func (r Range) Type() Type {
	return TypeBoolean
}

type RangeQueryValue struct {
	GreaterThanValue          *dynamic.StringNumberOrTime `json:"gt,omitempty" bson:"gt,omitempty"`
	GreaterThanOrEqualToValue *dynamic.StringNumberOrTime `json:"gte,omitempty" bson:"gt,omitempty"`
	LessThanValue             *dynamic.StringNumberOrTime `json:"lt,omitempty" bson:"lt,omitempty"`
	LessThanOrEqualToValue    *dynamic.StringNumberOrTime `json:"lte,omitempty" bson:"lte,omitempty"`
	FormatParam               `json:",inline" bson:",inline"`
	TimeZoneParam             `json:",inline" bson:",inline"`
	BoostParam                `json:",inline" bson:",inline"`
}

func (r *RangeQueryValue) GreaterThan() *dynamic.StringNumberOrTime {
	return r.GreaterThanValue
}
func (r *RangeQueryValue) GreaterThanOrEqualTo() *dynamic.StringNumberOrTime {
	return r.GreaterThanValue
}

func (r *RangeQueryValue) LessThan() *dynamic.StringNumberOrTime {
	return r.LessThanValue
}

func (r *RangeQueryValue) LessThanOrEqualTo() *dynamic.StringNumberOrTime {
	return r.LessThanOrEqualToValue
}

func (r *RangeQueryValue) SetGreaterThan(value interface{}) error {
	r.GreaterThanValue = nil
	v, err := dynamic.NewStringNumberOrTime(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.GreaterThanValue = v
	return nil
}

func (r *RangeQueryValue) SetGreaterThanOrEqualTo(value interface{}) error {
	r.GreaterThanOrEqualToValue = nil
	v, err := dynamic.NewStringNumberOrTime(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.GreaterThanOrEqualToValue = v
	return nil
}

func (r *RangeQueryValue) SetLessThan(value interface{}) error {
	r.LessThanValue = nil
	v, err := dynamic.NewStringNumberOrTime(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.LessThanValue = v
	return nil
}

func (r *RangeQueryValue) SetLessThanOrEqualTo(value interface{}) error {
	r.LessThanOrEqualToValue = nil
	v, err := dynamic.NewStringNumberOrTime(value)
	if err != nil {
		return err
	}
	if v.IsNilOrZero() {
		return nil
	}
	r.LessThanOrEqualToValue = v
	return nil
}

func (r RangeQueryValue) Type() Type {
	return TypeBoolean
}

type RangeQuery struct {
	RangeValue map[string]RangeQueryValue `json:"range,omitempty" bson:"range,omitempty"`
}

func (r *RangeQuery) AddRange(field string, value Range) error {
	if r.RangeValue == nil {
		r.RangeValue = map[string]RangeQueryValue{}
	}
	if _, exists := r.RangeValue[field]; exists {
		return QueryError{
			Field: field,
			Err:   ErrFieldExists,
			Type:  TypeRange,
		}
	}
	return r.SetRange(field, value)
}

func (r *RangeQuery) SetRange(field string, value Range) error {
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeRange)
	}
	if r.RangeValue == nil {
		r.RangeValue = map[string]RangeQueryValue{}
	}
	qv, err := value.QueryValue()
	if err != nil {
		return QueryError{
			Field: field,
			Err:   err,
			Type:  TypeRange,
		}
	}
	r.RangeValue[field] = qv
	return nil
}

func (r *RangeQuery) RemoveRange(field string) {
	delete(r.RangeValue, field)
}
