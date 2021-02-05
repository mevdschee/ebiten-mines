package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mevdschee/minesweeper.go/layers"
)

// Scene is a set of layers
type Scene struct {
	name   string
	layers map[string]*layers.Layer
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
	}
}

// Add adds a layers to the scene
func (s *Scene) Add(layer *layers.Layer) {
	s.layers[layer.GetName()] = layer
}

// Draw draws the scene
func (s *Scene) Draw(screen *ebiten.Image) {

}
