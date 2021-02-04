package sprites

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/png"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteMap is a set of sprites
type SpriteMap struct {
	image  *ebiten.Image
	single map[string]*SingleSprite `json:"single"`
	series map[string]*SeriesSprite `json:"series"`
	sliced map[string]*SlicedSprite `json:"sliced"`
}

// Sprite is the base struct for any sprite
type Sprite struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

// SingleSprite is a single rectangular frames
type SingleSprite struct {
	Sprite
	Width  int `json:"width"`
	Height int `json:"height"`
}

// SeriesSprite is a set of rectangular frames
type SeriesSprite struct {
	SingleSprite
	Count int `json:"count"`
	Gap   int `json:"gap"`
}

// SlicedSprite is a 9 sliced sprite
type SlicedSprite struct {
	Sprite
	Widths  [3]int `json:"widths"`
	Heights [3]int `json:"heights"`
	Gap     int    `json:"gap"`
}

func loadImageFromString(b64 string) (*ebiten.Image, error) {
	b64 = strings.ReplaceAll(b64, "\n", "")
	b64 = strings.ReplaceAll(b64, "\t", "")
	b64 = strings.ReplaceAll(b64, " ", "")
	bin, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(bytes.NewReader(bin))
	if err != nil {
		return nil, err
	}
	img2 := ebiten.NewImageFromImage(img)
	return img2, err
}

func loadJSONFromString(loadJSONFromString string) (*interface{}, error) {
	/*data := &interface{}
		err := json.Unmarshal([]byte(data), &data)
	    if err != nil {
	        log.Fatal("Unmarshal failed", err)
	    }
	    fmt.Println("Done", a)*/
}

// NewSpriteMap creates a new sprite map
func NewSpriteMap(base64image, jsondata string) *SpriteMap {
	return &SpriteMap{
		image:   spriteMap,
		sprites: map[string]*Sprite{},
	}
}

// AddSprite adds a series of rectangular sprites
func (s *SpriteMap) AddSprite(name string, x, y, width, height, gap, count int) error {
	if _, exists := s.sprites[name]; exists {
		return errors.New("duplicate sprite name")
	}
	s.sprites[name] = &Sprite{
		initialFrame: image.Rect(x, y, x+width, y+height),
		gap:          gap,
		count:        count,
	}
	return nil
}

// AddSlicedSprite adds a 9 sliced sprite
func (s *SpriteMap) AddSlicedSprite(name string, w0, w1, w2, h0, h1, h2, horizontalGap, verticalGap int) error {
	if _, exists := s.sprites[name]; exists {
		return errors.New("duplicate sprite name")
	}
	s.slices[name] = &SlicedSprite{
		widths:        [3]int{w0, w1, w2},
		heights:       [3]int{h0, h1, h2},
		horizontalGap: horizontalGap,
		verticalGap:   verticalGap,
	}
	return nil
}
