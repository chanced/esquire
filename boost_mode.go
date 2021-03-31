package picker

type BoostMode string

func (bm BoostMode) String() string {
	return string(bm)
}

const (
	BoostModeUnspecified BoostMode = ""
	// query score and function score is multiplied (default)
	BoostModeMultiply BoostMode = "multiply"
	// only function score is used, the query score is ignored
	BoostModeReplace BoostMode = "replace"
	// query score and function score are added
	BoostModeSum BoostMode = "sum"
	// average
	BoostModeAvg BoostMode = "avg"
	// max of query score and function score
	BoostModeMax BoostMode = "max"
	//min of query score and function score
	BoostModeMin BoostMode = "min"
)
