package mapping

// FieldTypeHandlers is a map of mapping Type to func that returns a Field instantiated with the appropriate Type
var FieldTypeHandlers = map[FieldType]func() Field{
	FieldTypeAlias:           func() Field { return &AliasField{} },
	FieldTypeBinary:          func() Field { return NewBinaryField() },
	FieldTypeBoolean:         func() Field { return NewBooleanField() },
	FieldTypeByte:            func() Field { return NewByteField() },
	FieldTypeCompletion:      func() Field { return NewCompletionField() },
	FieldTypeConstant:        func() Field { return NewConstantField() },
	FieldTypeDate:            func() Field { return NewDateField() },
	FieldTypeDateNanos:       func() Field { return NewDateNanoSecField() },
	FieldTypeDateRange:       func() Field { return NewDateRangeField() },
	FieldTypeDenseVector:     func() Field { return NewDenseVectorField() },
	FieldTypeDouble:          func() Field { return NewDoubleField() },
	FieldTypeDoubleRange:     func() Field { return NewDoubleRangeField() },
	FieldTypeFlattened:       func() Field { return NewFlattenedField() },
	FieldTypeFloat:           func() Field { return NewFloatField() },
	FieldTypeFloatRange:      func() Field { return NewFloatRangeField() },
	FieldTypeGeoPoint:        func() Field { return NewGeoPointField() },
	FieldTypeGeoShape:        func() Field { return NewGeoShapeField() },
	FieldTypeHalfFloat:       func() Field { return NewHalfFloatField() },
	FieldTypeIP:              func() Field { return NewIPField() },
	FieldTypeIPRange:         func() Field { return NewIPRangeField() },
	FieldTypeInteger:         func() Field { return NewIntegerField() },
	FieldTypeIntegerRange:    func() Field { return NewIntegerRangeField() },
	FieldTypeJoin:            func() Field { return NewJoinField() },
	FieldTypeKeyword:         func() Field { return NewKeywordField() },
	FieldTypeLong:            func() Field { return NewLongField() },
	FieldTypeLongRange:       func() Field { return NewLongRangeField() },
	FieldTypeNested:          func() Field { return NewNestedField() },
	FieldTypeObject:          func() Field { return NewObjectField() },
	FieldTypePercolator:      func() Field { return NewPercolatorField() },
	FieldTypeRankFeature:     func() Field { return NewRankFeatureField() },
	FieldTypeRankFeatures:    func() Field { return NewRankFeaturesField() },
	FieldTypeScaledFloat:     func() Field { return NewScaledFloatField() },
	FieldTypeSearchAsYouType: func() Field { return NewSearchAsYouTypeField() },
	FieldTypeShort:           func() Field { return NewShortField() },
	FieldTypeText:            func() Field { return NewTextField() },
	FieldTypeTokenCount:      func() Field { return NewTokenCountField() },
	FieldTypeUnsignedLong:    func() Field { return NewUnsignedLongField() },
	FieldTypeWildcardKeyword: func() Field { return NewWildcardField() },
}

// NewField returns an instantiated Mapping
func NewField(t FieldType) (Field, error) {
	handler, ok := FieldTypeHandlers[t]
	if !ok {
		return nil, ErrInvalidType
	}
	return handler(), nil
}

// Field is an elasticsearch field mapping
type Field interface {
	Type() FieldType
}
