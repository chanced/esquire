package picker

import (
	"errors"
)

var (
	// ErrInvalidType indicates an invalid type
	ErrInvalidType = errors.New("invalid Type")
)

// FieldType of ElasticSearch Mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html#mapping-types
type FieldType string

//https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html

const (
	// FieldTypeBinary is a value encoded as a Base64 string.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/binary.html
	FieldTypeBinary FieldType = "binary"

	// FieldTypeBoolean is for true and false values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/boolean.html
	FieldTypeBoolean FieldType = "boolean"

	// FieldTypeKeyword is used for structured content such as
	// IDs, email addresses, hostnames, status codes, zip codes, or tags.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#keyword-field-type
	FieldTypeKeyword FieldType = "keyword"

	// FieldTypeConstant is for keyword fields that always contain the same value.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-field-type
	FieldTypeConstant FieldType = "constant_keyword"

	// FieldTypeWildcardKeyword optimizes log lines and similar keyword values for grep-like wildcard queries.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#wildcard-field-type
	FieldTypeWildcardKeyword FieldType = "wildcard"

	// FieldTypeLong is a signed 64-bit integer with a minimum value of -263 and a maximum value of 263-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeLong FieldType = "long"

	// FieldTypeInteger is a signed 32-bit integer with a minimum value of -231 and a maximum value of 231-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeInteger FieldType = "integer"

	// FieldTypeShort is a signed 16-bit integer with a minimum value of -32,768 and a maximum value of 32,767.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeShort FieldType = "short"

	// FieldTypeByte is a signed 8-bit integer with a minimum value of -128 and a maximum value of 127.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeByte FieldType = "byte"

	// FieldTypeDouble is a double-precision 64-bit IEEE 754 floating point number, restricted to finite values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeDouble FieldType = "double"

	// FieldTypeFloat is a single-precision 32-bit IEEE 754 floating point number, restricted to finite values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeFloat FieldType = "float"

	// FieldTypeHalfFloat is a half-precision 16-bit IEEE 754 floating point number, restricted to finite values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeHalfFloat FieldType = "half_float"

	// FieldTypeScaledFloat is a floating point number that is backed by a long, scaled by a fixed double scaling factor.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeScaledFloat FieldType = "scaled_float"

	// FieldTypeUnsignedLong is an unsigned 64-bit integer with a minimum value of 0 and a maximum value of 264-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	FieldTypeUnsignedLong FieldType = "unsigned_long"

	// FieldTypeDateNanos is an addition to the date data type. However there is an important
	// distinction between the two. The existing date data type stores dates in millisecond
	// resolution. The date_nanos data type stores dates in nanosecond resolution, which
	// limits its range of dates from roughly 1970 to 2262, as dates are still stored as a
	// long representing nanoseconds since the epoch.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/date_nanos.html
	FieldTypeDateNanos FieldType = "date_nanos"

	// FieldTypeDate is a date in one of the following formats:
	//
	//      - strings containing formatted dates, e.g. "2015-01-01" or "2015/01/01 12:10:30".
	//
	//      - a long number representing milliseconds-since-the-epoch.
	//
	//      - an integer representing seconds-since-the-epoch.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
	FieldTypeDate FieldType = "date"

	// FieldTypeAlias defines an alias for an existing field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/alias.html
	FieldTypeAlias FieldType = "alias"

	// FieldTypeObject is a JSON object.
	//
	// Options:
	//
	// dynamic - Whether or not new properties should be added dynamically to an
	// existing object. Accepts true (default), false and strict.
	//
	// enabled - Whether the JSON value given for the object field should be parsed
	// and indexed (true, default) or completely ignored (false).
	//
	// properties - The fields within the object, which can be of any data type,
	// including object. New properties may be added to an existing object.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/object.html
	FieldTypeObject FieldType = "object"

	// FieldTypeFlattened is an entire JSON object as a single field value.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/flattened.html
	FieldTypeFlattened FieldType = "flattened"

	// FieldTypeNested is a JSON object that preserves the relationship between its subfields.
	//
	// The nested type is a specialised version of the object data type that allows arrays
	// of objects to be indexed in a way that they can be queried independently of each other.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html
	FieldTypeNested FieldType = "nested"

	//FieldTypeJoin defines a parent/child relationship for documents in the same index.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
	FieldTypeJoin FieldType = "join"

	// FieldTypeLongRange is a range of signed 64-bit integers with a minimum value of -263
	// and maximum of 263-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	FieldTypeLongRange FieldType = "long_range"

	// FieldTypeIntegerRange is range of signed 32-bit integers with a minimum value of -231
	// and maximum of 231-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	FieldTypeIntegerRange FieldType = "integer_range"

	// FieldTypeFloatRange is a range of single-precision 32-bit IEEE 754 floating point values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	FieldTypeFloatRange FieldType = "float_range"

	// FieldTypeDoubleRange is a range of double-precision 64-bit IEEE 754 floating point values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	FieldTypeDoubleRange FieldType = "double_range"

	// FieldTypeDateRange is a range of date values. Date ranges support various date formats
	// through the format mapping parameter. Regardless of the format used, date values
	// are parsed into an unsigned 64-bit integer representing milliseconds since the
	// Unix epoch in UTC. Values containing the now date math expression are not supported.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	FieldTypeDateRange FieldType = "date_range"

	// FieldTypeIPRange is a range of ip values supporting either IPv4 or IPv6 (or mixed) addresses.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	FieldTypeIPRange FieldType = "ip_range"

	// FieldTypeIP is IPv4 and IPv6 addresses.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ip.html
	FieldTypeIP FieldType = "ip"

	// FieldTypeVersion is a software version. Supports Semantic Versioning precedence rules.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/version.html
	FieldTypeVersion FieldType = "version"

	// FieldTypeMurmur3 compute and stores hashes of values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/plugins/7.11/mapper-murmur3.html
	FieldTypeMurmur3 FieldType = "murmur3"

	// FieldTypeAggregateMetricDouble is a pre-aggregated metric values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/aggregate-metric-double.html
	FieldTypeAggregateMetricDouble FieldType = "aggregate_metric_double"

	// FieldTypeHistogram is a pre-aggregated numerical values in the form of a histogram.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/histogram.html
	FieldTypeHistogram FieldType = "histogram"

	// FieldTypeText is analyzed, unstructured text.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html
	FieldTypeText FieldType = "text"

	// FieldTypeAnnotatedText is text containing special markup. Used for identifying named entities.
	//
	// https://www.elastic.co/guide/en/elasticsearch/plugins/7.11/mapper-annotated-text.html
	FieldTypeAnnotatedText FieldType = "annotated-text"

	// FieldTypeCompletion is used for auto-complete suggestions.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
	FieldTypeCompletion FieldType = "completion"

	// FieldTypeSearchAsYouType is text-like type for as-you-type completion.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html
	FieldTypeSearchAsYouType FieldType = "search_as_you_type"

	// FieldTypeTokenCount is a count of tokens in a text.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/token-count.html
	FieldTypeTokenCount FieldType = "token_count"

	// FieldTypeDenseVector is a Document Ranking Type that records dense vectors of float values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
	FieldTypeDenseVector FieldType = "dense_vector"

	//float64 FieldTypeRankFeature is a Document Ranking Type that can index numbers so that they can later
	// be used to boost documents in queries with a rank_feature query
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-feature.html
	FieldTypeRankFeature FieldType = "rank_feature"

	// FieldTypeRankFeatures is a Document Ranking Type that can index numeric feature vectors,
	// so that they can later be used to boost documents in queries with a rank_feature query.
	//
	// It is analogous to the rank_feature data type but is better suited when the list of
	// features is sparse so that it wouldn’t be reasonable to add one field to the mappings
	// for each of them.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-features.html
	FieldTypeRankFeatures FieldType = "rank_features"

	// FieldTypeGeoPoint is  a latitude and longitude points.
	//
	// Fields of type geo_point accept latitude-longitude pairs, which can be used:
	//
	//      - to find geo-points within a bounding box, within a certain distance of a central point, or within a polygon or within a geo_shape query.
	//
	//      - to aggregate documents geographically or by distance from a central point.
	//
	//      - to integrate distance into a document’s relevance score.
	//
	//      - to sort documents by distance.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-point.html
	FieldTypeGeoPoint FieldType = "geo_point"

	// FieldTypeGeoShape are complex shapes, such as polygons.
	//
	// The geo_shape data type facilitates the indexing of and searching with arbitrary
	// geo shapes such as rectangles and polygons. It should be used when either the
	// data being indexed or the queries being executed contain shapes other than just
	// points.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html
	FieldTypeGeoShape FieldType = "geo_shape"
	// FieldTypePoint is an arbitrary cartesian points.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point.html
	FieldTypePoint FieldType = "point"

	// FieldTypeShape data type facilitates the indexing of and searching with arbitrary
	// x, y cartesian shapes such as rectangles and polygons. It can be used to
	// index and query geometries whose coordinates fall in a 2-dimensional planar
	// coordinate system.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/shape.html
	FieldTypeShape FieldType = "shape"

	// FieldTypePercolator indexes queries written in Query DSL.
	//
	// The percolator field type parses a json structure into a native query and stores that
	// query, so that the percolate query can use it to match provided documents.
	//
	// Any field that contains a json object can be configured to be a percolator field.
	// The percolator field type has no settings. Just configuring the percolator field type
	// is sufficient to instruct Elasticsearch to treat a field as a query.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/percolator.html
	FieldTypePercolator FieldType = "percolator"
)
