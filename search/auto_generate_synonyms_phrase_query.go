package search

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
		return true
	}
	return *agspq.AutoGenerateSynonymsPhraseQueryValue
}

// SetAutoGenerateSynonymsPhraseQuery sets AutoGenerateSynonymsPhraseQueryValue to v
func (agspq *AutoGenerateSynonymsPhraseQueryParam) SetAutoGenerateSynonymsPhraseQuery(v bool) {
	if agspq.AutoGenerateSynonymsPhraseQuery() != v {
		agspq.AutoGenerateSynonymsPhraseQueryValue = &v
	}
}
