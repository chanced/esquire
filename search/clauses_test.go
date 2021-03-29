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
	assert.Equal(search.KindTerm, clauses[0].Kind())
	assert.Equal("chanced", clauses[0].(*search.TermQuery).Value())
	assert.NoError(err)
	res1, err := json.Marshal(clauses)
	assert.NoError(err)
	fmt.Println(string(res1))

	clauses = search.Clauses{}
	err = json.Unmarshal(res1, &clauses)
	assert.NotEmpty(clauses)
	assert.Equal(search.KindTerm, clauses[0].Kind())
	assert.Equal("chanced", clauses[0].(*search.TermQuery).Value())
	assert.NoError(err)

	json2 := []byte(`[{
        "term" : { "user.id" : "chanced" }
      }]`)

	err = json.Unmarshal(json2, &clauses)
	assert.Len(clauses, 1)
	assert.Equal(search.KindTerm, clauses[0].Kind())
	assert.Equal("chanced", clauses[0].(*search.TermQuery).Value())
	assert.NoError(err)
	d, err := json.Marshal(clauses)
	assert.NoError(err)
	fmt.Println(string(d))
}
