package mapping

// FieldTypeHandlers is a map of mapping Type to func that returns a Field instantiated with the appropriate Type
var FieldTypeHandlers = map[Type]func() Field{
	TypeAlias:           func() Field { return NewAliasField() },
	TypeBinary:          func() Field { return NewBinaryField() },
	TypeBoolean:         func() Field { return NewBooleanField() },
	TypeByte:            func() Field { return NewByteField() },
	TypeCompletion:      func() Field { return NewCompletionField() },
	TypeConstant:        func() Field { return NewConstantField() },
	TypeDate:            func() Field { return NewDateField() },
	TypeDateNanos:       func() Field { return NewDateNanoSecField() },
	TypeDateRange:       func() Field { return NewDateRangeField() },
	TypeDenseVector:     func() Field { return NewDenseVectorField() },
	TypeDouble:          func() Field { return NewDoubleField() },
	TypeDoubleRange:     func() Field { return NewDoubleRangeField() },
	TypeFlattened:       func() Field { return NewFlattenedField() },
	TypeFloat:           func() Field { return NewFloatField() },
	TypeFloatRange:      func() Field { return NewFloatRangeFIeld() },
	TypeGeoPoint:        func() Field { return NewGeoPointField() },
	TypeGeoShape:        func() Field { return NewGeoShapeField() },
	TypeHalfFloat:       func() Field { return NewHalfFloatField() },
	TypeIP:              func() Field { return NewIPField() },
	TypeIPRange:         func() Field { return NewIPRangeField() },
	TypeInteger:         func() Field { return NewIntegerField() },
	TypeIntegerRange:    func() Field { return NewIntegerRangeField() },
	TypeJoin:            func() Field { return NewJoinField() },
	TypeKeyword:         func() Field { return NewKeywordField() },
	TypeLong:            func() Field { return NewLongField() },
	TypeLongRange:       func() Field { return NewLongRangeField() },
	TypeNested:          func() Field { return NewNestedField() },
	TypeObject:          func() Field { return NewObjectField() },
	TypePercolator:      func() Field { return NewPerculatorField() },
	TypeRankFeature:     func() Field { return NewRankFeatureField() },
	TypeRankFeatures:    func() Field { return NewRankFeaturesField() },
	TypeScaledFloat:     func() Field { return NewScaledFloatField() },
	TypeSearchAsYouType: func() Field { return NewSearchAsYouTypeField() },
	TypeShort:           func() Field { return NewShortField() },
	TypeText:            func() Field { return NewTextField() },
	TypeTokenCount:      func() Field { return NewTokenCountField() },
	TypeUnsignedLong:    func() Field { return NewUnsignedLongField() },
	TypeWildcardKeyword: func() Field { return NewWildcardField() },
}

// NewField returns an instantiated Mapping
func NewField(t Type) (Field, error) {
	handler, ok := FieldTypeHandlers[t]
	if !ok {
		return nil, ErrInvalidType
	}
	return handler(), nil
}

// Field is an elasticsearch field mapping
type Field interface {
	Type() Type
}

// BaseField Mapping
type BaseField struct {
	MappingType Type `bson:"type" json:"type"`
}

func (b BaseField) Type() Type {
	return b.MappingType
}
