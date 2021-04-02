package picker

import "encoding/json"

type JoinFieldParams struct {
	// Relations defines a set of possible relations within the documents, each
	// relation being a parent name and a child name.
	//
	// Use like a map[string]string:
	//  Relations{"myField": []string{"myValue"}}
	Relations Relations `json:"relations,omitempty"`
	// To support aggregations and other operations that require looking up field
	// values on a per-document basis, Elasticsearch uses a data structure called
	// doc values. Term-based field types such as keyword store their doc values
	// using an ordinal mapping for a more compact representation. This mapping
	// works by assigning each term an incremental integer or ordinal based on its
	// lexicographic order. The field’s doc values store only the ordinals for each
	// document instead of the original terms, with a separate lookup structure to
	// convert between ordinals and terms.
	//
	// When used during aggregations, ordinals can greatly improve performance. As
	// an example, the terms aggregation relies only on ordinals to collect
	// documents into buckets at the shard-level, then converts the ordinals back to
	// their original term values when combining results across shards.
	//
	// Each index segment defines its own ordinal mapping, but aggregations collect
	// data across an entire shard. So to be able to use ordinals for shard-level
	// operations like aggregations, Elasticsearch creates a unified mapping called
	// global ordinals. The global ordinal mapping is built on top of segment
	// ordinals, and works by maintaining a map from global ordinal to the local
	// ordinal for each segment.
	//
	// Global ordinals are used if a search contains any of the following
	// components:
	//
	// - Certain bucket aggregations on keyword, ip, and flattened fields. This
	// includes terms aggregations as mentioned above, as well as composite,
	// diversified_sampler, and significant_terms.
	//
	// - Bucket aggregations on text fields that require fielddata to be enabled.
	//
	// - Operations on parent and child documents from a join field, including
	// has_child queries and parent aggregations.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eager-global-ordinals.html
	EagerGlobalOrdinals interface{} `json:"eager_global_ordinals"`
}

func (JoinFieldParams) Type() FieldType {
	return FieldTypeJoin
}

func (p JoinFieldParams) Field() (Field, error) {
	return p.Join()
}

func (p JoinFieldParams) Join() (*JoinField, error) {
	f := &JoinField{}
	f.SetRelations(p.Relations)
	err := f.SetEagerGlobalOrdinals(p.EagerGlobalOrdinals)
	if err != nil {
		return f, err
	}
	return f, nil
}

// A JoinField is a special field that creates parent/child relation within
// documents of the same index. The relations section defines a set of possible
// relations within the documents, each relation being a parent name and a child
// name.
//
// To index a document with a join, the name of the relation and the optional
// parent of the document must be provided in the source
//
// When indexing parent documents, you can choose to specify just the name of
// the relation as a shortcut instead of encapsulating it in the normal object
// notation
//
// When indexing a child, the name of the relation as well as the parent id of
// the document must be added in the _source.
//
// WARNING
//
// It is required to index the lineage of a parent in the same shard so you must
// always route child documents using their greater parent id.
//
// Parent-join and performance
//
// The join field shouldn’t be used like joins in a relation database. In
// Elasticsearch the key to good performance is to de-normalize your data into
// documents. Each join field, has_child or has_parent query adds a significant
// tax to your query performance. It can also trigger global ordinals to be
// built.
//
// The only case where the join field makes sense is if your data contains a
// one-to-many relationship where one entity significantly outnumbers the other
// entity. An example of such case is a use case with products and offers for
// these products. In the case that offers significantly outnumbers the number
// of products then it makes sense to model the product as parent document and
// the offer as child document.
//
// Parent-join restrictions
//
// Only one join field mapping is allowed per index.
//
// - Parent and child documents must be indexed on the same shard. This means
// that the same routing value needs to be provided when getting, deleting, or
// updating a child document.
//
// - An element can have multiple children but only one parent.
//
// - It is possible to add a new relation to an existing join field.
//
// - It is also possible to add a child to an existing element but only if the
// element is already a parent.
//
// Searching with parent-join
//
// The parent-join creates one field to index the name of the relation within
// the document (my_parent, my_child, …​).
//
// It also creates one field per parent/child relation. The name of this field
// is the name of the join field followed by # and the name of the parent in the
// relation. So for instance for the my_parent → [my_child, another_child]
// relation, the join field creates an additional field named
// my_join_field#my_parent.
//
// This field contains the parent _id that the document links to if the document
// is a child (my_child or another_child) and the _id of document if it’s a
// parent (my_parent).
//
// When searching an index that contains a join field, these two fields are
// always returned in the search response
//
// Parent-join queries and aggregations
//
// See the has_child and has_parent queries, the children aggregation, and inner
// hits for more information.
//
// The value of the join field is accessible in aggregations and scripts, and
// may be queried with the parent_id query
//
// Global ordinals
//
// The join field uses global ordinals to speed up joins. Global ordinals need
// to be rebuilt after any change to a shard. The more parent id values are
// stored in a shard, the longer it takes to rebuild the global ordinals for the
// join field.
//
// Global ordinals, by default, are built eagerly: if the index has changed,
// global ordinals for the join field will be rebuilt as part of the refresh.
// This can add significant time to the refresh. However most of the times this
// is the right trade-off, otherwise global ordinals are rebuilt when the first
// parent-join query or aggregation is used. This can introduce a significant
// latency spike for your users and usually this is worse as multiple global
// ordinals for the join field may be attempt rebuilt within a single refresh
// interval when many writes are occurring.
//
// When the join field is used infrequently and writes occur frequently it may
// make sense to disable eager loading:
//
// Multiple children per parent
//
// It is also possible to define multiple children for a single parent
//
// Using multiple levels of relations to replicate a relational model is not
// recommended. Each level of relation adds an overhead at query time in terms
// of memory and computation. You should de-normalize your data if you care
// about performance.
//
//
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/parent-join.html
type JoinField struct {
	eagerGlobalOrdinalsParam
	relationsParam
}

func (JoinField) Type() FieldType {
	return FieldTypeJoin
}

func (j *JoinField) UnmarshalJSON(data []byte) error {
	var p JoinFieldParams
	err := json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	f, err := p.Join()
	if err != nil {
		return err
	}
	*j = *f
	return nil
}

func (j *JoinField) MarshalJSON() ([]byte, error) {
	return json.Marshal(JoinFieldParams{
		Relations:           j.relations,
		EagerGlobalOrdinals: j.eagerGlobalOrdinals.Value(),
	})
}

func NewJoinField(params JoinFieldParams) (*JoinField, error) {
	return params.Join()
}
