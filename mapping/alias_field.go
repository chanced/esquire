package mapping

type Aliaser interface {
	Alias() (*AliasField, error)
}

type AliasFieldParams struct {
	Path string
}

func (a AliasFieldParams) Alias() (*AliasField, error) {
	f := &AliasField{}
	err := f.SetPath(a.Path)
	return f, err

}

func (a AliasFieldParams) Field() (Field, error) {
	return a.Alias()
}

// An AliasField mapping defines an alternate name for a field in the index. The
// alias can be used in place of the target field in search requests, and
// selected other APIs like field capabilities.
//
// Alias targets
//
// There are a few restrictions on the target of an alias:
//
// - The target must be a concrete field, and not an object or another field
// alias.
//
// - The target field must exist at the time the alias is created.
//
// - If nested objects are defined, a field alias must have the same nested
// scope as its target.
//
// - Additionally, a field alias can only have one target. This means that it is
// not possible to use a field alias to query over multiple target fields in a
// single clause.
//
// An alias can be changed to refer to a new target through a mappings update. A
// known limitation is that if any stored percolator queries contain the field
// alias, they will still refer to its original target. More information can be
// found in the percolator documentation.
//
// Currently only the search and field capabilities APIs will accept and resolve
// field aliases. Other APIs that accept field names, such as term vectors,
// cannot be used with field aliases.
//
// Finally, some queries, such as terms, geo_shape, and more_like_this, allow
// for fetching query information from an indexed document. Because field
// aliases aren’t supported when fetching documents, the part of the query that
// specifies the lookup path cannot refer to a field by its alias.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/alias.html
type AliasField struct {
	path string
}

// Path is the path for an alias
func (a AliasField) Path() string {
	return a.path
}

func (AliasField) Type() FieldType {
	return FieldTypeAlias
}

// SetPath sets path to path
func (p *AliasField) SetPath(path string) error {
	if len(path) == 0 {
		return ErrPathRequired
	}
	p.path = path
	return nil
}

func (AliasFieldParams) Kind()

func NewAliasField(params AliasFieldParams) (*AliasField, error) {
	return params.Alias()
}
