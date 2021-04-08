
package picker

type TypeQueryParams struct {
  Name string
completeClause
}
func (TypeQueryParams) Kind() QueryKind {
return QueryKindType
}

func (p TypeQueryParams) Clause() (QueryClause, error) {
  return p.Type()
}
func (p TypeQueryParams) Type() (*TypeQuery, error) {
  q := &TypeQuery{}
  _=q
  panic("not implemented")
  // return q, nil
}
type TypeQuery struct {
  nameParam
  completeClause
}
func (TypeQuery) Kind() QueryKind {
return QueryKindType
}
func (q *TypeQuery) Clause() (QueryClause, error) {
return q, nil
}
func (q *TypeQuery) Type() (QueryClause, error) {
return q, nil
}
func (q *TypeQuery) Clear() {
if q == nil {
return 
}
*q = TypeQuery{}
}
func (q *TypeQuery) UnmarshalJSON(data []byte) error {
panic("not implemented")
}
func (q TypeQuery) MarshalJSON() ([]byte, error) {
panic("not implemented")
}
func (q *TypeQuery) IsEmpty() bool {
panic("not implemented")
}

type typeQuery struct{
Name string `json:"_name,omitempty"`
}

