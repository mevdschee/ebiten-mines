package layers

import "github.com/mevdschee/minesweeper.go/clips"

// Layer is a set of layers
type Layer struct {
	name   string
	movies []clips.Clip
}

// New creates a new layer
func New(name string) *Layer {
	return &Layer{name: name}
}
