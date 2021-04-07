package picker

type QueryKind string

func (t QueryKind) String() string {
	return string(t)
}

func (t QueryKind) IsValid() bool {
	_, ok := queryKindHandlers[t]
	return ok
}

const (
	QueryKindPrefix            QueryKind = "prefix"
	QueryKindMatch             QueryKind = "match"
	QueryKindMatchAll          QueryKind = "match_all"
	QueryKindMatchNone         QueryKind = "match_none"
	QueryKindTerm              QueryKind = "term"
	QueryKindExists            QueryKind = "exists"
	QueryKindTerms             QueryKind = "terms"
	QueryKindRange             QueryKind = "range"
	QueryKindBoosting          QueryKind = "boosting"
	QueryKindBoolean           QueryKind = "bool"
	QueryKindConstantScore     QueryKind = "constant_score"
	QueryKindFunctionScore     QueryKind = "function_score"
	QueryKindDisjunctionMax    QueryKind = "dis_max"
	QueryKindAllOf             QueryKind = "all_of"
	QueryKindFuzzy             QueryKind = "fuzzy"
	QueryKindScriptScore       QueryKind = "script_score"
	QueryKindScript            QueryKind = "script"
	QueryKindIDs               QueryKind = "ids"
	QueryKindIntervals         QueryKind = "intervals"
	QueryKindMatchBoolPrefix   QueryKind = "match_bool_prefix"
	QueryKindMatchPhrase       QueryKind = "match_phrase"
	QueryKindMatchPhrasePrefix QueryKind = "match_phrase_prefix"
	QueryKindMultiMatch        QueryKind = "multi_match"
	QueryKindQueryString       QueryKind = "query_string"
	QueryKindSimpleQueryString QueryKind = "simple_query_string"
	QueryKindGeoBoundingBox    QueryKind = "geo_bounding_box"

	// TODO:
	QueryKindGeoDistance      QueryKind = "geo_distance"
	QueryKindGeoPolygon       QueryKind = "geo_polygon"
	QueryKindGeoShape         QueryKind = "geo_shape"
	QueryKindShape            QueryKind = "shape"
	QueryKindNested           QueryKind = "nested"
	QueryKindHasChild         QueryKind = "has_child"
	QueryKindHasParent        QueryKind = "has_parent"
	QueryKindParentID         QueryKind = "parent_id"
	QueryKindDistantFeature   QueryKind = "distance_feature"
	QueryKindMoreLikeThis     QueryKind = "more_like_this"
	QueryKindPercolate        QueryKind = "percolate"
	QueryKindRankFeature      QueryKind = "rank_feature"
	QueryKindWrapper          QueryKind = "wrapper"
	QueryKindPinned           QueryKind = "pinned"
	QueryKindSpanContaining   QueryKind = "span_containing"
	QueryKindFieldMaskingSpan QueryKind = "field_masking_span"
	QueryKindSpanFirst        QueryKind = "span_first"
	QueryKindSpanMulti        QueryKind = "span_multi"
	QueryKindSpanNear         QueryKind = "span_near"
	QueryKindSpanNot          QueryKind = "span_not"
	QueryKindSpanOr           QueryKind = "span_or"
	QueryKindSpanTerm         QueryKind = "span_term"
	QueryKindSpanWithin       QueryKind = "span_within"
)

var queryKindHandlers = map[QueryKind]func() QueryClause{
	QueryKindPrefix:            func() QueryClause { return &PrefixQuery{} },
	QueryKindMatch:             func() QueryClause { return &MatchQuery{} },
	QueryKindTerm:              func() QueryClause { return &TermQuery{} },
	QueryKindTerms:             func() QueryClause { return &TermsQuery{} },
	QueryKindBoolean:           func() QueryClause { return &BoolQuery{} },
	QueryKindExists:            func() QueryClause { return &ExistsQuery{} },
	QueryKindRange:             func() QueryClause { return &RangeQuery{} },
	QueryKindMatchAll:          func() QueryClause { return &MatchAllQuery{} },
	QueryKindMatchNone:         func() QueryClause { return &MatchNoneQuery{} },
	QueryKindScript:            func() QueryClause { return &ScriptQuery{} },
	QueryKindBoosting:          func() QueryClause { return &BoostingQuery{} },
	QueryKindConstantScore:     func() QueryClause { return &ConstantScoreQuery{} },
	QueryKindIDs:               func() QueryClause { return &IDsQuery{} },
	QueryKindIntervals:         func() QueryClause { return &IntervalsQuery{} },
	QueryKindMatchPhrase:       func() QueryClause { return &MatchPhraseQuery{} },
	QueryKindMatchPhrasePrefix: func() QueryClause { return &MatchPhrasePrefixQuery{} },
	QueryKindMultiMatch:        func() QueryClause { return &MultiMatchQuery{} },
	QueryKindQueryString:       func() QueryClause { return &QueryStringQuery{} },
	QueryKindSimpleQueryString: func() QueryClause { return &SimpleQueryStringQuery{} },
	QueryKindGeoBoundingBox:    func() QueryClause { return &GeoBoundingBoxQuery{} },
	// QueryKindGeoDistance:       func() QueryClause { return &GeoDistanceQuery{} },
	// QueryKindGeoPolygon:        func() QueryClause { return &GeoPolygonQuery{} },
	// QueryKindGeoShape:          func() QueryClause { return &GeoShapeQuery{} },
	// QueryKindShape:             func() QueryClause { return &ShapeQuery{} },
	// QueryKindNested:            func() QueryClause { return &NestedQuery{} },
	// QueryKindHasChild:          func() QueryClause { return &HasChildQuery{} },
	// QueryKindHasParent:         func() QueryClause { return &HasParentQuery{} },
	// QueryKindParentID:          func() QueryClause { return &ParentIDQuery{} },
	// QueryKindDistantFeature:    func() QueryClause { return &DistantFeatureQuery{} },
	// QueryKindMoreLikeThis:      func() QueryClause { return &MoreLikeThisQuery{} },
	// QueryKindPercolate:         func() QueryClause { return &PercolateQuery{} },
	// QueryKindRankFeature:       func() QueryClause { return &RankFeatureQuery{} },
	// QueryKindWrapper:           func() QueryClause { return &WrapperQuery{} },
	// QueryKindPinned:            func() QueryClause { return &PinnedQuery{} },
	// QueryKindSpanContaining:    func() QueryClause { return &SpanContainingQuery{} },
	// QueryKindFieldMaskingSpan:  func() QueryClause { return &FieldMaskingSpanQuery{} },
	// QueryKindSpanFirst:         func() QueryClause { return &SpanFirstQuery{} },
	// QueryKindSpanMulti:         func() QueryClause { return &SpanMultiQuery{} },
	// QueryKindSpanNear:          func() QueryClause { return &SpanNearQuery{} },
	// QueryKindSpanNot:           func() QueryClause { return &SpanNotQuery{} },
	// QueryKindSpanOr:            func() QueryClause { return &SpanOrQuery{} },
	// QueryKindSpanTerm:          func() QueryClause { return &SpanTermQuery{} },
	// QueryKindSpanWithin:        func() QueryClause { return &SpanWithinQuery{} },
}
