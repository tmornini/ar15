// Package components models components
package components

// Extractor is the interface for extractor components
type Extractor interface{}

// NewExtractor returns a new extractor
func NewExtractor() Extractor {
	return extractor{}
}

type extractor struct{}
