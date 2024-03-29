package scenes

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/clips"
	"github.com/mevdschee/minesweeper.go/layers"
	"github.com/mevdschee/minesweeper.go/sprites"
)

// Scene is a set of layers
type Scene struct {
	name   string
	layers map[string]*layers.Layer
	order  []string
}

// SceneJSON is a set of layers in JSON
type SceneJSON struct {
	Name   string
	Layers []layers.LayerJSON
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

// FromJSON creates a new scene from JSON
func FromJSON(spriteMap sprites.SpriteMap, sceneJSON SceneJSON, parameters map[string]interface{}) (*Scene, error) {
	scene := Scene{
		name:   sceneJSON.Name,
		layers: map[string]*layers.Layer{},
		order:  []string{},
	}
	for _, layerJSON := range sceneJSON.Layers {
		layer, err := layers.FromJSON(spriteMap, layerJSON, parameters)
		if err != nil {
			return nil, err
		}
		scene.Add(layer)
	}
	return &scene, nil
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

// GetClip gets a clip from the scene
func (s *Scene) GetClip(layer, clip string, i int) (*clips.Clip, error) {
	if l, ok := s.layers[layer]; ok {
		return l.GetClip(clip, i)
	}
	return nil, fmt.Errorf("GetClip: layer '%s' not found", layer)
}
