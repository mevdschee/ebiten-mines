package layers

import "github.com/mevdschee/minesweeper.go/clips"

// Layer is a set of layers
type Layer struct {
	name  string
	clips map[string]*clips.Clip
}

// GetName gets the name of the scene
func (l *Layer) GetName() string {
	return l.name
}

// New creates a new layer
func New(name string) *Layer {
	return &Layer{name: name}
}

// Add adds a layers to the scene
func (l *Layer) Add(clip *clips.Clip) {
	l.clips[clip.GetName()] = clip
}
