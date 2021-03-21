package search

// IndicesBoost boosts the _score of documents from specified indices
type IndicesBoost map[string]float32

func NewIndicesBoost() map[string]float32 {
	return map[string]float32{}
}
