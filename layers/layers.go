package layers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/clips"
)

// Layer is a set of layers
type Layer struct {
	name  string
	clips map[string]*clips.Clip
	order []string
}

// GetName gets the name of the scene
func (l *Layer) GetName() string {
	return l.name
}

// New creates a new layer
func New(name string) *Layer {
	return &Layer{
		name:  name,
		clips: map[string]*clips.Clip{},
		order: []string{},
	}
}

// Add adds a layers to the scene
func (l *Layer) Add(clip *clips.Clip) {
	name := clip.GetName()
	l.clips[name] = clip
	l.order = append(l.order, name)
}

// Draw draws the layer
func (l *Layer) Draw(screen *ebiten.Image) {
	for _, name := range l.order {
		l.clips[name].Draw(screen)
	}
}

// Update updates the layer
func (l *Layer) Update() (err error) {
	for _, name := range l.order {
		err = l.clips[name].Update()
		if err != nil {
			break
		}
	}
	return err
}
