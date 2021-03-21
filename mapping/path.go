package mapping

type WithPath interface {
	// Path is the path for an alias
	Path() string
	// SetPath sets the Path Value to v
	SetPath(v string)
}

type FieldWithPath interface {
	Field
	WithPath
}

type PathParam struct {
	PathValue string `bson:"path,omitempty" json:"path,omitempty"`
}

// Path is the path for an alias
func (p PathParam) Path() string {
	return p.PathValue
}

// SetPath sets the Path Value to v
func (p *PathParam) SetPath(v string) {
	if p.Path() != v {
		p.PathValue = v
	}
}
