package picker

type IndexPrefixes struct {
	MinimumChars *uint `bson:"min_chars,omitempty" json:"min_chars,omitempty"`
	MaximumChars *uint `bson:"max_chars,omitempty" json:"max_chars,omitempty"`
}

func (ip *IndexPrefixes) Clone() *IndexPrefixes {
	if ip == nil {
		return nil
	}
	v := IndexPrefixes{}
	if ip.MaximumChars != nil {
		max := *ip.MaximumChars
		v.MaximumChars = &max
	}
	if ip.MinimumChars != nil {
		min := *ip.MinimumChars
		v.MinimumChars = &min
	}
	return &v
}

// WithIndexPrefixes is a mapping with the index_prefixes parameter
//
// The index_prefixes parameter enables the indexing of term prefixes to speed up
// prefix searches. It accepts the following optional settings
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-prefixes.html
type WithIndexPrefixes interface {
	// IndexPrefixesMinChars is the minimum prefix length to index. Must be
	// greater than 0, and defaults to 2. The value is inclusive.
	IndexPrefixesMinChars() uint
	// SetIndexPrefixesMinChars sets the MinChars value to v
	SetIndexPrefixesMinChars(v uint) error
	// IndexPrefixesMaxChars is the maximum prefix length to index. Must be less
	// than 20, and defaults to 5. The value is inclusive.
	IndexPrefixesMaxChars() uint
	// SetIndexPrefixesMaxChars sets the maximum
	SetIndexPrefixesMaxChars(v uint) error
}

// FieldWithIndexPrefixes is a Field with index_prefix parameter
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-prefixes.html
type FieldWithIndexPrefixes interface {
	Field
	WithIndexPrefixes
}

// IndexPrefixesParams is a mixin that adds the index_prefixes param
//
// The index_prefixes parameter enables the indexing of term prefixes to speed
// up prefix searches. It accepts the following optional settings
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-prefixes.html
type IndexPrefixesParams struct {
	IndexPrefixesValue *IndexPrefixes `json:"index_prefixes,omitempty"`
}

// IndexPrefixesMinChars is the minimum prefix length to index. Must be greater
// than 0, and defaults to 2. The value is inclusive.
func (ip IndexPrefixesParams) IndexPrefixesMinChars() uint {
	if ip.IndexPrefixesValue == nil || ip.IndexPrefixesValue.MinimumChars == nil {
		return 2
	}
	return *ip.IndexPrefixesValue.MinimumChars
}

// SetIndexPrefixesMinChars sets the minimum prefix length to index. Must be
// greater than 0, and defaults to 2. The value is inclusive.
func (ip *IndexPrefixesParams) SetIndexPrefixesMinChars(v uint) error {
	if v == 0 {
		return ErrInvalidIndexPrefixMinChars
	}
	if ip.IndexPrefixesValue == nil {
		ip.IndexPrefixesValue = &IndexPrefixes{}
	}
	if ip.IndexPrefixesMinChars() != v {
		ip.IndexPrefixesValue.MinimumChars = &v
	}
	return nil
}

// IndexPrefixesMaxChars is the maximum prefix length to index. Must be less
// than 20, and defaults to 5. The value is inclusive.
func (ip IndexPrefixesParams) IndexPrefixesMaxChars() uint {
	if ip.IndexPrefixesValue == nil || ip.IndexPrefixesValue.MaximumChars == nil {
		return 5
	}
	return *ip.IndexPrefixesValue.MaximumChars
}

// SetIndexPrefixesMaxChars sets the maximum prefix length to index. Must be
// less than 20, and defaults to 5. The value is inclusive.
func (ip *IndexPrefixesParams) SetIndexPrefixesMaxChars(v uint) error {

	if v >= 20 {
		return ErrInvalidIndexPrefixMaxChars
	}
	if v == 0 {
		return ErrInvalidIndexPrefixMaxChars
	}
	if ip.IndexPrefixesMaxChars() == v {
		return nil
	}

	if ip.IndexPrefixesValue == nil {
		ip.IndexPrefixesValue = &IndexPrefixes{}
	}
	ip.IndexPrefixesValue.MaximumChars = &v
	return nil
}
