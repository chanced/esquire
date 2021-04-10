package picker

type Percolater interface {
	Percolate() (*PercolateQuery, error)
}

type PercolateDocumentQueryParams struct {
	// The field of type percolator that holds the indexed queries. This is a
	// required parameter.
	Field string
	// The suffix to be used for the _percolator_document_slot field in case
	// multiple percolate queries have been specified. This is an optional
	// parameter.
	Name     string
	Document interface{}
	// The type / mapping of the document being percolated.
	//
	// Deprecated and will be removed in Elasticsearch 8.0.
	DocumentType string
	completeClause
}

func (PercolateDocumentQueryParams) Kind() QueryKind {
	return QueryKindPercolate
}
func (p PercolateDocumentQueryParams) Clause() (QueryClause, error) {
	return p.Percolate()
}
func (p PercolateDocumentQueryParams) Percolate() (*PercolateQuery, error) {
	q := &PercolateQuery{}
	err := q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	err = q.SetDocument(p.Document)
	if err != nil {
		return q, err
	}

	q.SetDocumentType(p.DocumentType)
	q.SetName(p.Name)
	return q, nil
}

type PercolateDocumentsQueryParams struct {
	// The field of type percolator that holds the indexed queries. This is a
	// required parameter.
	Field string
	// The suffix to be used for the _percolator_document_slot field in case
	// multiple percolate queries have been specified. This is an optional
	// parameter.
	Name string
	// Like the document parameter, but accepts multiple documents via a json
	// array.
	Documents interface{}
	// The type / mapping of the document being percolated.
	//
	// Deprecated and will be removed in Elasticsearch 8.0.
	DocumentType string
	completeClause
}

func (PercolateDocumentsQueryParams) Kind() QueryKind {
	return QueryKindPercolate
}

func (p PercolateDocumentsQueryParams) Clause() (QueryClause, error) {
	return p.Percolate()
}
func (p PercolateDocumentsQueryParams) Percolate() (*PercolateQuery, error) {
	q := &PercolateQuery{}
	err := q.SetField(p.Field)
	if err != nil {
		return q, err
	}
	err = q.SetDocuments(p.Documents)
	if err != nil {
		return q, err
	}
	q.SetDocumentType(p.DocumentType)
	q.SetName(p.Name)
	return q, nil
	// return q, nil
}

type PercolateStoredDocumentQuery struct {
	// The suffix to be used for the _percolator_document_slot field in case
	// multiple percolate queries have been specified. This is an optional
	// parameter.
	Name string
	// The field of type percolator that holds the indexed queries. This is a required parameter.
	Field string
	// The id of the document to fetch. This is a required parameter.
	ID string
	// The index the document resides in. This is a required parameter.
	Index string
	// Optionally, routing to be used to fetch document to percolate.
	Routing string
	// The type of the document to fetch. This parameter is deprecated and will
	// be removed in Elasticsearch 8.0.
	Type string
	// Optionally, preference to be used to fetch document to percolate.
	Preference string
	// Optionally, the expected version of the document to be fetched.
	Version int
	completeClause
}

func (PercolateStoredDocumentQuery) Kind() QueryKind {
	return QueryKindPercolate
}

func (p PercolateStoredDocumentQuery) Clause() (QueryClause, error) {
	return p.Percolate()
}
func (p PercolateStoredDocumentQuery) Percolate() (*PercolateQuery, error) {
	q := &PercolateQuery{
		id:         p.ID,
		index:      p.Index,
		routing:    p.Routing,
		typ:        p.Type,
		version:    p.Version,
		preference: p.Preference,
		nameParam:  nameParam{name: p.Name},
	}
	err := q.SetField(p.Field)
	return q, err

	// return q, nil
}

type PercolateQuery struct {
	nameParam
	fieldParam
	document     interface{}
	documents    interface{}
	documentType string
	id           string
	index        string
	routing      string
	typ          string
	preference   string
	version      int
	completeClause
}

func (q PercolateQuery) DocumentType() string {
	return q.documentType
}
func (q *PercolateQuery) SetDocumentType(typ string) {
	q.documentType = typ
}
func (q PercolateQuery) ID() string {
	return q.id
}
func (q *PercolateQuery) SetID(id string) {
	q.id = id
}
func (q PercolateQuery) Preference() string {
	return q.preference
}
func (q *PercolateQuery) SetPreference(pref string) {
	q.preference = pref
}
func (q PercolateQuery) Index() string {
	return q.index
}
func (q *PercolateQuery) SetIndex(index string) {
	q.index = index
}
func (q PercolateQuery) Routing() string {
	return q.routing
}
func (q *PercolateQuery) SetRouting(routing string) {
	q.routing = routing
}
func (q PercolateQuery) Type() string {
	return q.typ
}
func (q *PercolateQuery) SetType(typ string) {
	q.typ = typ
}
func (q PercolateQuery) Version() int {
	return q.version
}
func (q *PercolateQuery) SetVersion(vers int) {
	q.version = vers
}

func (q PercolateQuery) Document() interface{} {
	return q.document
}
func (q *PercolateQuery) SetDocument(doc interface{}) error {
	q.document = doc
	return nil
}
func (q PercolateQuery) Documents() interface{} {
	return q.documents
}
func (q *PercolateQuery) SetDocuments(docs interface{}) error {
	q.documents = docs
	return nil
}

func (PercolateQuery) Kind() QueryKind {
	return QueryKindPercolate
}
func (q *PercolateQuery) Clause() (QueryClause, error) {
	return q, nil
}
func (q *PercolateQuery) Percolate() (*PercolateQuery, error) {
	return q, nil
}
func (q *PercolateQuery) Clear() {
	if q == nil {
		return
	}
	*q = PercolateQuery{}
}
func (q *PercolateQuery) UnmarshalJSON(data []byte) error {
	*q = PercolateQuery{}
	p := percolateQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*q = PercolateQuery{
		nameParam:    nameParam{name: p.Name},
		document:     p.Document,
		documents:    p.Documents,
		documentType: p.DocumentType,
		id:           p.ID,
		index:        p.Index,
		routing:      p.Routing,
		typ:          p.Type,
		preference:   p.Preference,
		version:      p.Version,
		fieldParam:   fieldParam{field: p.Field},
	}
	return nil
}
func (q PercolateQuery) MarshalJSON() ([]byte, error) {
	return percolateQuery{
		Name:         q.name,
		Field:        q.field,
		UName:        q.name,
		Document:     q.document,
		Documents:    q.documents,
		DocumentType: q.documentType,
		ID:           q.id,
		Index:        q.index,
		Routing:      q.routing,
		Type:         q.typ,
		Preference:   q.preference,
		Version:      q.version,
	}.MarshalJSON()
}
func (q *PercolateQuery) IsEmpty() bool {
	switch {
	case q == nil:
		return true
	case q.document != nil:
		return false
	case q.documents != nil:
		return false
	case len(q.id) > 0 && len(q.index) > 0:
		return false
	default:
		return true
	}
}

//easyjson:json
type percolateQuery struct {
	Field        string      `json:"field"`
	Name         string      `json:"name,omitempty"`
	UName        string      `json:"_name,omitempty"`
	Document     interface{} `json:"document,omitempty"`
	Documents    interface{} `json:"documents,omitempty"`
	DocumentType string      `json:"document_type,omitempty"`
	ID           string      `json:"id,omitempty"`
	Index        string      `json:"index,omitempty"`
	Routing      string      `json:"routing,omitempty"`
	Type         string      `json:"type,omitempty"`
	Preference   string      `json:"preference,omitempty"`
	Version      int         `json:"version,omitempty"`
}
