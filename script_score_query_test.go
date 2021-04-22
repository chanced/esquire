package picker_test

import (
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

type MyStruct struct {
	Color string `json:"color"`
	Size  int    `json:"size"`
}

func TestScriptScoreQuery(t *testing.T) {
	assert := require.New(t)
	s, err := picker.NewSearch(picker.SearchParams{
		Query: &picker.QueryParams{
			ScriptScore: &picker.ScriptScoreQueryParams{
				Query: &picker.QueryParams{
					Term: &picker.TermQueryParams{
						Field:           "ss_term",
						Value:           "val",
						Boost:           0.34,
						CaseInsensitive: true,
						Name:            "ss_term_query",
					},
				},
				MinScore: 3,
				Boost:    3,
				Name:     "script",
				Script: &picker.Script{
					Lang:   "painless",
					Source: "doc['num1'].value > 1",
					Params: MyStruct{Color: "red", Size: 34},
				},
			},
		},
	})
	assert.NoError(err)
	_ = s

}
