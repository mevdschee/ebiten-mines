package scenes

import "github.com/mevdschee/minesweeper.go/layers"

// Scene is a set of layers
type Scene struct {
	name   string
	layers []layers.Layer
}

// GetName gets the name of the scene
func (s *Scene) GetName() string {
	return s.name
}

// GetLayers gets the layers of the scene
func (s *Scene) GetLayers() []layers.Layer {
	return s.layers
}

// New creates a new scene
func New(name string) *Scene {
	return &Scene{
		name:   name,
		layers: []layers.Layer{},
	}
}
