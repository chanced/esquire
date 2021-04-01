package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

// RandomScoreFunc generates scores that are uniformly distributed from 0 up to but
// not including 1. By default, it uses the internal Lucene doc ids as a source
// of randomness, which is very efficient but unfortunately not reproducible
// since documents might be renumbered by merges.
//
// In case you want scores to be reproducible, it is possible to provide a seed
// and field. The final score will then be computed based on this seed, the
// minimum value of field for the considered document and a salt that is
// computed based on the index name and shard id so that documents that have the
// same value but are stored in different indexes get different scores. Note
// that documents that are within the same shard and have the same value for
// field will however get the same score, so it is usually desirable to use a
// field that has unique values for all documents. A good default choice might
// be to use the _seq_no field, whose only drawback is that scores will change
// if the document is updated since update operations also update the value of
// the _seq_no field.
type RandomScoreFunc struct {
	Field string
	// float (Optional)
	Weight interface{}
	// In case you want scores to be reproducible, it is possible to provide a
	// seed and field. The final score will then be computed based on this seed,
	// the minimum value of field for the considered document and a salt that is
	// computed based on the index name and shard id so that documents that have
	// the same value but are stored in different indexes get different scores.
	// Note that documents that are within the same shard and have the same
	// value for field will however get the same score, so it is usually
	// desirable to use a field that has unique values for all documents. A good
	// default choice might be to use the _seq_no field, whose only drawback is
	// that scores will change if the document is updated since update
	// operations also update the value of the _seq_no field.
	//
	// It was possible to set a seed without setting a field, but this has been
	// deprecated as this requires loading fielddata on the _id field which
	// consumes a lot of memory.
	Seed interface{}

	Filter CompleteClauser
}

func (RandomScoreFunc) FuncKind() FuncKind {
	return FuncKindRandomScore
}

func (rs RandomScoreFunc) Function() (Function, error) {
	return rs.RandomScoreFunction()
}
func (rs RandomScoreFunc) RandomScoreFunction() (*RandomScoreFunction, error) {
	f := &RandomScoreFunction{field: rs.Field}
	err := f.SetWeight(rs.Weight)
	if err != nil {
		return f, err
	}
	err = f.SetSeed(rs.Seed)
	if err != nil {
		return f, err
	}
	err = f.SetFilter(rs.Filter)
	if err != nil {
		return f, err
	}
	return f, nil
}

type RandomScoreFunction struct {
	weightParam
	seed   dynamic.Number
	field  string
	filter QueryClause
}

func (rs *RandomScoreFunction) SetSeed(seed interface{}) error {
	if rs == nil {
		*rs = RandomScoreFunction{}
	}
	return rs.seed.Set(seed)
}
func (rs *RandomScoreFunction) SetField(field string) {
	if rs == nil {
		*rs = RandomScoreFunction{}
	}
	rs.field = field

}

func (rs *RandomScoreFunction) SetFilter(filter CompleteClauser) error {
	if rs == nil {
		*rs = RandomScoreFunction{}
	}
	c, err := filter.Clause()
	if err != nil {
		return err
	}
	rs.filter = c
	return nil
}

func (rs *RandomScoreFunction) Filter() QueryClause {
	if rs == nil {
		return nil
	}
	return rs.filter
}
func (rs *RandomScoreFunction) FuncKind() FuncKind {
	return FuncKindRandomScore
}

type randomScoreParams struct {
	Field string   `json:"field,omitempty"`
	Seed  *float64 `json:"seed,omitempty"`
}

func (rs RandomScoreFunction) MarshalJSON() ([]byte, error) {
	return marshalFunction(&rs)
}

func (rs *RandomScoreFunction) marshalParams(data dynamic.JSONObject) error {
	params := randomScoreParams{Field: rs.field}
	if f, ok := rs.seed.Float(); ok {
		params.Seed = &f
	}
	pd, err := json.Marshal(params)
	if err != nil {
		return err
	}
	data["random_score"] = pd
	return nil
}
func (rs *RandomScoreFunction) unmarshalParams(data dynamic.JSON) error {
	params := randomScoreParams{}
	err := json.Unmarshal(data, &params)
	if err != nil {
		return err
	}
	return rs.seed.Set(params.Seed)
}
