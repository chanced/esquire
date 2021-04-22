package picker

import "time"

// PointInTime is a lightweight view into the state of the data as it existed
// when initiated.
//
// ! X-Pack
//
// A search request by default executes against the most recent visible data of
// the target indices, which is called point in time. Elasticsearch pit (point
// in time) is a lightweight view into the state of the data as it existed when
// initiated. In some cases, it’s preferred to perform multiple search requests
// using the same point in time. For example, if refreshes happen between
// search_after requests, then the results of those requests might not be
// consistent as changes happening between searches are only visible to the more
// recent point in time.
//
// Prerequisites
//
// - If the Elasticsearch security features are enabled, you must have the read
// index privilege for the target data stream, index, or index alias.
//
// - To search a point in time (PIT) for an index alias, you must have the read
// index privilege for the alias’s concrete indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
//easyjson:json
type PointInTime struct {
	ID        string     `bson:"id" json:"id"`
	KeepAlive *time.Time `bson:"keep_alive,omitempty" json:"keep_alive,omitempty"`
}

func (pit *PointInTime) Clone() *PointInTime {
	if pit == nil {
		return nil
	}
	n := PointInTime{}
	if pit.ID != "" {
		n.ID = pit.ID
	}
	if pit.KeepAlive != nil {
		t := *pit.KeepAlive
		n.KeepAlive = &t
	}
	return &n
}
