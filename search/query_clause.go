package search

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/chanced/dynamic"
)

type clause struct{}

func (clause) _Complete() {}

type Clause interface {
	Kind() Kind
}

type CompleteClauser interface {
	CompleteClause
	Clause() (QueryClause, error)
}

type CompleteClause interface {
	Clause
	_Complete()
}

type WithField interface {
	Field() string
}

type withfield interface {
	field() string
}

type Clauser interface {
	Clause
	Clause() (QueryClause, error)
}

// type Rules []Clause

// func (r *Rules) Add(rule Clause) error {
//     var err error
//     if v, ok := rule.(Clauser); ok {
//         rule, err = v.Rule()
//         if err != nil {
//             return err
//         }
//     }
//     *r = append(*r, rule)
//     return nil
// }

func marshalClauseParams(source Clause) (dynamic.Map, error) {
	return marshalParams(source)
}

func unmarshalParams(data []byte, target Clause) (map[string]dynamic.JSON, error) {
	var raw map[string]dynamic.JSON
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}
	res := map[string]dynamic.JSON{}
	for key, value := range raw {
		isParam, err := unmarshalParam(key, value, target)
		if err != nil {
			return nil, err
		}
		if !isParam {
			res[key] = value
		}
	}
	return res, nil
}

type Clauses []CompleteClause

func (c *Clauses) Add(clause CompleteClause) error {
	var err error
	if err != nil {
		return err
	}
	if c == nil {
		*c = Clauses{clause}
		return nil
	}
	*c = append(*c, clause)
	return nil
}
func (c Clauses) unpack() ([]QueryClause, error) {
	if c == nil {
		return []QueryClause{}, nil
	}

	res := make([]QueryClause, len(c))

	for i, ce := range c {
		qc, err := unpackClause(ce)
		if err != nil {
			return res, err
		}
		res[i] = qc
	}
	return res, nil
}

type QueryClause interface {
	Kind() Kind
	Name() string
	Clear()
	IsEmpty() bool
	json.Marshaler
	json.Unmarshaler
}

type QueryClauses struct {
	clauses []QueryClause
}

func (qc *QueryClauses) add(clauses Clauses) error {
	qcl, err := clauses.unpack()
	if err != nil {
		return err
	}
	qc.clauses = append(qc.clauses, qcl...)
	return nil
}

func (qc *QueryClauses) Set(clauses Clauses) error {
	qc.clauses = make([]QueryClause, len(clauses))
	qcl, err := clauses.unpack()
	if err != nil {
		return err
	}
	copy(qc.clauses, qcl)
	return nil
}
func (qc QueryClauses) IsEmpty() bool {
	return qc.Len() == 0
}
func (qc *QueryClauses) AddMany(clauses Clauses) error {
	return qc.add(clauses)
}

func (qc *QueryClauses) Add(clause CompleteClause, clauses ...CompleteClause) error {
	err := qc.add(append(clauses, clause))
	return err
}

func (qc *QueryClauses) Clauses() []QueryClause {
	return qc.clauses
}

func (qc QueryClauses) Len() int {
	return len(qc.clauses)
}

// ForEach calls fn for each QueryClause in the QueryClauses set or until an
// error is returned. If the error is dynamic.Done or an error which returns
// "done" from Error(), iteration is haulted and nil is returned.
func (qc QueryClauses) ForEach(fn func(query QueryClause) error) error {
	for _, c := range qc.clauses {
		err := fn(c)
		if err != nil {
			if errors.Is(err, dynamic.Done) || err.Error() == "done" {
				return nil
			}
			return err
		}
	}
	return nil
}
func (qc *QueryClauses) RemoveAllWithName(name string) []QueryClause {
	rem := []QueryClause{}
	if qc == nil || qc.clauses == nil {
		*qc = QueryClauses{
			clauses: []QueryClause{},
		}
	}
	for i, v := range qc.clauses {
		if wn, ok := v.(WithName); ok {
			if wn.Name() == name {
				rem = append(rem, v)
				qc.removeIndex(i)
			}
		}
	}
	return rem
}
func (qc *QueryClauses) RemoveAllForField(field string) []QueryClause {
	rem := []QueryClause{}
	if qc == nil || qc.clauses == nil {
		*qc = QueryClauses{
			clauses: []QueryClause{},
		}
	}
	for i, v := range qc.clauses {
		if wf, ok := v.(WithField); ok {
			if wf.Field() == field {
				rem = append(rem, v)
				qc.removeIndex(i)
			}
		}
	}
	return rem
}
func (qc *QueryClauses) RemoveAllOfKind(kind Kind) []QueryClause {
	rem := []QueryClause{}
	if qc == nil || qc.clauses == nil {
		*qc = QueryClauses{
			clauses: []QueryClause{},
		}
	}
	for i, v := range qc.clauses {
		if v.Kind() == kind {
			rem = append(rem, v)
			qc.removeIndex(i)
		}
	}
	return rem
}

func (qc *QueryClauses) removeIndex(i int) Clause {
	v := (qc.clauses)[i]
	qc.clauses = append(qc.clauses[:i], qc.clauses[i+1:]...)
	return v
}

func (qc QueryClauses) MarshalJSON() ([]byte, error) {
	res := make([]dynamic.JSONObject, len(qc.clauses))
	for i, c := range qc.clauses {
		obj := dynamic.JSONObject{}
		cd, err := c.MarshalJSON()
		if err != nil {
			return nil, err
		}
		obj[c.Kind().String()] = cd
		res[i] = obj
	}
	return json.Marshal(res)
}

func (qc *QueryClauses) UnmarshalJSON(data []byte) error {
	*qc = QueryClauses{}
	d := dynamic.JSON(data)
	var cm []map[Kind]dynamic.JSON
	if d.IsNull() || len(data) == 0 {
		qc.clauses = make([]QueryClause, 0)
		return nil
	} else if d.IsObject() {
		var o map[Kind]dynamic.JSON
		err := json.Unmarshal(d, &o)
		if err != nil {
			return err
		}
		cm = []map[Kind]dynamic.JSON{o}
	} else {
		err := json.Unmarshal(d, &cm)
		if err != nil {
			return err
		}
	}
	for _, cd := range cm {
		for t, d := range cd {
			handler, ok := clauseHandlers[t]
			if !ok {
				return fmt.Errorf("%w <%s>", ErrUnsupportedKind, t)
			}
			ce := handler()
			err := json.Unmarshal(d, &ce)
			if err != nil {
				return err
			}
			fmt.Println(ce)
			qc.clauses = append(qc.clauses, ce)
		}
	}
	return nil
}

func unpackClause(clause CompleteClause) (QueryClause, error) {
	if v, ok := clause.(Clauser); ok {
		qc, err := v.Clause()
		if err != nil {
			return nil, err
		}
		return qc, nil
	}
	if v, ok := clause.(QueryClause); ok {
		return v, nil
	}

	return nil, errors.New("invalid query")
}

func unmarshalQueryClause(data []byte) (QueryClause, error) {
	var v map[Kind]dynamic.JSON
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	for kind, cd := range v {
		handler := clauseHandlers[kind]
		if handler == nil {
			return nil, fmt.Errorf("%w <%s>", ErrUnsupportedKind, kind)
		}
		c := handler()
		err = c.UnmarshalJSON(cd)
		return c, err
	}
	return nil, nil
}
