package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var skin = `
	iVBORw0KGgoAAAANSUhEUgAAAJAAAAB6CAMAAABnRypuAAADAFBMVEUAAACA
	AAAAgACAgAAAAICAAIAAgIDAwMCAgID/AAAA/wD//wAAAP//AP8A//////8Q
	EBARERESEhITExMUFBQVFRUWFhYXFxcYGBgZGRkaGhobGxscHBwdHR0eHh4f
	Hx8gICAhISEiIiIjIyMkJCQlJSUmJiYnJycoKCgpKSkqKiorKyssLCwtLS0u
	Li4vLy8wMDAxMTEyMjIzMzM0NDQ1NTU2NjY3Nzc4ODg5OTk6Ojo7Ozs8PDw9
	PT0+Pj4/Pz9AQEBBQUFCQkJDQ0NERERFRUVGRkZHR0dISEhJSUlKSkpLS0tM
	TExNTU1OTk5PT09QUFBRUVFSUlJTU1NUVFRVVVVWVlZXV1dYWFhZWVlaWlpb
	W1tcXFxdXV1eXl5fX19gYGBhYWFiYmJjY2NkZGRlZWVmZmZnZ2doaGhpaWlq
	ampra2tsbGxtbW1ubm5vb29wcHBxcXFycnJzc3N0dHR1dXV2dnZ3d3d4eHh5
	eXl6enp7e3t8fHx9fX1+fn5/f3+AgICBgYGCgoKDg4OEhISFhYWGhoaHh4eI
	iIiJiYmKioqLi4uMjIyNjY2Ojo6Pj4+QkJCRkZGSkpKTk5OUlJSVlZWWlpaX
	l5eYmJiZmZmampqbm5ucnJydnZ2enp6fn5+goKChoaGioqKjo6OkpKSlpaWm
	pqanp6eoqKipqamqqqqrq6usrKytra2urq6vr6+wsLCxsbGysrKzs7O0tLS1
	tbW2tra3t7e4uLi5ubm6urq7u7u8vLy9vb2+vr6/v7/AwMDBwcHCwsLDw8PE
	xMTFxcXGxsbHx8fIyMjJycnKysrLy8vMzMzNzc3Ozs7Pz8/Q0NDR0dHS0tLT
	09PU1NTV1dXW1tbX19fY2NjZ2dna2trb29vc3Nzd3d3e3t7f39/g4ODh4eHi
	4uLj4+Pk5OTl5eXm5ubn5+fo6Ojp6enq6urr6+vs7Ozt7e3u7u7v7+/w8PDx
	8fHy8vLz8/P09PT19fX29vb39/f4+Pj5+fn6+vr7+/v8/Pz9/f3+/v7///9Q
	NrN6AAAE5UlEQVR4nO2bgZKjIAxAY7ezVu2M//+3pxAgCUHQUnFvzM61QGJ4
	Bgjo7sHrYgKvnkvr+t8Ber9D/YGC9cGJt38+n+afrXdeUP+Lgnrwgno3XFtA
	bwL0eBAiCjSUAf1mgMIESgO9BVDfu69gH4AMCwNi/hCFj4Dh8kB97750oDcD
	QmkJxIfM4TxCnc2hp5eeD1kAYkNGeQ7NIcR5kHoZkJxDv7WAAo4+ZEbCkKFQ
	oLXOgZDn0JARnFZA77ef1kXLngOVDJnjKVz2hUAsIgqQ0yuTmgIVzaFm9Rso
	CzRzITNMl9r2oxCI7DN3kLUfuOTsFSDmLgkEHki3d3oHBB5o238tILvmjb0t
	ArbqQOA2ijpAYI8wplNj75KQtV/KVr82vgyMqRsoY7/S2L3d+XeEByMUzlMG
	aOBAS90CmMxtwgPYkALyDQcjtFg5IiVC/dLzvHYPCOTt6Rz6KhCbQ6vezBnY
	Aorn0GEgoOJXGXh717XVvwZm74AApP9P5hA9lMdARm+BfITCJEoBuc21EhBf
	xvGQRUCR/4+A5CoLfGhvGPA4oq0ybcj6TyIk8pAEkss+ykO1h8w4wG/F3oBA
	3/vEWLJ1pIHk/psDiuwzW0fGfwQUnQcSQPQmN+3Fbp+zj4BkQ2sB8PcCS6XD
	n6VM2/falJSJHwbkBxuGsXNWpkzbQ3kgNuHajpW5vV4OfXUCCK2WzxHQCsu0
	3ZeHIdj4aztgZWGvln1fnYyQTyIrdZctU3tX7tYIkXLOnvrs4ggFI7TZLGsd
	2AiFcs6e+rwj9D9ECNztBOqNMrX3ZbvKoggl7TcjxHJPSR7SyjIPablK5qFg
	zzM1zbxfyMIl9kURSmXeksilIqRfWziHUnOiZG6VtO9eZa5drpqS1VfSvjsP
	uXaZV0ryU0n7HaH/KUKAK5Gfh1LnG2oDiXOP6ic+Y4F+HsrmD5mH9Gy+Pw85
	G5mpBycVz9Q7bRjQeDEB+dgUnp/GM1UEaNalf431VUnJAy0PnCnXH6iix2D3
	uMuBVBPrurLqrwOZzKm7rqTaBQQwrRLeLI31VQSIwqpA9mrjQbqupyKvYCms
	BhQutw6I64oq//aLwypA9HJjFVwvqh/4mfBTqmBaLzWfUqU4fEkeFz4VCMKH
	7BVVWq8Ye/0qKxIogo2BYGJAEwTXMLEISRWJkFBZFk8kgNjdJ4AYNev1iMqy
	4AcDiu/+DCAz82ynP2Y1CSA+ZgeAACCl2sG6C8jZxbNhWu92MjNIUeEsqQuU
	WS/YOmmqjatilQfid5/KQ/5eYyAfoTgPec/KbUiHZJVRniOZOnj+PFNHsM33
	soJMHe13393t83sZejjtPJTf7aV8/cS4cR6KfoHkngTqq4pOjKnLl+urq0qA
	kpe/xjNVAWi8mLA/aBqXH/Jot/GkyaUfo99UfgFoffSFMpmr4WwAmUdfIC9N
	IF2pDeQn+xJ5njeAvFaCdOVEIPLiDdKVUyPkX01CunJyhLBfSFfuCLWJkCEI
	2037CCFQPzeO0JaybR6yL2f7mUfo/Ey9pWwcoR73squssn5WgO48dGfqT4Eu
	GaGLnYdaZmplc215ptY31yL5JhDby8T/KklLTaCKvqoIj9A4imoDoCa9bgg/
	D80XiNDW8aMtkHL8aAwUZ+ob6Aa6geoAhZe1I/+LmkZA5OXxJYDoy+wrAP0D
	7BQA5IVg6IYAAAAASUVORK5CYII=`

const (
	iconClosed = iota
	iconOpened
	iconBomb
	iconMarked
	iconAnswerNoBomb
	iconAnswerIsBomb
	iconQuestionMark
	iconQuestionPressed
)

const (
	buttonPlaying = iota
	buttonEvaluate
	buttonLost
	buttonWon
	buttonPressed
)

const (
	stateThinking = iota
	statePlaying
	stateLost
	stateWon
)

type tiles struct {
	numbers    []*ebiten.Image
	icons      []*ebiten.Image
	digits     []*ebiten.Image
	buttons    []*ebiten.Image
	background *ebiten.Image
}

type game struct {
	scale   int
	width   int
	height  int
	bombs   int
	holding int
	tiles   tiles
	field   [][]int
}

type transform struct {
	src, dst image.Rectangle
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

func (g *game) Update() error {
	return nil
}

func (g *game) loadBackgroundTile(tilesImage *ebiten.Image) {
	w := g.width
	h := g.height
	width, height := g.getSize()
	background := ebiten.NewImage(width, height)
	transforms := []transform{
		{
			image.Rectangle{image.Point{0, 82}, image.Point{12, 93}},
			image.Rectangle{image.Point{0, 0}, image.Point{12, 11}},
		},
		{
			image.Rectangle{image.Point{13, 82}, image.Point{14, 93}},
			image.Rectangle{image.Point{12, 0}, image.Point{12 + w*16, 11}},
		},
		{
			image.Rectangle{image.Point{15, 82}, image.Point{27, 93}},
			image.Rectangle{image.Point{12 + w*16, 0}, image.Point{12 + w*16 + 12, 11}},
		},
		{
			image.Rectangle{image.Point{0, 94}, image.Point{12, 95}},
			image.Rectangle{image.Point{0, 11}, image.Point{12, 11 + 33}},
		},
		{
			image.Rectangle{image.Point{15, 94}, image.Point{27, 95}},
			image.Rectangle{image.Point{12 + w*16, 11}, image.Point{12 + w*16 + 12, 11 + 33}},
		},
		{
			image.Rectangle{image.Point{0, 96}, image.Point{12, 107}},
			image.Rectangle{image.Point{0, 11 + 33}, image.Point{12, 11 + 33 + 11}},
		},

		{
			image.Rectangle{image.Point{13, 96}, image.Point{14, 107}},
			image.Rectangle{image.Point{12, 11 + 33}, image.Point{12 + w*16, 11 + 33 + 11}},
		},
		{
			image.Rectangle{image.Point{15, 96}, image.Point{27, 107}},
			image.Rectangle{image.Point{12 + w*16, 11 + 33}, image.Point{12 + w*16 + 12, 11 + 33 + 11}},
		},
		{
			image.Rectangle{image.Point{0, 108}, image.Point{12, 109}},
			image.Rectangle{image.Point{0, 11 + 33 + 11}, image.Point{12, 11 + 33 + 11 + h*16}},
		},
		{
			image.Rectangle{image.Point{15, 108}, image.Point{27, 109}},
			image.Rectangle{image.Point{12 + w*16, 11 + 33 + 11}, image.Point{12 + w*16 + 12, 11 + 33 + 11 + h*16}},
		},
		{
			image.Rectangle{image.Point{0, 110}, image.Point{12, 121}},
			image.Rectangle{image.Point{0, 11 + 33 + 11 + h*16}, image.Point{12, 11 + 33 + 11 + h*16 + 11}},
		},
		{
			image.Rectangle{image.Point{13, 110}, image.Point{14, 121}},
			image.Rectangle{image.Point{12, 11 + 33 + 11 + h*16}, image.Point{12 + w*16, 11 + 33 + 11 + h*16 + 11}},
		},
		{
			image.Rectangle{image.Point{15, 110}, image.Point{27, 121}},
			image.Rectangle{image.Point{12 + w*16, 11 + 33 + 11 + h*16}, image.Point{12 + w*16 + 12, 11 + 33 + 11 + h*16 + 11}},
		},
		{
			image.Rectangle{image.Point{28, 82}, image.Point{69, 107}},
			image.Rectangle{image.Point{12 + 4, 11 + 4}, image.Point{12 + 4 + 41, 11 + 4 + 25}},
		},
		{
			image.Rectangle{image.Point{28, 82}, image.Point{69, 107}},
			image.Rectangle{image.Point{12 + w*16 - 4 - 41, 11 + 4}, image.Point{12 + w*16 - 4, 11 + 4 + 25}},
		},
	}
	for _, t := range transforms {
		source := tilesImage.SubImage(t.src).(*ebiten.Image)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(t.dst.Size().X)/float64(t.src.Size().X), float64(t.dst.Size().Y)/float64(t.src.Size().Y))
		op.GeoM.Translate(float64(t.dst.Min.X), float64(t.dst.Min.Y))
		background.DrawImage(source, op)
	}
	g.tiles.background = background
}

func (g *game) init() *game {
	tilesImage, err := loadImageFromString(skin)
	if err != nil {
		log.Fatalln("Could not load skin")
	}
	g.tiles.numbers = make([]*ebiten.Image, 9)
	for x := range g.tiles.numbers {
		g.tiles.numbers[x] = tilesImage.SubImage(image.Rect(x*16, 0, x*16+16, 16)).(*ebiten.Image)
	}
	g.tiles.icons = make([]*ebiten.Image, 8)
	for x := range g.tiles.icons {
		g.tiles.icons[x] = tilesImage.SubImage(image.Rect(x*16, 16, x*16+16, 32)).(*ebiten.Image)
	}
	g.tiles.digits = make([]*ebiten.Image, 11)
	for x := range g.tiles.digits {
		g.tiles.digits[x] = tilesImage.SubImage(image.Rect(x*12, 33, x*12+11, 54)).(*ebiten.Image)
	}
	g.tiles.buttons = make([]*ebiten.Image, 5)
	for x := range g.tiles.buttons {
		g.tiles.buttons[x] = tilesImage.SubImage(image.Rect(x*27, 55, x*27+26, 81)).(*ebiten.Image)
	}
	g.loadBackgroundTile(tilesImage)
	g.initField()
	return g
}

func (g *game) initField() {
	g.field = make([][]int, g.width)
	for x := range g.field {
		g.field[x] = make([]int, g.height)
	}
}

func (g *game) drawField(screen *ebiten.Image) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(12+x*16), float64(11+33+11+y*16))
			screen.DrawImage(g.tiles.icons[g.field[x][y]], op)
		}
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{0xcc, 0xcc, 0xcc, 0xcc})
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.tiles.background, op)
	g.drawField(screen)
}

func (g *game) getSize() (int, int) {
	return g.width*16 + 12*2, g.height*16 + 11*3 + 33
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.getSize()
}

func main() {
	ebiten.SetWindowTitle("Minesweeper.go")
	game := &game{
		scale:   3,
		width:   9,
		height:  9,
		bombs:   10,
		holding: 15,
	}
	width, height := game.getSize()
	ebiten.SetWindowSize(game.scale*width, game.scale*height)
	if err := ebiten.RunGame(game.init()); err != nil {
		log.Fatalf("%v\n", err)
	}
}
