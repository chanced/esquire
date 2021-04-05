package picker

type IndexPrefixes struct {
	MinimumChars interface{} `json:"min_chars,omitempty"`
	MaximumChars interface{} `json:"max_chars,omitempty"`
}

// WithIndexPrefixes is a mapping with the index_prefixes parameter
//
// The index_prefixes parameter enables the indexing of term prefixes to speed up
// prefix searches. It accepts the following optional settings
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-prefixes.html
type WithIndexPrefixes interface {
	IndexPrefixes() *IndexPrefixes
	SetIndexPrefixes(v *IndexPrefixes) error
}

// indexPrefixesParams is a mixin that adds the index_prefixes param
//
// The index_prefixes parameter enables the indexing of term prefixes to speed
// up prefix searches. It accepts the following optional settings
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-prefixes.html
type indexPrefixesParams struct {
	indexPrefixes *IndexPrefixes
}

func (ip indexPrefixesParams) IndexPrefixes() *IndexPrefixes {
	return ip.indexPrefixes
}
func (ip *indexPrefixesParams) SetIndexPrefixes(v *IndexPrefixes) error {
	// TODO: validate
	ip.indexPrefixes = v
	return nil
}
