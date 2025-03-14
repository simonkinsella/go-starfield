package main

import (
	"image/color"
	"math/rand"
)

const (
	velocityScalar = 1000 // higher values mean slower stars
	minIntensity   = 20
	maxIntensity   = 256
)

type Star struct {
	x                float32
	y                float32
	z                float32 // lower values mean further away
	apparentVelocity float32
	intensity        uint8
}

func NewStar() *Star {
	z := rand.Float32()
	intensity := uint8(rand.Float32() * z * maxIntensity)
	if intensity < minIntensity {
		intensity = minIntensity
	}
	return &Star{
		x:                rand.Float32(),
		y:                rand.Float32(),
		z:                z,
		apparentVelocity: rand.Float32() * z / velocityScalar,
		intensity:        intensity,
	}
}

func (s *Star) Tick() {
	s.x += s.apparentVelocity
	if s.x > 1.0 {
		s.x = 0
	}
}

func (s *Star) Get() (x, y float32, col color.Color) {
	return s.x, s.y, color.RGBA{
		R: s.intensity,
		G: s.intensity,
		B: s.intensity,
		A: 0,
	}
}
