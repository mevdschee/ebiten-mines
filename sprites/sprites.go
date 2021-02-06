package sprites

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/png"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteMap is a map of sprites
type SpriteMap = map[string]*Sprite

// Sprite is the base struct for any sprite
type Sprite struct {
	Image   *ebiten.Image `json:"-"`
	Name    string        `json:"name"`
	X       int           `json:"x"`
	Y       int           `json:"y"`
	Width   int           `json:"width,omitempty"`
	Height  int           `json:"height,omitempty"`
	Widths  [3]int        `json:"widths,omitempty"`
	Heights [3]int        `json:"heights,omitempty"`
	Count   int           `json:"count"`
	Grid    int           `json:"grid"`
	Gap     int           `json:"gap,omitempty"`
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
func NewSpriteMap(base64image, jsondata string) (SpriteMap, error) {
	image, err := loadImageFromString(base64image)
	if err != nil {
		return nil, err
	}
	sprites := []*Sprite{}
	spriteMap := SpriteMap{}
	err = json.Unmarshal([]byte(jsondata), &sprites)
	if err != nil {
		return nil, err
	}
	for _, sprite := range sprites {
		sprite.Image = image
		spriteMap[sprite.Name] = sprite
	}
	return spriteMap, nil
}
