package movies

import (
	"github.com/hajimehoshi/ebiten"
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
}

// Draw draws the movie
func (m *Movie) Draw(screen *ebiten.Image) {
	if currentScene != nil {
		currentScene.Draw(screen)
	}
}
