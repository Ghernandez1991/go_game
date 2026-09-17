package main

import (
	"log"

	"github.com/Ghernandez1991/go_game/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &ui.Game{}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("My Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
