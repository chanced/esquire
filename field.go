package picker

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// TODO TypeVersion
// TODO TypeMurmur3
// TODO TypeAggregateMetricDouble
// TODO TypeHistogram
// TODO TypeAnnotatedText

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
	FieldTypeHistogram:       func() Field { return &HistogramField{} },
}

func UnmarshalFieldJSON(data []byte, field *Field) error {
	if len(data) == 0 {
		return errors.New("picker: JSON is empty; can not unmarshal Field")
	}
	g := gjson.GetBytes(data, "type")
	if !g.Exists() {
		return fmt.Errorf("%w; can not unmarshal JSON", ErrMissingType)
	}
	str := g.String()
	if len(str) == 0 {
		return fmt.Errorf("%w; can not unmarshal JSON", ErrMissingType)
	}
	handler := FieldTypeHandlers[FieldType(str)]
	if handler == nil {
		handler = FieldTypeHandlers[FieldType(strings.ToLower(str))]
	}
	if handler == nil {
		return fmt.Errorf("%w <%s>", ErrUnsupportedType, str)
	}
	*field = handler()
	return (*field).UnmarshalJSON(data)
}

// Field is an elasticsearch field mapping
type Field interface {
	Fielder
	json.Marshaler
	json.Unmarshaler
}

type Fielder interface {
	Field() (Field, error)
	Type() FieldType
}
