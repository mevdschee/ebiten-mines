package sprites

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/png"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteMap is a set of sprites
type SpriteMap struct {
	Single map[string]*SingleSprite `json:"single"`
	Series map[string]*SeriesSprite `json:"series"`
	Sliced map[string]*SlicedSprite `json:"sliced"`
}

// Sprite is the base struct for any sprite
type Sprite struct {
	Image *ebiten.Image
	Name  string `json:"name"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
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

// NewSpriteMap creates a new sprite map
func NewSpriteMap(base64image, jsondata string) (*SpriteMap, error) {
	image, err := loadImageFromString(base64image)
	if err != nil {
		return nil, err
	}
	data := SpriteMap{
		Single: map[string]*SingleSprite{},
		Series: map[string]*SeriesSprite{},
		Sliced: map[string]*SlicedSprite{},
	}
	err = json.Unmarshal([]byte(jsondata), &data)
	if err != nil {
		return nil, err
	}
	for _, sprite := range data.Single {
		sprite.Image = image
	}
	for _, sprite := range data.Series {
		sprite.Image = image
	}
	for _, sprite := range data.Sliced {
		sprite.Image = image
	}
	return &data, nil
}
