package movies

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/scenes"
)

// Movie is a set of scenes
type Movie struct {
	currentScene *scenes.Scene
	scenes       map[string]*scenes.Scene
}

// New creates a new movie
func New() *Movie {
	return &Movie{
		currentScene: nil,
		scenes:       map[string]*scenes.Scene{},
	}
}

// Add adds a scene to the movie
func (m *Movie) Add(scene *scenes.Scene) {
	m.scenes[scene.GetName()] = scene
	if len(m.scenes) == 1 {
		m.currentScene = scene
	}
}

// Draw draws the movie
func (m *Movie) Draw(screen *ebiten.Image) {
	if m.currentScene != nil {
		m.currentScene.Draw(screen)
	}
}

// Update updates the movie
func (m *Movie) Update() (err error) {
	if m.currentScene != nil {
		err = m.currentScene.Update()
	}
	return err
}
