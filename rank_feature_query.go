package picker

import (
	"fmt"
	"strings"

	"github.com/chanced/dynamic"
)

type RankFeaturer interface {
	RankFeature() (*RankFeatureQuery, error)
}

type RankFeatureQueryParams struct {
	// (Required) rank_feature or rank_features field used to boost
	// relevance scores.
	Field string
	// (Optional, float) Floating point number used to decrease or increase
	// relevance scores. Defaults to 1.0.
	//
	// Boost values are relative to the default value of 1.0. A boost value
	// between 0 and 1.0 decreases the relevance score. A value greater than 1.0
	// increases the relevance score.
	Boost interface{}
	// (Optional) Saturation function used to boost relevance
	// scores based on the value of the rank feature field. If no function is
	// provided, the rank_feature query defaults to the saturation function.
	//
	// Only one function saturation, log, sigmoid or linear can be provided.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-rank-feature-query.html#rank-feature-query-logarithm
	Saturation SaturationFunctioner

	// (Optional) Logarithmic function used to boost relevance scores based on
	// the value of the rank feature field.
	//
	// Only one function saturation, log, sigmoid or linear can be provided.
	Log LogFunctioner

	// (Optional) Sigmoid function used to boost relevance scores based on the
	// value of the rank feature field. See Sigmoid for more information.
	//
	// Only one function saturation, log, sigmoid or linear can be provided.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-rank-feature-query.html#rank-feature-query-sigmoid
	Sigmoid SigmoidFunctioner
	// (Optional) Linear function used to boost relevance scores based on the
	// value of the rank feature field. See Linear for more information.
	//
	// Only one function saturation, log, sigmoid or linear can be provided.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-rank-feature-query.html#rank-feature-query-linear
	Linear LinearFunctioner
	Name   string
	completeClause
}

func (RankFeatureQueryParams) Kind() QueryKind {
	return QueryKindRankFeature
}

func (p RankFeatureQueryParams) Clause() (QueryClause, error) {
	return p.RankFeature()
}
func (p RankFeatureQueryParams) RankFeature() (*RankFeatureQuery, error) {
	q := &RankFeatureQuery{}
	q.SetName(p.Name)
	err := q.SetField(p.Field)
	if err != nil {
		return q, newQueryError(err, QueryKindRankFeature)
	}

	err = q.SetBoost(p.Boost)
	if err != nil {
		return q, newQueryError(err, QueryKindRankFeature, q.field)
	}

	err = q.SetLinear(p.Linear)
	if err != nil {
		return q, newQueryError(err, QueryKindRankFeature, q.field)
	}
	err = q.SetLog(p.Log)
	if err != nil {
		return q, newQueryError(err, QueryKindRankFeature, q.field)
	}

	err = q.SetSaturation(p.Saturation)
	if err != nil {
		return q, newQueryError(err, QueryKindRankFeature, q.field)
	}

	err = q.SetSigmoid(p.Sigmoid)
	if err != nil {
		return q, newQueryError(err, QueryKindRankFeature, q.field)
	}

	return q, nil
}

type RankFeatureQuery struct {
	log        *LogFunction
	sigmoid    *SigmoidFunction
	saturation *SaturationFunction
	linear     *LinearFunction
	boostParam
	fieldParam
	nameParam
	completeClause
}

func (q RankFeatureQuery) Funcs() []string {
	res := []string{}
	if q.saturation != nil {
		res = append(res, "saturation")
	}
	if q.sigmoid != nil {
		res = append(res, "sigmoid")
	}
	if q.linear != nil {
		res = append(res, "linear")
	}
	if q.log != nil {
		res = append(res, "log")
	}
	return res
}

func (q RankFeatureQuery) checkFuncs(target string) error {
	funcs := q.Funcs()
	if len(funcs) > 1 {
		return fmt.Errorf("%w, has: [%s]", ErrMultiRankFeatureFunctions, strings.Join(funcs, ", "))
	}
	if len(funcs) > 0 && funcs[0] != target {
		funcs = append(funcs, target)
		return fmt.Errorf("%w, has: [%s]", ErrMultiRankFeatureFunctions, strings.Join(funcs, ", "))
	}
	return nil
}

func (q *RankFeatureQuery) Sigmoid() *SigmoidFunction {
	return q.sigmoid
}

func (q *RankFeatureQuery) SetSigmoid(sig SigmoidFunctioner) error {
	if sig == nil {
		q.sigmoid = nil
		return nil
	}
	err := q.checkFuncs("sigmoid")
	if err != nil {
		return err
	}

	f, err := sig.Sigmoid()
	if err != nil {
		return err
	}
	q.sigmoid = f
	return nil
}
func (q *RankFeatureQuery) Saturation() *SaturationFunction {
	return q.saturation
}

func (q *RankFeatureQuery) SetSaturation(sat SaturationFunctioner) error {
	if sat == nil {
		q.saturation = nil
		return nil
	}
	err := q.checkFuncs("saturation")
	if err != nil {
		return err
	}

	f, err := sat.Saturation()
	if err != nil {
		return err
	}
	q.saturation = f
	return nil
}

func (q *RankFeatureQuery) Log() *LogFunction {
	return q.log
}

func (q *RankFeatureQuery) Linear() *LinearFunction {
	return q.linear
}

func (q *RankFeatureQuery) SetLinear(sat LinearFunctioner) error {
	if sat == nil {
		q.linear = nil
		return nil
	}
	err := q.checkFuncs("linear")
	if err != nil {
		return err
	}

	f, err := sat.Linear()
	if err != nil {
		return err
	}
	q.linear = f
	return nil
}
func (q *RankFeatureQuery) SetLog(log LogFunctioner) error {
	if log == nil {
		q.log = nil
		return nil
	}
	err := q.checkFuncs("log")
	if err != nil {
		return err
	}

	f, err := log.Log()
	if err != nil {
		return err
	}
	q.log = f
	return nil
}

func (RankFeatureQuery) Kind() QueryKind {
	return QueryKindRankFeature
}
func (q *RankFeatureQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *RankFeatureQuery) RankFeature() (*RankFeatureQuery, error) {
	return q, nil
}
func (q *RankFeatureQuery) Clear() {
	if q == nil {
		return
	}
	*q = RankFeatureQuery{}
}

func (q *RankFeatureQuery) UnmarshalJSON(data []byte) error {
	*q = RankFeatureQuery{}
	r := rankFeatureQuery{}
	err := r.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	err = q.SetBoost(r.Boost)
	if err != nil {
		return err
	}
	err = q.SetField(r.Field)
	if err != nil {
		return err
	}
	q.linear = r.Linear
	q.log = r.Log
	q.saturation = r.Saturation
	q.sigmoid = r.Sigmoid
	return nil
}
func (q RankFeatureQuery) MarshalJSON() ([]byte, error) {
	return rankFeatureQuery{
		Name:       q.name,
		Boost:      q.boost.Value(),
		Field:      q.field,
		Log:        q.log,
		Linear:     q.linear,
		Sigmoid:    q.sigmoid,
		Saturation: q.saturation,
	}.MarshalJSON()
}
func (q *RankFeatureQuery) UnmarshalBSON(data []byte) error {
	return q.UnmarshalJSON(data)
}
func (q RankFeatureQuery) MarshalBSON() ([]byte, error) {
	return q.MarshalJSON()
}
func (q *RankFeatureQuery) IsEmpty() bool {
	return q == nil || len(q.field) == 0
}

//easyjson:json
type rankFeatureQuery struct {
	Name       string              `json:"_name,omitempty"`
	Boost      interface{}         `json:"boost,omitempty"`
	Field      string              `json:"field,omitempty"`
	Log        *LogFunction        `json:"log,omitempty"`
	Linear     *LinearFunction     `json:"linear,omitempty"`
	Sigmoid    *SigmoidFunction    `json:"sigmoid,omitempty"`
	Saturation *SaturationFunction `json:"saturation,omitempty"`
}

type SaturationFunctioner interface {
	Saturation() (*SaturationFunction, error)
}

// SaturationFunctionParams gives a score equal to S / (S + pivot), where S is
// the value of the rank feature field and pivot is a configurable pivot value
// so that the result will be less than 0.5 if S is less than pivot and greater
// than 0.5 otherwise. Scores are always (0,1).
//
// If the rank feature has a negative score impact then the function will be
// computed as pivot / (S + pivot), which decreases when S increases.
//
// If a pivot value is not provided, Elasticsearch computes a default value
// equal to the approximate geometric mean of all rank feature values in the
// index.
//
// If a pivot value is not provided, Elasticsearch computes a default value
// equal to the approximate geometric mean of all rank feature values in the
// index. We recommend using this default value if you haven’t had the
// opportunity to train a good pivot value.
type SaturationFunctionParams struct {
	// (Optional, number)
	Pivot interface{}
}

//easyjson:json
type saturationFunction struct {
	Pivot interface{} `json:"pivot,omitempty"`
}

func (p SaturationFunctionParams) Saturation() (*SaturationFunction, error) {
	f := &SaturationFunction{}
	return f, f.SetPivot(p.Pivot)
}

// SaturationFunction gives a score equal to S / (S + pivot), where S is the
// value of the rank feature field and pivot is a configurable pivot value so
// that the result will be less than 0.5 if S is less than pivot and greater
// than 0.5 otherwise. Scores are always (0,1).
//
// If the rank feature has a negative score impact then the function will be
// computed as pivot / (S + pivot), which decreases when S increases.
//
// If a pivot value is not provided, Elasticsearch computes a default value
// equal to the approximate geometric mean of all rank feature values in the
// index.
type SaturationFunction struct {
	pivot dynamic.Number
}

func (s SaturationFunction) MarshalJSON() ([]byte, error) {
	return saturationFunction{
		Pivot: s.pivot.Value(),
	}.MarshalJSON()
}
func (s *SaturationFunction) UnmarshalJSON(data []byte) error {
	*s = SaturationFunction{}
	v := saturationFunction{}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	err = s.pivot.Set(v.Pivot)
	if err != nil {
		return err
	}
	return nil
}

func (s SaturationFunction) MarshalBSON() ([]byte, error) {
	return s.MarshalJSON()
}
func (s *SaturationFunction) UnmarshalBSON(data []byte) error {
	return s.UnmarshalJSON(data)
}

func (s *SaturationFunction) Saturation() (*SaturationFunction, error) {
	return s, nil
}

func (s SaturationFunction) Pivot() float64 {
	if f, ok := s.pivot.Float64(); ok {
		return f
	}
	return 0
}

func (s *SaturationFunction) SetPivot(pivot interface{}) error {
	return s.pivot.Set(pivot)
}

type LogFunctioner interface {
	Log() (*LogFunction, error)
}

// LogFunction is a logarithmic function used to boost relevance scores based on
// the value of the rank feature field.
type LogFunction struct {
	scalingFactorParam
}

//easyjson:json
type logFunction struct {
	ScalingFactor interface{} `json:"scaling_factor,omitempty"`
}

func (l *LogFunction) Log() (*LogFunction, error) {
	return l, nil
}

func (l LogFunction) MarshalJSON() ([]byte, error) {
	return logFunction{
		ScalingFactor: l.scalingFactor.Value(),
	}.MarshalJSON()
}
func (l *LogFunction) UnmarshalJSON(data []byte) error {
	*l = LogFunction{}
	v := logFunction{}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	return l.SetScalingFactor(v.ScalingFactor)
}

type LogFunctionParams struct {
	ScalingFactor interface{}
}

func (p LogFunctionParams) Log() (*LogFunction, error) {
	f := &LogFunction{}
	return f, f.SetScalingFactor(p.ScalingFactor)
}

type SigmoidFunctionParams struct {
	Pivot    interface{}
	Exponent interface{}
}

func (p SigmoidFunctionParams) Sigmoid() (*SigmoidFunction, error) {
	f := &SigmoidFunction{}
	err := f.SetExponent(p.Exponent)
	if err != nil {
		return f, err
	}
	err = f.SetPivot(p.Pivot)
	return f, err
}

type SigmoidFunctioner interface {
	Sigmoid() (*SigmoidFunction, error)
}

type SigmoidFunction struct {
	pivot    dynamic.Number
	exponent dynamic.Number
}

func (s SigmoidFunction) MarshalJSON() ([]byte, error) {
	return sigmoidFunction{
		Pivot:    s.pivot.Value(),
		Exponent: s.exponent.Value(),
	}.MarshalJSON()
}
func (s *SigmoidFunction) UnmarshalJSON(data []byte) error {
	*s = SigmoidFunction{}
	v := sigmoidFunction{}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	err = s.exponent.Set(v.Exponent)
	if err != nil {
		return err
	}
	err = s.pivot.Set(v.Pivot)
	if err != nil {
		return err
	}
	return nil
}

func (s SigmoidFunction) MarshalBSON() ([]byte, error) {
	return s.MarshalJSON()
}
func (s *SigmoidFunction) UnmarshalBSON(data []byte) error {
	return s.UnmarshalJSON(data)
}

//easyjson:json
type sigmoidFunction struct {
	Pivot    interface{} `json:"pivot,omitempty"`
	Exponent interface{} `json:"exponent,omitempty"`
}

func (s *SigmoidFunction) Sigmoid() (*SigmoidFunction, error) {
	return s, nil
}

func (s SigmoidFunction) Pivot() float64 {
	if f, ok := s.pivot.Float64(); ok {
		return f
	}
	return 0
}
func (s *SigmoidFunction) SetPivot(pivot interface{}) error {
	err := s.pivot.Set(pivot)
	if err != nil {
		return err
	}
	f, ok := s.pivot.Float64()
	if ok && f > 0 {
		return nil
	}
	if ok {
		return fmt.Errorf("%w, received %f", ErrInvalidPivot, f)
	}
	return ErrPivotRequired
}
func (s SigmoidFunction) Exponent() float64 {
	if f, ok := s.exponent.Float64(); ok {
		return f
	}
	return 0
}
func (s *SigmoidFunction) SetExponent(exponent interface{}) error {
	err := s.exponent.Set(exponent)
	if err != nil {
		return err
	}
	f, ok := s.exponent.Float64()
	if ok && f > 0 {
		return nil
	}
	if ok {
		return fmt.Errorf("%w, received %f", ErrInvalidExponent, f)
	}
	return ErrExponentRequired
}

type LinearFunctioner interface {
	Linear() (*LinearFunction, error)
}

type LinearFunctionParams struct{}

func (l LinearFunction) MarshalJSON() ([]byte, error) {
	return []byte(`{}`), nil
}
func (l *LinearFunction) UnmarshalJSON(data []byte) error {
	return nil
}
func (l *LinearFunction) UnmarshalBSON(data []byte) error {
	return nil
}

func (l LinearFunction) MarshalBSON() ([]byte, error) {
	return l.MarshalJSON()
}

func (p LinearFunctionParams) Linear() (*LinearFunction, error) {
	return &LinearFunction{}, nil
}

// LinearFunction for RankFeatureQuery.
type LinearFunction struct{}

func (l *LinearFunction) Linear() (*LinearFunction, error) {
	return l, nil
}
