package picker

import (
	"bytes"
	"encoding/json"
)

type Conflicts string

const (
	ConflictsNotSpecified = ""
	ConflictsProceed      = "proceed"
	ConflictsAbort        = "abort"
)

var DefaultConflicts = ConflictsAbort

type UpdateByQuerier interface {
	UpdateByQuery() (*UpdateByQuery, error)
}
type UpdateByQueryParams struct {
	Query     Querier
	Script    *Script
	Conflicts Conflicts
}

func (p UpdateByQueryParams) UpdateByQuery() (*UpdateByQuery, error) {
	u := &UpdateByQuery{}
	if p.Query == nil {
		return u, nil
	}
	q, err := p.Query.Query()
	if err != nil {
		return u, err
	}
	u.query = q
	return u, nil
}

type UpdateByQuery struct {
	query     *Query
	script    *Script
	conflicts Conflicts
}

func (u UpdateByQuery) Encode() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(u)
	return buf, err
}

func (u UpdateByQuery) MarshalJSON() ([]byte, error) {
	return updateByQuery{
		Query:     u.query,
		Script:    u.script,
		Conflicts: u.conflicts,
	}.MarshalJSON()
}

func (u *UpdateByQuery) UnmarshalJSON(data []byte) error {
	*u = UpdateByQuery{}
	p := updateByQuery{}
	err := p.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	u.query = p.Query
	u.conflicts = p.Conflicts
	u.script = p.Script
	return nil
}

func (u UpdateByQuery) MarshalBSON() ([]byte, error) {
	return u.MarshalJSON()
}

func (u *UpdateByQuery) UnmarshalBSON(data []byte) error {
	return u.UnmarshalJSON(data)
}

//easyjson:json
type updateByQuery struct {
	Query     *Query    `json:"query,omitempty"`
	Script    *Script   `json:"script,omitempty"`
	Conflicts Conflicts `json:"conflicts,omitempty"`
}

func (u *UpdateByQuery) Conflicts() Conflicts {
	return u.conflicts
}

func (u *UpdateByQuery) SetConflicts(c Conflicts) {
	u.conflicts = c
}

func (u *UpdateByQuery) Script() *Script {
	return u.script
}

func (u *UpdateByQuery) SetScript(script *Script) {
	u.script = script
}
func (u *UpdateByQuery) SetQuery(query Querier) error {
	if query == nil {
		u.query = nil
		return nil
	}
	q, err := query.Query()
	if err != nil {
		return err
	}
	u.query = q
	return nil
}

func (u *UpdateByQuery) Query() *Query {
	if u.query == nil {
		u.query = &Query{}
	}
	return u.query
}

func NewUpdateByQuery(params UpdateByQueryParams) (*UpdateByQuery, error) {
	return params.UpdateByQuery()
}

type UpdateByQueryResponse struct {
	Took                 int64         `json:"took"`
	TimedOut             bool          `json:"timed_out"`
	Total                int64         `json:"total"`
	Updated              int64         `json:"updated"`
	Deleted              int64         `json:"deleted"`
	Batches              int64         `json:"batches"`
	VersionConflicts     int64         `json:"version_conflicts"`
	Noops                int64         `json:"noops"`
	Retries              Retries       `json:"retries"`
	ThrottledMillis      int64         `json:"throttled_millis"`
	RequestsPerSecond    float64       `json:"requests_per_second"`
	Failures             []interface{} `json:"failures"`
	ThrottledUntilMillis int64         `json:"throttled_until_millis"`
}

type Retries struct {
	Bulk   int64 `json:"bulk"`
	Search int64 `json:"search"`
}
