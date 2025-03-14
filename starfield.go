package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Starfield struct {
	stars []*Star
}

func NewStarfield(numStars int) *Starfield {
	sf := Starfield{
		stars: make([]*Star, 0),
	}
	sf.initialiseStars(numStars)
	return &sf
}

func (s *Starfield) initialiseStars(numStars int) {
	for i := 0; i < numStars; i++ {
		s.stars = append(s.stars, NewStar())
	}
}

func (s *Starfield) Tick() {
	for _, star := range s.stars {
		star.Tick()
	}
}

func (s *Starfield) Render(img *ebiten.Image) {
	size := img.Bounds().Size()
	for _, star := range s.stars {
		x, y, col := star.Get()
		img.Set(scaleToInt(x, size.X), scaleToInt(y, size.Y), col)
	}
}

func scaleToInt(v float32, max int) int {
	return int(v * float32(max))
}
