package mapping

// A WildcardField stores values optimised for wildcard grep-like queries. Wildcard queries are possible on other field types but suffer from constraints:
//
//  	- text fields limit matching of any wildcard expressions to individual tokens rather than the original whole value held in a field
//
//  	- keyword fields are untokenized but slow at performing wildcard queries (especially patterns with leading wildcards).
//
// Internally the wildcard field indexes the whole field value using ngrams and stores the full string. The index is used as a rough filter to cut down the number of values that are then checked by retrieving and checking the full values. This field is especially well suited to run grep-like queries on log lines. Storage costs are typically lower than those of keyword fields but search speeds for exact matches on full terms are slower.
//
// Limitations
//
// wildcard fields are untokenized like keyword fields, so do not support queries that rely on word positions such as phrase queries.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/keyword.html#wildcard-field-type
type WildcardField struct {
	BaseField        `bson:",inline" json:",inline"`
	NullValueParam   `bson:",inline" json:",inline"`
	IgnoreAboveParam `bson:",inline" json:",inline"`
}

func (f WildcardField) Clone() Field {
	n := NewWildcardField()
	n.SetNullValue(f.NullValue())
	n.SetIgnoreAbove(f.IgnoreAbove())
	return n
}

func NewWildcardField() *WildcardField {
	return &WildcardField{BaseField: BaseField{MappingType: TypeWildcardKeyword}}
}
