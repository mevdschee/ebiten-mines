package clips

import "github.com/hajimehoshi/ebiten/v2"

// Clip is a set of frames
type Clip struct {
	name   string
	frames []ebiten.Image
}

// New creates a new movie
func New(name string) *Clip {
	return &Clip{name: name}
}
