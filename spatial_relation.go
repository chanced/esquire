package picker

const DefaultSpatialRelation = SpatialRelationIntersects

type SpatialRelation string

const (

	// SpatialRelationIntersects - (default) Return all documents whose
	// geo_shape or geo_point field intersects the query geometry.
	SpatialRelationIntersects SpatialRelation = "INTERSECTS"
	// SpatialRelationDisjoint - Return all documents whose geo_shape or
	// geo_point field has nothing in common with the query geometry.
	SpatialRelationDisjoint SpatialRelation = "DISJOINT"
	// SpatialRelationWithin  - Return all documents whose geo_shape or
	// geo_point field is within the query geometry. Line geometries are not
	// supported.
	SpatialRelationWithin SpatialRelation = "WITHIN"
	// SpatialRelationContains - Return all documents whose geo_shape or
	// geo_point field contains the query geometry.
	SpatialRelationContains SpatialRelation = "CONTAINS"
)
