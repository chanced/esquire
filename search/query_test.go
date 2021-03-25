package search

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshal(t *testing.T) {
	assert := require.New(t)
	f1, err := os.Open("./testdata/query_terms_1.json")
	assert.NoError(err)
	terms1, err := io.ReadAll(f1)
	assert.NoError(err)
	_ = terms1
	// f1.Close()
	// var q Query
	// err = json.Unmarshal(terms1, &q)
	// assert.NoError(err)

	// assert.Equal("user.id", q.TermsField)
	// fmt.Println(q)
}
