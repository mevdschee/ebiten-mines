package layers

import (
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/clips"
	"github.com/mevdschee/minesweeper.go/sprites"
)

// Layer is a set of layers
type Layer struct {
	name  string
	clips []*clips.Clip
}

// LayerJSON is a set of layers in JSON
type LayerJSON struct {
	Name  string
	Clips []clips.ClipJSON
}

// GetName gets the name of the scene
func (l *Layer) GetName() string {
	return l.name
}

// New creates a new layer
func New(name string) *Layer {
	return &Layer{
		name:  name,
		clips: []*clips.Clip{},
	}
}

func eval(expression string, parameters map[string]interface{}) (int, error) {
	if len(expression) == 0 {
		return 0, nil
	}
	value, err := expr.Eval(expression, parameters)
	if err != nil {
		return 0, err
	}
	return value.(int), nil
}

// FromJSON creates a new layer from JSON
func FromJSON(spriteMap sprites.SpriteMap, layerJSON LayerJSON, parameters map[string]interface{}) (*Layer, error) {
	layer := Layer{
		name:  layerJSON.Name,
		clips: []*clips.Clip{},
	}
	for _, clipJSON := range layerJSON.Clips {
		sprite, ok := spriteMap[clipJSON.Sprite]
		if !ok {
			return nil, fmt.Errorf("Could not find sprite '%s' for clip with name '%s'", clipJSON.Sprite, clipJSON.Name)
		}
		repeat, err := eval(clipJSON.Repeat, parameters)
		if err != nil {
			return nil, fmt.Errorf("Repeat in '%s': %v", clipJSON.Repeat, err)
		}
		if repeat == 0 {
			repeat = 1
		}
		for i := 0; i < repeat; i++ {
			parameters["i"] = i
			x, err := eval(clipJSON.X, parameters)
			if err != nil {
				return nil, fmt.Errorf("X in '%s': %v", clipJSON.X, err)
			}
			y, err := eval(clipJSON.Y, parameters)
			if err != nil {
				return nil, fmt.Errorf("Y in '%s': %v", clipJSON.Y, err)
			}
			width, err := eval(clipJSON.Width, parameters)
			if err != nil {
				return nil, fmt.Errorf("Width in '%s': %v", clipJSON.Width, err)
			}
			height, err := eval(clipJSON.Height, parameters)
			if err != nil {
				return nil, fmt.Errorf("Height in '%s': %v", clipJSON.Height, err)
			}
			if width == 0 {
				layer.Add(clips.New(sprite, clipJSON.Name, x, y))
			} else {
				layer.Add(clips.NewScaled(sprite, clipJSON.Name, x, y, width, height))
			}
		}
	}
	return &layer, nil
}

// Add adds a layers to the scene
func (l *Layer) Add(clip *clips.Clip) {
	l.clips = append(l.clips, clip)
}

// Draw draws the layer
func (l *Layer) Draw(screen *ebiten.Image) {
	for _, clip := range l.clips {
		clip.Draw(screen)
	}
}

// Update updates the layer
func (l *Layer) Update() (err error) {
	for _, clip := range l.clips {
		err = clip.Update()
		if err != nil {
			break
		}
	}
	return err
}

// GetClip gets a clip from the layer
func (l *Layer) GetClip(clip string, i int) (*clips.Clip, error) {
	n := 0
	for _, c := range l.clips {
		if c.GetName() == clip {
			if n == i {
				return c, nil
			}
			n++
		}
	}
	return nil, fmt.Errorf("GetClip: clip '%s(%d)' not found", clip, i)
}
