# picker

Utilities that pair with the official Elasticsearch Go package

- [Compound queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/compound-queries.html)
  - [x] [Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html)
  - [ ] [Boosting](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-boosting-query.html)
  - [ ] [Constant score](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-constant-score-query.html)
  - [ ] [Disjunction max](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-dis-max-query.html)
  - [x] [Function score](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html)
- [Fulltext queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/full-text-queries.html)
  - **[ ] [Intervals](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-intervals-query.html)**\
    A full text query that allows fine-grained control of the ordering and proximity of matching terms.
  - **[x] [Match](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html)**\
    The standard query for performing full text queries, including fuzzy matching and phrase or proximity queries.
  - [ ] [Match bool prefix](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-bool-prefix-query.html)\
        Creates a bool query that matches each term as a term query, except for the last term, which is matched as a prefix query
  - [ ] [Match phrase](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase.html)
  - [ ] [Match phrase prefix](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query-phrase-prefix.html)
  - [ ] [Multi-match](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-multi-match-query.html)
  - [ ] [Common Terms](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-common-terms-query.html) [Deprecated]
  - [ ] [Query string](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html)
  - [ ] [Simple query string](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-simple-query-string-query.html)
- [Geo queries](https://www.elastic.co/guide/en/elasticsearch/reference/current/geo-queries.html)
  geo_bounding_box query
  Finds documents with geo-points that fall into the specified rectangle.
  geo_distance query
  Finds documents with geo-points within the specified distance of a central point.
  geo_polygon query
  Find documents with geo-points within the specified polygon.
  geo_shape query
  Finds documents with:

geo-shapes which either intersect, are contained by, or do not intersect with the specified geo-shape
geo-points which intersect the specified geo-shape
