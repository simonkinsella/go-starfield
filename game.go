package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	starfield *Starfield
}

func (g *Game) Update() error {
	g.starfield.Tick()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.starfield.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
