package mapping

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
//      - Parent and child documents must be indexed on the same shard. This means
// that the same routing value needs to be provided when getting, deleting, or
// updating a child document.
//
//      - An element can have multiple children but only one parent.
//
//      - It is possible to add a new relation to an existing join field.
//
//      - It is also possible to add a child to an existing element but only if the
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
	BaseField                `bson:",inline" json:",inline"`
	EagerGlobalOrdinalsParam `bson:",inline" json:",inline"`
	RelationsParam           `bson:",inline" json:",inline"`
}

func (f JoinField) Clone() Field {
	n := NewJoinField()
	n.SetEagerGlobalOrdinals(f.EagerGlobalOrdinals())
	n.SetRelations(f.Relations().Clone())
	return n
}

func NewJoinField() *JoinField {
	return &JoinField{BaseField: BaseField{MappingType: TypeJoin}}
}
