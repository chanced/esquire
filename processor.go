package picker

// Processor is a standin until I can get around to writing out all the various
// processors
type Processor interface {
	Kind() ProcessorKind
}

// not sure how to deal with this silly naming convetion

type Processorer interface {
	Processor() (Processor, error)
}

type Processors []Processor

func (ps Processors) UnmarshalJSON([]byte) error {
	panic("not impl")
}
