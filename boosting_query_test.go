package picker_test

import (
	"encoding/json"
	"testing"

	"github.com/chanced/cmpjson"
	"github.com/chanced/picker"
	"github.com/stretchr/testify/require"
)

func TestBoostingQuery(t *testing.T) {

	assert := require.New(t)
	_ = assert
	data := []byte(`{
		"query": {
		  "boosting": {
			"positive": {
			  "term": {
				"text":{
					"value": "apple"
				}
			  }
			},
			"negative": {
			  "term": {
				"text": {
					"value": "pie tart fruit crumble tree"
				}
			  }
			},
			"negative_boost": 0.5
		  }
		}
	  }`)

	s, err := picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Boosting: picker.BoostingQueryParams{
				Negative: &picker.QueryParams{
					Term: picker.TermQueryParams{Field: "text", Value: "pie tart fruit crumble tree"},
				},
				Positive: &picker.QueryParams{
					Term: picker.TermQueryParams{Field: "text", Value: "apple"},
				},
				NegativeBoost: 0.5,
			},
		},
	})
	assert.NoError(err)
	jsonRes, err := json.Marshal(s)
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, jsonRes), cmpjson.Diff(data, jsonRes))
	// diff, str := jsondiff.Compare(data, res, &jsondiff.Options{})
	// assert.Equal(jsondiff.FullMatch, diff, str)
	var res picker.Search
	err = json.Unmarshal(jsonRes, &res)
	assert.NoError(err)

	jsonRes, err = json.Marshal(res)
	assert.NoError(err)
	assert.True(cmpjson.Equal(data, jsonRes), cmpjson.Diff(data, jsonRes))

	_, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Boosting: picker.BoostingQueryParams{
				Positive: &picker.QueryParams{
					Term: picker.TermQueryParams{Field: "f"},
				},
			},
		},
	})

	assert.ErrorIs(err, picker.ErrNegativeRequired)
	_, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Boosting: picker.BoostingQueryParams{
				Negative: &picker.QueryParams{
					Term: picker.TermQueryParams{Field: "f", Value: "val"},
				},
			},
		},
	})
	assert.ErrorIs(err, picker.ErrPositiveRequired)
	_, err = picker.NewSearch(picker.SearchParams{
		Query: picker.QueryParams{
			Boosting: picker.BoostingQueryParams{
				Negative: &picker.QueryParams{
					Term: picker.TermQueryParams{Field: "f", Value: "negval"},
				},
				Positive: &picker.QueryParams{
					Term: picker.TermQueryParams{Field: "f", Value: "posval"},
				},
			},
		},
	})
	assert.ErrorIs(err, picker.ErrInvalidNegativeBoost)

	// assert.NoError(err)
	// jsondiff.Compare(test1Data, test1Res)
}
