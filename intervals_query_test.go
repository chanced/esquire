package picker_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestIntervalsQuery(t *testing.T) {
	assert := require.New(t)
	data := []byte(`{
		"query": {
		  "intervals" : {
			"my_text" : {
			  "all_of" : {
				"ordered" : true,
				"intervals" : [
				  {
					"match" : {
					  "query" : "my favorite food",
					  "max_gaps" : 0,
					  "ordered" : true
					}
				  },
				  {
					"any_of" : {
					  "intervals" : [
						{ "match" : { "query" : "hot water" } },
						{ "match" : { "query" : "cold porridge" } }
					  ]
					}
				  }
				]
			  }
			}
		  }
		}
	  }`)

	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Intervals: picker.IntervalsQueryParams{
				Field: "my_text",
				Rule: picker.AllOfRuleParams{
					Ordered: true,
					Intervals: picker.Ruleset{
						picker.MatchRuleParams{
							Query:   "my favorite food",
							MaxGaps: 0,
							Ordered: true,
						},
						picker.AnyOfRuleParams{
							Intervals: picker.Ruleset{
								picker.MatchRuleParams{Query: "hot water"},
								picker.MatchRuleParams{Query: "cold porridge"},
							},
						},
					},
				},
			},
		},
	})
	assert.NoError(err)
	sd, err := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(sd))
	assert.NoError(err)
	assert.NoError(compareJSONObject(data, sd))
}
