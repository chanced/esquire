package picker

import "github.com/chanced/dynamic"

type Intervalser interface {
	Intervals() (*IntervalsQuery, error)
}

type IntervalsQueryParams struct {
	Field string `json:"field"`
	Name  string `json:"_name,omitempty"`
	Rule  Ruler  `json:"rule"`
}

func (IntervalsQueryParams) Kind() QueryKind {
	return QueryKindIntervals
}
func (p IntervalsQueryParams) Clause() (QueryClause, error) {
	return p.Intervals()
}
func (p IntervalsQueryParams) Intervals() (*IntervalsQuery, error) {
	q := &IntervalsQuery{}
	q.SetName(p.Name)
	err := q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	err = q.SetRule(p.Rule)
	if err != nil {
		return q, err
	}
	return q, nil
}

type IntervalsQuery struct {
	fieldParam
	nameParam
	rule QueryRule
	completeClause
}

func (q *IntervalsQuery) Rule() QueryRule {

	return q.rule
}

func (q *IntervalsQuery) SetRule(rule Ruler) error {
	r, err := rule.Rule()
	if err != nil {
		return err
	}
	q.rule = r
	return nil
}
func (q *IntervalsQuery) Set(field string, rule Ruler) error {
	*q = IntervalsQuery{}
	if len(field) == 0 {
		return ErrFieldRequired
	}
	q.field = field
	r, err := rule.Rule()
	if err != nil {
		return err
	}
	q.rule = r
	return nil
}
func (IntervalsQuery) Kind() QueryKind {
	return QueryKindIntervals
}
func (q *IntervalsQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *IntervalsQuery) Intervals() (*IntervalsQuery, error) {
	return q, nil
}

func (q IntervalsQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q IntervalsQuery) MarshalJSON() ([]byte, error) {
	rd, err := MarshalRule(q.rule)
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		q.field: rd,
	}
	return obj.MarshalJSON()
}
func (q *IntervalsQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *IntervalsQuery) UnmarshalJSON(data []byte) error {
	*q = IntervalsQuery{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for fld, rd := range obj {
		q.field = fld
		r, err := UnmarshalRule(rd)
		if err != nil {
			return err
		}
		q.rule = r
		return nil
	}
	return nil
}
func (q *IntervalsQuery) Clear() {
	*q = IntervalsQuery{}
}
func (q *IntervalsQuery) IsEmpty() bool {
	return q == nil || q.rule == nil || len(q.field) == 0
}
