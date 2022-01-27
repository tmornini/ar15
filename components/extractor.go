package components

type Extractor interface{}

func NewExtractor() Extractor {
	return extractor{}
}

type extractor struct{}
