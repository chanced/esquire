package picker

import (
	"encoding/json"
	"reflect"

	"github.com/chanced/dynamic"
)

type Ranger interface {
	Range() (*RangeClause, error)
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
	completeClause
}

func (r Range) field() string {
	return r.Field
}

func (r Range) Clause() (QueryClause, error) {
	return r.Range()
}
func (r Range) Range() (*RangeClause, error) {
	q := &RangeClause{field: r.Field}
	err := q.setGreaterThan(r.GreaterThan)
	if err != nil {
		return q, NewQueryError(err, KindRange, r.Field)
	}
	err = q.setGreaterThanOrEqualTo(r.GreaterThanOrEqualTo)
	if err != nil {
		return q, NewQueryError(err, KindRange, r.Field)
	}
	err = q.setLessThan(r.LessThan)
	if err != nil {
		return q, NewQueryError(err, KindRange, r.Field)
	}
	err = q.setLessThanOrEqualTo(r.LessThanOrEqualTo)
	if err != nil {
		return q, NewQueryError(err, KindRange, r.Field)
	}
	err = q.SetRelation(r.Relation)
	if err != nil {
		return q, NewQueryError(err, KindRange, r.Field)
	}
	err = q.SetBoost(r.Boost)
	if err != nil {
		return q, NewQueryError(err, KindRange, r.Field)
	}
	q.SetFormat(r.Format)
	q.SetTimeZone(r.TimeZone)
	return q, nil
}

func (r Range) Kind() QueryKind {
	return KindBoolean
}

type RangeClause struct {
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

func (r *RangeClause) Clause() (QueryClause, error) {
	return r, nil
}
func (RangeClause) Kind() QueryKind {
	return KindRange
}

func (r *RangeClause) Set(field string, ranger Ranger) error {
	q, err := ranger.Range()
	if err != nil {
		return NewQueryError(err, KindRange, field)
	}
	*r = *q
	return nil
}

func (r RangeClause) GreaterThan() dynamic.StringNumberOrTime {
	return r.greaterThan
}

func (r *RangeClause) setGreaterThan(value interface{}) error {
	err := r.greaterThan.Set(value)
	if err != nil {
		return NewQueryError(err, KindRange, r.field)
	}
	return nil
}

func (r RangeClause) GreaterThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.greaterThan
}

func (r *RangeClause) setGreaterThanOrEqualTo(value interface{}) error {
	err := r.greaterThanOrEqualTo.Set(value)
	if err != nil {
		return NewQueryError(err, KindRange, r.field)
	}
	return nil
}
func (r RangeClause) LessThan() dynamic.StringNumberOrTime {
	return r.lessThan
}

func (r *RangeClause) setLessThan(value interface{}) error {
	err := r.lessThan.Set(value)
	if err != nil {
		return NewQueryError(err, KindRange, r.field)
	}
	return nil
}

func (r RangeClause) LessThanOrEqualTo() dynamic.StringNumberOrTime {
	return r.lessThanOrEqualTo
}

func (r *RangeClause) setLessThanOrEqualTo(value interface{}) error {
	err := r.lessThanOrEqualTo.Set(value)
	if err != nil {

		return NewQueryError(err, KindRange, r.field)
	}
	return nil
}
func (r *RangeClause) IsEmpty() bool {
	return r == nil || !(!r.greaterThan.IsNilOrEmpty() ||
		!r.greaterThanOrEqualTo.IsNilOrEmpty() ||
		!r.lessThanOrEqualTo.IsNilOrEmpty() ||
		!r.lessThan.IsNilOrEmpty())
}
func (r RangeClause) values() map[string]dynamic.StringNumberOrTime {
	return map[string]dynamic.StringNumberOrTime{
		"gt":  r.greaterThan,
		"gte": r.greaterThanOrEqualTo,
		"lt":  r.lessThanOrEqualTo,
		"lte": r.lessThanOrEqualTo,
	}
}
func (r RangeClause) MarshalJSON() ([]byte, error) {
	if r.IsEmpty() {
		return dynamic.Null, nil
	}
	data, err := r.marshalClauseJSON()

	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.Map{r.field: data})
}

func (r RangeClause) marshalClauseJSON() (dynamic.JSON, error) {
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

func (r *RangeClause) UnmarshalJSON(data []byte) error {
	*r = RangeClause{}
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

func (r *RangeClause) unmarshalClauseJSON(data dynamic.JSON) error {
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
func (r *RangeClause) Clear() {
	*r = RangeClause{}
}
