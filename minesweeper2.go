package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mevdschee/minesweeper.go/clips"
	"github.com/mevdschee/minesweeper.go/layers"
	"github.com/mevdschee/minesweeper.go/movies"
	"github.com/mevdschee/minesweeper.go/scenes"
	"github.com/mevdschee/minesweeper.go/sprites"
)

var spriteMapImage = `
	iVBORw0KGgoAAAANSUhEUgAAAJAAAAB6BAMAAACit8dvAAAAMFBMVEUAAACAAAAAgACAgAAAAICA
	AIAAgIDAwMCAgID/AAAA/wD//wAAAP//AP8A//////9PEyZJAAAEtklEQVRo3u2YO2/bSBCARxVB
	s+FVAVymSeverYH8k6RlAiNLGQGcToqaUyfQTUoJacyef8H/R8A1unnsi7skQ4vM4QJ7JSw5D36a
	3Z0dLgTrmRqslLSp198DeuLrW2x4XVbYSL69ub3B66e/sL1B+QM2lAvAlqBMA2uDvjLoM4HeadAu
	Bn0JQCWBVi3QkwYpRSBSMujmvQaJ35cPbkgCUqoN+iogEqaB9NCI85avMke3N9je66EpNzT0K0AN
	zBFx3vWBzBx9HANijjc0pXhoJAtIKQ3CkQ0M7TNzpoBwsp+6ll+D4qHRyLqWvwukApByCWlAHXP0
	P9xrk0AnacewvozV/9TNgMrwF0L990paqHcgkdc+KCGQp080KCVQ4D8GtNyRviJ5ufNAAONACZaZ
	BEHLikHLHcoVRpQCpCmCiqRI2L+gzB4AUcWiiCoBqQoBOwRtUV/5IMDPEAhOVPlsRCr9AVWCINKn
	PEfPBOk5UlCl2zbIztEgCLglJ5n09YqeRvm76NMTByT+xfAcSVG3IEUgjgh0RBYEahxIL7MdmgFZ
	/1+AzKrRg6RfbqFS3qq5iEbmkQGZ5Td5NHpowRbBtUtU1blFWiCzjUOQ1QdbJPS3oLIDRC3Um91f
	9oF+ztSAwq0uIMd2Ifd9mrgXHw3aomb7kC8WC3jI5Z77LWvYeim91ns9P3WpQTSD6cMC4djLPfcV
	a8i6uJBe672enlqYiCjPqoc8D3vRU59fSh/qxTO/bIEWEPbuAYyI+1Avni8iIioOzG73le0xIjAR
	+fq0HZHkS0ceud7kkcssk0es1yDJzvF5HOs7IvIzOI7OjyiPI6qCeTFzEc9XfN+7av7qxCsY3/fm
	kZ8vcU7F9y8mIq8e+XXHr1Chj61cfj2KcsTkkZfrvXnkV8jxNbtPM/dbxD/6rg6+dDfWJKC/T66V
	B084bgZMraN3BDqufe9yM2BaK9dKDTo6DXo7+3ozYPrtoIKOVtabDmwW1G3qARVZ04D1hqapEwPq
	MQlIsB4InZsm0d4FCbUB9ZgYpLEOxHZ0EG9+FJ9lUJ+JQAbrQGJvEvYumqtr/NYCyuq6wa8xaayA
	DNYHQVNn1htQzgyoQXfPhC3TIIu1INQwqGbvTCJqEgIVjUSkTTUAfsnEoJojaIGErUF6oAJqmTIC
	ZZmAMgFhtM8GQXOFj17XBmSmthdUQ09E7jd6QDX5Ou+r+tqBaFJGgsKlaUjuXrWaf1BpEKGyNoh/
	1oIoIptHHKw14eo3dvk1tj+zaY5+ndkWO3mvxZl95u6P99q59Sja/WdXyHY9Kr33wcH/c2YzYOqo
	kCv/xXLwH90MmDpAvn2994XNWNPMr2x55T0+8uWu67/zu7tng46bbxC3f54fUYkgOqnggY5OPnKY
	Hw3iZNnvJT++8dkpp6MYdSycB+LTHBACchDhzIjofEkYIBQL50aECHycJghE+KMjwr19WM0REYHu
	p0bUFifmUbna708S0fmZPW9E97jXZsmjkwW9ZvZ/GtFc9WhaHrlNO61me5t2wltEQHqv3Z/idtZe
	m/Km5Yj2eotMAM0b0epxekTtMjIR5MrIVJDN7FfQHwLik+5B/h6aApID/QygcibQv7SdAex29G+U
	AAAAAElFTkSuQmCC`

var spriteMapMeta = `
	[
		{
			"name": "display",
			"x": 28,
			"y": 82,
			"width": 41,
			"height": 25,
			"count": 1,
			"gap": 0
		},
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
		},
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
	]`

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
	w, h := g.getSize()
	spriteMap, err := sprites.NewSpriteMap(spriteMapImage, spriteMapMeta)
	if err != nil {
		log.Fatalln(err)
	}
	g.movie = movies.New()
	game := scenes.New("game")
	bg := layers.New("bg")
	g.movie.Add(game)
	game.Add(bg)
	bg.Add(clips.NewScaled(spriteMap, "controls", 0, 0, w, 55))
	bg.Add(clips.NewScaled(spriteMap, "field", 0, 44, w, h-44))
	bg.Add(clips.New(spriteMap, "display", 16, 15))
	//bg.Add(clips.New(spriteMap, "display", w*16-33, 15))
	return g
}

func (g *game) Update() error {
	return g.movie.Update()
}

func (g *game) Draw(screen *ebiten.Image) {
	g.movie.Draw(screen)
}

func main() {
	ebiten.SetWindowTitle("Minesweeper.go")
	g := &game{c: config{
		scale:   3,
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
