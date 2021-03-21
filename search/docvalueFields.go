package search

// DocValueFields is used to return doc values for one or more fields in the search response.
//
// (Optional, array of strings and objects) Array of wildcard (*) patterns. The request returns doc values for field names matching these patterns in the hits.fields property of the response.
// You can use the docvalue_fields parameter to return doc values for one or more fields in the search response.
//
// Doc values store the same values as the _source but in an on-disk, column-based structure that’s optimized for sorting and aggregations. Since each field is stored separately, Elasticsearch only reads the field values that were requested and can avoid loading the whole document _source.
//
// Doc values are stored for supported fields by default. However, doc values are not supported for text or text_annotated fields.
//
// The following search request uses the docvalue_fields parameter to retrieve doc values for the user.id field, all fields starting with http.response., and the @timestamp field:
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-fields.html#docvalue-fields
//
// See also:
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/doc-values.html
type DocValueFields Fields
