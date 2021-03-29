package search

const (
	// ScoreExp is an exponential decay function
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay
	FuncKindExp FunctionKind = "exp"
	// ScoreGause is normal decay
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay
	FuncKindGauss FunctionKind = "gauss"
	// FunctionKindLinearDecay decay
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay
	FuncKindLinear FunctionKind = "linear"
	// FuncKindWeight allows you to multiply the score by the provided weight. This
	// can sometimes be desired since boost value set on specific queries gets
	// normalized, while for this score function it does not. The number value
	// is of type float.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-weight
	FuncKindWeight FunctionKind = "weight"
	// FuncKindScriptScore allows you to wrap another query and customize the
	// scoring of it optionally with a computation derived from other numeric
	// field values in the doc using a script expression.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-script-score
	FuncKindScriptScore FunctionKind = "script_score"
	// FuncKindRandomScore generates scores that are uniformly distributed from 0 up to
	// but not including 1. By default, it uses the internal Lucene doc ids as a
	// source of randomness, which is very efficient but unfortunately not
	// reproducible since documents might be renumbered by merges.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-random
	FuncKindRandomScore FunctionKind = "random_score"
	// FuncKindFieldValueFactor function allows you to use a field from a document
	// to influence the score. It’s similar to using the script_score function,
	// however, it avoids the overhead of scripting. If used on a multi-valued
	// field, only the first value of the field is used in calculations
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-field-value-factor
	FuncKindFieldValueFactor FunctionKind = "field_value_factor"
)

var scoreFunctionHandlers = map[FunctionKind]func() Function{
	FuncKindExp:              func() Function { return nil },
	FuncKindGauss:            func() Function { return nil },
	FuncKindLinear:           func() Function { return nil },
	FuncKindWeight:           func() Function { return &WeightFunction{} },
	FuncKindScriptScore:      func() Function { return nil },
	FuncKindRandomScore:      func() Function { return nil },
	FuncKindFieldValueFactor: func() Function { return nil },
}
