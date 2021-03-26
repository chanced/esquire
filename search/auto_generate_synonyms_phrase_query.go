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
	autoGenerateSynonymsPhraseQuery dynamic.Bool
}

// AutoGenerateSynonymsPhraseQuery determines if match phrase queries are
// automatically created for multi-term synonyms. Defaults to true.
func (agspq autoGenerateSynonymsPhraseQueryParam) AutoGenerateSynonymsPhraseQuery() bool {
	if v, ok := agspq.autoGenerateSynonymsPhraseQuery.Bool(); ok {
		return v
	}
	return DefaultAutoGenerateSynonymsPhraseQuery
}

// SetAutoGenerateSynonymsPhraseQuery sets AutoGenerateSynonymsPhraseQueryValue to v
func (agspq *autoGenerateSynonymsPhraseQueryParam) SetAutoGenerateSynonymsPhraseQuery(v bool) {
	agspq.autoGenerateSynonymsPhraseQuery.Set(v)
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

func marshalAutoGenerateSynonymsPhraseQueryParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithAutoGenerateSynonymsPhraseQuery); ok {
		if !b.AutoGenerateSynonymsPhraseQuery() {
			data[paramAutoGenerateSynonymsPhraseQuery] = b.AutoGenerateSynonymsPhraseQuery()
		}
	}
	return data, nil
}
