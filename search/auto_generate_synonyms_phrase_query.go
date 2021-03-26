package search

import (
	"github.com/chanced/dynamic"
)

const DefaultAutoGenerateSynonymsPhraseQuery = true

// WithAutoGenerateSynonymsPhraseQuery is an interface for the query mixin that
// adds auto_generate_synonyms_phrase_query param
type WithAutoGenerateSynonymsPhraseQuery interface {
	// AutoGenerateSynonymsPhraseQuery determines if match phrase queries are
	// automatically created for multi-term synonyms. Defaults to true.
	AutoGenerateSynonymsPhraseQuery() bool
	// SetAutoGenerateSynonymsPhraseQuery sets AutoGenerateSynonymsPhraseQueryValue to v
	SetAutoGenerateSynonymsPhraseQuery(v bool)
}

type autoGenerateSynonymsPhraseQueryParam struct {
	autoGenerateSynonymsPhraseQuery *bool `bson:"auto_generate_synonyms_phrase_query,omitempty" json:"auto_generate_synonyms_phrase_query,omitempty"`
}

// AutoGenerateSynonymsPhraseQuery determines if match phrase queries are
// automatically created for multi-term synonyms. Defaults to true.
func (agspq autoGenerateSynonymsPhraseQueryParam) AutoGenerateSynonymsPhraseQuery() bool {
	if agspq.autoGenerateSynonymsPhraseQuery == nil {
		return DefaultAutoGenerateSynonymsPhraseQuery
	}
	return *agspq.autoGenerateSynonymsPhraseQuery
}

// SetAutoGenerateSynonymsPhraseQuery sets AutoGenerateSynonymsPhraseQueryValue to v
func (agspq *autoGenerateSynonymsPhraseQueryParam) SetAutoGenerateSynonymsPhraseQuery(v bool) {
	if agspq.AutoGenerateSynonymsPhraseQuery() != v {
		agspq.autoGenerateSynonymsPhraseQuery = &v
	}
}
func unmarshalAutoGenerateSynonymsPhraseQueryParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithAutoGenerateSynonymsPhraseQuery); ok {
		b := dynamic.NewBool(data.UnquotedString())
		if v, ok := b.Bool(); ok {
			a.SetAutoGenerateSynonymsPhraseQuery(v)
		}
	}
	return nil
}
func marshalAutoGenerateSynonymsPhraseQueryParam(data M, source interface{}) (M, error) {
	if b, ok := source.(WithAutoGenerateSynonymsPhraseQuery); ok {
		if b.AutoGenerateSynonymsPhraseQuery() != DefaultAutoGenerateSynonymsPhraseQuery {
			data[paramAutoGenerateSynonymsPhraseQuery] = b.AutoGenerateSynonymsPhraseQuery()
		}
	}
	return data, nil
}
