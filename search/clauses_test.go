package search_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/picker/search"
	"github.com/stretchr/testify/require"
)

func TestClauses(t *testing.T) {
	assert := require.New(t)
	_ = assert
	var clauses search.Clauses

	json1 := []byte(`{
        "term" : { "user.id" : "chanced" }
      }`)

	err := json.Unmarshal(json1, &clauses)
	assert.NotEmpty(clauses)
	assert.Equal(search.TypeTerm, clauses[0].Type())
	assert.Equal("chanced", clauses[0].(*search.TermQuery).TermValue)
	assert.NoError(err)
	json2 := []byte(`[{
        "term" : { "user.id" : "chanced" }
      }]`)

	err = json.Unmarshal(json2, &clauses)
	assert.Len(clauses, 1)
	assert.Equal(search.TypeTerm, clauses[0].Type())
	assert.Equal("chanced", clauses[0].(*search.TermQuery).TermValue)
	assert.NoError(err)
	d, err := json.Marshal(clauses)
	assert.NoError(err)
	fmt.Println(string(d))
}
