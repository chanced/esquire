package picker

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/chanced/dynamic"
)

type completeClause struct{}

func (completeClause) Complete() {}

type Clause interface {
	Kind() QueryKind
}

type CompleteClauser interface {
	CompleteClause
	Clause() (QueryClause, error)
}

type CompleteClause interface {
	Clause
	Complete()
}

type WithField interface {
	Field() string
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
	Name() string
	Clear()
	IsEmpty() bool
	CompleteClauser
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
func (qc *QueryClauses) RemoveAllOfKind(kind QueryKind) []QueryClause {
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
	var cm []map[QueryKind]dynamic.JSON
	if d.IsNull() || len(data) == 0 {
		qc.clauses = make([]QueryClause, 0)
		return nil
	} else if d.IsObject() {
		var o map[QueryKind]dynamic.JSON
		err := json.Unmarshal(d, &o)
		if err != nil {
			return err
		}
		cm = []map[QueryKind]dynamic.JSON{o}
	} else {
		err := json.Unmarshal(d, &cm)
		if err != nil {
			return err
		}
	}
	for _, cd := range cm {
		for t, d := range cd {
			handler, ok := queryKindHandlers[t]
			if !ok {
				return fmt.Errorf("%w <%s>", ErrUnsupportedType, t)
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

func marshalSingleQueryClause(clause QueryClause) (dynamic.JSON, error) {
	if clause == nil {
		return nil, nil
	}
	cd, err := clause.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(dynamic.JSONObject{clause.Kind().String(): cd})
}

func unmarshalSingleQueryClause(data dynamic.JSON) (QueryClause, error) {
	var cd map[QueryKind]dynamic.JSON
	err := json.Unmarshal(data, &cd)
	if err != nil {
		return nil, err
	}
	for t, d := range cd {
		handler, ok := queryKindHandlers[t]
		if !ok {
			return nil, fmt.Errorf("%w <%s>", ErrUnsupportedType, t)
		}
		ce := handler()
		err := json.Unmarshal(d, &ce)
		if err != nil {
			return nil, err
		}
		return ce, nil
	}
	return nil, nil
}

func unmarshalQueryClause(data []byte) (QueryClause, error) {
	if len(data) == 0 {
		return nil, nil
	}
	var v map[QueryKind]dynamic.JSON
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	for kind, cd := range v {
		handler := queryKindHandlers[kind]
		if handler == nil {
			return nil, fmt.Errorf("%w <%s>", ErrUnsupportedType, kind)
		}
		c := handler()
		err = c.UnmarshalJSON(cd)
		return c, err
	}
	return nil, nil
}
