package search

import "github.com/tidwall/gjson"

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

type AutoGenerateSynonymsPhraseQueryParam struct {
	AutoGenerateSynonymsPhraseQueryValue *bool `bson:"auto_generate_synonyms_phrase_query,omitempty" json:"auto_generate_synonyms_phrase_query,omitempty"`
}

// AutoGenerateSynonymsPhraseQuery determines if match phrase queries are
// automatically created for multi-term synonyms. Defaults to true.
func (agspq AutoGenerateSynonymsPhraseQueryParam) AutoGenerateSynonymsPhraseQuery() bool {
	if agspq.AutoGenerateSynonymsPhraseQueryValue == nil {
		return DefaultAutoGenerateSynonymsPhraseQuery
	}
	return *agspq.AutoGenerateSynonymsPhraseQueryValue
}

// SetAutoGenerateSynonymsPhraseQuery sets AutoGenerateSynonymsPhraseQueryValue to v
func (agspq *AutoGenerateSynonymsPhraseQueryParam) SetAutoGenerateSynonymsPhraseQuery(v bool) {
	if agspq.AutoGenerateSynonymsPhraseQuery() != v {
		agspq.AutoGenerateSynonymsPhraseQueryValue = &v
	}
}
func unmarshalAutoGenerateSynonymsPhraseQueryParam(value gjson.Result, target interface{}) error {
	if a, ok := target.(WithAutoGenerateSynonymsPhraseQuery); ok {
		a.SetAutoGenerateSynonymsPhraseQuery(value.Bool())
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
