package search

// IndicesBoost boosts the _score of documents from specified indices
type IndicesBoost map[string]float64

func (ib IndicesBoost) Clone() IndicesBoost {
	res := IndicesBoost{}
	for k, v := range ib {
		res[k] = v
	}
	return res
}

func NewIndicesBoost() map[string]float64 {
	return map[string]float64{}
}
