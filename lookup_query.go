package picker

import (
	"encoding/json"
)

type Lookup struct {
	Field string
	// Name of the index from which to fetch field
	// values.(Required)
	Index string

	// ID of the document from which to fetch field
	// values. (Required)
	ID string

	// Path of the field from which to fetch field values. Elasticsearch
	// uses these values as search terms for the query.(Required)
	//
	// If the field values include an array of nested inner objects, you can
	// access those objects using dot notation syntax.
	Path string

	// Routing value of the document from which to fetch term values. If a
	// custom routing value was provided when the document was indexed, this
	// parameter is required. (Optional)
	Routing string

	Boost interface{}

	CaseInsensitive bool

	QueryName string
}

func (l Lookup) Name() string {
	return l.QueryName
}

func (l Lookup) Clause() (QueryClause, error) {
	return l.Terms()
}
func (l Lookup) Terms() (*TermsQuery, error) {
	v := LookupValues{}
	err := v.SetID(l.ID)
	if err != nil {
		return nil, err
	}
	err = v.SetIndex(l.Index)
	if err != nil {
		return nil, err
	}
	err = v.SetPath(l.Path)
	if err != nil {
		return nil, err
	}

	q := &TermsQuery{}

	err = q.setLookup(v)
	if err != nil {
		return nil, err
	}
	err = q.SetBoost(l.Boost)
	if err != nil {
		return nil, err
	}

	q.SetCaseInsensitive(l.CaseInsensitive)
	return q, nil
}

func (l Lookup) Kind() QueryKind {
	return KindTerms
}

type LookupValues struct {
	id      string
	index   string
	path    string
	routing string
}

func (l LookupValues) ID() string {
	return l.id
}
func (l *LookupValues) SetID(id string) error {
	if len(id) == 0 {
		return ErrIDRequired
	}
	l.id = id
	return nil
}
func (l *LookupValues) Index() string {
	return l.index
}
func (l *LookupValues) SetIndex(index string) error {
	if len(index) == 0 {
		return ErrIndexRequired
	}
	l.index = index
	return nil
}

func (l *LookupValues) Path() string {
	return l.path
}
func (l *LookupValues) SetPath(path string) error {
	if len(path) == 0 {
		return ErrPathRequired
	}
	l.path = path
	return nil
}
func (l LookupValues) Routing() string {
	return l.routing
}
func (l LookupValues) SetRouting(routing string) {
	l.routing = routing
}

func (l LookupValues) IsEmpty() bool {
	return len(l.id) == 0 && len(l.index) == 0 && len(l.path) == 0 && len(l.routing) == 0
}

func (l *LookupValues) UnmarshalJSON(data []byte) error {
	*l = LookupValues{}
	var m map[string]string
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	l.id = m["id"]
	l.index = m["index"]
	l.path = m["path"]
	l.routing = m["routing"]
	return nil
}
func (l LookupValues) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"id":      l.id,
		"index":   l.index,
		"path":    l.path,
		"routing": l.routing,
	}
	return json.Marshal(m)
}

type TermsLookup = LookupValues
