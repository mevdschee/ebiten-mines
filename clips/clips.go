package clips

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/sprites"
)

// Clip is a set of frames
type Clip struct {
	name          string
	x, y          int
	width, height int
	frames        []*ebiten.Image
}

// GetName gets the name of the clip
func (c *Clip) GetName() string {
	return c.name
}

// NewSlice creates a new slice sprite based clip
func NewSlice(sprite sprites.SlicedSprite, x, y, width, height int) *Clip {
	frame0 := ebiten.NewImage(width, height)

	img := sprite.Image.SubImage(image.Rect(sprite.X, sprite.Y, sprite.X+sprite.Widths[0], sprite.Y+sprite.Heights[0])).(*ebiten.Image)
	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Scale(float64(t.dst.Size().X)/float64(t.src.Size().X), float64(t.dst.Size().Y)/float64(t.src.Size().Y))
	op.GeoM.Translate(float64(x), float64(y))
	frame0.DrawImage(img, op)

	return &Clip{
		name:   sprite.Name,
		x:      x,
		y:      y,
		width:  width,
		height: height,
		frames: []*ebiten.Image{frame0},
	}
}
