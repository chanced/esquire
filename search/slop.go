package search

type WithSlop interface {
	Slop() int
	SetSlop(v int)
}

type SlopParam struct {
	SlopValue *int `json:"slop,omitempty" bson:"slop,omitempty"`
}

func (s SlopParam) Slop() int {
	if s.SlopValue == nil {
		return 0
	}
	return *s.SlopValue
}

func (s *SlopParam) SetSlop(v int) {
	s.SlopValue = &v
}
