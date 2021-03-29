package search

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

type Clause interface {
	Kind() Kind
}

type CompleteClause interface {
	Clause
	_Complete()
}

type clause interface {
	Name() string
	Clause()
}

type WithField interface {
	Field() string
}

type withfield interface {
	field() string
}

type Clauser interface {
	Clause
	Clause() (CompleteClause, error)
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

func (c *Clauses) unpack() (Clauses, error) {
	if c == nil {
		*c = Clauses{}
	}
	for i, ce := range *c {
		if v, ok := ce.(Clauser); ok {
			r, err := unpackClause(v)
			if err != nil {
				return nil, err
			}
			(*c)[i] = r
		}
	}
	return *c, nil
}

type QueryClause interface {
	Clause
	Name() string
	json.Marshaler
}

type QueryClauses struct {
	clauses []QueryClause
}

func (qc QueryClauses) MarshalJSON() ([]byte, error) {
	return json.Marshal(qc.clauses)
}

func (qc *QueryClauses) UnmarshalJSON(data []byte) error {
	*qc = QueryClauses{}
	d := dynamic.JSON(data)
	var cm []map[Kind]dynamic.JSON
	if d.IsObject() {
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

func (c *QueryClauses) RemoveIndex(i int) Clause {
	v := (c.clauses)[i]
	c.clauses = append(c.clauses[:i], c.clauses[i+1:]...)
	return v
}

func (c *QueryClauses) RemoveAllWithName(name string) []Clause {
	rem := []Clause{}
	if c == nil {
		*c = Clauses{}
	}
	for i, v := range *c {
		if wn, ok := v.(WithName); ok {
			if wn.Name() == name {
				rem = append(rem, v)
				c.RemoveIndex(i)
			}
		} else if wn, ok := v.(withName); ok {
			if wn.name() == name {
				rem = append(rem, v)
				c.RemoveIndex(i)
			}
		}
	}
	return rem
}
func (c *Clauses) Add(clause CompleteClause) error {

	var err error
	clause, err = unpackClause(clause)
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
func unpackClause(clause CompleteClause) (CompleteClause, error) {
	var err error
	if v, ok := clause.(Clauser); ok {
		clause, err = v.Clause()
		if err != nil {
			return nil, err
		}
	}
	return clause, nil
}
