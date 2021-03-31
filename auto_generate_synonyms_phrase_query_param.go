package picker

import (
	"encoding/json"

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
func unmarshalAutoGenerateSynonymsPhraseQueryParam(data dynamic.JSON, target interface{}) error {
	if a, ok := target.(WithAutoGenerateSynonymsPhraseQuery); ok {
		var b dynamic.Bool
		var err error
		if data.IsBool() {
			b, err = dynamic.NewBool(data.UnquotedString())
			if err != nil {
				return err
			}
		} else if data.IsString() {
			var str String
			err := json.Unmarshal(data, &str)
			if err != nil {
				return err
			}

		}

		if v, ok := b.Bool(); ok {
			a.SetAutoGenerateSynonymsPhraseQuery(v)
		}
	}
	return nil
}

func marshalAutoGenerateSynonymsPhraseQueryParam(source interface{}) (dynamic.JSON, error) {
	if b, ok := source.(WithAutoGenerateSynonymsPhraseQuery); ok {
		if !b.AutoGenerateSynonymsPhraseQuery() {
			return json.Marshal(b.AutoGenerateSynonymsPhraseQuery())
		}
	}
	return nil, nil
}
