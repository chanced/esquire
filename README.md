# picker

Utilities that pair with the official Elasticsearch Go package

## Todo

### [Field Mappings](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/mapping-types.html)

- #### Common types

  - [ ] **[Binary](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/binary.html)**\
         Binary value encoded as a Base64 string.
  - [ ] **[Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/boolean.html)**\
         true and false values.
  - [ ] **[Keyword](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/keyword.html#keyword-field-type)**\
         used for structured content such as IDs, email addresses, hostnames, status codes, zip codes, or tags.
  - [ ] **[Constant keyword](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/keyword.html#constant-keyword-field-type)** [X-Pack]\
         Constant keyword is a specialization of the keyword field for the case that all documents in the index have the same value.
  - [ ] **[Wildcard](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/keyword.html#wildcard-field-type)** [X-Pack]\
         The wildcard field type is a specialized keyword field for unstructured machine-generated content you plan to search using grep-like wildcard and regexp queries. The wildcard type is optimized for fields with large values or high cardinality.
  - [ ] **[Numbers](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/number.html)**\
         Numeric types, such as long and double, used to express amounts.
  - [ ] **[Date](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/date.html)**\
         Date field type
  - [ ] **[Date nanoseconds](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/date_nanos.html)**\
         Date nanoseconds field type
  - [ ] **[Alias](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/alias.html)**\
         Defines an alias for an existing field.

- #### Objects and relational types

  - [ ] **[Object](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/object.html)**\
         A JSON object.
  - [ ] **[Flattened](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/flattened.html)**\
         An entire JSON object as a single field value.
  - [ ] **[Nested](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/nested.html)**\
         A JSON object that preserves the relationship between its subfields.
  - [ ] **[Join](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/parent-join.html)**\
         Defines a parent/child relationship for documents in the same index.

- #### Structured data types

  - [ ] **[Range](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/range.html)**\
         Range types, such as long_range, double_range, date_range, and ip_range.
  - [ ] **[IP](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/ip.html)**\
         IPv4 and IPv6 addresses.
  - [ ] **[Version](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/version.html)** [X-Pack]\
         Software versions. Supports Semantic Versioning precedence rules.
  - [ ] **[Murmur3](https://www.elastic.co/guide/en/elasticsearch/plugins/7.12/mapper-murmur3.html)** [X-Pack]\
         Compute and stores hashes of values.

- #### Aggregate data types

  - [ ] **[Aggregate metric double](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/aggregate-metric-double.html)** [X-Pack]\
         Pre-aggregated metric values.
  - [ ] **[Histogram](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/histogram.html)** [X-Pack]\
         Pre-aggregated numerical values in the form of a histogram.

- #### Text search types

  - [ ] **[Text](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/text.html)**\
         Analyzed, unstructured text.
  - [ ] **[Annotated-text](https://www.elastic.co/guide/en/elasticsearch/plugins/7.12/mapper-annotated-text.html)**\
         Text containing special markup. Used for identifying named entities.
  - [ ] **[Completion](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/search-suggesters.html#completion-suggester)**\
         Used for auto-complete suggestions.
  - [ ] **[Search as you type](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/search-as-you-type.html)**\
         text-like type for as-you-type completion.
  - [ ] **[Token count](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/token-count.html)**\
         A count of tokens in a text.

- #### Document ranking types

  - [ ] **[Dense vector](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/dense-vector.html)** [X-Pack]\
         Records dense vectors of float values.
  - [ ] **[Sparse vector](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/sparse-vector.html)** [X-Pack] [Deprecated]\
         Records sparse vectors of float values.
  - [ ] **[Rank feature](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/rank-feature.html)**\
         Records a numeric feature to boost hits at query time.
  - [ ] **[Rank features](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/rank-features.html)**\
         Records numeric features to boost hits at query time.

- #### Spatial data types

  - [ ] **[Geo point](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/geo-point.html)**\
         Latitude and longitude points.
  - [ ] **[Geo shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/geo-shape.html)**\
         Complex shapes, such as polygons.
  - [ ] **[Point](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/point.html)**\
         Arbitrary cartesian points.
  - [ ] **[Shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/shape.html)**\
         Arbitrary cartesian geometries.

- #### Other types
  - [ ] **[Percolator](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/percolator.html)**\
         Indexes queries written in Query DSL.

### Queries

- #### [Compound queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/compound-queries.html)
  Compound queries wrap other compound or leaf queries, either to combine their results and scores, to change their behaviour, or to switch from query to filter context.
  - [x] **[Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-bool-query.html)**\
         The default query for combining multiple leaf or compound query clauses, as must, should, must_not, or filter clauses. The must and should clauses have their scores combined — the more matching clauses, the better — while the must_not and filter clauses are executed in filter context.
  - [ ] **[Boosting](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-boosting-query.html)**\
         Return documents which match a positive query, but reduce the score of documents which also match a negative query.
  - [ ] **[Constant score](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-constant-score-query.html)**\
         A query which wraps another query, but executes it in filter context. All matching documents are given the same “constant” \_score.
  - [ ] **[Disjunction max](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-dis-max-query.html)**\
         A query which accepts multiple queries, and returns any documents which match any of the query clauses. While the bool query combines the scores from all matching queries, the dis_max query uses the score of the single best- matching query clause.
  - [x] **[Function score](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-function-score-query.html)**\
         Modify the scores returned by the main query with functions to take into account factors like popularity, recency, distance, or custom algorithms implemented with scripting.
- #### [Fulltext queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/full-text-queries.html)
  The full text queries enable you to search analyzed text fields such as the body of an email. The query string is processed using the same analyzer that was applied to the field during indexing.
  - [ ] **[Intervals](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-intervals-query.html)**\
         A full text query that allows fine-grained control of the ordering and proximity of matching terms.
  - [x] **[Match](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query.html)**\
         The standard query for performing full text queries, including fuzzy matching and phrase or proximity queries.
  - [ ] **[Match bool prefix](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-bool-prefix-query.html)**\
         Creates a bool query that matches each term as a term query, except for the last term, which is matched as a prefix query
  - [ ] **[Match phrase](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query-phrase.html)**\
         Like the match query but used for matching exact phrases or word proximity matches.
  - [ ] **[Match phrase prefix](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-query-phrase-prefix.html)**\
         Like the match_phrase query, but does a wildcard search on the final word.
  - [ ] **[Multi-match](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-multi-match-query.html)**\
         The multi-field version of the match query.
  - [ ] **[Common Terms](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-common-terms-query.html) [Deprecated]**\
         A more specialized query which gives more preference to uncommon words.
  - [ ] **[Query string](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-query-string-query.html)**\
         Supports the compact Lucene query string syntax, allowing you to specify AND|OR|NOT conditions and multi-field search within a single query string. For expert users only.
  - [ ] **[Simple query string](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-simple-query-string-query.html)**\
         A simpler, more robust version of the query_string syntax suitable for exposing directly to users.
- #### [Geo queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/geo-queries.html)
  Elasticsearch supports two types of geo data: geo_point fields which support lat/lon pairs, and geo_shape fields, which support points, lines, circles, polygons, multi-polygons, etc.
  - [ ] **[Geo bounding box](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-bounding-box-query.html)**\
         Finds documents with geo-points that fall into the specified rectangle.
  - [ ] **[Geo distance](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-distance-query.html)**\
         Finds documents with geo-points within the specified distance of a central point.
  - [ ] **[Geo polygon](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-polygon-query.html) [Deprecated]**\
         Find documents with geo-points within the specified polygon.
  - [ ] **[Geo shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-geo-shape-query.html)**
    - geo-shapes which either intersect, are contained by, or do not intersect with the specified geo-shape
    - geo-points which intersect the specified geo-shape
- #### [Shape queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/shape-queries.html) [X-Pack]
  Like geo_shape Elasticsearch supports the ability to index arbitrary two dimension (non Geospatial) geometries making it possible to map out virtual worlds, sporting venues, theme parks, and CAD diagrams.
  Elasticsearch supports two types of cartesian data: point fields which support x/y pairs, and shape fields, which support points, lines, circles, polygons, multi-polygons, etc.
  - [ ] **[Shape](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-shape-query.html)**
    - shapes which either intersect, are contained by, are within or do not intersect with the specified shape
    - points which intersect the specified shape
- #### [Joining queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/joining-queries.html)
  - [ ] **[Nested](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-nested-query.html)**\
         Documents may contain fields of type nested. These fields are used to index arrays of objects, where each object can be queried (with the nested query) as an independent document.
  - [ ] **[Has child](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-has-child-query.html)**\
         A join field relationship can exist between documents within a single index. The has_child query returns parent documents whose child documents match the specified query, while the has_parent query returns child documents whose parent document matches the specified query.
  - [ ] **[Has parent](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-has-parent-query.html)**\
         Returns child documents whose joined parent document matches a provided query. You can create parent-child relationships between documents in the same index using a join field mapping.
  - [ ] **[Parent ID](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-parent-id-query.html)**\
         Returns child documents joined to a specific parent document. You can use a join field mapping to create parent-child relationships between documents in the same index.
- [x] **[Match all](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-all-query.html)**\
       The most simple query, which matches all documents, giving them all a \_score of 1.0.
- [x] **[Match none](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-match-all-query.html)**\
       This is the inverse of the match_all query, which matches no documents.
- #### [Span queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/span-queries.html)
  Span queries are low-level positional queries which provide expert control over the order and proximity of the specified terms. These are typically used to implement very specific queries on legal documents or patents.
  It is only allowed to set boost on an outer span query. Compound span queries, like span_near, only use the list of matching spans of inner span queries in order to find their own spans, which they then use to produce a score. Scores are never computed on inner span queries, which is the reason why boosts are not allowed: they only influence the way scores are computed, not spans.
  Span queries cannot be mixed with non-span queries (with the exception of the span_multi query).
  - [ ] **[Span containing](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-containing-query.html)**\
         Accepts a list of span queries, but only returns those spans which also match a second span query.
  - [ ] **[Field masking span](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-field-masking-query.html)**\
         Allows queries like span-near or span-or across different fields.
  - [ ] **[Span first](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-first-query.html)**\
         Accepts another span query whose matches must appear within the first N positions of the field.
  - [ ] **[Span multi](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-multi-term-query.html)**\
         Wraps a term, range, prefix, wildcard, regexp, or fuzzy query.
  - [ ] **[Span near](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-near-query.html)**\
         Accepts multiple span queries whose matches must be within the specified distance of each other, and possibly in the same order.
  - [ ] **[Span not](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-not-query.html)**\
         Wraps another span query, and excludes any documents which match that query.
  - [ ] **[Span or](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-or-query.html)**\
         Combines multiple span queries — returns documents which match any of the specified queries.
  - [ ] **[Span term](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-term-query.html)**\
         The equivalent of the term query but for use with other span queries.
  - [ ] **[Span within](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-span-within-query.html)**\
         The result from a single span query is returned as long is its span falls within the spans returned by a list of other span queries.
- #### [Specialized queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/specialized-queries.html)
  - [ ] **[Distance feature](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-distance-feature-query.html)**\
         A query that computes scores based on the dynamically computed distances between the origin and documents' date, date_nanos and geo_point fields. It is able to efficiently skip non-competitive hits.
  - [ ] **[More like this](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-mlt-query.html)**\
         This query finds documents which are similar to the specified text, document, or collection of documents.
  - [ ] **[Percolate](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-percolate-query.html)**\
         This query finds queries that are stored as documents that match with the specified document.
  - [ ] **[Rank feature](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-rank-feature-query.html)**\
         A query that computes scores based on the values of numeric features and is able to efficiently skip non-competitive hits.
  - [x] **[Script](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-script-query.html)**\
         This query allows a script to act as a filter. Also see the function_score query.
  - [x] **[Script score](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-script-score-query.html)**\
         A query that allows to modify the score of a sub-query with a script.
  - [ ] **[Wrapper](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-wrapper-query.html)**\
         A query that accepts other queries as json or yaml string.
  - [ ] **[Pinned](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-pinned-query.html)**\
         A query that promotes selected documents over others matching a given query.
- #### [Term-level queries](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/term-level-queries.html)
  You can use term-level queries to find documents based on precise values in structured data. Examples of structured data include date ranges, IP addresses, prices, or product IDs.
  Unlike full-text queries, term-level queries do not analyze search terms. Instead, term-level queries match the exact terms stored in a field.
  - [x] **[Exists](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-exists-query.html)**\
         Returns documents that contain any indexed value for a field.
  - [x] **[Fuzzy](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-fuzzy-query.html)**\
         Returns documents that contain terms similar to the search term. Elasticsearch measures similarity, or fuzziness, using a Levenshtein edit distance.
  - [ ] **[IDs](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-ids-query.html)**\
         Returns documents based on their document IDs.
  - [x] **[Prefix](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-prefix-query.html)**\
         Returns documents that contain a specific prefix in a provided field.
  - [x] **[Range](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-range-query.html)**\
         Returns documents that contain terms within a provided range.
  - [ ] **[Regexp](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-regexp-query.html)**\
         Returns documents that contain terms matching a regular expression.
  - [x] **[Term](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-term-query.html)**\
         Returns documents that contain an exact term in a provided field.
  - [x] **[Terms](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-terms-query.html)**\
         Returns documents that contain one or more exact terms in a provided field.
  - [ ] **[Terms set](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-terms-set-query.html)**\
         Returns documents that contain a minimum number of exact terms in a provided field. You can define the minimum number of matching terms using a field or script.
  - [ ] **[Type](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-type-query.html) [Deprecated]**\
         Returns documents of the specified type.
  - [ ] **[Wildcard](https://www.elastic.co/guide/en/elasticsearch/reference/7.12/query-dsl-wildcard-query.html)**\
         Returns documents that contain terms matching a wildcard pattern.
