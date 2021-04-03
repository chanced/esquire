package picker

// WithCopyTo is a Field mapping with a CopyTo param
//
// The copy_to parameter allows you to copy the values of multiple fields into
// a group field, which can then be queried as a single field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/copy-to.html
type WithCopyTo interface {
	CopyTo() string
	SetCopyTo(v string)
}

// CopyToParam is a Field mixin for CopyTo
//
// The copy_to parameter allows you to copy the values of multiple fields into
// a group field, which can then be queried as a single field.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/copy-to.html
type CopyToParam struct {
	CopyToValue string `bson:"copy_to,omitempty" json:"copy_to,omitempty"`
}

// CopyTo parameter allows you to copy the values of multiple fields into a group
// field, which can then be queried as a single field.
func (ctp CopyToParam) CopyTo() string {

	return ctp.CopyToValue
}

// SetCopyTo sets CopyToParam.Value to v
func (ctp *CopyToParam) SetCopyTo(v string) {
	ctp.CopyToValue = v
}
