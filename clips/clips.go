package clips

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mevdschee/minesweeper.go/sprites"
)

type InputHandlerFunc func(id int)

// Clip is a set of frames
type Clip struct {
	name             string
	x, y             int
	width, height    int
	frame            int
	frames           []*ebiten.Image
	onPress          InputHandlerFunc
	onLongPress      InputHandlerFunc
	onRelease        InputHandlerFunc
	onReleaseOutside InputHandlerFunc
}

// ClipJSON is a clip in JSON
type ClipJSON struct {
	Name          string
	Sprite        string
	Repeat        string
	X, Y          string
	Width, Height string
}

// GetName gets the name of the clip
func (c *Clip) GetName() string {
	return c.name
}

// New creates a new sprite based clip
func New(sprite *sprites.Sprite, name string, x, y int) *Clip {
	frames := []*ebiten.Image{}

	srcWidth, srcHeight := sprite.Width, sprite.Height
	for i := 0; i < sprite.Count; i++ {
		grid := sprite.Grid
		if grid == 0 {
			grid = sprite.Count
		}
		srcX := sprite.X + (i%grid)*(srcWidth+sprite.Gap)
		srcY := sprite.Y + (i/grid)*(srcHeight+sprite.Gap)
		r := image.Rect(srcX, srcY, srcX+srcWidth, srcY+srcHeight)
		frame := sprite.Image.SubImage(r).(*ebiten.Image)
		frames = append(frames, frame)
	}

	return &Clip{
		name:   name,
		x:      x,
		y:      y,
		width:  srcWidth,
		height: srcHeight,
		frame:  0,
		frames: frames,
	}
}

// NewScaled creates a new 9 slice scaled sprite based clip
func NewScaled(sprite *sprites.Sprite, name string, x, y, width, height int) *Clip {
	frame0 := ebiten.NewImage(width, height)

	srcY := sprite.Y
	dstY := 0
	for h := 0; h < 3; h++ {
		srcHeight := sprite.Heights[h]
		dstHeight := sprite.Heights[h]
		if h == 1 {
			dstHeight = height - sprite.Heights[0] - sprite.Heights[2]
		}
		srcX := sprite.X
		dstX := 0
		for w := 0; w < 3; w++ {
			srcWidth := sprite.Widths[w]
			dstWidth := sprite.Widths[w]
			if w == 1 {
				dstWidth = width - sprite.Widths[0] - sprite.Widths[2]
			}

			r := image.Rect(srcX, srcY, srcX+srcWidth, srcY+srcHeight)
			img := sprite.Image.SubImage(r).(*ebiten.Image)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(float64(dstWidth)/float64(srcWidth), float64(dstHeight)/float64(srcHeight))
			op.GeoM.Translate(float64(dstX), float64(dstY))
			frame0.DrawImage(img, op)

			srcX += srcWidth + sprite.Gap
			dstX += dstWidth
		}
		srcY += srcHeight + sprite.Gap
		dstY += dstHeight
	}

	return &Clip{
		name:   name,
		x:      x,
		y:      y,
		width:  width,
		height: height,
		frame:  0,
		frames: []*ebiten.Image{frame0},
	}
}

// Draw draws the clip
func (c *Clip) Draw(screen *ebiten.Image) {
	img := c.frames[c.frame]
	srcWidth, srcHeight := img.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(c.width)/float64(srcWidth), float64(c.height)/float64(srcHeight))
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(img, op)
}

// GotoFrame goes to a frame of the clip
func (c *Clip) GotoFrame(frame int) {
	c.frame = frame
}

// OnPress sets the click handler function
func (c *Clip) OnPress(handler InputHandlerFunc) {
	c.onPress = handler
}

// OnLongPress sets the click handler function
func (c *Clip) OnLongPress(handler InputHandlerFunc) {
	c.onLongPress = handler
}

// OnRelease sets the click handler function
func (c *Clip) OnRelease(handler InputHandlerFunc) {
	c.onRelease = handler
}

// OnReleaseOutside sets the click handler function
func (c *Clip) OnReleaseOutside(handler InputHandlerFunc) {
	c.onReleaseOutside = handler
}

// IsHovered returns whether or not the cursor is hovering the clip
func (c *Clip) IsHovered() bool {
	cursorX, cursorY := ebiten.CursorPosition()
	cursor := image.Point{cursorX, cursorY}
	rect := image.Rect(c.x, c.y, c.x+c.width, c.y+c.height)
	return cursor.In(rect)
}

// Update updates the clip
func (c *Clip) Update() (err error) {
	log.Println(inpututil.JustPressedTouchIDs())
	if c.IsHovered() {
		if c.onPress != nil {
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				c.onPress(-1)
			}
		}
		if c.onLongPress != nil {
			if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == ebiten.MaxTPS()/2 {
				c.onLongPress(-1)
			}
		}
		if c.onRelease != nil {
			if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
				c.onRelease(-1)
			}
		}
	} else {
		if c.onReleaseOutside != nil {
			if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
				c.onReleaseOutside(-1)
			}
		}
	}
	//if c.fps > 0 && c.ticks%(60/c.fps) == 0 {
	//	c.frame = (c.frame + 1) % len(c.frames)
	//}
	//if moving do:
	//c.x++
	//if resizing do:
	//c.height--
	return nil
}
