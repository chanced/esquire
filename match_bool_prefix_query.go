package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type MatchBoolPrefixer interface {
	MatchBoolPrefix() (*MatchBoolPrefixQuery, error)
}

type MatchBoolPrefixQueryParams struct {
	Query               string
	Field               string
	MinimumShouldMatch  string
	Operator            Operator
	Analyzer            string
	Fuzziness           string
	PrefixLength        interface{}
	FuzzyTranspositions interface{}
	FuzzyRewrite        Rewrite
	MaxExpansions       interface{}
	Name                string
	completeClause
}

func (MatchBoolPrefixQueryParams) Kind() QueryKind {
	return QueryKindMatchBoolPrefix
}
func (p MatchBoolPrefixQueryParams) Clause() (QueryClause, error) {
	return p.MatchBoolPrefix()
}

func (p MatchBoolPrefixQueryParams) MatchBoolPrefix() (*MatchBoolPrefixQuery, error) {
	q := &MatchBoolPrefixQuery{}
	err := q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	q.SetAnalyzer(p.Analyzer)
	q.SetFuzziness(p.Fuzziness)
	err = q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyRewrite(p.FuzzyRewrite)
	if err != nil {
		return q, err
	}
	err = q.SetFuzzyTranspositions(p.FuzzyTranspositions)
	if err != nil {
		return q, err
	}
	err = q.SetMaxExpansions(p.MaxExpansions)
	if err != nil {
		return q, err
	}
	q.SetMinimumShouldMatch(p.MinimumShouldMatch)
	q.SetName(p.Name)
	err = q.SetOperator(p.Operator)
	if err != nil {
		return q, err
	}
	err = q.SetPrefixLength(p.PrefixLength)
	if err != nil {
		return q, err
	}
	return q, nil
}

type MatchBoolPrefixQuery struct {
	query string
	minimumShouldMatchParam
	operatorParam
	analyzerParam
	fuzzinessParam
	prefixLengthParam
	fuzzyTranspositionsParam
	maxExpansionsParam
	fieldParam
	completeClause
	nameParam
}

func (MatchBoolPrefixQuery) Kind() QueryKind {
	return QueryKindMatchBoolPrefix
}
func (m *MatchBoolPrefixQuery) UnmarshalJSON(data []byte) error {
	*m = MatchBoolPrefixQuery{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for fld, md := range obj {
		m.field = fld
		var mq matchBoolPrefixQuery
		if md.IsString() {
			var str string
			err := json.Unmarshal(md, &str)
			if err != nil {
				return err
			}
			mq = matchBoolPrefixQuery{Query: str}
		} else {
			mq := matchBoolPrefixQuery{}
			err := mq.UnmarshalJSON(md)
			if err != nil {
				return err
			}
		}
		m.query = mq.Query
		m.analyzer = mq.Analyzer
		m.fuzziness = mq.Fuzziness
		m.fuzzyRewrite = mq.FuzzyRewrite
		err := m.SetQuery(mq.Query)
		if err != nil {
			return err
		}
		err = m.fuzzyTranspositions.Set(mq.FuzzyTranspositions)
		if err != nil {
			return err
		}
		err = m.maxExpansions.Set(mq.MaxExpansions)
		if err != nil {
			return err
		}
		m.minimumShouldMatch = mq.MinimumShouldMatch
		m.name = mq.Name
		err = m.SetOperator(mq.Operator)
		if err != nil {
			return err
		}
		_ = m.prefixLength.Set(mq.PrefixLength)
		return nil
	}
	return nil
}
func (m MatchBoolPrefixQuery) MarshalJSON() ([]byte, error) {
	qd, err := matchBoolPrefixQuery{
		Query:               m.query,
		MinimumShouldMatch:  m.minimumShouldMatch,
		Operator:            m.operator,
		Analyzer:            m.analyzer,
		Fuzziness:           m.fuzziness,
		PrefixLength:        m.prefixLength.Value(),
		FuzzyTranspositions: m.fuzzyTranspositions.Value(),
		Name:                m.name,
		FuzzyRewrite:        m.fuzzyRewrite,
	}.MarshalJSON()
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		m.field: qd,
	}
	return obj.MarshalJSON()
}
func (m *MatchBoolPrefixQuery) SetQuery(query string) error {
	if len(query) == 0 {
		return ErrQueryRequired
	}
	m.query = query
	return nil
}
func (m MatchBoolPrefixQuery) Query() string {
	return m.query
}
func (m *MatchBoolPrefixQuery) Clear() {
	if m == nil {
		return
	}
	*m = MatchBoolPrefixQuery{}
}
func (m *MatchBoolPrefixQuery) IsEmpty() bool {
	return m == nil || len(m.query) == 0 || len(m.field) == 0
}
func (m *MatchBoolPrefixQuery) Clause() (QueryClause, error) {
	return m, nil
}
func (m *MatchBoolPrefixQuery) MatchBoolPrefix() (*MatchBoolPrefixQuery, error) {
	return m, nil

}

//easyjson:json
type matchBoolPrefixQuery struct {
	Query               string      `json:"query"`
	MinimumShouldMatch  string      `json:"minimum_should_match,omitempty"`
	Operator            Operator    `json:"operator,omitempty"`
	Analyzer            string      `json:"analyzer,omitempty"`
	Fuzziness           string      `json:"fuzziness,omitempty"`
	PrefixLength        interface{} `json:"prefix_length,omitempty"`
	FuzzyTranspositions interface{} `json:"fuzzy_transpositions,omitempty"`
	Name                string      `json:"_name,omitempty"`
	MaxExpansions       interface{} `json:"max_expansions,omitempty"`
	FuzzyRewrite        Rewrite     `json:"fuzzy_rewrite,omitempty"`
}
