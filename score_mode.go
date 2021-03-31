package picker

type ScoreMode string

func (sm ScoreMode) String() string {
	return string(sm)
}

const (
	ScoreModeUnspecified ScoreMode = ""
	// ScoreModeMultiply - scores are multiplied (default)
	ScoreModeMultiply ScoreMode = "multiply"
	// ScoreModeSum - scores are summed
	ScoreModeSum ScoreMode = "sum"
	// ScoreModeAvg scores are averaged
	ScoreModeAvg ScoreMode = "avg"
	// ScoreModeFirst - the first function that has a matching filter is applied
	ScoreModeFirst ScoreMode = "first"
	// ScoreModeMax - maximum score is used
	ScoreModeMax ScoreMode = "max"
	// ScoreModeMin - minimum score is used
	ScoreModeMin ScoreMode = "min"
)
