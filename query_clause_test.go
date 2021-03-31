package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestClauses(t *testing.T) {
	assert := require.New(t)
	_ = assert
	var clauses picker.QueryClauses

	json1 := []byte(`{
        "term" : { "user.id" : "chanced" }
      }`)

	err := json.Unmarshal(json1, &clauses)
	assert.NoError(err)
	assert.False(clauses.IsEmpty())
	assert.Equal(picker.KindTerm, clauses.Clauses()[0].Kind())

	assert.Equal("chanced", clauses.Clauses()[0].(*picker.TermClause).Value())
	assert.NoError(err)
	res1, err := json.Marshal(clauses)
	assert.NoError(err)
	fmt.Println(string(res1))

	clauses = picker.QueryClauses{}
	err = json.Unmarshal(res1, &clauses)
	assert.NotEmpty(clauses.Clauses())
	assert.Equal(picker.KindTerm, clauses.Clauses()[0].Kind())
	assert.Equal("chanced", clauses.Clauses()[0].(*picker.TermClause).Value())
	assert.NoError(err)

	json2 := []byte(`[{
	    "term" : { "user.id" : "chanced" }
	  }]`)

	err = json.Unmarshal(json2, &clauses)
	assert.Equal(clauses.Len(), 1)
	assert.Equal(picker.KindTerm, clauses.Clauses()[0].Kind())
	assert.Equal("chanced", clauses.Clauses()[0].(*picker.TermClause).Value())
	assert.NoError(err)
	d, err := json.Marshal(clauses)
	assert.NoError(err)
	fmt.Println(string(d))
}
