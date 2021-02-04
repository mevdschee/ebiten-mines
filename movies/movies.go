package movies

import (
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
