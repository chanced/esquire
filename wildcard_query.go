package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type Wildcarder interface {
	Wildcard() (*WildcardQuery, error)
}
type WildcardQueryParams struct {
	Name            string
	Boost           interface{}
	Value           string
	CaseInsensitive bool
	Field           string
	Rewrite         Rewrite
	completeClause
}

func (WildcardQueryParams) Kind() QueryKind {
	return QueryKindWildcard
}

func (p WildcardQueryParams) Clause() (QueryClause, error) {
	return p.Wildcard()
}
func (p WildcardQueryParams) Wildcard() (*WildcardQuery, error) {
	q := &WildcardQuery{}
	q.SetCaseInsensitive(p.CaseInsensitive)
	err := q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindWildcard)
	}
	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindWildcard, q.field)
	}
	err = q.SetValue(p.Value)
	if err != nil {
		return q, newQueryError(err, QueryKindWildcard, q.field)
	}
	q.SetName(p.Name)
	err = q.SetRewrite(p.Rewrite)
	if err != nil {
		return q, newQueryError(err, QueryKindWildcard, q.field)
	}
	return q, nil
}

type WildcardQuery struct {
	value string
	nameParam
	boostParam
	rewriteParam
	caseInsensitiveParam
	fieldParam
	completeClause
}

func (WildcardQuery) Kind() QueryKind {
	return QueryKindWildcard
}
func (q *WildcardQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *WildcardQuery) Wildcard() (*WildcardQuery, error) {
	return q, nil
}
func (q *WildcardQuery) Clear() {
	if q == nil {
		return
	}
	*q = WildcardQuery{}
}
func (q *WildcardQuery) SetValue(value string) error {
	if len(value) == 0 {
		return ErrValueRequired
	}
	q.value = value
	return nil
}
func (q WildcardQuery) Value() string {
	return q.value
}
func (q *WildcardQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}

func (q *WildcardQuery) UnmarshalJSON(data []byte) error {
	*q = WildcardQuery{}
	rd := dynamic.JSONObject{}
	err := rd.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for field, d := range rd {
		q.field = field
		obj, err := unmarshalClauseParams(d, q)
		if err != nil {
			return err
		}
		var value string
		err = json.Unmarshal(obj["value"], &value)
		if err != nil {
			return err
		}
		q.value = value
		return nil
	}
	return nil
}
func (q WildcardQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}

func (q WildcardQuery) MarshalJSON() ([]byte, error) {
	data, err := marshalClauseParams(&q)
	if err != nil {
		return nil, err
	}
	data["value"] = q.value
	qd, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.JSONObject{q.field: qd})
}
func (q *WildcardQuery) IsEmpty() bool {
	return q == nil || len(q.value) == 0 || len(q.field) == 0
}
