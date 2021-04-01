# picker

Utilities that pair with the official Elasticsearch Go package

## Todo

### Queries

- #### [Compound queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/compound-queries.html)
  Compound queries wrap other compound or leaf queries, either to combine their results and scores, to change their behaviour, or to switch from query to filter context.
  - [x] **[Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html)**\
         The default query for combining multiple leaf or compound query clauses, as must, should, must_not, or filter clauses. The must and should clauses have their scores combined — the more matching clauses, the better — while the must_not and filter clauses are executed in filter context.
  - [ ] **[Boosting](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-boosting-query.html)**\
         Return documents which match a positive query, but reduce the score of documents which also match a negative query.
  - [ ] **[Constant score](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html)**\
         A query which wraps another query, but executes it in filter context. All matching documents are given the same “constant” \_score.
  - [ ] **[Disjunction max](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-dis-max-query.html)**\
         A query which accepts multiple queries, and returns any documents which match any of the query clauses. While the bool query combines the scores from all matching queries, the dis_max query uses the score of the single best- matching query clause.
  - [x] **[Function score](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html)**\
         Modify the scores returned by the main query with functions to take into account factors like popularity, recency, distance, or custom algorithms implemented with scripting.
- #### [Fulltext queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/full-text-queries.html)
  The full text queries enable you to search analyzed text fields such as the body of an email. The query string is processed using the same analyzer that was applied to the field during indexing.
  - [ ] **[Intervals](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-intervals-query.html)**\
         A full text query that allows fine-grained control of the ordering and proximity of matching terms.
  - [x] **[Match](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html)**\
         The standard query for performing full text queries, including fuzzy matching and phrase or proximity queries.
  - [ ] **[Match bool prefix](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-bool-prefix-query.html)**\
         Creates a bool query that matches each term as a term query, except for the last term, which is matched as a prefix query
  - [ ] **[Match phrase](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase.html)**\
         Like the match query but used for matching exact phrases or word proximity matches.
  - [ ] **[Match phrase prefix](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase-prefix.html)**\
         Like the match_phrase query, but does a wildcard search on the final word.
  - [ ] **[Multi-match](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html)**\
         The multi-field version of the match query.
  - [ ] **[Common Terms](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-common-terms-query.html) [Deprecated]**\
         A more specialized query which gives more preference to uncommon words.
  - [ ] **[Query string](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html)**\
         Supports the compact Lucene query string syntax, allowing you to specify AND|OR|NOT conditions and multi-field search within a single query string. For expert users only.
  - [ ] **[Simple query string](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-simple-query-string-query.html)**\
         A simpler, more robust version of the query_string syntax suitable for exposing directly to users.
- #### [Geo queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-queries.html)
  Elasticsearch supports two types of geo data: geo_point fields which support lat/lon pairs, and geo_shape fields, which support points, lines, circles, polygons, multi-polygons, etc.
  - [ ] **[Geo bounding box](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-geo-bounding-box-query.html)**\
         Finds documents with geo-points that fall into the specified rectangle.
  - [ ] **[Geo distance](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-geo-distance-query.html)**\
         Finds documents with geo-points within the specified distance of a central point.
  - [ ] **[Geo polygon](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-geo-polygon-query.html) [Deprecated]**\
         Find documents with geo-points within the specified polygon.
  - [ ] **[Geo shape](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-geo-shape-query.html)**
    - geo-shapes which either intersect, are contained by, or do not intersect with the specified geo-shape
    - geo-points which intersect the specified geo-shape
- #### [Shape queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/shape-queries.html) [X-Pack]
  Like geo_shape Elasticsearch supports the ability to index arbitrary two dimension (non Geospatial) geometries making it possible to map out virtual worlds, sporting venues, theme parks, and CAD diagrams.
  Elasticsearch supports two types of cartesian data: point fields which support x/y pairs, and shape fields, which support points, lines, circles, polygons, multi-polygons, etc.
  - [ ] **[Shape](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-shape-query.html)**
    - shapes which either intersect, are contained by, are within or do not intersect with the specified shape
    - points which intersect the specified shape
- #### [Joining queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/joining-queries.html)
  - [ ] **[Nested](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html)**\
         Documents may contain fields of type nested. These fields are used to index arrays of objects, where each object can be queried (with the nested query) as an independent document.
  - [ ] **[Has child](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-has-child-query.html)**\
         A join field relationship can exist between documents within a single index. The has_child query returns parent documents whose child documents match the specified query, while the has_parent query returns child documents whose parent document matches the specified query.
  - [ ] **[Has parent](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-has-parent-query.html)**\
         Returns child documents whose joined parent document matches a provided query. You can create parent-child relationships between documents in the same index using a join field mapping.
  - [ ] **[Parent ID](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-parent-id-query.html)**\
         Returns child documents joined to a specific parent document. You can use a join field mapping to create parent-child relationships between documents in the same index.
- [x] **[Match all](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html)**\
       The most simple query, which matches all documents, giving them all a \_score of 1.0.
- [x] **[Match none](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-all-query.html)**\
       This is the inverse of the match_all query, which matches no documents.
- #### [Span queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/span-queries.html)
  Span queries are low-level positional queries which provide expert control over the order and proximity of the specified terms. These are typically used to implement very specific queries on legal documents or patents.
  It is only allowed to set boost on an outer span query. Compound span queries, like span_near, only use the list of matching spans of inner span queries in order to find their own spans, which they then use to produce a score. Scores are never computed on inner span queries, which is the reason why boosts are not allowed: they only influence the way scores are computed, not spans.
  Span queries cannot be mixed with non-span queries (with the exception of the span_multi query).
  - [ ] **[Span containing](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-containing-query.html)**\
         Accepts a list of span queries, but only returns those spans which also match a second span query.
  - [ ] **[Field masking span](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-field-masking-query.html)**\
         Allows queries like span-near or span-or across different fields.
  - [ ] **[Span first](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-first-query.html)**\
         Accepts another span query whose matches must appear within the first N positions of the field.
  - [ ] **[Span multi](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-multi-term-query.html)**\
         Wraps a term, range, prefix, wildcard, regexp, or fuzzy query.
  - [ ] **[Span near](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-near-query.html)**\
         Accepts multiple span queries whose matches must be within the specified distance of each other, and possibly in the same order.
  - [ ] **[Span not](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-not-query.html)**\
         Wraps another span query, and excludes any documents which match that query.
  - [ ] **[Span or](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-or-query.html)**\
         Combines multiple span queries — returns documents which match any of the specified queries.
  - [ ] **[Span term](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-term-query.html)**\
         The equivalent of the term query but for use with other span queries.
  - [ ] **[Span within](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-span-within-query.html)**\
         The result from a single span query is returned as long is its span falls within the spans returned by a list of other span queries.
- #### [Specialized queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/specialized-queries.html)
  - [ ] **[Distance feature](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-distance-feature-query.html)**\
         A query that computes scores based on the dynamically computed distances between the origin and documents' date, date_nanos and geo_point fields. It is able to efficiently skip non-competitive hits.
  - [ ] **[More like this](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-mlt-query.html)**\
         This query finds documents which are similar to the specified text, document, or collection of documents.
  - [ ] **[Percolate](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-percolate-query.html)**\
         This query finds queries that are stored as documents that match with the specified document.
  - [ ] **[Rank feature](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-rank-feature-query.html)**\
         A query that computes scores based on the values of numeric features and is able to efficiently skip non-competitive hits.
  - [x] **[Script](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-query.html)**\
         This query allows a script to act as a filter. Also see the function_score query.
  - [x] **[Script score](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-script-score-query.html)**\
         A query that allows to modify the score of a sub-query with a script.
  - [ ] **[Wrapper](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wrapper-query.html)**\
         A query that accepts other queries as json or yaml string.
  - [ ] **[Pinned](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-pinned-query.html)**\
         A query that promotes selected documents over others matching a given query.
- #### [Term-level queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/term-level-queries.html)
  You can use term-level queries to find documents based on precise values in structured data. Examples of structured data include date ranges, IP addresses, prices, or product IDs.
  Unlike full-text queries, term-level queries do not analyze search terms. Instead, term-level queries match the exact terms stored in a field.
  - [x] **[Exists]()**\
        Returns documents that contain any indexed value for a field.
  - [x] **[Fuzzy]()**\
        Returns documents that contain terms similar to the search term. Elasticsearch measures similarity, or fuzziness, using a Levenshtein edit distance.
  - [ ] **[Ids]()**\
        Returns documents based on their document IDs.
  - [x] **[Prefix]()**\
         Returns documents that contain a specific prefix in a provided field.
  - [ ] **[Range]()**\
        Returns documents that contain terms within a provided range.
  - [ ] **[Regexp]()**\
        Returns documents that contain terms matching a regular expression.
  - [x] **[Term]()**\
        Returns documents that contain an exact term in a provided field.
  - [x] **[Terms]()**\
        Returns documents that contain one or more exact terms in a provided field.
  - [ ] **[Terms set]()**\
        Returns documents that contain a minimum number of exact terms in a provided field. You can define the minimum number of matching terms using a field or script.
  - [ ] **[Type]()**\
        Returns documents of the specified type.
  - [ ] **[Wildcard]()**\
        Returns documents that contain terms matching a wildcard pattern.
