package picker

import (
	"encoding/json"

	"github.com/chanced/dynamic"
)

type FunctionScorer interface {
	FunctionScore() (*FunctionScoreQuery, error)
}

// FunctionScoreQueryParams  allows you to modify the score of documents that are retrieved
// by a query. This can be useful if, for example, a score function is
// computationally expensive and it is sufficient to compute the score on a
// filtered set of documents.
//
// To use function_score, the user has to define a query and one or more
// functions, that compute a new score for each document returned by the query.
type FunctionScoreQueryParams struct {
	Query *QueryParams
	Boost interface{}
	// Documents with a score lower than this floating point number are excluded
	// from the search results. (Optional)
	MinScore  float64
	MaxBoost  float64
	BoostMode BoostMode
	ScoreMode ScoreMode
	Functions Funcs
	Name      string
}

func (fs *FunctionScoreQueryParams) FunctionScore() (*FunctionScoreQuery, error) {
	if fs == nil {
		return nil, nil
	}

	c := &FunctionScoreQuery{}
	c.SetName(fs.Name)
	err := c.SetQuery(fs.Query)
	if err != nil {
		return c, err
	}
	err = c.SetBoostMode(fs.BoostMode)
	if err != nil {
		return c, err
	}
	err = c.SetScoreMode(fs.ScoreMode)
	if err != nil {
		return c, err
	}
	c.SetMinScore(fs.MinScore)
	c.SetMaxBoost(fs.MaxBoost)
	err = c.SetFunctions(fs.Functions)
	if err != nil {
		return c, err
	}
	err = c.SetBoost(fs.Boost)
	if err != nil {
		return c, err
	}

	return c, nil
}
func (fs *FunctionScoreQueryParams) Clause() (QueryClause, error) {
	return fs.FunctionScore()
}

// FunctionScoreQuery allows you to modify the score of documents that are retrieved
// by a query. This can be useful if, for example, a score function is
// computationally expensive and it is sufficient to compute the score on a
// filtered set of documents.
//
// To use function_score, the user has to define a query and one or more
// functions, that compute a new score for each document returned by the query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html
type FunctionScoreQuery struct {
	query *Query
	boostParam
	nameParam
	boostModeParam
	scoreModeParam
	maxBoostParam
	minScoreParam
	functions Functions
	completeClause
}

func (fs *FunctionScoreQuery) Clause() (QueryClause, error) {
	return fs, nil
}
func (fs *FunctionScoreQuery) Query() *Query {
	return fs.query
}

func (fs *FunctionScoreQuery) SetFunctions(funcs Funcs) error {
	if fs == nil {
		*fs = FunctionScoreQuery{}
	}
	f, err := funcs.functions()
	if err != nil {
		return err
	}
	fs.functions = f
	return nil
}
func (fs *FunctionScoreQuery) Functions() Functions {
	if fs == nil {
		return nil
	}
	return fs.functions
}
func (fs *FunctionScoreQuery) SetQuery(query *QueryParams) error {
	if fs == nil {
		*fs = FunctionScoreQuery{}
	}
	q, err := query.Query()
	if err != nil {
		return err
	}
	fs.query = q
	return nil
}
func (fs *FunctionScoreQuery) Clear() {
	*fs = FunctionScoreQuery{}
}

func (fs *FunctionScoreQuery) UnmarshalJSON(data []byte) error {
	*fs = FunctionScoreQuery{}
	params, err := unmarshalClauseParams(data, fs)
	if err != nil {
		return err
	}
	fd := params["functions"]
	if len(fd) > 0 {
		err = fs.functions.UnmarshalJSON(fd)
		if err != nil {
			return err
		}
	}
	qd := params["query"]
	if len(qd) > 0 {
		var q Query
		err = json.Unmarshal(qd, &q)
		if err != nil {
			return err
		}
		fs.query = &q
	}
	return nil
}

func (fs *FunctionScoreQuery) Name() string {
	if fs == nil {
		*fs = FunctionScoreQuery{}
	}
	return fs.name
}

func (fs FunctionScoreQuery) MarshalJSON() ([]byte, error) {
	data, err := marshalClauseParams(fs)
	if err != nil {
		return nil, err
	}
	if len(fs.functions) > 0 {
		fd, err := json.Marshal([]Function(fs.functions))
		if err != nil {
			return nil, err
		}
		data["functions"] = dynamic.JSON(fd)
	}
	if !fs.query.IsEmpty() {
		qd, err := fs.query.MarshalJSON()
		if err != nil {
			return nil, err
		}

		data["query"] = dynamic.JSON(qd)
	}
	return json.Marshal(data)
}

func (fs *FunctionScoreQuery) IsEmpty() bool {
	return fs == nil || len(fs.functions) == 0
}

func (FunctionScoreQuery) Kind() QueryKind {
	return QueryKindFunctionScore
}
