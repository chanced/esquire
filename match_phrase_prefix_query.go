package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type MatchPhrasePrefixer interface {
	MatchPhrasePrefix() (*MatchPhrasePrefixQuery, error)
}

type MatchPhrasePrefixQueryParams struct {
	// (Required, string) Text you wish to find in the provided <field>.
	//
	// The match_phrase_prefix query analyzes any provided text into tokens before
	// performing a search. The last term of this text is treated as a prefix,
	// matching any words that begin with that term.
	Query string
	// Field to query
	Field string
	// (Optional, integer) Maximum number of terms to which the last provided
	// term of the query value will expand. Defaults to 50.
	MaxExpansions interface{}
	// (Optional, integer) Maximum number of positions allowed between matching
	// tokens. Defaults to 0. Transposed terms have a slop of 2.
	Slop interface{}
	// (Optional, string) Indicates whether no documents are returned if the
	// analyzer removes all tokens, such as when using a stop filter. Valid
	// values are:
	ZeroTermsQuery ZeroTerms
	// (Optional, string) Analyzer used to convert text in the query value into
	// tokens. Defaults to the index-time analyzer mapped for the <field>. If no
	// analyzer is mapped, the index’s default analyzer is used.
	Analyzer string
	Name     string
	completeClause
}

func (MatchPhrasePrefixQueryParams) Kind() QueryKind {
	return QueryKindMatchPhrasePrefix
}
func (p MatchPhrasePrefixQueryParams) Clause() (QueryClause, error) {
	return p.MatchPhrasePrefix()
}

func (p MatchPhrasePrefixQueryParams) MatchPhrasePrefix() (*MatchPhrasePrefixQuery, error) {
	q := &MatchPhrasePrefixQuery{}
	err := q.SetQuery(p.Query)
	if err != nil {
		return q, err
	}
	q.SetAnalyzer(p.Analyzer)
	err = q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	err = q.SetMaxExpansions(p.MaxExpansions)
	if err != nil {
		return q, err
	}
	q.SetName(p.Name)
	err = q.SetSlop(p.Slop)
	if err != nil {
		return q, err
	}
	err = q.SetZeroTermsQuery(p.ZeroTermsQuery)
	if err != nil {
		return q, err
	}
	return q, nil
}

// MatchPhrasePrefixQuery returns documents that contain the words of a provided
// text, in the same order as provided. The last term of the provided text is
// treated as a prefix, matching any words that begin with that term.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query-phrase-prefix.html
type MatchPhrasePrefixQuery struct {
	query string
	analyzerParam
	maxExpansionsParam
	slopParam
	fieldParam
	zeroTermsQueryParam
	nameParam
	completeClause
}

func (MatchPhrasePrefixQuery) Kind() QueryKind {
	return QueryKindMatchPhrasePrefix
}
func (m *MatchPhrasePrefixQuery) UnmarshalBSON(data []byte) error {
	return m.UnmarshalJSON(data)
}

func (m *MatchPhrasePrefixQuery) UnmarshalJSON(data []byte) error {
	*m = MatchPhrasePrefixQuery{}
	obj := dynamic.JSONObject{}
	err := obj.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	for fld, md := range obj {
		m.field = fld
		var mq matchPhrasePrefixQuery
		if md.IsString() {
			var str string
			err := json.Unmarshal(md, &str)
			if err != nil {
				return err
			}
			mq = matchPhrasePrefixQuery{Query: str}
		} else {
			mq = matchPhrasePrefixQuery{}
			_ = mq
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
		err = m.maxExpansions.Set(mq.MaxExpansions)
		if err != nil {
			return err
		}
		m.name = mq.Name
		return nil
	}
	return nil
}
func (m MatchPhrasePrefixQuery) MarshalBSON() ([]byte, error) {
	return m.MarshalJSON()
}

func (m MatchPhrasePrefixQuery) MarshalJSON() ([]byte, error) {
	qd, err := matchPhrasePrefixQuery{
		Query:    m.query,
		Analyzer: m.analyzer,
		Name:     m.name,
	}.MarshalJSON()
	if err != nil {
		return nil, err
	}
	obj := dynamic.JSONObject{
		m.field: qd,
	}
	return obj.MarshalJSON()
}
func (m *MatchPhrasePrefixQuery) SetQuery(query string) error {
	if len(query) == 0 {
		return ErrQueryRequired
	}
	m.query = query
	return nil
}
func (m MatchPhrasePrefixQuery) Query() string {
	return m.query
}
func (m *MatchPhrasePrefixQuery) Clear() {
	if m == nil {
		return
	}
	*m = MatchPhrasePrefixQuery{}
}
func (m *MatchPhrasePrefixQuery) IsEmpty() bool {
	return m == nil || len(m.query) == 0 || len(m.field) == 0
}
func (m *MatchPhrasePrefixQuery) Clause() (QueryClause, error) {
	return m, nil
}
func (m *MatchPhrasePrefixQuery) MatchPhrasePrefix() (*MatchPhrasePrefixQuery, error) {
	return m, nil

}

//easyjson:json
type matchPhrasePrefixQuery struct {
	Query              string      `json:"query"`
	MinimumShouldMatch string      `json:"minimum_should_match,omitempty"`
	Analyzer           string      `json:"analyzer,omitempty"`
	Name               string      `json:"_name,omitempty"`
	MaxExpansions      interface{} `json:"max_expansions,omitempty"`
	ZeroTermsQuery     ZeroTerms   `json:"zero_terms_query,omitempty"`
	Slop               interface{} `json:"slop,omitempty"`
}
