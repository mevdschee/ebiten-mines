package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/mevdschee/minesweeper.go/clips"
	"github.com/mevdschee/minesweeper.go/layers"
	"github.com/mevdschee/minesweeper.go/movies"
	"github.com/mevdschee/minesweeper.go/scenes"
	"github.com/mevdschee/minesweeper.go/sprites"
)

var spriteMapImage = `
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

var spriteMapMeta = `
{
	"single": [
		{
			"name": "display",
			"x": 28,
			"y": 82,
			"width": 41,
			"height": 25
		}
	],
	"series": [
		{
			"name": "numbers",
			"x": 0,
			"y": 0,
			"width": 16,
			"height": 16,
			"count": 9,
			"gap": 0
		},
		{
			"name": "icons",
			"x": 0,
			"y": 16,
			"width": 16,
			"height": 16,
			"count": 8,
			"gap": 0
		},
		{
			"name": "digits",
			"x": 0,
			"y": 33,
			"width": 11,
			"height": 21,
			"count": 11,
			"gap": 1
		},
		{
			"name": "buttons",
			"x": 0,
			"y": 55,
			"width": 26,
			"height": 26,
			"count": 5,
			"gap": 1
		}
	],
	"sliced": [
		{
			"name": "controls",
			"x": 0,
			"y": 82,
			"widths": [12,1,12],
			"heights": [11,1,11],
			"gap": 1
		},
		{
			"name": "field",
			"x": 0,
			"y": 96,
			"widths": [12,1,12],
			"heights": [11,1,11],
			"gap": 1
		}
	]
}`

type config struct {
	scale   int
	width   int
	height  int
	bombs   int
	holding int
}

type game struct {
	c     config
	movie *movies.Movie
}

func (g *game) getSize() (int, int) {
	return g.c.width*16 + 12*2, g.c.height*16 + 11*3 + 33
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.getSize()
}

func (g *game) init() *game {
	spriteMap, err := sprites.NewSpriteMap(spriteMapImage, spriteMapMeta)
	if err != nil {
		log.Fatalln(err)
	}
	g.movie = movies.New()
	gameScene := scenes.New("game")
	backgroundLayer := layers.New("bg")
	controlsClip := clips.NewSlice(spriteMap.Sliced["controls"], 0, 0, 100, 40)
	g.movie.Add(gameScene)
	gameScene.Add(backgroundLayer)
	backgroundLayer.Add(controlsClip)
	return g
}

func main() {
	ebiten.SetWindowTitle("Minesweeper.go")
	g := &game{c: config{
		scale:   5,
		width:   9,
		height:  9,
		bombs:   10,
		holding: 15,
	}}
	width, height := g.getSize()
	ebiten.SetMaxTPS(30)
	ebiten.SetWindowSize(g.c.scale*width, g.c.scale*height)
	if err := ebiten.RunGame(g.init()); err != nil {
		log.Fatalf("%v\n", err)
	}
}
