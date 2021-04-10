package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type MatchPhraser interface {
	MatchPhrase() (*MatchPhraseQuery, error)
}

type MatchPhraseQueryParams struct {
	Query          string
	Field          string
	Analyzer       string
	ZeroTermsQuery ZeroTerms
	Name           string
	completeClause
}

func (MatchPhraseQueryParams) Kind() QueryKind {
	return QueryKindMatchPhrase
}
func (p MatchPhraseQueryParams) Clause() (QueryClause, error) {
	return p.MatchPhrase()
}

func (p MatchPhraseQueryParams) MatchPhrase() (*MatchPhraseQuery, error) {
	q := &MatchPhraseQuery{}
	err := q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	q.SetAnalyzer(p.Analyzer)
	err = q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	err = q.SetZeroTermsQuery(p.ZeroTermsQuery)
	if err != nil {
		return q, err
	}
	return q, nil
}

type MatchPhraseQuery struct {
	query string
	fieldParam
	analyzerParam
	zeroTermsQueryParam
	completeClause
	nameParam
}

func (MatchPhraseQuery) Kind() QueryKind {
	return QueryKindMatchPhrase
}
func (m *MatchPhraseQuery) UnmarshalJSON(data []byte) error {
	*m = MatchPhraseQuery{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for fld, md := range obj {
		m.field = fld
		var mq matchPhraseQuery
		if md.IsString() {
			var str string
			err := json.Unmarshal(md, &str)
			if err != nil {
				return err
			}
			mq = matchPhraseQuery{Query: str}
		} else {
			mq = matchPhraseQuery{}
			err := mq.UnmarshalJSON(md)
			if err != nil {
				return err
			}
		}
		m.query = mq.Query
		m.analyzer = mq.Analyzer
		err := m.SetQuery(mq.Query)
		if err != nil {
			return err
		}
		m.zeroTermsQuery = mq.ZeroTerms
		m.name = mq.Name
		return nil
	}
	return nil
}
func (m MatchPhraseQuery) MarshalJSON() ([]byte, error) {
	qd, err := matchPhraseQuery{
		Query:     m.query,
		Analyzer:  m.analyzer,
		Name:      m.name,
		ZeroTerms: m.zeroTermsQuery,
	}.MarshalJSON()
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		m.field: qd,
	}
	return obj.MarshalJSON()
}
func (m *MatchPhraseQuery) SetQuery(query string) error {
	if len(query) == 0 {
		return ErrQueryRequired
	}
	m.query = query
	return nil
}
func (m MatchPhraseQuery) Query() string {
	return m.query
}
func (m *MatchPhraseQuery) Clear() {
	if m == nil {
		return
	}
	*m = MatchPhraseQuery{}
}
func (m *MatchPhraseQuery) IsEmpty() bool {
	return m == nil || len(m.query) == 0 || len(m.field) == 0
}
func (m *MatchPhraseQuery) Clause() (QueryClause, error) {
	return m, nil
}
func (m *MatchPhraseQuery) MatchPhrase() (*MatchPhraseQuery, error) {
	return m, nil

}

//easyjson:json
type matchPhraseQuery struct {
	Query     string    `json:"query"`
	Analyzer  string    `json:"analyzer,omitempty"`
	Name      string    `json:"_name,omitempty"`
	ZeroTerms ZeroTerms `json:"zero_terms_query,omitempty"`
}
