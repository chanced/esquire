package search

import (
	"encoding/json"
	"fmt"

	"github.com/chanced/dynamic"
)

type Clauses []Clause

func (c Clauses) MarshalJSON() ([]byte, error) {
	u, err := c.unpack()
	if err != nil {
		return nil, err
	}
	s := []Clause(u)
	return json.Marshal(s)
}

func (c *Clauses) UnmarshalJSON(data []byte) error {
	*c = Clauses{}

	d := dynamic.JSON(data)
	var cm []map[Type]dynamic.JSON
	if d.IsObject() {
		var j map[Type]dynamic.JSON
		err := json.Unmarshal(d, &j)
		if err != nil {
			return err
		}
		cm = []map[Type]dynamic.JSON{j}
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
				return fmt.Errorf("%w <%s>", ErrUnsupportedType, t)
			}
			ce := handler()
			err := json.Unmarshal(d, &ce)
			if err != nil {
				return err
			}
			fmt.Println(ce)
			*c = append(*c, ce)
		}
	}
	return nil
}

func (c *Clauses) Validate() error {
	if c == nil {
		*c = Clauses{}
	}
	_, err := c.unpack()
	return err
}

func (c *Clauses) unpack() (Clauses, error) {
	if c == nil {
		*c = Clauses{}
	}
	for i, ce := range *c {
		if v, ok := ce.(clauser); ok {
			r, err := unpackClause(v)
			if err != nil {
				return nil, err
			}
			(*c)[i] = r
		}
	}
	return *c, nil

}

func (c *Clauses) RemoveIndex(i int) Clause {
	v := (*c)[i]
	*c = append((*c)[:i], (*c)[i+1:]...)
	return v
}

func (c *Clauses) RemoveAllForField(field string) []Clause {
	rem := []Clause{}
	if c == nil {
		*c = Clauses{}
	}
	for i, v := range *c {
		if wn, ok := v.(WithField); ok {
			if wn.Field() == field {
				rem = append(rem, v)
				c.RemoveIndex(i)
			}
		} else if wn, ok := v.(withfield); ok {
			if wn.field() == field {
				rem = append(rem, v)
				c.RemoveIndex(i)
			}
		}
	}
	return rem
}
func (c *Clauses) RemoveAllWithName(name string) []Clause {
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
func (c *Clauses) Add(clause Clause) error {

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
func unpackClause(clause Clause) (Clause, error) {
	var err error
	if v, ok := clause.(clauser); ok {
		clause, err = v.Clause()
		if err != nil {
			return nil, err
		}
	}
	return clause, nil
}
func unpackClauses(clauses Clauses) (Clauses, error) {
	if clauses == nil {
		return Clauses{}, nil
	}
	return clauses.unpack()
}
