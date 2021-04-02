package picker

import "encoding/json"

// TODO TypeVersion
// TODO TypeMurmur3
// TODO TypeAggregateMetricDouble
// TODO TypeHistogram
// TODO TypeAnnotatedText
// TODO TypePoint
// TODO TypeShape

// FieldTypeHandlers is a map of mapping Type to func that returns a Field instantiated with the appropriate Type
var FieldTypeHandlers = map[FieldType]func() Field{
	FieldTypeAlias:           func() Field { return &AliasField{} },
	FieldTypeBinary:          func() Field { return &BinaryField{} },
	FieldTypeBoolean:         func() Field { return &BooleanField{} },
	FieldTypeByte:            func() Field { return &ByteField{} },
	FieldTypeCompletion:      func() Field { return &CompletionField{} },
	FieldTypeConstant:        func() Field { return &ConstantField{} },
	FieldTypeDate:            func() Field { return &DateField{} },
	FieldTypeDateNanos:       func() Field { return &DateNanoSecField{} },
	FieldTypeDateRange:       func() Field { return &DateRangeField{} },
	FieldTypeDenseVector:     func() Field { return &DenseVectorField{} },
	FieldTypeDouble:          func() Field { return &DoubleField{} },
	FieldTypeDoubleRange:     func() Field { return &DoubleRangeField{} },
	FieldTypeFlattened:       func() Field { return &FlattenedField{} },
	FieldTypeFloat:           func() Field { return &FloatField{} },
	FieldTypeFloatRange:      func() Field { return &FloatRangeField{} },
	FieldTypeGeoPoint:        func() Field { return &GeoPointField{} },
	FieldTypeGeoShape:        func() Field { return &GeoShapeField{} },
	FieldTypeHalfFloat:       func() Field { return &HalfFloatField{} },
	FieldTypeIP:              func() Field { return &IPField{} },
	FieldTypeIPRange:         func() Field { return &IPRangeField{} },
	FieldTypeInteger:         func() Field { return &IntegerField{} },
	FieldTypeIntegerRange:    func() Field { return &IntegerRangeField{} },
	FieldTypeJoin:            func() Field { return &JoinField{} },
	FieldTypeKeyword:         func() Field { return &KeywordField{} },
	FieldTypeLong:            func() Field { return &LongField{} },
	FieldTypeLongRange:       func() Field { return &LongRangeField{} },
	FieldTypeNested:          func() Field { return &NestedField{} },
	FieldTypeObject:          func() Field { return &ObjectField{} },
	FieldTypePercolator:      func() Field { return &PercolatorField{} },
	FieldTypeRankFeature:     func() Field { return &RankFeatureField{} },
	FieldTypeRankFeatures:    func() Field { return &RankFeaturesField{} },
	FieldTypeScaledFloat:     func() Field { return &ScaledFloatField{} },
	FieldTypeSearchAsYouType: func() Field { return &SearchAsYouTypeField{} },
	FieldTypeShort:           func() Field { return &ShortField{} },
	FieldTypeText:            func() Field { return &TextField{} },
	FieldTypeTokenCount:      func() Field { return &TokenCountField{} },
	FieldTypeUnsignedLong:    func() Field { return &UnsignedLongField{} },
	FieldTypeWildcardKeyword: func() Field { return &WildcardField{} },
}

// Field is an elasticsearch field mapping
type Field interface {
	Type() FieldType
	json.Marshaler
	json.Unmarshaler
}
