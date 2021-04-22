package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
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
		Query: &picker.QueryParams{
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
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd), cmpjson.Diff(data, sd))
	var sr *picker.Search
	err = json.Unmarshal(data, &sr)
	assert.NoError(err)
	sd2, err := sr.MarshalJSON()
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, sd2), cmpjson.Diff(data, sd2))

}
