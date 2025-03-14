package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	numStars := flag.Int("n", 100, "number of stars (default 100)")
	width := flag.Int("width", 640, "window width (default 640)")
	height := flag.Int("height", 480, "window height (default 480)")
	flag.Parse()

	game := Game{
		starfield: NewStarfield(*numStars),
	}

	ebiten.SetWindowSize(*width, *height)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
