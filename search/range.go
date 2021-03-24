package search

import "github.com/chanced/dynamic"

// Range returns documents that contain terms within a provided range.
type Range struct {
	Field                string
	GreaterThan          interface{}
	GreaterThanOrEqualTo interface{}
	LessThan             interface{}
	LessThanOrEqualTo    interface{}
	Format               string
	TimeZone             string
	Boost                float64
}

func (r Range) Rule() (Rule, error) {
	return r.Range()
}
func (r Range) Range() (*RangeRule, error) {
	q := &RangeRule{}
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

type RangeRule struct {
	GreaterThanValue          *dynamic.StringNumberOrTime `json:"gt,omitempty" bson:"gt,omitempty"`
	GreaterThanOrEqualToValue *dynamic.StringNumberOrTime `json:"gte,omitempty" bson:"gt,omitempty"`
	LessThanValue             *dynamic.StringNumberOrTime `json:"lt,omitempty" bson:"lt,omitempty"`
	LessThanOrEqualToValue    *dynamic.StringNumberOrTime `json:"lte,omitempty" bson:"lte,omitempty"`
	FormatParam               `json:",inline" bson:",inline"`
	TimeZoneParam             `json:",inline" bson:",inline"`
	BoostParam                `json:",inline" bson:",inline"`
}

func (r *RangeRule) GreaterThan() *dynamic.StringNumberOrTime {
	return r.GreaterThanValue
}
func (r *RangeRule) GreaterThanOrEqualTo() *dynamic.StringNumberOrTime {
	return r.GreaterThanValue
}

func (r *RangeRule) LessThan() *dynamic.StringNumberOrTime {
	return r.LessThanValue
}

func (r *RangeRule) LessThanOrEqualTo() *dynamic.StringNumberOrTime {
	return r.LessThanOrEqualToValue
}

func (r *RangeRule) SetGreaterThan(value interface{}) error {
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

func (r *RangeRule) SetGreaterThanOrEqualTo(value interface{}) error {
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

func (r *RangeRule) SetLessThan(value interface{}) error {
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

func (r *RangeRule) SetLessThanOrEqualTo(value interface{}) error {
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

func (r RangeRule) Type() Type {
	return TypeBoolean
}

type RangeQuery struct {
	RangeValue map[string]*RangeRule `json:"range,omitempty" bson:"range,omitempty"`
}

func (r *RangeQuery) AddRange(field string, value Range) error {
	if r.RangeValue == nil {
		r.RangeValue = map[string]*RangeRule{}
	}
	if _, exists := r.RangeValue[field]; exists {
		return QueryError{
			Field: field,
			Err:   ErrFieldExists,
			Type:  TypeRange,
		}
	}
	return r.AssignRange(field, value)
}

func (r *RangeQuery) SetRange(ranges map[string]Range) error {
	r.RangeValue = map[string]*RangeRule{}
	for k, v := range ranges {
		err := r.AssignRange(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RangeQuery) AssignRange(field string, value Range) error {
	if field == "" {
		return NewQueryError(ErrFieldRequired, TypeRange)
	}
	if r.RangeValue == nil {
		r.RangeValue = map[string]*RangeRule{}
	}
	qv, err := value.Range()
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
