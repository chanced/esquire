package mapping

import (
	"errors"
)

var (
	// ErrInvalidType indicates an invalid type
	ErrInvalidType = errors.New("invalid Type")
)

// Type of ElasticSearch Mapping
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html#mapping-types
type Type string

//https://www.elastic.co/guide/en/elasticsearch/reference/current/multi-fields.html

const (
	// TypeBinary is a value encoded as a Base64 string.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/binary.html
	TypeBinary Type = "binary"

	// TypeBoolean is for true and false values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/boolean.html
	TypeBoolean Type = "boolean"

	// TypeKeyword is used for structured content such as
	// IDs, email addresses, hostnames, status codes, zip codes, or tags.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#keyword-field-type
	TypeKeyword Type = "keyword"

	// TypeConstant is for keyword fields that always contain the same value.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#constant-keyword-field-type
	TypeConstant Type = "constant_keyword"

	// TypeWildcardKeyword optimizes log lines and similar keyword values for grep-like wildcard queries.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#wildcard-field-type
	TypeWildcardKeyword Type = "wildcard"

	// TypeLong is a signed 64-bit integer with a minimum value of -263 and a maximum value of 263-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeLong Type = "long"

	// TypeInteger is a signed 32-bit integer with a minimum value of -231 and a maximum value of 231-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeInteger Type = "integer"

	// TypeShort is a signed 16-bit integer with a minimum value of -32,768 and a maximum value of 32,767.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeShort Type = "short"

	// TypeByte is a signed 8-bit integer with a minimum value of -128 and a maximum value of 127.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeByte Type = "byte"

	// TypeDouble is a double-precision 64-bit IEEE 754 floating point number, restricted to finite values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeDouble Type = "double"

	// TypeFloat is a single-precision 32-bit IEEE 754 floating point number, restricted to finite values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeFloat Type = "float"

	// TypeHalfFloat is a half-precision 16-bit IEEE 754 floating point number, restricted to finite values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeHalfFloat Type = "half_float"

	// TypeScaledFloat is a floating point number that is backed by a long, scaled by a fixed double scaling factor.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeScaledFloat Type = "scaled_float"

	// TypeUnsignedLong is an unsigned 64-bit integer with a minimum value of 0 and a maximum value of 264-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/number.html
	TypeUnsignedLong Type = "unsigned_long"

	// TypeDateNanos is an addition to the date data type. However there is an important
	// distinction between the two. The existing date data type stores dates in millisecond
	// resolution. The date_nanos data type stores dates in nanosecond resolution, which
	// limits its range of dates from roughly 1970 to 2262, as dates are still stored as a
	// long representing nanoseconds since the epoch.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/date_nanos.html
	TypeDateNanos Type = "date_nanos"

	// TypeDate is a date in one of the following formats:
	//
	//      - strings containing formatted dates, e.g. "2015-01-01" or "2015/01/01 12:10:30".
	//
	//      - a long number representing milliseconds-since-the-epoch.
	//
	//      - an integer representing seconds-since-the-epoch.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/date.html
	TypeDate Type = "date"

	// TypeAlias defines an alias for an existing field.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/alias.html
	TypeAlias Type = "alias"

	// TypeObject is a JSON object.
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
	TypeObject Type = "object"

	// TypeFlattened is an entire JSON object as a single field value.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/flattened.html
	TypeFlattened Type = "flattened"

	// TypeNested is a JSON object that preserves the relationship between its subfields.
	//
	// The nested type is a specialised version of the object data type that allows arrays
	// of objects to be indexed in a way that they can be queried independently of each other.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/nested.html
	TypeNested Type = "nested"

	//TypeJoin defines a parent/child relationship for documents in the same index.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
	TypeJoin Type = "join"

	// TypeLongRange is a range of signed 64-bit integers with a minimum value of -263
	// and maximum of 263-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	TypeLongRange Type = "long_range"

	// TypeIntegerRange is range of signed 32-bit integers with a minimum value of -231
	// and maximum of 231-1.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	TypeIntegerRange Type = "integer_range"

	// TypeFloatRange is a range of single-precision 32-bit IEEE 754 floating point values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	TypeFloatRange Type = "float_range"

	// TypeDoubleRange is a range of double-precision 64-bit IEEE 754 floating point values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	TypeDoubleRange Type = "double_range"

	// TypeDateRange is a range of date values. Date ranges support various date formats
	// through the format mapping parameter. Regardless of the format used, date values
	// are parsed into an unsigned 64-bit integer representing milliseconds since the
	// Unix epoch in UTC. Values containing the now date math expression are not supported.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	TypeDateRange Type = "date_range"

	// TypeIPRange is a range of ip values supporting either IPv4 or IPv6 (or mixed) addresses.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/range.html
	TypeIPRange Type = "ip_range"

	// TypeIP is IPv4 and IPv6 addresses.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ip.html
	TypeIP Type = "ip"

	// TypeVersion is a software version. Supports Semantic Versioning precedence rules.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/version.html
	TypeVersion Type = "version"

	// TypeMurmur3 compute and stores hashes of values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/plugins/7.11/mapper-murmur3.html
	TypeMurmur3 Type = "murmur3"

	// TypeAggregateMetricDouble is a pre-aggregated metric values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/aggregate-metric-double.html
	TypeAggregateMetricDouble Type = "aggregate_metric_double"

	// TypeHistogram is a pre-aggregated numerical values in the form of a histogram.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/histogram.html
	TypeHistogram Type = "histogram"

	// TypeText is analyzed, unstructured text.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/text.html
	TypeText Type = "text"

	// TypeAnnotatedText is text containing special markup. Used for identifying named entities.
	//
	// https://www.elastic.co/guide/en/elasticsearch/plugins/7.11/mapper-annotated-text.html
	TypeAnnotatedText Type = "annotated-text"

	// TypeCompletion is used for auto-complete suggestions.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-suggesters.html#completion-suggester
	TypeCompletion Type = "completion"

	// TypeSearchAsYouType is text-like type for as-you-type completion.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html
	TypeSearchAsYouType Type = "search_as_you_type"

	// TypeTokenCount is a count of tokens in a text.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/token-count.html
	TypeTokenCount Type = "token_count"

	// TypeDenseVector is a Document Ranking Type that records dense vectors of float values.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/dense-vector.html
	TypeDenseVector Type = "dense_vector"

	// TypeRankFeature is a Document Ranking Type that can index numbers so that they can later
	// be used to boost documents in queries with a rank_feature query
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-feature.html
	TypeRankFeature Type = "rank_feature"

	// TypeRankFeatures is a Document Ranking Type that can index numeric feature vectors,
	// so that they can later be used to boost documents in queries with a rank_feature query.
	//
	// It is analogous to the rank_feature data type but is better suited when the list of
	// features is sparse so that it wouldn’t be reasonable to add one field to the mappings
	// for each of them.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rank-features.html
	TypeRankFeatures Type = "rank_features"

	// TypeGeoPoint is  a latitude and longitude points.
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
	TypeGeoPoint Type = "geo_point"

	// TypeGeoShape are complex shapes, such as polygons.
	//
	// The geo_shape data type facilitates the indexing of and searching with arbitrary
	// geo shapes such as rectangles and polygons. It should be used when either the
	// data being indexed or the queries being executed contain shapes other than just
	// points.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-shape.html
	TypeGeoShape Type = "geo_shape"
	// TypePoint is an arbitrary cartesian points.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point.html
	TypePoint Type = "point"

	// TypeShape data type facilitates the indexing of and searching with arbitrary
	// x, y cartesian shapes such as rectangles and polygons. It can be used to
	// index and query geometries whose coordinates fall in a 2-dimensional planar
	// coordinate system.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/shape.html
	TypeShape Type = "shape"

	// TypePercolator indexes queries written in Query DSL.
	//
	// The percolator field type parses a json structure into a native query and stores that
	// query, so that the percolate query can use it to match provided documents.
	//
	// Any field that contains a json object can be configured to be a percolator field.
	// The percolator field type has no settings. Just configuring the percolator field type
	// is sufficient to instruct Elasticsearch to treat a field as a query.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/percolator.html
	TypePercolator Type = "percolator"
)

// TODO TypeVersion
// TODO TypeMurmur3
// TODO TypeAggregateMetricDouble
// TODO TypeHistogram
// TODO TypeAnnotatedText
// TODO TypePoint
// TODO TypeShape
