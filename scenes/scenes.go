package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/layers"
)

// Scene is a set of layers
type Scene struct {
	name   string
	layers map[string]*layers.Layer
	order  []string
}

var (
	currentScene *Scene
	allScenes    map[string]*Scene
)

func init() {
	allScenes = map[string]*Scene{}
}

// GetName gets the name of the scene
func (s *Scene) GetName() string {
	return s.name
}

// GetLayers gets the layers of the scene
func (s *Scene) GetLayers() map[string]*layers.Layer {
	return s.layers
}

// New creates a new scene
func New(name string) *Scene {
	return &Scene{
		name:   name,
		layers: map[string]*layers.Layer{},
		order:  []string{},
	}
}

// Add adds a layers to the scene
func (s *Scene) Add(layer *layers.Layer) {
	name := layer.GetName()
	s.layers[name] = layer
	s.order = append(s.order, name)
}

// Draw draws the scene
func (s *Scene) Draw(screen *ebiten.Image) {
	for _, name := range s.order {
		s.layers[name].Draw(screen)
	}
}

// Update updates the scene
func (s *Scene) Update() (err error) {
	for _, name := range s.order {
		err = s.layers[name].Update()
		if err != nil {
			break
		}
	}
	return err
}
