package picker

import "github.com/chanced/dynamic"

const DefaultPhraseSlop = 0

type WithPhraseSlop interface {
	PhraseSlop() int
	SetPhraseSlop(v interface{}) error
}

type phraseSlopParam struct {
	phraseSlop dynamic.Number
}

func (mds phraseSlopParam) PhraseSlop() int {
	if i, ok := mds.phraseSlop.Int(); ok {
		return i
	}
	if f, ok := mds.phraseSlop.Float64(); ok {
		return int(f)
	}
	return DefaultPhraseSlop
}
func (mds *phraseSlopParam) SetPhraseSlop(v interface{}) error {
	return mds.phraseSlop.Set(v)
}
